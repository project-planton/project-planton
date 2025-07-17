############################
# DynamoDB table – outputs #
############################

# Fully-qualified Amazon Resource Name of the table.
output "table_arn" {
  description = "Fully-qualified Amazon Resource Name (ARN) of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

# Name of the DynamoDB table.
output "table_name" {
  description = "Name of the DynamoDB table (may include runtime suffixes)."
  value       = aws_dynamodb_table.this.name
}

# AWS-assigned unique identifier of the table.
output "table_id" {
  description = "AWS-assigned unique identifier of the DynamoDB table."
  value       = aws_dynamodb_table.this.id
}

# Current (latest) stream information.
output "stream" {
  description = "Current stream information, present only when streams are enabled. Returns null when streams are disabled."
  value = aws_dynamodb_table.this.stream_enabled ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

# ARN of the customer-managed KMS key (only populated when SSE uses a CMK).
output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when server-side encryption uses a CMK. Empty when SSE is disabled or uses AWS-owned keys."
  value       = aws_dynamodb_table.this.kms_key_arn
}

# Names of provisioned global secondary indexes.
output "global_secondary_index_names" {
  description = "Names of provisioned global secondary indexes (GSIs). Empty when no GSIs are defined."
  value       = [for g in aws_dynamodb_table.this.global_secondary_index : g.name]
}

# Names of provisioned local secondary indexes.
output "local_secondary_index_names" {
  description = "Names of provisioned local secondary indexes (LSIs). Empty when no LSIs are defined."
  value       = [for l in aws_dynamodb_table.this.local_secondary_index : l.name]
}
