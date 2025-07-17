############################################
# DynamoDB table – exported stack outputs  #
############################################

# NOTE: the resource is assumed to be defined as
#   resource "aws_dynamodb_table" "this" { ... }
# in the root module.  Update the references below
# if a different address is used.

# -----------------------------------------------------------------------------
# Core identifiers
# -----------------------------------------------------------------------------

output "table_arn" {
  description = "Fully-qualified Amazon Resource Name (ARN) of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

output "table_name" {
  description = "Name of the DynamoDB table (may include runtime-generated suffixes)."
  value       = aws_dynamodb_table.this.name
}

output "table_id" {
  description = "AWS-assigned unique identifier of the DynamoDB table."
  value       = aws_dynamodb_table.this.id
}

# -----------------------------------------------------------------------------
# Stream information (only populated when streams are enabled)
# -----------------------------------------------------------------------------

output "stream" {
  description = "Current (latest) DynamoDB Stream identifiers. Empty when streams are disabled."
  value = {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  }
}

# -----------------------------------------------------------------------------
# Encryption information
# -----------------------------------------------------------------------------

output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when server-side encryption uses a CMK; empty otherwise."
  value       = aws_dynamodb_table.this.kms_key_arn
}

# -----------------------------------------------------------------------------
# Index names
# -----------------------------------------------------------------------------

output "global_secondary_index_names" {
  description = "Names of provisioned global secondary indexes (GSIs)."
  value       = [for gsi in aws_dynamodb_table.this.global_secondary_index : gsi.name]
}

output "local_secondary_index_names" {
  description = "Names of provisioned local secondary indexes (LSIs)."
  value       = [for lsi in aws_dynamodb_table.this.local_secondary_index : lsi.name]
}
