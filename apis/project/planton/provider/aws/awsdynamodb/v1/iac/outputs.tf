############################################
# DynamoDB table – exported stack outputs
############################################

/*
   NOTE:  The surrounding module is expected to provision a single
   "aws_dynamodb_table" resource named "this".  The outputs defined
   here surface the identifiers declared in the AwsDynamodbStackOutputs
   protobuf message so that they can be consumed by upstream modules or
   external systems.
*/

###################################################
# Core table identifiers
###################################################

output "table_arn" {
  description = "Fully-qualified Amazon Resource Name (ARN) of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

output "table_name" {
  description = "Name of the DynamoDB table (may include runtime suffixes)."
  value       = aws_dynamodb_table.this.name
}

output "table_id" {
  description = "AWS-assigned unique identifier of the table."
  value       = aws_dynamodb_table.this.id
}

###################################################
# Stream information (only present when enabled)
###################################################

output "stream" {
  description = "Current (latest) DynamoDB Stream identifiers – only set when streams are enabled."
  value = aws_dynamodb_table.this.stream_enabled ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

###################################################
# Server-side encryption (KMS)
###################################################

output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when server-side encryption uses a CMK; null otherwise."
  value       = try(aws_dynamodb_table.this.server_side_encryption[0].kms_key_arn, null)
}

###################################################
# Secondary index names
###################################################

output "global_secondary_index_names" {
  description = "Names of provisioned Global Secondary Indexes (GSIs)."
  value       = [for g in try(aws_dynamodb_table.this.global_secondary_index, []) : g.name]
}

output "local_secondary_index_names" {
  description = "Names of provisioned Local Secondary Indexes (LSIs)."
  value       = [for l in try(aws_dynamodb_table.this.local_secondary_index, []) : l.name]
}
