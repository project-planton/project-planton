variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(object({
      name = optional(string),
      id = optional(string),
    })),
    labels = optional(object({
      key = string, value = string
    })),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # Table name. If provided, the bucket will be created with this name instead of generating the name from the context
    table_name = string

    # Controls how you are charged for read and write throughput and how you manage
    # capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults
    # to `PROVISIONED`.
    billing_mode = string

    # Attribute to use as the hash (partition) key. Must also be defined as an `attribute`.
    hash_key = object({

      # Name of the attribute
      name = string

      # Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
      type = string
    })

    # Attribute to use as the range (sort) key. Must also be defined as an `attribute`, see below.
    range_key = object({

      # Name of the attribute
      name = string

      # Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
      type = string
    })

    # Whether Streams are enabled.
    enable_streams = bool

    # When an item in the table is modified, StreamViewType determines what information
    # is written to the table's stream. Valid values are
    # `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
    stream_view_type = string

    # Encryption at rest options. AWS DynamoDB tables are automatically
    # encrypted at rest with an AWS-owned Customer Master Key if this argument
    # isn't specified.
    server_side_encryption = object({

      # Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK).
      # If `enabled` is `false` then server-side encryption is set to
      # AWS-_owned_ key (shown as `DEFAULT` in the AWS console).
      # Potentially confusingly, if `enabled` is `true` and no `kmsKeyArn` is specified then
      # server-side encryption is set to the _default_ KMS-_managed_ key (shown as `KMS` in the AWS console).
      # The [AWS KMS documentation](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html)
      # explains the difference between AWS-_owned_ and KMS-_managed_ keys.
      is_enabled = bool

      # ARN of the CMK that should be used for the AWS KMS encryption.
      # This argument should only be used if the key is different from the default KMS-managed DynamoDB key,
      # `alias/aws/dynamodb`.
      # **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
      kms_key_arn = string
    })

    # Enable point-in-time recovery options.
    point_in_time_recovery = object({

      # Whether to enable point-in-time recovery. It can take 10 minutes to enable for
      # new tables. If the `pointInTimeRecovery` block is not provided,
      # this defaults to `false`.
      is_enabled = bool
    })

    # Configuration block for TTL.
    ttl = object({

      # Whether TTL is enabled. Default value is `false`.
      is_enabled = bool

      # Name of the table attribute to store the TTL timestamp in.
      # Required if `enabled` is `true`, must not be set otherwise.
      attribute_name = string
    })

    # Dynamodb auto scale config
    auto_scale = object({

      # Description for is_enabled
      is_enabled = bool

      # auto scale capacity for read
      read_capacity = object({

        # Min capacity of the scalable target.
        min_capacity = number

        # Max capacity of the scalable target.
        max_capacity = number

        # target capacity utilization percentage
        target_utilization = number
      })

      # auto scale capacity for write
      write_capacity = object({

        # Min capacity of the scalable target.
        min_capacity = number

        # Max capacity of the scalable target.
        max_capacity = number

        # target capacity utilization percentage
        target_utilization = number
      })
    })

    # Set of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes.
    attributes = list(object({

      # Name of the attribute
      name = string

      # Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
      type = string
    }))

    # Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc.
    global_secondary_indexes = list(object({

      # Name of the index.
      name = string

      # One of `ALL`, `INCLUDE` or `KEYS_ONLY` where
      # `ALL` projects every attribute into the index,
      # `KEYS_ONLY` projects  into the index only the table and index hashKey and sortKey attributes ,
      # `INCLUDE` projects into the index all of the attributes that are defined in `nonKeyAttributes`
      # in addition to the attributes that that`KEYS_ONLY` project.
      projection_type = string

      # Only required with `INCLUDE` as a projection type; a list of attributes to project into the index.
      # These do not need to be defined as attributes on the table.
      non_key_attributes = list(string)

      # Name of the hash key in the index; must be defined as an attribute in the resource.
      hash_key = string

      # Name of the range key; must be defined
      range_key = string

      # Number of read units for this index. Must be set if billingMode is set to PROVISIONED.
      read_capacity = number

      # Number of write units for this index. Must be set if billingMode is set to PROVISIONED.
      write_capacity = number
    }))

    # Describe an LSI on the table; these can only be allocated _at creation_
    # so you cannot change this definition after you have created the resource.
    local_secondary_indexes = list(object({

      # Name of the index.
      name = string

      # One of `ALL`, `INCLUDE` or `KEYS_ONLY` where
      # `ALL` projects every attribute into the index,
      # `KEYS_ONLY` projects  into the index only the table and index hashKey and sortKey attributes ,
      # `INCLUDE` projects into the index all of the attributes that are defined in `nonKeyAttributes` in addition to
      # the attributes that that`KEYS_ONLY` project.
      projection_type = string

      # Only required with `INCLUDE` as a projection type; a list of attributes to project into the index.
      # These do not need to be defined as attributes on the table.
      non_key_attributes = list(string)

      # Name of the range key; must be defined
      range_key = string
    }))

    # Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)]
    # (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations.
    replica_region_names = list(string)

    # Import Amazon S3 data into a new table. See below.
    import_table = object({

      # Type of compression to be used on the input coming from the imported table.
      # Valid values are `GZIP`, `ZSTD` and `NONE`.
      input_compression_type = string

      # The format of the source data.
      # Valid values are `CSV`, `DYNAMODB_JSON`, and `ION`.
      input_format = string

      # Describe the format options for the data that was imported into the target table.
      # There is one value, `csv`.
      input_format_options = object({

        # This block contains the processing options for the CSV file being imported:
        csv = object({

          # The delimiter used for separating items in the CSV file being imported.
          delimiter = string

          # List of the headers used to specify a common header for all source CSV files being imported.
          headers = list(string)
        })
      })

      # Values for the S3 bucket the source file is imported from.
      s3_bucket_source = object({

        # The S3 bucket that is being imported from.
        bucket = string

        # The account number of the S3 bucket that is being imported from.
        bucket_owner = string

        # The key prefix shared by all S3 Objects that are being imported.
        key_prefix = string
      })
    })
  })
}
