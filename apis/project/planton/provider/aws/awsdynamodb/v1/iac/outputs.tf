############################################
# DynamoDB – public outputs
############################################

# Fully-qualified Amazon Resource Name of the table
output "table_arn" {
  description = "Fully-qualified Amazon Resource Name (ARN) of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

# Name of the table (may include runtime suffixes)
output "table_name" {
  description = "Name of the DynamoDB table (may include runtime suffixes)."
  value       = aws_dynamodb_table.this.name
}

# AWS-assigned unique identifier of the table
output "table_id" {
  description = "AWS-assigned unique identifier of the table."
  value       = aws_dynamodb_table.this.id
}

# Current (latest) stream information – only relevant when streams are enabled
output "stream" {
  description = "Current (latest) stream information, present only when streams are enabled."
  value = aws_dynamodb_table.this.stream_enabled ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

# ARN of the customer-managed KMS key (only when SSE uses a CMK)
output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when SSE uses a CMK."
  value       = aws_dynamodb_table.this.kms_key_arn
}

# Names of the provisioned global secondary indexes (GSIs)
output "global_secondary_index_names" {
  description = "Names of provisioned global secondary indexes (GSIs)."
  value       = [for gsi in aws_dynamodb_table.this.global_secondary_index : gsi.name]
}

# Names of the provisioned local secondary indexes (LSIs)
output "local_secondary_index_names" {
  description = "Names of provisioned local secondary indexes (LSIs)."
  value       = [for lsi in try(aws_dynamodb_table.this.local_secondary_index, []) : lsi.name]
}
