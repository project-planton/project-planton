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
  - aws → `project/planton/credential/awscredential/v1/spec.proto` → type `project.planton.credential.awscredential.v1.AwsCredentialSpec`
  - gcp → `project/planton/credential/gcpcredential/v1/spec.proto` → type `project.planton.credential.gcpcredential.v1.GcpCredentialSpec`
  - azure → `project/planton/credential/azurecredential/v1/spec.proto` → type `project.planton.credential.azurecredential.v1.AzureCredentialSpec`
  - kubernetes → `project/planton/credential/kubernetesclustercredential/v1/spec.proto` → type `project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpec`
  - cloudflare → `project/planton/credential/cloudflarecredential/v1/spec.proto` → type `project.planton.credential.cloudflarecredential.v1.CloudflareCredentialSpec`
  - digitalocean → `project/planton/credential/digitaloceancredential/v1/spec.proto` → type `project.planton.credential.digitaloceancredential.v1.DigitalOceanCredentialSpec`
  - confluent → `project/planton/credential/confluentcredential/v1/spec.proto` → type `project.planton.credential.confluentcredential.v1.ConfluentCredentialSpec`
  - snowflake → `project/planton/credential/snowflakecredential/v1/spec.proto` → type `project.planton.credential.snowflakecredential.v1.SnowflakeCredentialSpec`
  - civo → `project/planton/credential/civocredential/v1/spec.proto` → type `project.planton.credential.civocredential.v1.CivoCredentialSpec`
  - mongodbatlas → `project/planton/credential/mongodbatlascredential/v1/spec.proto` → type `project.planton.credential.mongodbatlascredential.v1.MongoDbAtlasCredentialSpec`

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
