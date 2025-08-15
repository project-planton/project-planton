output "cert_arn" {
  description = "The ARN of the ACM certificate"
  value       = aws_acm_certificate.this.arn
}

output "certificate_domain_name" {
  description = "The primary domain name for the certificate"
  value       = aws_acm_certificate.this.domain_name
}


