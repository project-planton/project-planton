#!/usr/bin/env bash
set -euo pipefail

go build -o goapp -gcflags "all=-N -l" .
exec dlv --listen=:2345 --headless=true --api-version=2 exec ./goapp


