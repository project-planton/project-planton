# Table name
output "table_name" {
  description = "Name of the DynamoDB table."
  value       = aws_dynamodb_table.this.name
}

# Table ARN
output "table_arn" {
  description = "ARN of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

# Table stream ARN
output "table_stream_arn" {
  description = "DynamoDB table stream ARN (if streams are enabled)."
  value       = aws_dynamodb_table.this.stream_arn
}

# Autoscaling read policy ARN for the table (null if autoscaling is disabled)
output "autoscaling_read_policy_arn" {
  description = "ARN of the read autoscaling policy for the table."
  value       = length(aws_appautoscaling_policy.table_read) > 0 ? aws_appautoscaling_policy.table_read[0].arn : null
}

# Autoscaling write policy ARN for the table (null if autoscaling is disabled)
output "autoscaling_write_policy_arn" {
  description = "ARN of the write autoscaling policy for the table."
  value       = length(aws_appautoscaling_policy.table_write) > 0 ? aws_appautoscaling_policy.table_write[0].arn : null
}

# List of autoscaling read policy ARNs for each GSI
output "autoscaling_index_read_policy_arn_list" {
  description = "List of read autoscaling policy ARNs for each global secondary index."
  value       = length(aws_appautoscaling_policy.gsi_read) > 0 ? [for _, r in aws_appautoscaling_policy.gsi_read : r.arn] : []
}

# List of autoscaling write policy ARNs for each GSI
output "autoscaling_index_write_policy_arn_list" {
  description = "List of write autoscaling policy ARNs for each global secondary index."
  value       = length(aws_appautoscaling_policy.gsi_write) > 0 ? [for _, r in aws_appautoscaling_policy.gsi_write : r.arn] : []
}
