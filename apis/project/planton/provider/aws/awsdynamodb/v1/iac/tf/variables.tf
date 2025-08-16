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

    # Table name. If provided, the bucket will be created with this name instead of generating the name from the context
    table_name = string

    # Description for aws_region
    aws_region = string

    # Controls how you are charged for read and write throughput and how you manage
    # capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults
    # to `PROVISIONED`.
    billing_mode = string

    # Description for partition_key_name
    partition_key_name = string

    # Description for partition_key_type
    partition_key_type = string

    # Description for sort_key_name
    sort_key_name = string

    # Description for sort_key_type
    sort_key_type = string

    # Description for read_capacity_units
    read_capacity_units = number

    # Description for write_capacity_units
    write_capacity_units = number

    # Description for point_in_time_recovery_enabled
    point_in_time_recovery_enabled = bool

    # Description for server_side_encryption_enabled
    server_side_encryption_enabled = bool
  })
}