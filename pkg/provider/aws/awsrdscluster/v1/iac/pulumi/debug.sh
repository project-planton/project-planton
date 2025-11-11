#!/usr/bin/env bash
set -euo pipefail

if ! command -v dlv >/dev/null 2>&1; then
  echo "Delve (dlv) not found. Install via: go install github.com/go-delve/delve/cmd/dlv@latest" >&2
  exit 1
fi

dlv debug --headless --listen=:2345 --api-version=2 -- "$@"


