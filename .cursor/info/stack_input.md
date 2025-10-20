# Stack Input Authoring Guide

Purpose: author `stack_input.proto` for a resource kind, defining how the CLI passes inputs to IaC modules.

## Location and Package
- Path: `apis/project/planton/provider/<provider>/<kindfolder>/v1/stack_input.proto`
- `syntax = "proto3";`
- `package project.planton.provider.<provider>.<kindfolder>.v1;`
- Do NOT include `go_package`.

## Imports
- `project/planton/shared/iac.proto`
- `project/planton/shared/iac/pulumi/pulumi.proto`
- `project/planton/shared/iac/terraform/terraform.proto`
- `project/planton/provider/<provider>/<kindfolder>/v1/api.proto`
- Provider credential spec import (choose one based on provider):
  - aws → `cloud/planton/apis/connect/awscredential/v1/spec.proto` → type `cloud.planton.apis.connect.awscredential.v1.AwsProviderConfig`
  - gcp → `cloud/planton/apis/connect/gcpcredential/v1/spec.proto` → type `cloud.planton.apis.connect.gcpcredential.v1.GcpProviderConfig`
  - azure → `cloud/planton/apis/connect/azurecredential/v1/spec.proto` → type `cloud.planton.apis.connect.azurecredential.v1.AzureProviderConfig`
  - kubernetes → `cloud/planton/apis/connect/kubernetesclustercredential/v1/spec.proto` → type `cloud.planton.apis.connect.kubernetesclustercredential.v1.KubernetesProviderConfig`
  - cloudflare → `cloud/planton/apis/connect/cloudflarecredential/v1/spec.proto` → type `cloud.planton.apis.connect.cloudflarecredential.v1.CloudflareProviderConfig`
  - digitalocean → `cloud/planton/apis/connect/digitaloceancredential/v1/spec.proto` → type `cloud.planton.apis.connect.digitaloceancredential.v1.DigitalOceanProviderConfig`
  - confluent → `cloud/planton/apis/connect/confluentcredential/v1/spec.proto` → type `cloud.planton.apis.connect.confluentcredential.v1.ConfluentProviderConfig`
  - snowflake → `cloud/planton/apis/connect/snowflakecredential/v1/spec.proto` → type `cloud.planton.apis.connect.snowflakecredential.v1.SnowflakeProviderConfig`
  - civo → `cloud/planton/apis/connect/civocredential/v1/spec.proto` → type `cloud.planton.apis.connect.civocredential.v1.CivoProviderConfig`
  - mongodbatlas → `cloud/planton/apis/connect/mongodbatlascredential/v1/spec.proto` → type `cloud.planton.apis.connect.mongodbatlascredential.v1.MongoDbAtlasCredentialSpec`

## Message
- Define `<Kind>StackInput` with fields (in this exact order/numbering):
  1. `project.planton.shared.IacProvisioner provisioner = 1;`
  2. `project.planton.shared.iac.pulumi.PulumiStackInfo pulumi = 2;`
  3. `project.planton.shared.iac.terraform.TerraformStackInfo terraform = 3;`
  4. `<Kind> target = 1;` (from `api.proto`)
  5. `provider_credential = 2;` (provider-specific credential spec type per import above)

## Notes
- Keep field ordering and numbers stable.
- No validation options/imports here; validations belong in `api.proto`/`spec.proto` rules.
