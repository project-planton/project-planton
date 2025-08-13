output "distribution_id" {
  value = aws_cloudfront_distribution.this.id
}

output "domain_name" {
  value = aws_cloudfront_distribution.this.domain_name
}

output "hosted_zone_id" {
  # CloudFront hosted zone ID is static
  value = "Z2FDTNDATAQYW2"
}


