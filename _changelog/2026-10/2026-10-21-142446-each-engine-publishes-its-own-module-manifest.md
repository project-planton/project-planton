# Each Engine Publishes Its Own Module Manifest

## Why

A release published one module manifest for both engines, and only after both module lanes finished. The OpenTofu lane zips sources and finishes about eight minutes after the tag; the Pulumi lane compiles a binary per kind per platform and took about fifty. So a platform that runs OpenTofu modules by default waited fifty minutes to take any release, for a build it does not run by default. The Pulumi lane's longest leg also ran 51 of its 60 allowed minutes, so a few more kinds would have failed every release.

## What a release publishes

- **One manifest per engine, beside that engine's artifacts:** `releases/<tag>/modules/terraform/manifest.json` and `releases/<tag>/modules/pulumi/manifest.json`. `release.yaml`'s `terraform-module-manifest` and `pulumi-module-manifest` jobs each wait only for their own lane and collect only its parts (`module-part-<engine>-*`). Each refuses an incomplete engine, so a manifest's presence means that engine's modules are complete.
- **Format 2.** A manifest names its engine (`"engine": "terraform" | "pulumi"`), and `modulemanifest.Parse` refuses a module directory that isn't that engine's. `modulemanifest.Engine` is the one engine vocabulary: `modulebuild` uses it, and `downloads.BuildModuleManifestURL(version, engine)` composes the URL. The single-manifest URL is gone.
- **The Pulumi lane builds three modules at a time.** `tools/ci/release/publish_pulumi_module.sh` builds, publishes and records one module, in its own scratch directory and part file. Each leg builds its first module alone, which puts the provider's Pulumi SDK (gigabytes of compiler memory for a large provider) in the build cache, then runs the rest through `xargs -P 3`. Three cold compiles of the SDK at once ran the cloudflare legs out of memory. The largest legs drop from about 50 minutes to about 20 (azure on linux_arm64: 185 modules in 11 minutes).
- **Releases v0.6.15 to v0.6.19** had their single manifests split by engine once (`backfill.engine-module-manifests.yaml`, deleted after it ran). The split carries the same digests and checksums, and is byte-identical to what `modulebuild` writes.

## Gates

- `workflows_test.go` holds each manifest job to its own lane and its own parts.
- `publish_pulumi_module_test.py` covers the published bytes, the recorded checksum, the windows name, the source zip on linux_amd64 only, a failed build, and runs side by side.
- `lint.release-packaging.yaml` runs both release-script tests.
