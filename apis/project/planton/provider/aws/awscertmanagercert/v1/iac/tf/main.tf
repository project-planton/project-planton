resource "aws_acm_certificate" "this" {
  domain_name       = var.spec.primary_domain_name
  validation_method = upper(try(var.spec.validation_method, "DNS"))

  subject_alternative_names = try(var.spec.alternate_domain_names, [])

  lifecycle {
    create_before_destroy = true
  }
}

# DNS validation records (only when DNS validation is selected)
resource "aws_route53_record" "validation" {
  for_each = local.is_dns_validation ? { for dvo in aws_acm_certificate.this.domain_validation_options : dvo.domain_name => {
    name  = dvo.resource_record_name
    type  = dvo.resource_record_type
    value = dvo.resource_record_value
  } } : {}

  zone_id = local.route53_zone_id
  name    = each.value.name
  type    = each.value.type
  ttl     = 60
  records = [each.value.value]
}

resource "aws_acm_certificate_validation" "this" {
  count = local.is_dns_validation ? 1 : 0

  certificate_arn         = aws_acm_certificate.this.arn
  validation_record_fqdns = [for r in aws_route53_record.validation : r.value.fqdn]
}


