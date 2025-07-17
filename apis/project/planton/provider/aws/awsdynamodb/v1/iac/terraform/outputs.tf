###############################
# DynamoDB table – stack outputs
###############################

# Fully-qualified Amazon Resource Name of the table.
output "table_arn" {
  description = "Fully-qualified Amazon Resource Name (ARN) of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

# Name of the table (may include runtime-generated suffixes).
output "table_name" {
  description = "Name of the DynamoDB table as created in AWS."
  value       = aws_dynamodb_table.this.name
}

# AWS-assigned unique identifier of the table.
output "table_id" {
  description = "Unique identifier of the DynamoDB table assigned by AWS."
  value       = aws_dynamodb_table.this.id
}

# Current stream identifiers – only present when streams are enabled on the table.
output "stream" {
  description = "Most-recent DynamoDB Stream information (available when streams are enabled)."
  value = {
    stream_arn   = try(aws_dynamodb_table.this.stream_arn, null)
    stream_label = try(aws_dynamodb_table.this.stream_label, null)
  }
}

# ARN of the customer-managed KMS key used for server-side encryption (when applicable).
output "kms_key_arn" {
  description = "ARN of the KMS key used for DynamoDB server-side encryption, when a customer-managed CMK is configured."
  value       = try(aws_dynamodb_table.this.sse_description[0].kms_master_key_arn, null)
}

# Names of provisioned Global Secondary Indexes (GSIs).
output "global_secondary_index_names" {
  description = "Names of all Global Secondary Indexes (GSIs) created on the table."
  value       = [for gsi in try(aws_dynamodb_table.this.global_secondary_index, []) : gsi.name]
}

# Names of provisioned Local Secondary Indexes (LSIs).
output "local_secondary_index_names" {
  description = "Names of all Local Secondary Indexes (LSIs) created on the table."
  value       = [for lsi in try(aws_dynamodb_table.this.local_secondary_index, []) : lsi.name]
}
