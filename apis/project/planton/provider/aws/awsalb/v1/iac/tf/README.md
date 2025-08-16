# Terraform Module to Deploy AwsAlb

This module provisions an AWS Application Load Balancer and, if SSL is enabled,
adds an HTTP->HTTPS redirect listener and an HTTPS listener. Optional Route53
records are created when DNS is enabled.

Generated `variables.tf` reflects the proto schema for `AwsAlb`.

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


