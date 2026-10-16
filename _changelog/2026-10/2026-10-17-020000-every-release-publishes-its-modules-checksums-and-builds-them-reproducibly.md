# Every Release Publishes Its Modules' Checksums and Builds Them Reproducibly

## Why

A platform that deploys Planton's official modules wants to re-run a resource on a new release only when that resource's module actually changed. The only identity a module had was the release tag, so every release looked like a new version of every module, and every resource re-ran. The bytes were not reproducible either: Go stamped VCS data into every binary, and `zip` and `gzip` recorded the checkout time.

## What a release publishes

- **`releases/<tag>/modules/manifest.json`, the module manifest.** It lists every official module by its catalog directory (`catalog/aws/awsrdsinstance/iac/tf`), each artifact a runner may execute with the sha256 of its published bytes, and the module's **digest**: sha256 over the `"<artifact> <sha256>"` lines in name order. Two releases carry the same digest for a module only when they published byte-identical artifacts for it. `release.yaml`'s `module-manifest` job publishes it once both module lanes have uploaded everything, and refuses an incomplete release. Its grammar and reader are `pkg/downloads/modulemanifest`; its URL is `downloads.BuildModuleManifestURL`.
- **Reproducible artifacts.** `tools/modulebuild` is the only place a release builds or packs a module:
  - Pulumi binaries build with the toolchain `go.mod` names, `GOWORK=off`, cgo off, default microarchitecture levels and `-trimpath -buildvcs=false -ldflags "-s -w -buildid="`, gzipped with no name or time.
  - `module.zip` holds the module's git-tracked files except Markdown, plus `LICENSE` and `NOTICE`, written by the deterministic zip writer `pkg/detzip`. A README edit never changes the bytes a runner executes.
  - `source.zip` holds every tracked file of a Pulumi module, written the same way.
- **Recorded checksums.** `publish_once.sh` records each upload's sha256 as object metadata and prints the sha256 of the bytes the key holds, including for an object a rerun kept, so the manifest always describes what is published.

## Gates

- `lint.reproducible-modules.yaml` builds an AWS Pulumi module, a cross-compiled Kubernetes Pulumi module and an OpenTofu module archive twice, in two checkouts with empty build caches, and fails unless the bytes match.
- `lint.release-packaging.yaml` runs the tooling's tests, including the checks that the release matrices cover every module's provider and every platform the manifest requires.

`defspack` and the catalog bundle builder write through `pkg/detzip`; skill archives are byte-identical to before.
