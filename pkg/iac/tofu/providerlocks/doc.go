// Package providerlocks keeps every OpenTofu module in the catalog on exact,
// checksummed provider versions.
//
// A module's required_providers block says which provider releases it can
// work with ("~> 6.58"); it does not say which one runs. Without a lock file,
// `tofu init` picks the newest release inside that range at the moment it
// runs, so the same module can run a different provider on two consecutive
// deploys -- and a module that "did not change" can change the infrastructure
// it manages. The lock file closes that: it names one version per provider
// and the checksums of its packages, and `tofu init` refuses anything else.
//
// # One canonical lock, one projection per module
//
// The catalog has one source of truth, catalog/_providers/.terraform.lock.hcl
// (CanonicalLockPath): one exact version, and its package checksums, for every
// provider any module requires. Each module's committed .terraform.lock.hcl is
// a projection of it -- exactly the providers that module requires, at the
// canonical versions, with the module's own constraints -- so every module
// that uses a provider runs the same release of it, and a provider moves for
// the whole catalog at once, deliberately.
//
// A projection is byte-for-byte what `tofu init` itself would write for that
// module given the canonical versions (see Render), so an init in a module
// directory never rewrites its lock, and `tofu init -lockfile=readonly` passes.
//
// # Keeping it current
//
// `make lock-providers` (cli/) is the only writer. It keeps every locked
// version unless a module's constraint now excludes it or the caller names the
// provider to upgrade; only then does it ask the registry for the newest release
// every module's constraints allow and record that release's checksums for the
// platforms a module can run on (TargetPlatforms). Everything else is offline:
// `cli check` (hack/guards/ensure_tf_provider_locks.sh) proves the canonical
// lock covers every requirement and every module's lock is its projection.
package providerlocks
