# Every OpenTofu Module Runs Exact, Checksummed Provider Releases

## Why

A module's `required_providers` block names a range of releases it can work with, such as `~> 6.58`. Without a lock file, `tofu init` took the newest release inside that range at the moment it ran. So an unchanged module could run a different provider from one deploy to the next, and change the infrastructure it manages when nothing about the module had changed.

## What changed

**One canonical lock.** `catalog/_providers/.terraform.lock.hcl` names one exact release of every provider any module requires, with the package checksums for each platform a module runs on: linux and darwin on amd64 and arm64, and windows on amd64. Every module that uses a provider runs the same release of it.

**Each module commits its projection.** Every published module's `iac/tf/.terraform.lock.hcl` holds exactly the providers it requires, at the canonical releases, with its own constraints. It is byte for byte what `tofu init` would write, so an init never rewrites it. It ships in the module's release zip and in `iac-source.zip`.

**Init honours the lock.** `tofumodule.Init`, shared by the CLI and the platform runner, runs OpenTofu with `-lockfile=readonly` whenever the module directory holds a lock. OpenTofu then installs exactly the locked releases, or refuses a package whose checksum doesn't match. A module without a lock (custom, generated or ejected) and the terraform binary, which resolves against its own registry, are unchanged.

**One way to change the locks.**
- `make lock-providers` keeps every locked release unless a module's constraint now excludes it. It records any provider no module required before, drops any provider no module requires any more, and rewrites every projection. It reaches the registry only for a provider it must record.
- `make lock-providers upgrade=aws` moves a provider to the newest release every module's constraint allows. Only the modules that require that provider change.

**An offline guard.** `hack/guards/ensure_tf_provider_locks.sh` runs on every module change in `lint.terraform-modules`. It proves the canonical lock covers every module's constraints and every module's lock is its projection. The validate lanes init read-only against the committed locks. The CI toolchain image pins OpenTofu to the release the platform runner executes modules with, and verifies its download.

## Locked releases

| Provider | Release |
|---|---|
| alekc/kubectl | 2.4.1 |
| auth0/auth0 | 1.59.0 |
| azure/azapi | 2.11.0 |
| cloudflare/cloudflare | 5.28.0 |
| digitalocean/digitalocean | 2.106.0 |
| hashicorp/aws | 6.69.1 |
| hashicorp/azurerm | 5.10.0 |
| hashicorp/google | 8.7.0 |
| hashicorp/google-beta | 8.7.0 |
| hashicorp/helm | 3.3.0 |
| hashicorp/http | 3.6.2 |
| hashicorp/kubernetes | 2.38.0 |
| hashicorp/random | 3.9.1 |
| hashicorp/time | 0.14.2 |
| hashicorp/tls | 4.4.1 |
| openfga/openfga | 0.5.1 |
| stripe/stripe | 0.3.0 |

## Verified

- A module per provider (17 providers) installs and validates under `tofu init -lockfile=readonly`, on OpenTofu 1.11.5 and on the runner's 1.13.1. A plain init leaves every lock unchanged.
- A tampered checksum is refused: "doesn't match any of the checksums".
- The guard, the package's unit tests (constraint normalization against OpenTofu's own output, a lock OpenTofu wrote rendered back byte for byte, projections, every stale-lock case, incremental moves), `tofumodule`, `anatomy` and the e2e runner pass.
