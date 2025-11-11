# API Proto Authoring Guide

Purpose: author `api.proto` for a resource kind, wiring Kubernetes-style envelope and linking existing `spec.proto` and `stack_outputs.proto`.

## Location and Package
- Path: `apis/project/planton/provider/<provider>/<kindfolder>/v1/api.proto`
- `syntax = "proto3";`
- `package org.project_planton.provider.<provider>.<kindfolder>.v1;`
- Do NOT include `go_package`.

## Imports
- `buf/validate/validate.proto`
- `project/planton/provider/<provider>/<kindfolder>/v1/spec.proto`
- `project/planton/provider/<provider>/<kindfolder>/v1/stack_outputs.proto`
- `project/planton/shared/status.proto`
- `project/planton/shared/metadata.proto`

## Messages
- `<Kind>`
  - `string api_version = 1` with const `"<provider>.project-planton.org/v1"`
  - `string kind = 2` with const `<Kind>` (PascalCase)
  - `org.project_planton.shared.CloudResourceMetadata metadata = 3` with `(buf.validate.field).required = true`
  - `<Kind>Spec spec = 4` with `(buf.validate.field).required = true`
  - `<Kind>Status status = 5` (optional)
- `<Kind>Status`
  - `org.project_planton.shared.ApiResourceLifecycle lifecycle = 99;`
  - `org.project_planton.shared.ApiResourceAudit audit = 98;`
  - `string stack_job_id = 97;`
  - `<Kind>StackOutputs outputs = 1;`

## Notes
- Keep `api_version` and `kind` constants exact.
- Do not rename/add/remove fields of `spec`/`status` here; only wire them.
