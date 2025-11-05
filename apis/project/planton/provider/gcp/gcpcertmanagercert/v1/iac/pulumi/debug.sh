#!/bin/bash

# This script helps debug the Pulumi module locally
# Usage: ./debug.sh

set -e

echo "Building Pulumi program..."
go build -o /tmp/gcp-cert-manager-cert-pulumi .

echo "Running Pulumi preview..."
pulumi preview

echo "To run Pulumi up, execute: pulumi up"

