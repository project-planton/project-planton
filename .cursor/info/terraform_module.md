# Terraform Module Authoring Guide

Purpose: implement the Terraform module under `iac/tf/` for a resource kind using multi-file layout.

## Inputs to read
- `api.proto`, `spec.proto`, `stack_input.proto`, `stack_outputs.proto`

## Target directory
- `apis/project/planton/provider/<provider>/<kindfolder>/v1/iac/tf/`

## Files (typical)
- `variables.tf` — generated via CLI (do not hand-edit)
- `provider.tf` — required_providers and minimal provider config
- `locals.tf` — safe_* locals, computed booleans, derived values
- `outputs.tf` — map to `<Kind>StackOutputs`
- Concern files: `security_group.tf`, `instance.tf`, `dns.tf`, `iam.tf`, `data.tf`, etc.
- Optional `main.tf` (keep minimal) and optional nested `modules/`

## Notes
- Use the CLI generator for variables.tf; derive convenience values in locals.tf.
- Split by concern; Terraform builds the graph from all files automatically.
