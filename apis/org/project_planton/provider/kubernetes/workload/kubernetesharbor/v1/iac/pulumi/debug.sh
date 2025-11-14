#!/bin/bash

# This script is used by Pulumi to run the binary in debug mode
# Uncomment the binary option in Pulumi.yaml to enable debugging

# Build the binary with debug symbols
go build -gcflags="all=-N -l" -o pulumi-harborkubernetes .

# Run with dlv (requires delve to be installed: go install github.com/go-delve/delve/cmd/dlv@latest)
dlv exec ./pulumi-harborkubernetes --headless --listen=:2345 --api-version=2 --accept-multiclient

