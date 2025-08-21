# Pulumi Entrypoint Authoring Guide

Purpose: scaffold entrypoint and project files under `iac/pulumi/` (outside module/).

## Target directory
- `apis/project/planton/provider/<provider>/<kindfolder>/v1/iac/pulumi/`

## Files
- `main.go` — loads `<Kind>StackInput` via `stackinput.LoadStackInput` and calls `module.Resources`
- `Pulumi.yaml` — Go runtime project skeleton
- `Makefile` — basic build/tidy targets

## Notes
- Do not add BUILD.bazel here in this rule; higher-level builds handle it.
- Program must import `<pkg>/iac/pulumi/module` and `<pkg>/pulumimodule/stackinput`.
