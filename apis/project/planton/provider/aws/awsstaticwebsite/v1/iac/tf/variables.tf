variable "metadata" {
  description = "metadata"
  type = object({

    # name of the resource
    name = string

    # id of the resource
    id = string

    # id of the organization to which the api-resource belongs to
    org = string

    # environment to which the resource belongs to
    env = string

    # labels for the resource
    labels = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # annotations for the resource
    annotations = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # tags for the resource
    tags = list(string)
  })
}

variable "spec" {
  description = "spec"
  type = object({

    # Description for enable_cdn
    enable_cdn = bool

    # Description for route53_zone_id
    route53_zone_id = object({

      # Description for value
      value = string

      # Description for value_from
      value_from = object({

        # Description for kind
        kind = string

        # Description for env
        env = string

        # Description for name
        name = string

        # Description for field_path
        field_path = string
      })
    })

    # Description for domain_aliases
    domain_aliases = list(string)

    # Description for certificate_arn
    certificate_arn = object({

      # Description for value
      value = string

      # Description for value_from
      value_from = object({

        # Description for kind
        kind = string

        # Description for env
        env = string

        # Description for name
        name = string

        # Description for field_path
        field_path = string
      })
    })

    # Description for content_bucket_arn
    content_bucket_arn = object({

      # Description for value
      value = string

      # Description for value_from
      value_from = object({

        # Description for kind
        kind = string

        # Description for env
        env = string

        # Description for name
        name = string

        # Description for field_path
        field_path = string
      })
    })

    # Description for content_prefix
    content_prefix = string

    # Description for is_spa
    is_spa = bool

    # Description for index_document
    index_document = string

    # Description for error_document
    error_document = string

    # Description for default_ttl_seconds
    default_ttl_seconds = number

    # Description for max_ttl_seconds
    max_ttl_seconds = number

    # Description for min_ttl_seconds
    min_ttl_seconds = number

    # Description for compress
    compress = bool

    # Description for ipv6_enabled
    ipv6_enabled = bool

    # Description for price_class
    price_class = string

    # Description for logging
    logging = object({

      # Description for s3_enabled
      s3_enabled = bool

      # Description for s3_target_bucket_arn
      s3_target_bucket_arn = object({

        # Description for value
        value = string

        # Description for value_from
        value_from = object({

          # Description for kind
          kind = string

          # Description for env
          env = string

          # Description for name
          name = string

          # Description for field_path
          field_path = string
        })
      })

      # Description for s3_target_prefix
      s3_target_prefix = string

      # Description for cdn_enabled
      cdn_enabled = bool
    })

    # Description for tags
    tags = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })
  })
}