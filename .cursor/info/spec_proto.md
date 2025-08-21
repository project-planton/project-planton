# Spec Proto Authoring Guide

This document describes how to author or update `spec.proto` for a Planton resource kind.

## Folder and Naming
- Kind (PascalCase): e.g., AwsCloudFront, GcpPubsubTopic, AzureKeyVault
- Kind keyword (snake_case): aws_cloudfront, gcp_pubsub_topic, azure_key_vault
- Folder name (lowercase, no underscores): awscloudfront, gcppubsubtopic, azurekeyvault
- Path: `apis/project/planton/provider/<provider>/<kindfolder>/v1/spec.proto`

## Syntax and Package
- `syntax = "proto3";`
- `package project.planton.provider.<provider>.<kindfolder>.v1;`
- Do NOT include `go_package` in new proto files.

## Imports
- No validations in this step (do not import `buf/validate/validate.proto`).
- Optional when needed for value-or-reference fields:
  - `import "project/planton/shared/foreignkey/v1/foreign_key.proto";`

## Message Structure
- Define a single top-level message named `<Kind>Spec`.
- Keep messages and enums minimal; prefer clarity over completeness.

## Field Guidelines (80/20)
- Include only essential fields most users need.
- For cross-resource identifiers (IDs/ARNs such as IAM role ARN, KMS key ARN, security group IDs, subnet IDs, Route53 zone IDs, etc.), prefer the shared foreign key wrappers:
  - `project.planton.shared.foreignkey.v1.StringValueOrRef`
  - `project.planton.shared.foreignkey.v1.Int32ValueOrRef`
- When using `StringValueOrRef` for well-known kinds, you may set default hints using field options to improve Canvas wiring later:
  - `(project.planton.shared.foreignkey.v1.default_kind) = <CloudResourceKind>`
  - `(project.planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.<field>"`
- Common examples for defaults:
  - Subnets: default_kind = AwsVpc
  - Security groups: default_kind = AwsSecurityGroup, default_kind_field_path = "status.outputs.security_group_id"
  - IAM role: default_kind = AwsIamRole, default_kind_field_path = "status.outputs.role_arn"
  - KMS key: default_kind = AwsKmsKey, default_kind_field_path = "status.outputs.key_arn"
- If a referenced kind does not yet exist (e.g., a future Lambda Layer), use plain `StringValueOrRef` without defaults.

## What to Avoid
- Do not add provider credentials here (those belong in stack input later).
- Avoid deep nesting unless essential.
- Keep comments brief and helpful.

## Example Skeleton (adapt kind/package)
```proto
syntax = "proto3";
package project.planton.provider.aws.awscloudfront.v1;

// Optional if you need value-or-ref wrappers
// import "project/planton/shared/foreignkey/v1/foreign_key.proto";

message AwsCloudFrontSpec {
  // Add essential 80/20 fields here
}
```

## Notes
- Use official provider docs as reference while keeping the draft minimal.
- Validations (buf/validate + CEL) are added by a later rule.
