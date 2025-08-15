# AWS Static Website Examples

Below are several examples demonstrating how to define an AWS Static Website component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic Static Website

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-basic-website
spec:
  enableCdn: true
  indexDocument: "index.html"
  errorDocument: "error.html"
  defaultTtlSeconds: 300
  maxTtlSeconds: 86400
  minTtlSeconds: 0
  compress: true
  ipv6Enabled: true
  priceClass: "PriceClass_100"
```

This example creates a basic static website:
• CloudFront CDN enabled for global content delivery.
• Standard index and error documents.
• Default cache TTL settings (5 minutes default, 24 hours max).
• Compression and IPv6 support enabled.
• Price Class 100 for cost optimization.
• Suitable for simple static websites.

---

## Static Website with Custom Domain

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-custom-domain-website
spec:
  enableCdn: true
  route53ZoneId:
    value: "Z1234567890ABC"
  domainAliases:
    - "www.example.com"
    - "example.com"
  certificateArn:
    value: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-ef56-7890-abcd-ef1234567890"
  indexDocument: "index.html"
  errorDocument: "error.html"
  compress: true
  ipv6Enabled: true
```

This example creates a static website with custom domain:
• CloudFront CDN with custom domain support.
• Route53 DNS zone integration.
• Multiple domain aliases (www and apex).
• ACM certificate for HTTPS.
• Compression and IPv6 support.
• Suitable for production websites with custom domains.

---

## Single Page Application (SPA)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-spa-website
spec:
  enableCdn: true
  route53ZoneId:
    value: "Z1234567890ABC"
  domainAliases:
    - "app.example.com"
  certificateArn:
    value: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-ef56-7890-abcd-ef1234567890"
  isSpa: true
  indexDocument: "index.html"
  errorDocument: "error.html"
  defaultTtlSeconds: 3600
  maxTtlSeconds: 86400
  minTtlSeconds: 0
  compress: true
  ipv6Enabled: true
```

This example creates a Single Page Application website:
• SPA mode enabled (404s redirect to index.html).
• Custom domain with HTTPS.
• Longer cache TTL for better performance.
• Compression and IPv6 support.
• Suitable for React, Angular, Vue.js applications.

---

## Static Website with Custom S3 Bucket

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-custom-bucket-website
spec:
  enableCdn: true
  contentBucketArn:
    value: "arn:aws:s3:::my-existing-content-bucket"
  contentPrefix: "website/"
  indexDocument: "index.html"
  errorDocument: "error.html"
  defaultTtlSeconds: 300
  maxTtlSeconds: 86400
  minTtlSeconds: 0
  compress: true
  ipv6Enabled: true
```

This example creates a static website with custom S3 bucket:
• Uses existing S3 bucket for content storage.
• Content prefix for organized file structure.
• CloudFront CDN for global delivery.
• Standard cache settings.
• Suitable for existing S3 content with CDN.

---

## High-Performance Static Website

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-high-performance-website
spec:
  enableCdn: true
  route53ZoneId:
    value: "Z1234567890ABC"
  domainAliases:
    - "cdn.example.com"
  certificateArn:
    value: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-ef56-7890-abcd-ef1234567890"
  indexDocument: "index.html"
  errorDocument: "error.html"
  defaultTtlSeconds: 7200
  maxTtlSeconds: 604800
  minTtlSeconds: 0
  compress: true
  ipv6Enabled: true
  priceClass: "PriceClass_All"
```

This example creates a high-performance static website:
• Extended cache TTL for better performance.
• Price Class All for maximum edge locations.
• Custom domain with HTTPS.
• Compression and IPv6 support.
• Suitable for high-traffic websites.

---

## Development Static Website

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-dev-website
spec:
  enableCdn: false
  indexDocument: "index.html"
  errorDocument: "error.html"
  defaultTtlSeconds: 0
  maxTtlSeconds: 0
  minTtlSeconds: 0
  compress: false
  ipv6Enabled: false
```

This example creates a development static website:
• CloudFront CDN disabled for direct S3 access.
• No caching for immediate updates.
• Compression and IPv6 disabled.
• Suitable for development and testing.
• Direct S3 website endpoint access.

---

## Static Website with Logging

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-logged-website
spec:
  enableCdn: true
  route53ZoneId:
    value: "Z1234567890ABC"
  domainAliases:
    - "logged.example.com"
  certificateArn:
    value: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-ef56-7890-abcd-ef1234567890"
  indexDocument: "index.html"
  errorDocument: "error.html"
  logging:
    s3Logging:
      bucket: "my-logging-bucket"
      prefix: "website-logs/"
    cloudFrontLogging:
      bucket: "my-cf-logs-bucket"
      prefix: "cloudfront-logs/"
  compress: true
  ipv6Enabled: true
```

This example creates a static website with comprehensive logging:
• S3 access logging for content requests.
• CloudFront access logging for CDN requests.
• Custom domain with HTTPS.
• Compression and IPv6 support.
• Suitable for production websites with monitoring.

---

## Static Website with Resource References

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-referenced-website
spec:
  enableCdn: true
  route53ZoneId:
    valueFrom:
      kind: "AwsRoute53Zone"
      name: "example.com"
      fieldPath: "status.outputs.zone_id"
  domainAliases:
    - "referenced.example.com"
  certificateArn:
    valueFrom:
      kind: "AwsCertManagerCert"
      name: "example-cert"
      fieldPath: "status.outputs.cert_arn"
  contentBucketArn:
    valueFrom:
      kind: "AwsS3Bucket"
      name: "content-bucket"
      fieldPath: "status.outputs.bucket_arn"
  indexDocument: "index.html"
  errorDocument: "error.html"
  compress: true
  ipv6Enabled: true
```

This example creates a static website with resource references:
• Route53 zone reference from existing resource.
• ACM certificate reference from existing resource.
• S3 bucket reference from existing resource.
• Custom domain with HTTPS.
• Suitable for integrated deployments.

---

## Static Website for Documentation

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-docs-website
spec:
  enableCdn: true
  route53ZoneId:
    value: "Z1234567890ABC"
  domainAliases:
    - "docs.example.com"
  certificateArn:
    value: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-ef56-7890-abcd-ef1234567890"
  contentPrefix: "docs/"
  indexDocument: "index.html"
  errorDocument: "404.html"
  defaultTtlSeconds: 1800
  maxTtlSeconds: 86400
  minTtlSeconds: 0
  compress: true
  ipv6Enabled: true
```

This example creates a documentation website:
• Content prefix for organized documentation.
• Custom 404 error page.
• Moderate cache TTL for documentation updates.
• Custom domain with HTTPS.
• Suitable for technical documentation sites.

---

## Static Website for Marketing

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsStaticWebsite
metadata:
  name: my-marketing-website
spec:
  enableCdn: true
  route53ZoneId:
    value: "Z1234567890ABC"
  domainAliases:
    - "marketing.example.com"
    - "www.marketing.example.com"
  certificateArn:
    value: "arn:aws:acm:us-east-1:123456789012:certificate/abcd1234-ef56-7890-abcd-ef1234567890"
  indexDocument: "index.html"
  errorDocument: "error.html"
  defaultTtlSeconds: 3600
  maxTtlSeconds: 604800
  minTtlSeconds: 0
  compress: true
  ipv6Enabled: true
  priceClass: "PriceClass_200"
```

This example creates a marketing website:
• Multiple domain aliases for marketing campaigns.
• Extended cache TTL for performance.
• Price Class 200 for broader coverage.
• Custom domain with HTTPS.
• Suitable for marketing and promotional sites.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the static website is active via the AWS console or by
using the AWS CLI:

```shell
aws s3 website s3://<your-bucket-name> --index-document index.html --error-document error.html
```

For CloudFront distribution information:

```shell
aws cloudfront list-distributions --query "DistributionList.Items[?Aliases.Items[?contains(@, '<your-domain>')]]"
```

To check Route53 DNS records:

```shell
aws route53 list-resource-record-sets --hosted-zone-id <your-zone-id>
```

This will show the static website details including S3 bucket, CloudFront distribution, and DNS configuration information.
