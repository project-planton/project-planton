#!/usr/bin/env bash
# =============================================================================
# Build one Pulumi module for one platform, publish it under the release tag,
# and record what was published for the Pulumi module manifest.
#
# release.pulumi-modules.yaml runs this once per module, several at a time: a
# leg builds every module of one provider for one platform, and the largest
# providers ship around two hundred modules, so building them one after
# another left the runner's other cores idle and the Pulumi lane far slower
# than it needs to be. Each run therefore keeps everything to itself -- its
# own scratch directory, its own part file -- so any number can run side by
# side, and a failure says which module failed.
#
# What it does, for a module directory (catalog/aws/awsvpc/iac/pulumi):
#   1. modulebuild pulumi builds the binary for PLATFORM;
#   2. publish_once.sh publishes it at
#      releases/<TAG>/modules/pulumi/<kind>/<PLATFORM>.gz (.exe.gz on windows);
#   3. "<module dir>\t<artifact>\t<published sha256>" is written to
#      PARTS_DIR/<kind>.tsv -- the leg joins these into the part the Pulumi
#      manifest job reads;
#   4. on linux_amd64 only, the module's source.zip is packed and published
#      (source is platform-independent, and nothing executes it, so it is no
#      part of the manifest).
#
# Usage:
#   publish_pulumi_module.sh <module dir>
#
# Reads from the environment: TAG, PLATFORM (<goos>_<goarch>), MODULEBUILD
# (the built modulebuild binary), PARTS_DIR, and what publish_once.sh reads
# (R2_BUCKET, AWS_ENDPOINT_URL and the AWS credentials). Runs from the
# repository root, as modulebuild does. Exits non-zero when the module could
# not be built or published.
# =============================================================================
set -euo pipefail

module_dir="$1"
kind="$(echo "${module_dir}" | cut -d/ -f3)"
here="$(cd "$(dirname "$0")" && pwd)"

if [ "${PLATFORM#windows}" != "${PLATFORM}" ]; then
  artifact="${PLATFORM}.exe.gz"
else
  artifact="${PLATFORM}.gz"
fi

scratch="$(mktemp -d)"
trap 'rm -rf "${scratch}"' EXIT

echo "BUILDING: ${kind} (${PLATFORM})"
if ! "${MODULEBUILD}" pulumi -module "${module_dir}" -platform "${PLATFORM}" -out "${scratch}/${artifact}"; then
  echo "✗ ERROR: Failed to build ${kind} for ${PLATFORM}"
  exit 1
fi
# Upload once (a tag's artifacts never change) and record the checksum of
# the bytes the key holds.
if ! sha="$("${here}/publish_once.sh" "${scratch}/${artifact}" "releases/${TAG}/modules/pulumi/${kind}/${artifact}" \
  --cache-control "public, max-age=31536000, immutable")"; then
  echo "✗ ERROR: Failed to upload ${artifact} of ${kind} to R2"
  exit 1
fi
printf '%s\t%s\t%s\n' "${module_dir}" "${artifact}" "${sha}" > "${PARTS_DIR}/${kind}.tsv"

# The module source, once per kind (the linux_amd64 leg writes it). LICENSE
# and NOTICE land at its root (Apache-2.0 section 4(d)).
if [ "${PLATFORM}" = "linux_amd64" ]; then
  if ! "${MODULEBUILD}" source -module "${module_dir}" -out "${scratch}/source.zip"; then
    echo "✗ ERROR: Failed to pack the source of ${kind}"
    exit 1
  fi
  if ! "${here}/publish_once.sh" "${scratch}/source.zip" "releases/${TAG}/modules/pulumi/${kind}/source.zip" \
    --cache-control "public, max-age=31536000, immutable" >/dev/null; then
    echo "✗ ERROR: Failed to upload the source of ${kind} to R2"
    exit 1
  fi
fi
echo "✓ ${kind} (${PLATFORM})"
