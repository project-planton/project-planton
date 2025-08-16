# Terraform Module to Deploy AwsEksNodeGroup

This module provisions an AWS EKS (Elastic Kubernetes Service) node group with support for auto-scaling, Spot instances, and SSH access.
It includes configurable instance types, disk sizes, and Kubernetes labels for worker node management.

Generated `variables.tf` reflects the proto schema for `AwsEksNodeGroup`.

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

