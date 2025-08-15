# Terraform Module to Deploy AwsDynamodb

This module provisions an AWS DynamoDB table with comprehensive configuration options including
billing modes, indexes, encryption, autoscaling, streams, and global table replication.

Generated `variables.tf` reflects the proto schema for `AwsDynamodb`.

## Usage

Use the ProjectPlanton CLI (tofu) with the default local backend:

```shell
project-planton tofu init --manifest hack/manifest.yaml
project-planton tofu plan --manifest hack/manifest.yaml
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

**Note**: Credentials are provided via stack input (CLI), not in the manifest `spec`.

For more examples, see [`examples.md`](./examples.md) and [`hack/manifest.yaml`](../hack/manifest.yaml).
