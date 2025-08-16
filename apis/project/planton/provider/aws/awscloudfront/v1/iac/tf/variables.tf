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

    # Description for enabled
    enabled = bool

    # Description for aliases
    aliases = list(string)

    # Description for certificate_arn
    certificate_arn = string

    # Description for price_class
    price_class = string

    # Description for origins
    origins = list(object({

      # Description for domain_name
      domain_name = string

      # Description for origin_path
      origin_path = string

      # Description for is_default
      is_default = bool
    }))

    # Description for default_root_object
    default_root_object = string
  })
}