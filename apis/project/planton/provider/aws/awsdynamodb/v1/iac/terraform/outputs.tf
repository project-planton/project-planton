############################################
# DynamoDB table – stack outputs
############################################

# Fully-qualified Amazon Resource Name of the table.
output "table_arn" {
  description = "Fully-qualified Amazon Resource Name (ARN) of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

# Name of the DynamoDB table (may include runtime suffixes).
output "table_name" {
  description = "Name of the DynamoDB table."
  value       = aws_dynamodb_table.this.name
}

# AWS-assigned unique identifier of the table.
output "table_id" {
  description = "Unique identifier of the DynamoDB table."
  value       = aws_dynamodb_table.this.id
}

# Current (latest) stream information, present only when streams are enabled.
output "stream" {
  description = "Object containing stream identifiers when DynamoDB Streams are enabled."
  value = {
    stream_arn   = try(aws_dynamodb_table.this.stream_arn, null)
    stream_label = try(aws_dynamodb_table.this.stream_label, null)
  }
}

# ARN of the customer-managed KMS key when SSE uses a CMK.
output "kms_key_arn" {
  description = "ARN of the KMS key used for server-side encryption, when applicable."
  value       = try(aws_dynamodb_table.this.kms_key_arn, null)
}

# Names of provisioned global secondary indexes (GSIs).
output "global_secondary_index_names" {
  description = "Names of global secondary indexes (GSIs) created for the table."
  value       = try([for gsi in aws_dynamodb_table.this.global_secondary_index : gsi.name], [])
}

# Names of provisioned local secondary indexes (LSIs).
output "local_secondary_index_names" {
  description = "Names of local secondary indexes (LSIs) created for the table."
  value       = try([for lsi in aws_dynamodb_table.this.local_secondary_index : lsi.name], [])
}
