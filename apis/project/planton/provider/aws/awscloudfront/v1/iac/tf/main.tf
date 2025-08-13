resource "aws_cloudfront_distribution" "this" {
  enabled = true

  dynamic "origin" {
    for_each = var.spec.origins
    content {
      domain_name = origin.value.domain_name
      origin_id   = origin.value.id
    }
  }

  default_cache_behavior {
    target_origin_id       = var.spec.default_cache_behavior.origin_id
    viewer_protocol_policy = local.viewer_protocol_policy
    compress               = var.spec.default_cache_behavior.compress
    cached_methods         = ["GET", "HEAD"]
    allowed_methods        = local.allowed_methods
    cache_policy_id        = try(var.spec.default_cache_behavior.cache_policy_id, null)

    dynamic "forwarded_values" {
      for_each = try(var.spec.default_cache_behavior.cache_policy_id, null) == null ? [1] : []
      content {
        query_string = false
        cookies {
          forward = "none"
        }
      }
    }
  }

  price_class = local.price_class

  viewer_certificate {
    acm_certificate_arn           = length(local.safe_certificate_arn) > 0 ? local.safe_certificate_arn : null
    ssl_support_method             = length(local.safe_certificate_arn) > 0 ? "sni-only" : null
    minimum_protocol_version       = length(local.safe_certificate_arn) > 0 ? "TLSv1.2_2021" : null
    cloudfront_default_certificate = length(local.safe_certificate_arn) == 0 ? true : null
  }

  aliases = local.safe_aliases

  logging_config {
    bucket          = local.safe_logging.enabled && length(local.safe_logging.bucket_name) > 0 ? "${local.safe_logging.bucket_name}.s3.amazonaws.com" : null
    include_cookies = local.safe_logging.enabled
    prefix          = try(local.safe_logging.prefix, null)
  }

  web_acl_id = try(var.spec.web_acl_arn, null)

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
}

resource "aws_route53_record" "alias" {
  count = local.safe_dns.enabled ? length(local.safe_aliases) : 0

  zone_id = local.safe_dns.route53_zone_id
  name    = local.safe_aliases[count.index]
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.this.domain_name
    zone_id                = aws_cloudfront_distribution.this.hosted_zone_id
    evaluate_target_health = true
  }
}

