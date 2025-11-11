# Terraform Module to Deploy AwsStaticWebsite

This module provisions AWS Static Websites with support for S3 hosting, CloudFront CDN, Route53 DNS, and ACM TLS certificates.
It includes configurable domain aliases, cache TTL settings, SPA routing, compression, IPv6 support, and comprehensive logging capabilities.

Generated `variables.tf` reflects the proto schema for `AwsStaticWebsite`.

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

