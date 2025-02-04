resource "aws_lambda_function" "this" {
  function_name = local.resource_id
  role          = aws_iam_role.lambda.arn

  architectures                  = local.lambda_architectures
  description                    = local.lambda_description
  kms_key_arn                    = local.lambda_kms_key_arn
  layers                         = local.lambda_layers
  publish                        = local.lambda_publish
  reserved_concurrent_executions = local.lambda_reserved_concurrent_executions
  runtime                        = local.lambda_runtime
  handler                        = local.lambda_handler
  package_type                   = local.lambda_package_type
  memory_size                    = local.lambda_memory_size
  source_code_hash               = local.lambda_source_code_hash
  timeout                        = local.lambda_timeout
  tags                           = local.final_tags

  ephemeral_storage {
    size = local.lambda_ephemeral_storage_size
  }

  environment {
    variables = local.lambda_variables
  }

  # S3-based deployment package
  s3_bucket = local.lambda_s3_bucket
  s3_key    = local.lambda_s3_key
  s3_object_version = local.lambda_s3_object_version

  # Container image
  image_uri = local.lambda_image_uri

  # Dead-letter queue or SNS topic
  dynamic "dead_letter_config" {
    for_each = (local.lambda_dead_letter_config_target_arn != null ? [local.lambda_dead_letter_config_target_arn] : [])

    content {
      target_arn = dead_letter_config.value
    }
  }

  # File system config (EFS)
  dynamic "file_system_config" {
    for_each = local.lambda_file_system_config != null ? [local.lambda_file_system_config] : []

    content {
      arn              = file_system_config.value.arn
      local_mount_path = file_system_config.value.local_mount_path
    }
  }

  # VPC config
  dynamic "vpc_config" {
    for_each = local.lambda_vpc_config != null ? [local.lambda_vpc_config] : []

    content {
      security_group_ids = vpc_config.value.security_group_ids
      subnet_ids         = vpc_config.value.subnet_ids
      # vpc_id is not a field on aws_lambda_function.vpc_config,
      # so if you need to use it, handle that separately.
      # The Pulumi code doesn't actually store vpc_id here; it might be informational only.
    }
  }

  # Image config for container images
  dynamic "image_config" {
    for_each = local.lambda_image_config != null ? [local.lambda_image_config] : []

    content {
      command           = image_config.value.commands
      entry_point       = image_config.value.entry_points
      working_directory = image_config.value.working_directory
    }
  }

  # X-Ray tracing config
  dynamic "tracing_config" {
    for_each = local.lambda_tracing_config_mode != null && local.lambda_tracing_config_mode != "" ? [local.lambda_tracing_config_mode] : []

    content {
      mode = tracing_config.value
    }
  }
}
