variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name    = string,
    id      = optional(string),
    org     = optional(string),
    env     = optional(string),
    labels  = optional(map(string)),
    tags    = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "AwsCloudFrontSpec configuration"
  type = object({
    aliases         = optional(list(string))
    certificate_arn = optional(string)
    price_class     = optional(string) # PRICE_CLASS_100 | PRICE_CLASS_200 | PRICE_CLASS_ALL

    logging = optional(object({
      enabled     = bool
      bucket_name = string
      prefix      = optional(string)
    }))

    origins = list(object({
      id                       = string
      domain_name              = string
      origin_access_control_id = optional(string)
    }))

    default_cache_behavior = object({
      origin_id               = string
      viewer_protocol_policy  = string  # ALLOW_ALL | HTTPS_ONLY | REDIRECT_TO_HTTPS
      compress                = bool
      cache_policy_id         = optional(string)
      allowed_methods         = string  # GET_HEAD | GET_HEAD_OPTIONS | ALL
    })

    web_acl_arn = optional(string)

    dns = optional(object({
      enabled        = bool
      route53_zone_id = optional(string)
    }))
  })
}


