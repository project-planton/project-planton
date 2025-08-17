locals {
  # Naming and tags
  resource_name = coalesce(try(var.metadata.name, null), coalesce(try(var.spec.function.handler, null), "aws-lambda-function"))
  tags          = merge({
    "Name" = local.resource_name
  }, try(var.metadata.labels, {}))

  # VPC networking
  safe_subnet_ids = try(var.spec.function.vpc_config.subnet_ids, [])
  safe_security_group_ids = try(var.spec.function.vpc_config.security_group_ids, [])

  # Package type and runtime/handler handling
  package_type = upper(coalesce(try(var.spec.function.package_type, null), "Zip"))
  is_image     = local.package_type == "IMAGE"
  runtime      = local.is_image ? null : try(var.spec.function.runtime, null)
  handler      = local.is_image ? null : try(var.spec.function.handler, null)

  # Code sources
  image_uri        = try(var.spec.function.image_uri, null)
  s3_bucket        = try(var.spec.function.s3_bucket, null)
  s3_key           = try(var.spec.function.s3_key, null)
  s3_object_version = try(var.spec.function.s3_object_version, null)

  # KMS for env vars
  kms_key_arn = try(var.spec.function.kms_key_arn, null)

  # Environment variables as single key/value → convert to map(string) if present
  env_var_object = try(var.spec.function.variables, null)
  env_var_map    = local.env_var_object != null ? tomap({ (local.env_var_object.key) = local.env_var_object.value }) : null

  # Layers
  layer_arns = try(var.spec.function.layers, [])

  # CloudWatch log group settings
  log_retention_days = try(var.spec.cloudwatch_log_group.retention_in_days, null)
  log_kms_key_arn    = try(var.spec.cloudwatch_log_group.kms_key_arn, null)

  # IAM role toggles
  enable_lambda_at_edge              = coalesce(try(var.spec.iam_role.lambda_at_edge_enabled, null), false)
  enable_cloudwatch_lambda_insights  = coalesce(try(var.spec.iam_role.cloudwatch_lambda_insights_enabled, null), false)
  custom_policy_arns                 = try(var.spec.iam_role.custom_iam_policy_arns, [])
  ssm_parameter_names                = try(var.spec.iam_role.ssm_parameter_names, [])
  inline_policy_json                 = try(var.spec.iam_role.inline_iam_policy, null)
}


