// -----------------------------------------------------------------------------
//  DynamoDB table module – implements the AwsDynamodbSpec protobuf contract
// -----------------------------------------------------------------------------
//  The configuration is highly-parameterised so that every feature described in
//  the spec (capacity modes, GSIs/LSIs, TTL, SSE, Streams, Tags …) can be
//  expressed through input variables.
// -----------------------------------------------------------------------------

terraform {
  required_version = ">= 1.3"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  # Configuration (region/profile) is expected to be supplied by the caller.
}

// -----------------------------------------------------------------------------
//  Input variables
// -----------------------------------------------------------------------------

variable "table_name" {
  description = "Name of the DynamoDB table"
  type        = string
}

variable "attribute_definitions" {
  description = "List of attribute definitions used by the table and indexes"
  type = list(object({
    name = string                     // Attribute name
    type = string                     // One of "S", "N", "B"
  }))
}

variable "hash_key" {
  description = "Primary (partition/HASH) key attribute name"
  type        = string
}

variable "range_key" {
  description = "Optional primary (sort/RANGE) key attribute name"
  type        = string
  default     = null
}

variable "billing_mode" {
  description = "Billing mode – either \"PROVISIONED\" or \"PAY_PER_REQUEST\""
  type        = string
  default     = "PAY_PER_REQUEST"
  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], var.billing_mode)
    error_message = "billing_mode must be either PROVISIONED or PAY_PER_REQUEST."
  }
}

variable "provisioned_throughput" {
  description = "Provisioned capacity units – required when billing_mode == PROVISIONED"
  type = object({
    read_capacity  = number
    write_capacity = number
  })
  default = null
}

variable "global_secondary_indexes" {
  description = "Definitions for global secondary indexes (GSIs)"
  type = list(object({
    name               = string
    hash_key           = string
    range_key          = optional(string)
    projection_type    = string               // ALL | KEYS_ONLY | INCLUDE
    non_key_attributes = optional(list(string))
    read_capacity      = optional(number)
    write_capacity     = optional(number)
  }))
  default = []
}

variable "local_secondary_indexes" {
  description = "Definitions for local secondary indexes (LSIs)"
  type = list(object({
    name               = string
    range_key          = string               // Must share HASH key with the table
    projection_type    = string               // ALL | KEYS_ONLY | INCLUDE
    non_key_attributes = optional(list(string))
  }))
  default = []
}

variable "ttl" {
  description = "Time-to-live (TTL) configuration"
  type = object({
    enabled        = bool
    attribute_name = optional(string)
  })
  default = {
    enabled        = false
    attribute_name = null
  }
  validation {
    condition     = (!var.ttl.enabled) || (var.ttl.enabled && var.ttl.attribute_name != null && length(var.ttl.attribute_name) > 0)
    error_message = "When ttl.enabled is true, ttl.attribute_name must be provided."
  }
}

variable "stream_enabled" {
  description = "Enable DynamoDB Streams"
  type        = bool
  default     = false
}

variable "stream_view_type" {
  description = "Stream view type – NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES or KEYS_ONLY"
  type        = string
  default     = null
  validation {
    condition     = (!var.stream_enabled) || contains(["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "KEYS_ONLY"], var.stream_view_type)
    error_message = "stream_view_type must be one of NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | KEYS_ONLY when streams are enabled."
  }
}

variable "server_side_encryption" {
  description = "Server-side encryption (SSE) settings"
  type = object({
    enabled      = bool
    kms_key_arn  = optional(string)
  })
  default = {
    enabled     = false
    kms_key_arn = null
  }
  validation {
    condition     = (!var.server_side_encryption.enabled) || (var.server_side_encryption.enabled && (var.server_side_encryption.kms_key_arn == null || length(var.server_side_encryption.kms_key_arn) > 0))
    error_message = "When SSE is enabled, kms_key_arn may be provided but must be a non-empty string."
  }
}

variable "tags" {
  description = "Key/value tags applied to the table"
  type        = map(string)
  default     = {}
}

// -----------------------------------------------------------------------------
//  DynamoDB table resource
// -----------------------------------------------------------------------------

resource "aws_dynamodb_table" "this" {
  name         = var.table_name
  billing_mode = var.billing_mode

  # Primary key
  hash_key  = var.hash_key
  range_key = var.range_key

  # Provisioned capacity – only when billing_mode == "PROVISIONED"
  read_capacity  = var.billing_mode == "PROVISIONED" ? var.provisioned_throughput.read_capacity  : null
  write_capacity = var.billing_mode == "PROVISIONED" ? var.provisioned_throughput.write_capacity : null

  # ---------------------------------------------------------------------------
  #  Attribute definitions
  # ---------------------------------------------------------------------------
  dynamic "attribute" {
    for_each = { for a in var.attribute_definitions : a.name => a }
    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  # ---------------------------------------------------------------------------
  #  Global secondary indexes
  # ---------------------------------------------------------------------------
  dynamic "global_secondary_index" {
    for_each = var.global_secondary_indexes
    content {
      name               = global_secondary_index.value.name
      hash_key           = global_secondary_index.value.hash_key
      range_key          = try(global_secondary_index.value.range_key, null)
      projection_type    = global_secondary_index.value.projection_type
      non_key_attributes = global_secondary_index.value.projection_type == "INCLUDE" ? try(global_secondary_index.value.non_key_attributes, null) : null

      read_capacity  = var.billing_mode == "PROVISIONED" ? try(global_secondary_index.value.read_capacity, null)  : null
      write_capacity = var.billing_mode == "PROVISIONED" ? try(global_secondary_index.value.write_capacity, null) : null
    }
  }

  # ---------------------------------------------------------------------------
  #  Local secondary indexes
  # ---------------------------------------------------------------------------
  dynamic "local_secondary_index" {
    for_each = var.local_secondary_indexes
    content {
      name               = local_secondary_index.value.name
      range_key          = local_secondary_index.value.range_key
      projection_type    = local_secondary_index.value.projection_type
      non_key_attributes = local_secondary_index.value.projection_type == "INCLUDE" ? try(local_secondary_index.value.non_key_attributes, null) : null
    }
  }

  # ---------------------------------------------------------------------------
  #  Time-to-live (TTL)
  # ---------------------------------------------------------------------------
  dynamic "ttl" {
    for_each = var.ttl.enabled ? [var.ttl] : []
    content {
      attribute_name = ttl.value.attribute_name
      enabled        = true
    }
  }

  # ---------------------------------------------------------------------------
  #  DynamoDB Streams
  # ---------------------------------------------------------------------------
  stream_enabled   = var.stream_enabled
  stream_view_type = var.stream_enabled ? var.stream_view_type : null

  # ---------------------------------------------------------------------------
  #  Server-side encryption
  # ---------------------------------------------------------------------------
  dynamic "server_side_encryption" {
    for_each = var.server_side_encryption.enabled ? [var.server_side_encryption] : []
    content {
      enabled     = true
      kms_key_arn = try(server_side_encryption.value.kms_key_arn, null)
    }
  }

  tags = var.tags
}

// -----------------------------------------------------------------------------
//  Outputs – surface identifiers required by AwsDynamodbStackOutputs proto
// -----------------------------------------------------------------------------

output "table_arn" {
  description = "Fully-qualified Amazon Resource Name of the table"
  value       = aws_dynamodb_table.this.arn
}

output "table_name" {
  description = "Name of the DynamoDB table (may include runtime suffixes)"
  value       = aws_dynamodb_table.this.name
}

output "table_id" {
  description = "AWS-assigned unique identifier of the table"
  value       = aws_dynamodb_table.this.id
}

output "stream" {
  description = "Current (latest) stream information – null when streams are disabled"
  value = aws_dynamodb_table.this.stream_enabled ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when SSE uses a CMK"
  value       = try(var.server_side_encryption.kms_key_arn, null)
}

output "global_secondary_index_names" {
  description = "Names of provisioned global secondary indexes (GSIs)"
  value       = [for g in aws_dynamodb_table.this.global_secondary_index : g["name"]]
}

output "local_secondary_index_names" {
  description = "Names of provisioned local secondary indexes (LSIs)"
  value       = [for l in aws_dynamodb_table.this.local_secondary_index : l["name"]]
}
