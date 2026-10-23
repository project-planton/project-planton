# Provider locks

`.terraform.lock.hcl` in this folder is the catalog's one source of truth for OpenTofu provider releases. For every provider any module requires, it names one exact release and the checksums of that release's packages for each platform a module runs on.

A module's `provider.tf` says which releases it can work with (`~> 6.58`). Its committed `iac/tf/.terraform.lock.hcl` says which one it does run: a projection of this file, holding exactly the providers that module requires, at these releases, with the module's own constraints. Every module that uses a provider therefore runs the same release of it, and `tofu init -lockfile=readonly` installs exactly that release or refuses.

- **Keep the locks current:** `make lock-providers`. It keeps every locked release unless a module's constraint now excludes it, records any provider no module required before, drops any provider no module requires any more, and rewrites every module's projection.
- **Move a provider to a newer release:** `make lock-providers upgrade=aws` (or `upgrade=aws,helm`). It takes the newest release every module's constraint allows. Only the modules that require that provider change. In the same change:
  - validate every module of that provider against the new release (`tofu init -backend=false -lockfile=readonly && tofu validate` in each), because a release can rename an argument outright and only validation sees it before an apply does; CI validates every module whose lock moved;
  - regenerate the parity schema with `make generate-provider-schemas`, which reads every version from this file, and work the parity findings it surfaces (a schema whose version differs from this lock fails `TestCommittedArtifactsMatchCanonicalLock`);
  - move the Pulumi SDK to the release that bridges the same provider version, so both engines run one provider.
- **Check without the network:** `bash hack/guards/ensure_tf_provider_locks.sh`, which CI runs on every module change.

Never edit a lock by hand, and never commit one a plain `tofu init` or `terraform init` wrote: such a lock records one platform, and the guard refuses it.
