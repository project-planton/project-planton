# AwsCloudFront

Provision and manage an AWS CloudFront distribution with optional Route53 DNS aliases. This module captures the core distribution settings (origins, cache behavior, TLS) and wires optional DNS records when requested.

## Spec fields (80/20)
- **aliases**: CNAMEs such as cdn.example.com. Requires `certificate_arn` when set.
- **certificate_arn**: ACM certificate ARN for custom domains.
- **price_class**: Edge coverage. One of `PRICE_CLASS_100`, `PRICE_CLASS_200`, `PRICE_CLASS_ALL`.
- **logging**:
  - **enabled**: Enable access logs.
  - **bucket_name**: Target S3 bucket name (no s3://).
  - **prefix**: Optional key prefix for log objects.
- **origins[]**:
  - **id**: Origin identifier used by cache behaviors.
  - **domain_name**: Origin domain (e.g., my-bucket.s3.amazonaws.com).
  - **origin_access_control_id**: Optional OAC id for private S3 origins.
- **default_cache_behavior**:
  - **origin_id**: Origin id to use for default behavior.
  - **viewer_protocol_policy**: One of `ALLOW_ALL`, `HTTPS_ONLY`, `REDIRECT_TO_HTTPS`.
  - **compress**: Enable edge compression.
  - **cache_policy_id**: Optional cache policy (managed or custom).
  - **allowed_methods**: One of `GET_HEAD`, `GET_HEAD_OPTIONS`, `ALL`.
- **web_acl_arn**: Optional AWS WAFv2 web ACL ARN to attach.
- **dns**:
  - **enabled**: If true, create Route53 alias records for each hostname in `aliases`.
  - **route53_zone_id**: Hosted zone id to manage records when enabled.

## Stack outputs
- **distribution_id**: CloudFront distribution id.
- **domain_name**: CloudFront domain name (e.g., d123.cloudfront.net).
- **hosted_zone_id**: Zone id for alias records.

## How it works
- Pulumi: `iac/pulumi/module` orchestrates the CloudFront distribution and optional Route53 alias records.
- Terraform: `iac/tf` provides an equivalent module with the same spec surface and outputs.

## References
- AWS CloudFront Distribution (Terraform): https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution
- AWS CloudFront Developer Guide: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/introduction.html
- AWS Route53 Alias records: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html


