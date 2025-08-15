locals {
  route53_zone_id = coalesce(try(var.spec.route53_hosted_zone_id.value, null), try(var.spec.route53_hosted_zone_id.value_from.name, null))
  is_dns_validation = upper(try(var.spec.validation_method, "DNS")) == "DNS"
}


