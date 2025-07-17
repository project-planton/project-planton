terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

############################
# Input specification
############################

variable "spec" {
  description = "Serialized AwsDynamodbSpec object decoded into a Terraform map/object."
  type        = any
}

############################
# Local helpers & look-ups
############################

locals {
  spec = var.spec

  # Enum → string look-ups
  billing_mode_map = {
    1 = "PROVISIONED"
    2 = "PAY_PER_REQUEST"
  }

  attribute_type_map = {
    1 = "S" # STRING
    2 = "N" # NUMBER
    3 = "B" # BINARY
  }

  projection_type_map = {
    1 = "ALL"
    2 = "KEYS_ONLY"
    3 = "INCLUDE"
  }

  stream_view_type_map = {
    1 = "NEW_IMAGE"
    2 = "OLD_IMAGE"
    3 = "NEW_AND_OLD_IMAGES"
    4 = "KEYS_ONLY"
  }

  gsi_list = try(local.spec.global_secondary_indexes, [])
  lsi_list = try(local.spec.local_secondary_indexes, [])
}

############################
# DynamoDB table resource
############################

resource "aws_dynamodb_table" "this" {
  name         = local.spec.table_name
  billing_mode = lookup(local.billing_mode_map, local.spec.billing_mode, null)

  # Primary keys -------------------------------------------------
  hash_key  = local.spec.key_schema[0].attribute_name
  range_key = length(local.spec.key_schema) > 1 ? local.spec.key_schema[1].attribute_name : null

  # Attribute definitions ---------------------------------------
  dynamic "attribute" {
    for_each = local.spec.attribute_definitions
    content {
      name = attribute.value.attribute_name
      type = lookup(local.attribute_type_map, attribute.value.attribute_type, "S")
    }
  }

  # Provisioned capacity (table-level) ---------------------------
  read_capacity  = local.spec.billing_mode == 1 ? local.spec.provisioned_throughput.read_capacity_units  : null
  write_capacity = local.spec.billing_mode == 1 ? local.spec.provisioned_throughput.write_capacity_units : null

  # Global secondary indexes ------------------------------------
  dynamic "global_secondary_index" {
    for_each = local.gsi_list
    content {
      name               = global_secondary_index.value.index_name
      hash_key           = global_secondary_index.value.key_schema[0].attribute_name
      range_key          = length(global_secondary_index.value.key_schema) > 1 ? global_secondary_index.value.key_schema[1].attribute_name : null
      projection_type    = lookup(local.projection_type_map, global_secondary_index.value.projection.projection_type, "ALL")
      non_key_attributes = global_secondary_index.value.projection.projection_type == 3 ? global_secondary_index.value.projection.non_key_attributes : null

      # Per-index capacity only when table is PROVISIONED
      read_capacity  = local.spec.billing_mode == 1 ? global_secondary_index.value.provisioned_throughput.read_capacity_units  : null
      write_capacity = local.spec.billing_mode == 1 ? global_secondary_index.value.provisioned_throughput.write_capacity_units : null
    }
  }

  # Local secondary indexes -------------------------------------
  dynamic "local_secondary_index" {
    for_each = local.lsi_list
    content {
      name               = local_secondary_index.value.index_name
      range_key          = local_secondary_index.value.key_schema[1].attribute_name
      projection_type    = lookup(local.projection_type_map, local_secondary_index.value.projection.projection_type, "ALL")
      non_key_attributes = local_secondary_index.value.projection.projection_type == 3 ? local_secondary_index.value.projection.non_key_attributes : null
    }
  }

  # Streams ------------------------------------------------------
  stream_enabled   = try(local.spec.stream_specification.stream_enabled, false)
  stream_view_type = try(lookup(local.stream_view_type_map, local.spec.stream_specification.stream_view_type, null), null)

  # Time-to-live -------------------------------------------------
  dynamic "ttl" {
    for_each = (try(local.spec.ttl_specification.ttl_enabled, false) || try(length(local.spec.ttl_specification.attribute_name), 0) > 0) ? [local.spec.ttl_specification] : []
    content {
      attribute_name = ttl.value.attribute_name
      enabled        = ttl.value.ttl_enabled
    }
  }

  # Point-in-time recovery --------------------------------------
  point_in_time_recovery {
    enabled = local.spec.point_in_time_recovery_enabled
  }

  # Server-side encryption --------------------------------------
  dynamic "server_side_encryption" {
    for_each = local.spec.sse_specification.enabled ? [local.spec.sse_specification] : []
    content {
      enabled     = true
      kms_key_arn = local.spec.sse_specification.sse_type == 2 ? local.spec.sse_specification.kms_master_key_id : null
    }
  }

  # Tags ---------------------------------------------------------
  tags = try(local.spec.tags, {})
}

############################
# Stack outputs
############################

output "table_arn" {
  description = "Fully-qualified Amazon Resource Name of the table."
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

output "stream" {
  description = "Current (latest) stream information, present only when streams are enabled."
  value = aws_dynamodb_table.this.stream_enabled ? {
    stream_arn  = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when SSE uses a CMK."
  value       = (local.spec.sse_specification.enabled && local.spec.sse_specification.sse_type == 2) ? local.spec.sse_specification.kms_master_key_id : null
}

output "global_secondary_index_names" {
  description = "Names of provisioned global secondary indexes (GSIs)."
  value       = try([for g in aws_dynamodb_table.this.global_secondary_index : g.name], [])
}

output "local_secondary_index_names" {
  description = "Names of provisioned local secondary indexes (LSIs)."
  value       = try([for l in aws_dynamodb_table.this.local_secondary_index : l.name], [])
}
