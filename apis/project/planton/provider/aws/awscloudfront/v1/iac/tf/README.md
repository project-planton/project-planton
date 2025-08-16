# Terraform Module to Deploy AwsCloudFront

This module provisions an AWS CloudFront distribution with origins, default cache behavior,
optional custom domain aliases with ACM certificate, and price class configuration.

Generated `variables.tf` reflects the proto schema for `AwsCloudFront`.

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


