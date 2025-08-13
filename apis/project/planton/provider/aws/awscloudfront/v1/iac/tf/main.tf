resource "aws_cloudfront_distribution" "this" {
  enabled             = var.spec.enabled
  aliases             = try(var.spec.aliases, null)
  price_class         = var.spec.price_class == "PRICE_CLASS_100" ? "PriceClass_100" : var.spec.price_class == "PRICE_CLASS_200" ? "PriceClass_200" : "PriceClass_All"
  default_root_object = try(var.spec.default_root_object, null)

  dynamic "origin" {
    for_each = { for o in var.spec.origins : o.id => o }
    content {
      domain_name = origin.value.domain_name
      origin_id   = origin.value.id
      origin_path = try(origin.value.origin_path, null)

      custom_origin_config {
        origin_protocol_policy = "https-only"
        http_port              = 80
        https_port             = 443
        origin_ssl_protocols   = ["TLSv1.2"]
      }
    }
  }

  default_cache_behavior {
    target_origin_id       = var.spec.default_origin_id
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
    min_ttl     = 0
    default_ttl = 3600
    max_ttl     = 86400
  }

  dynamic "viewer_certificate" {
    for_each = var.spec.certificate_arn != null && var.spec.certificate_arn != "" ? [1] : []
    content {
      acm_certificate_arn      = var.spec.certificate_arn
      ssl_support_method       = "sni-only"
      minimum_protocol_version = "TLSv1.2_2021"
    }
  }

  dynamic "viewer_certificate" {
    for_each = var.spec.certificate_arn == null || var.spec.certificate_arn == "" ? [1] : []
    content {
      cloudfront_default_certificate = true
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
}


