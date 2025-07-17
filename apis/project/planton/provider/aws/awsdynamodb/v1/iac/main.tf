terraform {
  required_version = ">= 1.3.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

###############################################################################
#  PROVIDER                                                                    
###############################################################################
provider "aws" {
  region = var.aws_region
}

###############################################################################
#  INPUT VARIABLES                                                             
###############################################################################
variable "aws_region" {
  description = "AWS region the resources will be created in."
  type        = string
  default     = "us-east-1"
}

variable "table_name" {
  description = "Name of the DynamoDB table."
  type        = string
  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "table_name must be 3–255 characters long."
  }
}

variable "hash_key" {
  description = "Partition (hash) key name for the table."
  type        = string
  validation {
    condition     = length(var.hash_key) >= 1 && length(var.hash_key) <= 255
    error_message = "hash_key must be between 1 and 255 characters long."
  }
}

variable "range_key" {
  description = "Optional sort (range) key name for the table."
  type        = string
  default     = ""
  validation {
    condition     = var.range_key == "" || (length(var.range_key) >= 1 && length(var.range_key) <= 255)
    error_message = "range_key must be empty or 1–255 characters long."
  }
}

variable "attributes" {
  description = <<EOT
Complete list of attribute definitions for the table and every index.
Each object requires:
  - name : Attribute name
  - type : One of "S", "N" or "B"
EOT
  type = list(object({
    name = string
    type = string
  }))
  validation {
    condition = length(var.attributes) > 0 && alltrue([
      for a in var.attributes : contains(["S", "N", "B"], a.type)
    ])
    error_message = "Each attribute must have type 'S', 'N' or 'B', and at least one attribute definition is required."
  }
}

variable "billing_mode" {
  description = "Billing mode for the table: PAY_PER_REQUEST or PROVISIONED"
  type        = string
  default     = "PAY_PER_REQUEST"
  validation {
    condition     = contains(["PAY_PER_REQUEST", "PROVISIONED"], upper(var.billing_mode))
    error_message = "billing_mode must be either PAY_PER_REQUEST or PROVISIONED."
  }
}

variable "read_capacity" {
  description = "Provisioned read capacity units (only when billing_mode = PROVISIONED)."
  type        = number
  default     = 5
  validation {
    condition     = var.billing_mode == "PROVISIONED" ? var.read_capacity > 0 : true
    error_message = "read_capacity must be > 0 when billing_mode is PROVISIONED."
  }
}

variable "write_capacity" {
  description = "Provisioned write capacity units (only when billing_mode = PROVISIONED)."
  type        = number
  default     = 5
  validation {
    condition     = var.billing_mode == "PROVISIONED" ? var.write_capacity > 0 : true
    error_message = "write_capacity must be > 0 when billing_mode is PROVISIONED."
  }
}

variable "global_secondary_indexes" {
  description = <<EOT
List of global secondary indexes (GSI).
Each item can contain:
  - name                (string, required)
  - hash_key            (string, required)
  - range_key           (string, optional)
  - projection_type     (string, one of ALL|KEYS_ONLY|INCLUDE, required)
  - non_key_attributes  (list(string), optional)
  - read_capacity       (number, required if billing_mode = PROVISIONED)
  - write_capacity      (number, required if billing_mode = PROVISIONED)
EOT
  type = list(object({
    name               = string
    hash_key           = string
    range_key          = optional(string)
    projection_type    = string
    non_key_attributes = optional(list(string), [])
    read_capacity      = optional(number)
    write_capacity     = optional(number)
  }))
  default = []
  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes : contains(["ALL", "KEYS_ONLY", "INCLUDE"], g.projection_type)
    ])
    error_message = "projection_type for each GSI must be ALL, KEYS_ONLY, or INCLUDE."
  }
}

variable "local_secondary_indexes" {
  description = <<EOT
List of local secondary indexes (LSI).
Each item can contain:
  - name                (string, required)
  - range_key           (string, required)
  - projection_type     (string, one of ALL|KEYS_ONLY|INCLUDE, required)
  - non_key_attributes  (list(string), optional)
EOT
  type = list(object({
    name               = string
    range_key          = string
    projection_type    = string
    non_key_attributes = optional(list(string), [])
  }))
  default = []
  validation {
    condition = alltrue([
      for l in var.local_secondary_indexes : contains(["ALL", "KEYS_ONLY", "INCLUDE"], l.projection_type)
    ])
    error_message = "projection_type for each LSI must be ALL, KEYS_ONLY, or INCLUDE."
  }
}

variable "stream_enabled" {
  description = "Enable DynamoDB Streams."
  type        = bool
  default     = false
}

variable "stream_view_type" {
  description = "When streams are enabled, what data is written to the stream. One of NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | KEYS_ONLY."
  type        = string
  default     = "NEW_AND_OLD_IMAGES"
  validation {
    condition     = !var.stream_enabled || contains(["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "KEYS_ONLY"], var.stream_view_type)
    error_message = "stream_view_type must be a valid value when stream_enabled is true."
  }
}

variable "ttl_enabled" {
  description = "Enable TTL for the table."
  type        = bool
  default     = false
}

variable "ttl_attribute_name" {
  description = "Name of the attribute that stores TTL epoch timestamp. Required when ttl_enabled = true."
  type        = string
  default     = ""
  validation {
    condition     = !var.ttl_enabled || (length(var.ttl_attribute_name) >= 1 && length(var.ttl_attribute_name) <= 255)
    error_message = "ttl_attribute_name must be between 1 and 255 chars when ttl_enabled is true."
  }
}

variable "sse_enabled" {
  description = "Enable server-side encryption (SSE)."
  type        = bool
  default     = true
}

variable "sse_type" {
  description = "Type of SSE. Allowed values: AES256, KMS. When sse_enabled = false, leave as empty string."
  type        = string
  default     = "KMS"
  validation {
    condition     = !var.sse_enabled || contains(["AES256", "KMS"], var.sse_type)
    error_message = "sse_type must be AES256 or KMS when SSE is enabled."
  }
}

variable "kms_key_arn" {
  description = "Existing KMS key ARN used when sse_type = KMS. Leave empty to create a new CMK automatically."
  type        = string
  default     = ""
}

variable "tags" {
  description = "Map of tags to apply to all resources."
  type        = map(string)
  default     = {}
}

###############################################################################
#  LOCALS                                                                      
###############################################################################
locals {
  billing_mode_upper = upper(var.billing_mode)
  sse_type_upper     = upper(var.sse_type)
  # True when we need provisioned throughput values
  use_provisioned_capacity = local.billing_mode_upper == "PROVISIONED"

  # Computed list of global secondary index names (for outputs)
  gsi_names = [for g in var.global_secondary_indexes : g.name]
  lsi_names = [for l in var.local_secondary_indexes : l.name]
}

###############################################################################
#  OPTIONAL KMS KEY (only when SSE enabled with KMS & no external ARN)         
###############################################################################
resource "aws_kms_key" "this" {
  count               = var.sse_enabled && local.sse_type_upper == "KMS" && var.kms_key_arn == "" ? 1 : 0
  description         = "Customer managed CMK for DynamoDB table ${var.table_name}"
  deletion_window_in_days = 10
  enable_key_rotation = true
  tags                = var.tags
}

###############################################################################
#  DYNAMODB TABLE                                                              
###############################################################################
resource "aws_dynamodb_table" "this" {
  name         = var.table_name
  hash_key     = var.hash_key
  billing_mode = local.billing_mode_upper
  tags         = var.tags

  dynamic "range_key" {
    # Terraform requires the attribute exists; we wrap in dynamic block to add only when given
    for_each = var.range_key != "" ? [var.range_key] : []
    content  = range_key.value
  }

  #######################
  # Attribute definitions
  #######################
  dynamic "attribute" {
    for_each = var.attributes
    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  ########################################
  # Provisioned throughput (conditional)  
  ########################################
  dynamic "provisioned_throughput" {
    for_each = local.use_provisioned_capacity ? [1] : []
    content {
      read_capacity  = var.read_capacity
      write_capacity = var.write_capacity
    }
  }

  #######################
  # Server-side encryption
  #######################
  server_side_encryption {
    enabled     = var.sse_enabled
    # kms_key_arn accepted only for KMS; otherwise omit with null
    kms_key_arn = var.sse_enabled && local.sse_type_upper == "KMS" ? (
      var.kms_key_arn != "" ? var.kms_key_arn : aws_kms_key.this[0].arn
    ) : null
  }

  #######################
  # Streams             
  #######################
  stream_enabled   = var.stream_enabled
  stream_view_type = var.stream_enabled ? var.stream_view_type : null

  #######################
  # TTL                 
  #######################
  dynamic "ttl" {
    for_each = var.ttl_enabled ? [1] : []
    content {
      enabled        = true
      attribute_name = var.ttl_attribute_name
    }
  }

  #######################
  # Global secondary indexes
  #######################
  dynamic "global_secondary_index" {
    for_each = var.global_secondary_indexes
    content {
      name               = global_secondary_index.value.name
      hash_key           = global_secondary_index.value.hash_key
      range_key          = lookup(global_secondary_index.value, "range_key", null)
      projection_type    = global_secondary_index.value.projection_type
      non_key_attributes = length(lookup(global_secondary_index.value, "non_key_attributes", [])) > 0 ? global_secondary_index.value.non_key_attributes : null

      dynamic "provisioned_throughput" {
        for_each = local.use_provisioned_capacity ? [1] : []
        content {
          read_capacity  = global_secondary_index.value.read_capacity
          write_capacity = global_secondary_index.value.write_capacity
        }
      }
    }
  }

  #######################
  # Local secondary indexes
  #######################
  dynamic "local_secondary_index" {
    for_each = var.local_secondary_indexes
    content {
      name               = local_secondary_index.value.name
      range_key          = local_secondary_index.value.range_key
      projection_type    = local_secondary_index.value.projection_type
      non_key_attributes = length(lookup(local_secondary_index.value, "non_key_attributes", [])) > 0 ? local_secondary_index.value.non_key_attributes : null
    }
  }
}

###############################################################################
#  OUTPUTS                                                                     
###############################################################################
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
  value = var.stream_enabled ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when SSE uses a CMK."
  value = var.sse_enabled && local.sse_type_upper == "KMS" ? (
    var.kms_key_arn != "" ? var.kms_key_arn : aws_kms_key.this[0].arn
  ) : null
}

output "global_secondary_index_names" {
  description = "Names of provisioned global secondary indexes (GSIs)."
  value       = local.gsi_names
}

output "local_secondary_index_names" {
  description = "Names of provisioned local secondary indexes (LSIs)."
  value       = local.lsi_names
}
