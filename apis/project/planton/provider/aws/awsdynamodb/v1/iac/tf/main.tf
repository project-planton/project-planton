###############################################################################
#  DynamoDB table – primary resource, indexes, capacity, encryption & tagging  #
###############################################################################

terraform {
  required_version = ">= 1.3"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

########################
# Provider configuration
########################
provider "aws" {
  region = var.aws_region
}

########################
# Input variables
########################

variable "aws_region" {
  description = "AWS region where the DynamoDB table will be provisioned"
  type        = string
}

# A single object that mirrors the AwsDynamodbSpec protobuf message.  Optional
# attributes use Terraform's `optional()` type so callers can omit them safely.
variable "dynamodb" {
  description = "Amazon DynamoDB table specification (loosely based on AwsDynamodbSpec)."

  type = object({
    table_name            = string

    attribute_definitions = list(object({
      name = string               # Attribute name
      type = string               # One of: STRING | NUMBER | BINARY
    }))

    key_schema = list(object({
      attribute_name = string     # Must reference one of the attribute_definitions
      key_type       = string     # HASH | RANGE
    }))

    billing_mode           = string                       # PROVISIONED | PAY_PER_REQUEST
    provisioned_throughput = optional(object({            # Required when PROVISIONED
      read  = number
      write = number
    }))

    global_secondary_indexes = optional(list(object({
      name                    = string
      key_schema              = list(object({
        attribute_name = string
        key_type       = string
      }))
      projection_type     = string                        # ALL | KEYS_ONLY | INCLUDE
      non_key_attributes  = optional(list(string))        # Required when INCLUDE
      provisioned_throughput = optional(object({          # Required when billing_mode == PROVISIONED
        read  = number
        write = number
      }))
    })), [])

    local_secondary_indexes = optional(list(object({
      name               = string
      key_schema         = list(object({
        attribute_name = string
        key_type       = string
      })) # Must contain exactly 2 elements (HASH + RANGE)
      projection_type    = string                       # ALL | KEYS_ONLY | INCLUDE
      non_key_attributes = optional(list(string))       # Required when INCLUDE
    })), [])

    stream_specification = optional(object({
      enabled   = bool
      view_type = optional(string)                      # NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | STREAM_KEYS_ONLY
    }))

    ttl_specification = optional(object({
      enabled        = bool
      attribute_name = optional(string)
    }))

    sse_specification = optional(object({
      enabled            = bool
      type               = optional(string)              # AES256 | KMS
      kms_master_key_id  = optional(string)
    }))

    point_in_time_recovery_enabled = optional(bool, false)
    tags                           = optional(map(string), {})
  })

  ############################################################
  # Semantic validations reflecting the CEL rules in the proto
  ############################################################
  validation {
    condition = (
      # Billing-mode / throughput consistency
      (
        var.dynamodb.billing_mode == "PROVISIONED" ? var.dynamodb.provisioned_throughput != null : var.dynamodb.provisioned_throughput == null
      ) &&
      alltrue([
        for g in var.dynamodb.global_secondary_indexes : (
          var.dynamodb.billing_mode == "PROVISIONED" ? try(g.provisioned_throughput != null, false) : try(g.provisioned_throughput == null, true)
        )
      ])
    )

    error_message = "When billing_mode is PROVISIONED, provisioned_throughput must be configured for the table and every GSI; when PAY_PER_REQUEST it must be unset everywhere."
  }

  # Projection rule (INCLUDE requires non_key_attributes)
  validation {
    condition = alltrue([
      for idx in concat(var.dynamodb.global_secondary_indexes, var.dynamodb.local_secondary_indexes) : (
        idx.projection_type == "INCLUDE" ? length(try(idx.non_key_attributes, [])) > 0 : length(try(idx.non_key_attributes, [])) == 0
      )
    ])

    error_message = "non_key_attributes must be set when projection_type is INCLUDE and must be empty otherwise."
  }

  # Stream specification rule
  validation {
    condition     = var.dynamodb.stream_specification == null || (var.dynamodb.stream_specification.enabled == false || var.dynamodb.stream_specification.view_type != null)
    error_message = "stream_view_type must be specified when streams are enabled."
  }

  # TTL specification rule
  validation {
    condition     = var.dynamodb.ttl_specification == null || (var.dynamodb.ttl_specification.enabled == false || (var.dynamodb.ttl_specification.attribute_name != null && var.dynamodb.ttl_specification.attribute_name != ""))
    error_message = "attribute_name must be provided when TTL is enabled."
  }

  # SSE rules
  validation {
    condition = (
      var.dynamodb.sse_specification == null || (
        (!var.dynamodb.sse_specification.enabled && var.dynamodb.sse_specification.type == null && var.dynamodb.sse_specification.kms_master_key_id == null) ||
        (var.dynamodb.sse_specification.enabled && var.dynamodb.sse_specification.type != null && (
          var.dynamodb.sse_specification.type == "KMS" ? var.dynamodb.sse_specification.kms_master_key_id != "" : var.dynamodb.sse_specification.kms_master_key_id == null
        ))
      )
    )
    error_message = "When SSE is enabled, sse_type must be set and kms_master_key_id is required only for KMS; when disabled they must be unset."
  }
}

########################
# Local helpers
########################

locals {
  # Convert proto enum strings to AWS shorthand (S / N / B)
  attribute_definitions = [
    for a in var.dynamodb.attribute_definitions : {
      name = a.name
      type = a.type == "STRING" ? "S" : a.type == "NUMBER" ? "N" : "B"
    }
  ]

  # Extract table HASH and (optional) RANGE keys from key_schema
  hash_key  = [for k in var.dynamodb.key_schema : k.key_type == "HASH" ? k.attribute_name : null][0]
  range_key = try([for k in var.dynamodb.key_schema : k.key_type == "RANGE" ? k.attribute_name : null][0], null)

  ########################
  # Secondary index helpers
  ########################
  gsi_configs = [
    for g in var.dynamodb.global_secondary_indexes : {
      name               = g.name
      hash_key           = [for k in g.key_schema : k.key_type == "HASH" ? k.attribute_name : null][0]
      range_key          = try([for k in g.key_schema : k.key_type == "RANGE" ? k.attribute_name : null][0], null)
      projection_type    = g.projection_type
      non_key_attributes = g.projection_type == "INCLUDE" ? g.non_key_attributes : null
      read_capacity      = var.dynamodb.billing_mode == "PROVISIONED" ? g.provisioned_throughput.read  : null
      write_capacity     = var.dynamodb.billing_mode == "PROVISIONED" ? g.provisioned_throughput.write : null
    }
  ]

  lsi_configs = [
    for l in var.dynamodb.local_secondary_indexes : {
      name               = l.name
      range_key          = [for k in l.key_schema : k.key_type == "RANGE" ? k.attribute_name : null][0]
      projection_type    = l.projection_type
      non_key_attributes = l.projection_type == "INCLUDE" ? l.non_key_attributes : null
    }
  ]

  ######################
  # SSE / KMS processing
  ######################
  use_kms = var.dynamodb.sse_specification != null && var.dynamodb.sse_specification.enabled && var.dynamodb.sse_specification.type == "KMS"

  kms_key_arn = local.use_kms ? (
    var.dynamodb.sse_specification.kms_master_key_id != null && var.dynamodb.sse_specification.kms_master_key_id != "" ?
    var.dynamodb.sse_specification.kms_master_key_id :
    try(aws_kms_key.this[0].arn, null)
  ) : null

  stream_enabled = var.dynamodb.stream_specification != null && var.dynamodb.stream_specification.enabled
}

########################
# Conditional KMS key
########################
resource "aws_kms_key" "this" {
  count = local.use_kms && (var.dynamodb.sse_specification.kms_master_key_id == null || var.dynamodb.sse_specification.kms_master_key_id == "") ? 1 : 0

  description             = "Customer-managed CMK for DynamoDB table ${var.dynamodb.table_name}"
  deletion_window_in_days = 7
  enable_key_rotation     = true
  policy                  = null # Rely on default key policy
  tags                    = var.dynamodb.tags
}

########################################
# DynamoDB table and all nested settings
########################################
resource "aws_dynamodb_table" "this" {
  name           = var.dynamodb.table_name
  billing_mode   = var.dynamodb.billing_mode
  hash_key       = local.hash_key
  range_key      = local.range_key
  attribute      = local.attribute_definitions
  tags           = var.dynamodb.tags

  # Provisioned capacity (only when PROVISIONED)
  dynamic "provisioned_throughput" {
    for_each = var.dynamodb.billing_mode == "PROVISIONED" ? [var.dynamodb.provisioned_throughput] : []
    content {
      read_capacity  = provisioned_throughput.value.read
      write_capacity = provisioned_throughput.value.write
    }
  }

  #############################
  # Global secondary indexes
  #############################
  dynamic "global_secondary_index" {
    for_each = local.gsi_configs
    iterator = gsi
    content {
      name               = gsi.value.name
      hash_key           = gsi.value.hash_key
      range_key          = gsi.value.range_key
      projection_type    = gsi.value.projection_type
      non_key_attributes = gsi.value.projection_type == "INCLUDE" ? gsi.value.non_key_attributes : null

      # Capacity only when PROVISIONED
      read_capacity  = var.dynamodb.billing_mode == "PROVISIONED" ? gsi.value.read_capacity  : null
      write_capacity = var.dynamodb.billing_mode == "PROVISIONED" ? gsi.value.write_capacity : null
    }
  }

  #############################
  # Local secondary indexes
  #############################
  dynamic "local_secondary_index" {
    for_each = local.lsi_configs
    iterator = lsi
    content {
      name               = lsi.value.name
      range_key          = lsi.value.range_key
      projection_type    = lsi.value.projection_type
      non_key_attributes = lsi.value.projection_type == "INCLUDE" ? lsi.value.non_key_attributes : null
    }
  }

  #############################
  # Streams
  #############################
  dynamic "stream_specification" {
    for_each = local.stream_enabled ? [var.dynamodb.stream_specification] : []
    content {
      stream_enabled   = true
      stream_view_type = stream_specification.value.view_type
    }
  }

  #############################
  # TTL configuration
  #############################
  dynamic "ttl" {
    for_each = var.dynamodb.ttl_specification != null ? [var.dynamodb.ttl_specification] : []
    content {
      enabled        = ttl.value.enabled
      attribute_name = ttl.value.attribute_name
    }
  }

  #############################
  # Server-side encryption
  #############################
  dynamic "server_side_encryption" {
    for_each = var.dynamodb.sse_specification != null && var.dynamodb.sse_specification.enabled ? [1] : []
    content {
      enabled     = true
      kms_key_arn = local.kms_key_arn
    }
  }

  #############################
  # Point-in-time recovery
  #############################
  point_in_time_recovery {
    enabled = var.dynamodb.point_in_time_recovery_enabled
  }
}

########################
# Stack-style outputs
########################
output "table_arn" {
  description = "Fully-qualified Amazon Resource Name of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

output "table_name" {
  description = "Name of the DynamoDB table (might include suffixes added by Terraform)."
  value       = aws_dynamodb_table.this.name
}

output "table_id" {
  description = "AWS-assigned unique identifier of the table."
  value       = aws_dynamodb_table.this.id
}

output "stream" {
  description = "Most-recent DynamoDB stream details – null when streams are disabled."
  value = local.stream_enabled ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

output "kms_key_arn" {
  description = "ARN of the customer-managed CMK when SSE uses KMS; null otherwise."
  value       = local.kms_key_arn
}

output "global_secondary_index_names" {
  description = "Names of all provisioned global secondary indexes."
  value       = [for idx in aws_dynamodb_table.this.global_secondary_index : idx.name]
}

output "local_secondary_index_names" {
  description = "Names of all provisioned local secondary indexes."
  value       = [for idx in aws_dynamodb_table.this.local_secondary_index : idx.name]
}
