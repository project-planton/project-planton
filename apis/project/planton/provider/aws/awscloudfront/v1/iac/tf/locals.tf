locals {
  safe_aliases         = try(var.spec.aliases, [])
  safe_certificate_arn = try(var.spec.certificate_arn, "")
  safe_logging         = try(var.spec.logging, { enabled = false, bucket_name = "", prefix = null })
  safe_dns             = try(var.spec.dns, { enabled = false, route53_zone_id = null })

  price_class_map = {
    PRICE_CLASS_100 = "PriceClass_100"
    PRICE_CLASS_200 = "PriceClass_200"
    PRICE_CLASS_ALL = "PriceClass_All"
  }
  price_class = lookup(local.price_class_map, coalesce(var.spec.price_class, "PRICE_CLASS_100"), "PriceClass_100")

  viewer_protocol_policy_map = {
    ALLOW_ALL         = "allow-all"
    HTTPS_ONLY        = "https-only"
    REDIRECT_TO_HTTPS = "redirect-to-https"
  }
  viewer_protocol_policy = lookup(local.viewer_protocol_policy_map, var.spec.default_cache_behavior.viewer_protocol_policy, "allow-all")

  allowed_methods_map = {
    GET_HEAD          = ["GET", "HEAD"]
    GET_HEAD_OPTIONS  = ["GET", "HEAD", "OPTIONS"]
    ALL               = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"]
  }
  allowed_methods = lookup(local.allowed_methods_map, var.spec.default_cache_behavior.allowed_methods, ["GET", "HEAD"])
}


