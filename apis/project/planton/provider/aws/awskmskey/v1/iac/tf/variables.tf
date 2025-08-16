variable "metadata" {
  description = "metadata for all resource objects on planton-cloud"
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
  description = "Specification for Deployment Component"
  type = object({

    # Description for key_spec
    key_spec = string

    # Description for description
    description = string

    # Description for disable_key_rotation
    disable_key_rotation = bool

    # Description for deletion_window_days
    deletion_window_days = number

    # Description for alias_name
    alias_name = string
  })
}