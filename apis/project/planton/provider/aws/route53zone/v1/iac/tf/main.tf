# Create the Route53 hosted zone
resource "aws_route53_zone" "r53_zone" {
  # var.metadata.name is expected to be something like "planton.net"
  name = var.metadata.name

  # Convert list(string) tags to map(string) if needed.
  # Assuming tags are just keys, you might assign a simple value like "true" to each:
  tags = var.metadata.tags
}

resource "aws_route53_record" "records" {
  for_each = local.normalized_records

  zone_id = aws_route53_zone.r53_zone.zone_id
  name    = each.value.name
  type    = each.value.record_type
  ttl     = each.value.ttl_seconds == 0 ? 300 : each.value.ttl_seconds
  records = each.value.values
}
