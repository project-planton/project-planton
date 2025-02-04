variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id   = optional(string),
    org  = optional(string),
    env  = optional(object({
      name = optional(string),
      id   = optional(string),
    })),

    labels = optional(object({
      key   = string,
      value = string
    })),

    tags    = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "spec"
  type = object({

    ###########################################################################
    # AWS Lambda Function Spec
    ###########################################################################
    function = object({
      # Instruction set architectures: valid values "x86_64" or "arm64".
      architectures = list(string)

      description = string

      file_system_config = object({
        arn             = string
        local_mount_path = string
      })

      handler = string

      # ECR image URI, optional
      image_uri = optional(string, "")

      # KMS key for environment encryption, optional
      kms_key_arn = optional(string, "")

      # List of Lambda Layer ARNs, defaults to an empty list
      layers = list(string)

      memory_size = number
      package_type = string
      publish      = bool

      reserved_concurrent_executions = number
      runtime                        = string

      s3_bucket = optional(string, "")
      s3_key    = string

      s3_object_version = string
      source_code_hash  = string

      timeout   = number
      variables = optional(map(string), {})

      dead_letter_config_target_arn = string

      image_config = object({
        commands          = list(string)
        entry_points      = list(string)
        working_directory = string
      })

      tracing_config_mode = string

      vpc_config = object({
        security_group_ids = list(string)
        subnet_ids         = list(string)
        vpc_id             = string
      })

      ephemeral_storage_size = number
    })

    ###########################################################################
    # IAM Role Spec
    ###########################################################################
    iam_role = object({
      permissions_boundary = string

      # Optional flags
      lambda_at_edge_enabled             = optional(bool, false)
      cloudwatch_lambda_insights_enabled = optional(bool, false)

      # Optional lists
      ssm_parameter_names    = optional(list(string), [])
      custom_iam_policy_arns = optional(list(string), [])

      # Inline policy JSON
      inline_iam_policy = optional(string, "")
    })

    ###########################################################################
    # CloudWatch Log Group
    ###########################################################################
    cloudwatch_log_group = object({
      # Optional KMS Key ARN for log encryption
      kms_key_arn = optional(string, "")

      # Log retention in days
      retention_in_days = number
    })

    ###########################################################################
    # Invoke Function Permissions
    ###########################################################################
    invoke_function_permissions = list(object({
      principal  = string
      source_arn = string
    }))
  })
}
