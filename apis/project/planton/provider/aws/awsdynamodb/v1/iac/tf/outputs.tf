# Table ARN
output "table_arn" {
  description = "The ARN of the DynamoDB table"
  value       = aws_dynamodb_table.table.arn
}

# Table ID
output "table_id" {
  description = "The ID of the DynamoDB table"
  value       = aws_dynamodb_table.table.id
}

# Table Name
output "table_name" {
  description = "The name of the DynamoDB table"
  value       = aws_dynamodb_table.table.name
}

# Stream ARN (only if point-in-time recovery is enabled)
output "stream_arn" {
  description = "The stream ARN if point-in-time recovery is enabled"
  value       = local.safe_point_in_time_recovery_enabled ? aws_dynamodb_table.table.stream_arn : null
}

# AWS Region
output "aws_region" {
  description = "The AWS region where the table is located"
  value       = local.safe_spec.aws_region
}
