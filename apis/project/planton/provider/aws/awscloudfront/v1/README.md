# AwsCloudFront

AWS CloudFront is a global content delivery network (CDN). This resource provisions a CloudFront distribution with a minimal 80/20 configuration: origins, default origin selection, optional custom domain aliases with ACM certificate, price class, and default root object.

## Spec fields
- enabled: Whether the distribution is enabled.
- aliases: Optional custom domain names (CNAMEs) like cdn.example.com.
- certificate_arn: ACM certificate ARN in us-east-1 required when using aliases.
- price_class: Edge location price class: PRICE_CLASS_100, PRICE_CLASS_200, or PRICE_CLASS_ALL.
- origins: List of origins with domain_name, optional origin_path, and is_default flag.
- default_root_object: Default object to serve when no object is specified (e.g., index.html).

Validation highlights:
- aliases unique; if aliases set, certificate_arn must be non-empty.
- origins must contain at least one item; exactly one origin must be marked as default.
- Enums are enforced to defined values only.

## Stack outputs
- distribution_id: CloudFront distribution ID.
- domain_name: CloudFront distribution domain (e.g., d123.cloudfront.net).
- hosted_zone_id: Route53 hosted zone ID for aliasing to CloudFront.

## How it works
This module can be provisioned with Pulumi or Terraform via the CLI. Stack inputs wire the chosen IaC backend, target manifest, and provider credentials.

## References
- CloudFront distributions: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-working-with.html
- Price classes: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html
