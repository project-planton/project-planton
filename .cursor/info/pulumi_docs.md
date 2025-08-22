# Pulumi Docs Authoring Guide

Purpose: author `README.md`, `examples.md`, and `debug.sh` for the Pulumi package under `iac/pulumi/`.

## Target directory
- `apis/project/planton/provider/<provider>/<kindfolder>/v1/iac/pulumi/`

## Files
- `README.md` — overview of the Pulumi program and how it wires the resource
- `examples.md` — runnable examples (manifests and CLI flows)
- `debug.sh` — helper script to run the Pulumi program locally (optional binary mode)

## Notes
- Do not lint here; focus on content.
- Align examples with the resource’s `iac/hack/manifest.yaml` and CLI flows.
