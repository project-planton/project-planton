###############################################################################
# locals.tf - parse and store user-supplied inputs for easy reference
###############################################################################

locals {
  # Derive a stable resource ID (same as Pulumi logic)
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "aws_lambda_function"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? { "organization" = var.metadata.org } : {}

  # Environment label only if var.metadata.env is non-empty
  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env, "") != ""
  ) ? { "environment" = var.metadata.env } : {}

  # Merge base, org, and environment labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Merge final_labels with any user-supplied labels in var.metadata.labels
  final_tags = merge(local.final_labels, try(var.metadata.labels, {}))

  # Decide whether to create the CloudWatch Log Group
  create_cloudwatch_log_group = var.spec.cloudwatch_log_group != null
  cloudwatch_log_group_name   = "/aws/lambda/${local.resource_id}"
  cloudwatch_log_group_retention_in_days = try(var.spec.cloudwatch_log_group.retention_in_days, 0)
  cloudwatch_log_group_kms_key_arn       = try(var.spec.cloudwatch_log_group.kms_key_arn, null)

  # Lambda function fields
  lambda_architectures = (
    length(try(var.spec.function.architectures, [])) > 0
    ? var.spec.function.architectures
    : ["x86_64"]
  )

  lambda_description = try(var.spec.function.description, "")
  lambda_handler     = try(var.spec.function.handler, null)
  lambda_image_uri   = try(var.spec.function.image_uri, null)
  lambda_kms_key_arn = try(var.spec.function.kms_key_arn, null)
  lambda_layers      = try(var.spec.function.layers, [])
  lambda_memory_size = (
    try(var.spec.function.memory_size, 0) != 0
    ? var.spec.function.memory_size
    : 128
  )
  lambda_package_type = (
    try(var.spec.function.package_type, "") != ""
    ? var.spec.function.package_type
    : "Zip"
  )
  lambda_publish = try(var.spec.function.publish, false)
  lambda_reserved_concurrent_executions = (
    try(var.spec.function.reserved_concurrent_executions, 0) != 0
    ? var.spec.function.reserved_concurrent_executions
    : -1
  )
  lambda_runtime           = try(var.spec.function.runtime, null)
  lambda_s3_bucket         = try(var.spec.function.s3_bucket, null)
  lambda_s3_key            = try(var.spec.function.s3_key, null)
  lambda_s3_object_version = try(var.spec.function.s3_object_version, null)
  lambda_source_code_hash  = try(var.spec.function.source_code_hash, null)
  lambda_timeout = (
    try(var.spec.function.timeout, 0) != 0
    ? var.spec.function.timeout
    : 3
  )
  lambda_dead_letter_config_target_arn = try(var.spec.function.dead_letter_config_target_arn, null)
  lambda_tracing_config_mode           = try(var.spec.function.tracing_config_mode, null)
  lambda_ephemeral_storage_size = (
    try(var.spec.function.ephemeral_storage_size, 0) != 0
    ? var.spec.function.ephemeral_storage_size
    : 512
  )
  lambda_variables = try(var.spec.function.variables, {})

  lambda_file_system_config = try(var.spec.function.file_system_config, null)
  lambda_vpc_config         = try(var.spec.function.vpc_config, null)
  lambda_image_config       = try(var.spec.function.image_config, null)

  # IAM role is always created, but we store whether the user specified config
  iam_role_provided = var.spec.iam_role != null

  iam_role_permissions_boundary               = try(var.spec.iam_role.permissions_boundary, null)
  iam_role_lambda_at_edge_enabled             = try(var.spec.iam_role.lambda_at_edge_enabled, false)
  iam_role_cloudwatch_lambda_insights_enabled = try(var.spec.iam_role.cloudwatch_lambda_insights_enabled, false)
  iam_role_ssm_parameter_names                = try(var.spec.iam_role.ssm_parameter_names, [])
  iam_role_custom_iam_policy_arns             = try(var.spec.iam_role.custom_iam_policy_arns, [])
  iam_role_inline_iam_policy                  = try(var.spec.iam_role.inline_iam_policy, null)

  # Lambda invoke function permissions
  invoke_function_permissions = try(var.spec.invoke_function_permissions, [])
}
