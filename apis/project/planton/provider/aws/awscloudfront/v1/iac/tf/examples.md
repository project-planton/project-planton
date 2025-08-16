# AWS CloudFront Examples

Below are several examples demonstrating how to define an AWS CloudFront component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic CloudFront Distribution

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  name: basic-cdn
spec:
  enabled: true
  priceClass: PRICE_CLASS_100
  origins:
    - id: origin-1
      domainName: my-bucket.s3.amazonaws.com
  defaultOriginId: origin-1
  defaultRootObject: index.html
```

This example creates a basic CloudFront distribution:
• Uses S3 bucket as origin with custom domain name.
• Sets price class to PRICE_CLASS_100 for cost optimization.
• Configures index.html as the default root object.

---

## CloudFront with Custom Domain

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  name: custom-domain-cdn
spec:
  enabled: true
  aliases:
    - cdn.example.com
    - static.example.com
  certificateArn: arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-5678-efgh-ijkl-123456abcdef
  priceClass: PRICE_CLASS_200
  origins:
    - id: origin-1
      domainName: my-bucket.s3.amazonaws.com
  defaultOriginId: origin-1
  defaultRootObject: index.html
```

This example includes custom domain configuration:
• Uses custom domain aliases (cdn.example.com, static.example.com).
• Requires ACM certificate ARN in us-east-1 region.
• Sets price class to PRICE_CLASS_200 for broader edge location coverage.

---

## CloudFront with Multiple Origins

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  name: multi-origin-cdn
spec:
  enabled: true
  priceClass: PRICE_CLASS_ALL
  origins:
    - id: s3-origin
      domainName: my-bucket.s3.amazonaws.com
    - id: api-origin
      domainName: api.example.com
      originPath: /v1
    - id: cdn-origin
      domainName: cdn.example.com
  defaultOriginId: s3-origin
  defaultRootObject: index.html
```

This example demonstrates multiple origins:
• S3 bucket as primary origin for static content.
• API origin with custom path for backend services.
• CDN origin for additional content sources.
• Uses PRICE_CLASS_ALL for maximum edge location coverage.

---

## CloudFront for Static Website

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  name: static-website-cdn
spec:
  enabled: true
  aliases:
    - www.example.com
    - example.com
  certificateArn: arn:aws:acm:us-east-1:123456789012:certificate/website-cert-1234
  priceClass: PRICE_CLASS_100
  origins:
    - id: website-origin
      domainName: website-bucket.s3-website-us-east-1.amazonaws.com
  defaultOriginId: website-origin
  defaultRootObject: index.html
```

This example is optimized for static websites:
• Uses S3 website endpoint as origin.
• Configures both apex and www domain aliases.
• Includes ACM certificate for HTTPS support.

---

## Minimal CloudFront Distribution

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsCloudFront
metadata:
  name: minimal-cdn
spec:
  enabled: true
  origins:
    - id: origin-1
      domainName: my-bucket.s3.amazonaws.com
  defaultOriginId: origin-1
```

A minimal configuration with:
• Only required fields specified.
• Uses default CloudFront certificate (no custom domain).
• Default price class and cache behavior.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the CloudFront distribution is active via the AWS console or by
using the AWS CLI:

```shell
aws cloudfront list-distributions
```

You should see your new CloudFront distribution with its domain name (e.g., d123.cloudfront.net) and any custom aliases configured.


