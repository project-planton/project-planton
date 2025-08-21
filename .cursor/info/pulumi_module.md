# Pulumi Module Authoring Guide

Purpose: implement the Pulumi module under `iac/pulumi/module` for a resource kind. Do not add BUILD.bazel here; `make build` will handle it.

## Inputs to read
- `api.proto`, `spec.proto`, `stack_input.proto`, `stack_outputs.proto`
- Provider credential proto referenced by `stack_input.proto`

## Target directory
- `apis/project/planton/provider/<provider>/<kindfolder>/v1/iac/pulumi/module/`

## Files (typical)
- `main.go` — controller function `Resources(ctx *pulumi.Context, in *<pkg>.<Kind>StackInput) error`
- `locals.go` — `initializeLocals(ctx, in)` returning a struct with ctx, input, target, spec, derived values
- `outputs.go` — constants for `<Kind>StackOutputs` names and helpers
- `resource_*.go` — one or more resource creators split by concern (e.g., dns.go, security_group.go)

## Controller pattern
- Initialize locals; init provider from credentials; orchestrate resources in order; guard optional flows; export outputs.

## Notes
- Use provider SDK imports matching the provider (e.g., pulumi-aws).
- Reflect outputs aligned to `<Kind>StackOutputs`.
