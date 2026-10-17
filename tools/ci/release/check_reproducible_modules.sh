#!/usr/bin/env bash
# =============================================================================
# Prove that a module build is reproducible: build representative modules
# twice -- in this checkout and in a second worktree of the same commit, each
# with its own empty Go build cache -- and fail unless every artifact is
# byte-identical.
#
# Why it matters: a platform recognises a module by the checksums its release
# published (the module manifest's digest). Correctness never depends on
# reproducibility -- a module that rebuilt differently would simply read as
# changed and re-run -- but every release would then re-run every resource
# built from it. This gate keeps "unchanged source, unchanged bytes" true, so
# a release re-runs only the modules it really changed.
#
# The modules cover each way a module is built: an AWS Pulumi module
# (natively), a Kubernetes Pulumi module (cross-compiled for windows) and an
# OpenTofu module archive.
#
# Usage: tools/ci/release/check_reproducible_modules.sh   (from anywhere in the repo)
# =============================================================================
set -euo pipefail

ROOT="$(git rev-parse --show-toplevel)"
cd "${ROOT}"
WORK="$(mktemp -d)"
cleanup() {
  git worktree remove --force "${WORK}/second" >/dev/null 2>&1 || true
  rm -rf "${WORK}"
}
trap cleanup EXIT

# modulebuild must run on the toolchain go.mod names; GOTOOLCHAIN makes
# `go build` fetch it if this machine's go is another release.
TOOLCHAIN="go$(awk '$1 == "go" {print $2}' go.mod)"
GOTOOLCHAIN="${TOOLCHAIN}" go build -o "${WORK}/modulebuild" ./tools/modulebuild
git worktree add --quiet --detach "${WORK}/second" HEAD

build_all() { # <checkout> <out dir>
  local checkout="$1" out="$2"
  mkdir -p "${out}"
  (
    cd "${checkout}"
    export GOCACHE="${out}.gocache"
    "${WORK}/modulebuild" pulumi -module catalog/aws/awselasticip/iac/pulumi -platform linux_amd64 -out "${out}/aws-pulumi-linux_amd64.gz"
    "${WORK}/modulebuild" pulumi -module catalog/kubernetes/kubernetesnamespace/iac/pulumi -platform windows_amd64 -out "${out}/kubernetes-pulumi-windows_amd64.exe.gz"
    "${WORK}/modulebuild" terraform -module catalog/aws/awselasticip/iac/tf -out "${out}/aws-tf-module.zip"
  )
}

build_all "${ROOT}" "${WORK}/first"
sleep 2 # a clock that leaked into an artifact would now differ
build_all "${WORK}/second" "${WORK}/second-out"

failed=0
for artifact in "${WORK}/first"/*; do
  name="$(basename "${artifact}")"
  a="$(shasum -a 256 "${artifact}" | cut -d' ' -f1)"
  b="$(shasum -a 256 "${WORK}/second-out/${name}" | cut -d' ' -f1)"
  if [ "${a}" = "${b}" ]; then
    echo "✓ ${name} reproduces: ${a}"
  else
    echo "✗ ${name} differs between two builds of one commit: ${a} vs ${b}"
    failed=1
  fi
done
exit "${failed}"
