#!/usr/bin/env bash
# =============================================================================
# Publish a release artifact under its tag exactly once.
#
# A tag names exactly one set of bytes for every artifact under it, forever:
# platforms that deploy a module record "this resource ran release vX.Y.Z" and
# trust the tag to mean those exact bytes, so an artifact that changed under
# its tag would let a resource read as unchanged while different code is what
# deploys. The first write under a key wins; a later write to the same key --
# a rerun of a failed lane, or a release re-dispatched by hand for an old tag --
# keeps what is already published and says so, which keeps reruns working
# without ever replacing a published artifact.
#
# Usage:
#   publish_once.sh <local-file> <object-key> [extra `aws s3 cp` arguments...]
#
# Reads R2_BUCKET and AWS_ENDPOINT_URL from the environment, like the
# workflows' own aws calls. Exits non-zero only when an upload fails.
# =============================================================================
set -euo pipefail

file="$1"
key="$2"
shift 2

if aws s3api head-object --bucket "${R2_BUCKET}" --key "${key}" \
    --endpoint-url "${AWS_ENDPOINT_URL}" >/dev/null 2>&1; then
  echo "= Kept: https://downloads.planton.ai/${key} (already published under this tag; a tag's artifacts never change)"
  exit 0
fi

aws s3 cp "${file}" "s3://${R2_BUCKET}/${key}" --endpoint-url "${AWS_ENDPOINT_URL}" "$@"
echo "✓ Uploaded: https://downloads.planton.ai/${key}"
