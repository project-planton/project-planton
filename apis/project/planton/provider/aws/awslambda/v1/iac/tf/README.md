# Terraform Module to Deploy AwsLambda

This module provisions an AWS Lambda function with support for multiple deployment types, IAM roles, VPC integration, and monitoring.
It includes configurable runtime environments, environment variables, CloudWatch logging, and comprehensive serverless function management.

Generated `variables.tf` reflects the proto schema for `AwsLambda`.

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
