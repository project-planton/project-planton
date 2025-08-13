# Terraform Module to Deploy AwsCloudFront

This module provisions an AWS CloudFront distribution. Use the ProjectPlanton CLI (tofu) to run it with a manifest, not raw Terraform.

## Quickstart

1. Prepare a minimal manifest at `hack/manifest.yaml` (one is scaffolded).
2. Run the following commands from this directory using the ProjectPlanton CLI:

```bash
project-planton tofu init --manifest hack/manifest.yaml --backend-type s3 \
  --backend-config="bucket=<your-tf-backend-bucket>" \
  --backend-config="dynamodb_table=<your-tf-lock-table>" \
  --backend-config="region=<your-region>" \
  --backend-config="key=project-planton/aws-stacks/<stack-key>.tfstate"

project-planton tofu plan --manifest hack/manifest.yaml

project-planton tofu apply --manifest hack/manifest.yaml --auto-approve

project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

## Notes

- Credentials are provided via stack input configured in the CLI flow; do not put credentials in `hack/manifest.yaml`.
- When using custom domains (`spec.aliases`), ensure `spec.certificate_arn` references a valid ACM certificate in the us-east-1 region for CloudFront.
- To automate Route53 alias records, set `spec.dns.enabled: true` and provide `spec.dns.route53_zone_id`.


