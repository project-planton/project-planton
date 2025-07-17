# -----------------------------------------------------------------------------
# DynamoDB table driven by the AwsDynamodbSpec proto definition                 
# -----------------------------------------------------------------------------

terraform {
  required_version = ">= 1.3"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# -----------------------------------------------------------------------------
# Provider & variables                                                           
# -----------------------------------------------------------------------------

provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  description = "AWS region where the table will be created"
  type        = string
}

variable "spec" {
  description = "AwsDynamodbSpec message converted to a Terraform map/object"
  type        = any
}

# -----------------------------------------------------------------------------
# Local helpers – maps that translate numeric enums from the proto into the      
# strings expected by the AWS provider.                                          
# -----------------------------------------------------------------------------

locals {
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

  billing_mode_map = {
    1 = "PROVISIONED"
    2 = "PAY_PER_REQUEST"
  }

  stream_view_type_map = {
    1 = "NEW_IMAGE"
    2 = "OLD_IMAGE"
    3 = "NEW_AND_OLD_IMAGES"
    4 = "KEYS_ONLY"
  }

  # ---------------------------------------------------------------------------
  # Primary key                                                               
  # ---------------------------------------------------------------------------

  hash_key  = one([for ks in var.spec.key_schema : ks.attribute_name if ks.key_type == 1])
  range_key = try(one([for ks in var.spec.key_schema : ks.attribute_name if ks.key_type == 2]), null)

  # ---------------------------------------------------------------------------
  # Global & local secondary indexes                                           
  # ---------------------------------------------------------------------------

  global_secondary_indexes = [
    for g in try(var.spec.global_secondary_indexes, []) : {
      name               = g.index_name
      hash_key           = one([for ks in g.key_schema : ks.attribute_name if ks.key_type == 1])
      range_key          = try(one([for ks in g.key_schema : ks.attribute_name if ks.key_type == 2]), null)
      projection_type    = local.projection_type_map[g.projection.projection_type]
      non_key_attributes = g.projection.projection_type == 3 ? g.projection.non_key_attributes : null
      read_capacity      = var.spec.billing_mode == 1 ? g.provisioned_throughput.read_capacity_units : null
      write_capacity     = var.spec.billing_mode == 1 ? g.provisioned_throughput.write_capacity_units : null
    }
  ]

  local_secondary_indexes = [
    for l in try(var.spec.local_secondary_indexes, []) : {
      name               = l.index_name
      range_key          = one([for ks in l.key_schema : ks.attribute_name if ks.key_type == 2])
      projection_type    = local.projection_type_map[l.projection.projection_type]
      non_key_attributes = l.projection.projection_type == 3 ? l.projection.non_key_attributes : null
    }
  ]
}

# -----------------------------------------------------------------------------
# DynamoDB table                                                                
# -----------------------------------------------------------------------------

resource "aws_dynamodb_table" "this" {
  name         = var.spec.table_name
  billing_mode = local.billing_mode_map[var.spec.billing_mode]

  hash_key  = local.hash_key
  range_key = local.range_key

  ############################################################
  # Attribute definitions                                    #
  ############################################################
  dynamic "attribute" {
    for_each = var.spec.attribute_definitions
    content {
      name = attribute.value.attribute_name
      type = local.attribute_type_map[attribute.value.attribute_type]
    }
  }

  ############################################################
  # Global secondary indexes                                 #
  ############################################################
  dynamic "global_secondary_index" {
    for_each = local.global_secondary_indexes
    iterator = gsi
    content {
      name            = gsi.value.name
      hash_key        = gsi.value.hash_key
      range_key       = gsi.value.range_key
      projection_type = gsi.value.projection_type

      # Only set when projection_type == INCLUDE
      non_key_attributes = gsi.value.non_key_attributes

      # Only relevant for PROVISIONED mode
      read_capacity  = gsi.value.read_capacity
      write_capacity = gsi.value.write_capacity
    }
  }

  ############################################################
  # Local secondary indexes                                  #
  ############################################################
  dynamic "local_secondary_index" {
    for_each = local.local_secondary_indexes
    iterator = lsi
    content {
      name            = lsi.value.name
      range_key       = lsi.value.range_key
      projection_type = lsi.value.projection_type
      non_key_attributes = lsi.value.non_key_attributes
    }
  }

  ############################################################
  # Capacity (only when PROVISIONED)                         #
  ############################################################
  read_capacity  = local.billing_mode_map[var.spec.billing_mode] == "PROVISIONED" ? var.spec.provisioned_throughput.read_capacity_units : null
  write_capacity = local.billing_mode_map[var.spec.billing_mode] == "PROVISIONED" ? var.spec.provisioned_throughput.write_capacity_units : null

  ############################################################
  # Streams                                                  #
  ############################################################
  stream_enabled   = var.spec.stream_specification.stream_enabled
  stream_view_type = var.spec.stream_specification.stream_enabled ? local.stream_view_type_map[var.spec.stream_specification.stream_view_type] : null

  ############################################################
  # TTL                                                      #
  ############################################################
  dynamic "ttl" {
    for_each = try(var.spec.ttl_specification, null) == null ? [] : [var.spec.ttl_specification]
    content {
      attribute_name = ttl.value.attribute_name
      enabled        = ttl.value.ttl_enabled
    }
  }

  ############################################################
  # Point-in-time recovery                                   #
  ############################################################
  point_in_time_recovery {
    enabled = var.spec.point_in_time_recovery_enabled
  }

  ############################################################
  # Server-side encryption                                   #
  ############################################################
  server_side_encryption {
    enabled     = var.spec.sse_specification.enabled
    kms_key_arn = var.spec.sse_specification.sse_type == 2 ? var.spec.sse_specification.kms_master_key_id : null
  }

  ############################################################
  # Tags                                                     #
  ############################################################
  tags = var.spec.tags
}

# -----------------------------------------------------------------------------
# Outputs – match AwsDynamodbStackOutputs                                       
# -----------------------------------------------------------------------------

output "table_arn" {
  description = "Fully-qualified ARN of the DynamoDB table"
  value       = aws_dynamodb_table.this.arn
}

output "table_name" {
  description = "Name of the DynamoDB table"
  value       = aws_dynamodb_table.this.name
}

output "table_id" {
  description = "Unique identifier of the table"
  value       = aws_dynamodb_table.this.id
}

output "stream" {
  description = "Stream identifiers (null when streams are disabled)"
  value = aws_dynamodb_table.this.stream_arn != "" ? {
    stream_arn   = aws_dynamodb_table.this.stream_arn
    stream_label = aws_dynamodb_table.this.stream_label
  } : null
}

output "kms_key_arn" {
  description = "ARN of the customer-managed KMS key when SSE uses KMS"
  value       = try(aws_dynamodb_table.this.kms_key_arn, null)
}

output "global_secondary_index_names" {
  description = "Names of provisioned GSIs"
  value       = [for g in aws_dynamodb_table.this.global_secondary_index : g.name]
}

output "local_secondary_index_names" {
  description = "Names of provisioned LSIs"
  value       = [for l in aws_dynamodb_table.this.local_secondary_index : l.name]
}
