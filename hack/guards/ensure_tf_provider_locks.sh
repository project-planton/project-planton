#!/usr/bin/env bash
set -euo pipefail

# Guard: every OpenTofu module runs exact, checksummed provider releases.
#
# A module's required_providers block names a RANGE ("~> 6.58"); without a lock
# file, `tofu init` takes the newest release inside it at the moment it runs, so
# an unchanged module could run a different provider from one deploy to the next.
# The catalog therefore keeps ONE canonical lock, catalog/_providers/.terraform.lock.hcl
# (one exact release, with its package checksums, per provider any module requires),
# and every module commits its projection of it as .terraform.lock.hcl: exactly the
# providers that module requires, at the canonical releases, with its own constraints.
#
# WHAT IT CHECKS (offline: no network, no registry, no credentials)
#   - the canonical lock holds every provider a module requires, at a release every
#     module's constraint allows, and nothing no module requires;
#   - every module's committed lock is byte-for-byte its projection (what `tofu init`
#     itself would write), and a module that requires no provider has no lock.
#
# FIX: `make lock-providers` rewrites the projections, and re-records only the
# providers a constraint change has excluded. Moving a provider to a newer release
# is deliberate: `make lock-providers upgrade=aws` (see ensure_tf_provider_pins.sh
# for the full procedure). Logic: pkg/iac/tofu/providerlocks.

repo_root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root_dir"

go run ./pkg/iac/tofu/providerlocks/cli check --repo-root "$repo_root_dir"
