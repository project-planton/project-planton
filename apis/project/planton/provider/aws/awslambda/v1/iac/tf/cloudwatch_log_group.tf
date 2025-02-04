resource "aws_cloudwatch_log_group" "lambda" {
  count = local.create_cloudwatch_log_group ? 1 : 0

  name              = local.cloudwatch_log_group_name
  retention_in_days = local.cloudwatch_log_group_retention_in_days

  # Ternary in one line or with parentheses
  kms_key_id = local.cloudwatch_log_group_kms_key_arn != null ? local.cloudwatch_log_group_kms_key_arn : null

  tags = local.final_tags
}
