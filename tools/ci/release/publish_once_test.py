#!/usr/bin/env python3
"""Tests for publish_once.sh: run with `python3 -m unittest tools/ci/release/publish_once_test.py`.

A fake `aws` on PATH stands in for R2: objects live as files in a temp
directory, with their metadata checksum beside them.
"""
import hashlib
import os
import subprocess
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
SCRIPT = os.path.join(HERE, "publish_once.sh")

FAKE_AWS = r'''#!/usr/bin/env bash
set -euo pipefail
store="${FAKE_R2}"
if [ "$1" = "s3api" ] && [ "$2" = "head-object" ]; then
  key=""; while [ $# -gt 0 ]; do [ "$1" = "--key" ] && key="$2"; shift; done
  obj="${store}/$(echo "${key}" | tr '/' '_')"
  [ -f "${obj}" ] || exit 254
  if [ -f "${obj}.sha256" ]; then cat "${obj}.sha256"; else echo None; fi
  exit 0
fi
if [ "$1" = "s3" ] && [ "$2" = "cp" ]; then
  src="$3"; dst="$4"; meta=""
  shift 4; while [ $# -gt 0 ]; do [ "$1" = "--metadata" ] && meta="${2#sha256=}"; shift; done
  if [[ "${dst}" == s3://* ]]; then
    obj="${store}/$(echo "${dst#s3://*/}" | tr '/' '_')"
    cp "${src}" "${obj}"; [ -n "${meta}" ] && echo "${meta}" > "${obj}.sha256"
  else
    cp "${store}/$(echo "${src#s3://*/}" | tr '/' '_')" "${dst}"
  fi
  exit 0
fi
exit 1
'''


def sha256(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


class PublishOnceTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.store = os.path.join(self.tmp.name, "r2")
        os.mkdir(self.store)
        bindir = os.path.join(self.tmp.name, "bin")
        os.mkdir(bindir)
        fake = os.path.join(bindir, "aws")
        with open(fake, "w") as f:
            f.write(FAKE_AWS)
        os.chmod(fake, 0o755)
        self.env = dict(os.environ, PATH=bindir + os.pathsep + os.environ["PATH"],
                        FAKE_R2=self.store, R2_BUCKET="bucket", AWS_ENDPOINT_URL="https://r2")

    def tearDown(self):
        self.tmp.cleanup()

    def publish(self, content: bytes, key="releases/v1/modules/terraform/x/module.zip"):
        path = os.path.join(self.tmp.name, "artifact")
        with open(path, "wb") as f:
            f.write(content)
        return subprocess.run([SCRIPT, path, key], env=self.env, capture_output=True, text=True, check=True)

    def object_path(self, key="releases/v1/modules/terraform/x/module.zip"):
        return os.path.join(self.store, key.replace("/", "_"))

    def test_first_publish_uploads_records_and_prints_the_checksum(self):
        result = self.publish(b"one")
        self.assertEqual(result.stdout.strip(), sha256(b"one"))
        self.assertIn("Uploaded", result.stderr)
        with open(self.object_path() + ".sha256") as f:
            self.assertEqual(f.read().strip(), sha256(b"one"))

    def test_a_rerun_keeps_the_object_and_prints_the_published_checksum(self):
        self.publish(b"one")
        result = self.publish(b"one")
        self.assertEqual(result.stdout.strip(), sha256(b"one"))
        self.assertIn("Kept", result.stderr)
        self.assertNotIn("::warning::", result.stderr)

    def test_a_rebuild_that_differs_warns_and_the_published_bytes_stand(self):
        self.publish(b"one")
        result = self.publish(b"two")
        self.assertEqual(result.stdout.strip(), sha256(b"one"))
        self.assertIn("::warning::", result.stderr)
        with open(self.object_path(), "rb") as f:
            self.assertEqual(f.read(), b"one")

    def test_a_kept_object_without_a_recorded_checksum_is_hashed(self):
        with open(self.object_path(), "wb") as f:
            f.write(b"published-before-checksums")
        result = self.publish(b"rebuilt")
        self.assertEqual(result.stdout.strip(), sha256(b"published-before-checksums"))


if __name__ == "__main__":
    unittest.main()
