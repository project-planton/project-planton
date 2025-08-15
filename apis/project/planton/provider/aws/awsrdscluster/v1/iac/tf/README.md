# Terraform Module to Deploy AwsRdsCluster

This module provisions an AWS RDS (Relational Database Service) cluster with support for multiple database engines, high availability, and scalability.
It includes configurable engine types, instance configurations, VPC integration, security features, and comprehensive database management capabilities.

Generated `variables.tf` reflects the proto schema for `AwsRdsCluster`.

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
