#!/usr/bin/env python3
"""Tests for publish_pulumi_module.sh: run with `python3 -m unittest tools/ci/release/publish_pulumi_module_test.py`.

The fake `aws` from publish_once_test.py stands in for R2, and a fake
modulebuild writes bytes derived from the module directory (and fails for a
module whose directory says "broken"), so what is published, and what the
part file records, can be checked exactly.
"""
import hashlib
import os
import subprocess
import sys
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
SCRIPT = os.path.join(HERE, "publish_pulumi_module.sh")
sys.path.insert(0, HERE)
from publish_once_test import FAKE_AWS  # noqa: E402

FAKE_MODULEBUILD = r'''#!/usr/bin/env bash
set -euo pipefail
cmd="$1"; shift
module=""; out=""; platform=""
while [ $# -gt 0 ]; do
  case "$1" in
    -module) module="$2"; shift ;;
    -out) out="$2"; shift ;;
    -platform) platform="$2"; shift ;;
  esac
  shift
done
case "${module}" in *broken*) echo "build failed" >&2; exit 1 ;; esac
printf '%s %s %s' "${cmd}" "${module}" "${platform}" > "${out}"
'''


def sha256(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


class PublishPulumiModuleTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.store = os.path.join(self.tmp.name, "r2")
        self.parts = os.path.join(self.tmp.name, "parts")
        bindir = os.path.join(self.tmp.name, "bin")
        for d in (self.store, self.parts, bindir):
            os.mkdir(d)
        for name, body in (("aws", FAKE_AWS), ("modulebuild", FAKE_MODULEBUILD)):
            path = os.path.join(bindir, name)
            with open(path, "w") as f:
                f.write(body)
            os.chmod(path, 0o755)
        self.env = dict(os.environ, PATH=bindir + os.pathsep + os.environ["PATH"],
                        FAKE_R2=self.store, R2_BUCKET="bucket", AWS_ENDPOINT_URL="https://r2",
                        TAG="v1.2.3", MODULEBUILD=os.path.join(bindir, "modulebuild"), PARTS_DIR=self.parts)

    def tearDown(self):
        self.tmp.cleanup()

    def run_script(self, module_dir, platform):
        return subprocess.run([SCRIPT, module_dir], env=dict(self.env, PLATFORM=platform),
                              capture_output=True, text=True)

    def published(self, key):
        path = os.path.join(self.store, key.replace("/", "_"))
        if not os.path.exists(path):
            return None
        with open(path, "rb") as f:
            return f.read()

    def part(self, kind):
        with open(os.path.join(self.parts, kind + ".tsv")) as f:
            return f.read()

    def test_publishes_the_binary_and_its_source_and_records_the_published_checksum(self):
        result = self.run_script("catalog/aws/awsvpc/iac/pulumi", "linux_amd64")
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        binary = self.published("releases/v1.2.3/modules/pulumi/awsvpc/linux_amd64.gz")
        self.assertEqual(binary, b"pulumi catalog/aws/awsvpc/iac/pulumi linux_amd64")
        self.assertIsNotNone(self.published("releases/v1.2.3/modules/pulumi/awsvpc/source.zip"))
        self.assertEqual(self.part("awsvpc"), f"catalog/aws/awsvpc/iac/pulumi\tlinux_amd64.gz\t{sha256(binary)}\n")

    def test_windows_carries_the_executable_suffix_and_only_linux_amd64_ships_the_source(self):
        result = self.run_script("catalog/aws/awsvpc/iac/pulumi", "windows_amd64")
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertIsNotNone(self.published("releases/v1.2.3/modules/pulumi/awsvpc/windows_amd64.exe.gz"))
        self.assertIsNone(self.published("releases/v1.2.3/modules/pulumi/awsvpc/source.zip"))
        self.assertIn("\twindows_amd64.exe.gz\t", self.part("awsvpc"))

    def test_a_failed_build_names_its_module_and_records_nothing(self):
        result = self.run_script("catalog/aws/broken/iac/pulumi", "linux_amd64")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("Failed to build broken for linux_amd64", result.stdout)
        self.assertFalse(os.path.exists(os.path.join(self.parts, "broken.tsv")))

    def test_modules_built_side_by_side_keep_to_themselves(self):
        kinds = [f"kind{i}" for i in range(6)]
        runs = [subprocess.Popen([SCRIPT, f"catalog/aws/{k}/iac/pulumi"], env=dict(self.env, PLATFORM="linux_arm64"),
                                 stdout=subprocess.PIPE, stderr=subprocess.PIPE) for k in kinds]
        for run in runs:
            self.assertEqual(run.wait(), 0)
            run.stdout.close()
            run.stderr.close()
        for k in kinds:
            binary = self.published(f"releases/v1.2.3/modules/pulumi/{k}/linux_arm64.gz")
            self.assertEqual(binary, f"pulumi catalog/aws/{k}/iac/pulumi linux_arm64".encode())
            self.assertEqual(self.part(k), f"catalog/aws/{k}/iac/pulumi\tlinux_arm64.gz\t{sha256(binary)}\n")


if __name__ == "__main__":
    unittest.main()
