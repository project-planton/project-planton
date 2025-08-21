#!/usr/bin/env bash
set -euo pipefail
: ${STACK:=organization/<project>/<stack>}
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack "$STACK" \
  --module-dir . | cat

project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack "$STACK" \
  --module-dir . \
  --yes | cat

project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack "$STACK" \
  --module-dir . | cat

project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack "$STACK" \
  --module-dir . | cat
