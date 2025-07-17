// -----------------------------------------------------------------------------
//  Project-Planton – AWS DynamoDB table module
//  This file translates the AwsDynamodbSpec proto into native Terraform HCL.
//  All proto fields are represented through a single variable  "aws_dynamodb".
// -----------------------------------------------------------------------------

terraform {
  required_version = ">= 1.3"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

provider "aws" {
  # The caller/stack is expected to configure region/profile/assume-role, etc.
}

# -----------------------------------------------------------------------------
#  Input – full in-memory representation of the AwsDynamodbSpec proto message.
# -----------------------------------------------------------------------------
variable "aws_dynamodb" {
  description = "Rendered AwsDynamodbSpec proto (converted to tf vars)."
  type = object({
    table_name                     = string
    attribute_definitions          = list(object({ attribute_name = string, attribute_type = string }))
    key_schema                     = list(object({ attribute_name = string, key_type = string }))
    billing_mode                   = string                                 # "PROVISIONED" | "PAY_PER_REQUEST"
    provisioned_throughput         = optional(object({ read_capacity_units = number, write_capacity_units = number }))
    global_secondary_indexes = optional(list(object({
      index_name              = string
      key_schema              = list(object({ attribute_name = string, key_type = string }))
      projection              = object({ projection_type = string, non_key_attributes = optional(list(string)) })
      provisioned_throughput  = optional(object({ read_capacity_units = number, write_capacity_units = number }))
    })))
    local_secondary_indexes = optional(list(object({
      index_name = string
      key_schema = list(object({ attribute_name = string, key_type = string }))
      projection = object({ projection_type = string, non_key_attributes = optional(list(string)) })
    })))
    stream_specification = optional(object({ stream_enabled = bool, stream_view_type = string }))
    ttl_specification    = optional(object({ ttl_enabled = bool, attribute_name = string }))
    sse_specification = optional(object({
      enabled           = bool
      sse_type          = string   # "AES256" | "KMS"
      kms_master_key_id = optional(string)
    }))
    point_in_time_recovery_enabled = optional(bool)
    tags                           = optional(map(string))
  })
}

# -----------------------------------------------------------------------------
#  Locals – helper calculations.
# -----------------------------------------------------------------------------
locals {
  primary_hash_key  = one([for k in var.aws_dynamodb.key_schema : k.attribute_name if upper(k.key_type) == "HASH"])
  primary_range_key = try(one([for k in var.aws_dynamodb.key_schema : k.attribute_name if upper(k.key_type) == "RANGE"]), null)

  # When SSE is enabled with KMS but no CMK id was passed, we create one.
  sse_kms_required = (
    var.aws_dynamodb.sse_specification != null &&
    var.aws_dynamodb.sse_specification.enabled &&
    upper(var.aws_dynamodb.sse_specification.sse_type) == "KMS" &&
    try(var.aws_dynamodb.sse_specification.kms_master_key_id, "") == ""
  )

  kms_key_arn = local.sse_kms_required ? aws_kms_key.this[0].arn : try(var.aws_dynamodb.sse_specification.kms_master_key_id, null)
}

# -----------------------------------------------------------------------------
#  Optional KMS CMK creation (only if required).
# -----------------------------------------------------------------------------
resource "aws_kms_key" "this" {
  count               = local.sse_kms_required ? 1 : 0
  description         = "CMK for DynamoDB table ${var.aws_dynamodb.table_name} (managed by Project-Planton)"
  enable_key_rotation = true
  deletion_window_in_days = 30
  tags = merge(
    {
      "Name" = "dynamodb/${var.aws_dynamodb.table_name}/cmk"
    },
    try(var.aws_dynamodb.tags, {})
  )
}

# -----------------------------------------------------------------------------
#  DynamoDB table definition.
# -----------------------------------------------------------------------------
resource "aws_dynamodb_table" "this" {
  name         = var.aws_dynamodb.table_name
  billing_mode = upper(var.aws_dynamodb.billing_mode)

  # ---------------------------------------------------------------------------
  #  Key schema & attribute definitions
  # ---------------------------------------------------------------------------
  hash_key  = local.primary_hash_key
  # range_key is optional. Terraform does not allow omitting attributes
  # dynamically, so we only set it when present.
  dynamic "range_key" {
    for_each = local.primary_range_key == null ? [] : [local.primary_range_key]
    content {
      range_key = range_key.value
    }
  }

  dynamic "attribute" {
    for_each = var.aws_dynamodb.attribute_definitions
    content {
      name = attribute.value.attribute_name
      type = upper(attribute.value.attribute_type)
    }
  }

  # ---------------------------------------------------------------------------
  #  Capacity – only valid for PROVISIONED mode
  # ---------------------------------------------------------------------------
  read_capacity  = upper(var.aws_dynamodb.billing_mode) == "PROVISIONED" ? var.aws_dynamodb.provisioned_throughput.read_capacity_units  : null
  write_capacity = upper(var.aws_dynamodb.billing_mode) == "PROVISIONED" ? var.aws_dynamodb.provisioned_throughput.write_capacity_units : null

  # ---------------------------------------------------------------------------
  #  Streams (optional)
  # ---------------------------------------------------------------------------
  stream_enabled   = try(var.aws_dynamodb.stream_specification.stream_enabled, false)
  stream_view_type = try(var.aws_dynamodb.stream_specification.stream_view_type, null)

  # ---------------------------------------------------------------------------
  #  Time-to-Live (optional)
  # ---------------------------------------------------------------------------
  dynamic "ttl" {
    for_each = var.aws_dynamodb.ttl_specification == null ? [] : [var.aws_dynamodb.ttl_specification]
    content {
      attribute_name = ttl.value.attribute_name
      enabled        = ttl.value.ttl_enabled
    }
  }

  # ---------------------------------------------------------------------------
  #  Server-side encryption (optional)
  # ---------------------------------------------------------------------------
  dynamic "server_side_encryption" {
    for_each = var.aws_dynamodb.sse_specification == null ? [] : [var.aws_dynamodb.sse_specification]
    content {
      enabled     = server_side_encryption.value.enabled
      kms_key_arn = server_side_encryption.value.enabled && upper(server_side_encryption.value.sse_type) == "KMS" ? local.kms_key_arn : null
    }
  }

  # ---------------------------------------------------------------------------
  #  Point-in-time recovery (optional)
  # ---------------------------------------------------------------------------
  dynamic "point_in_time_recovery" {
    for_each = try(var.aws_dynamodb.point_in_time_recovery_enabled, false) ? [1] : []
    content {
      enabled = true
    }
  }

  # ---------------------------------------------------------------------------
  #  Global Secondary Indexes (optional)
  # ---------------------------------------------------------------------------
  dynamic "global_secondary_index" {
    for_each = try(var.aws_dynamodb.global_secondary_indexes, [])
    content {
      name            = global_secondary_index.value.index_name
      hash_key        = one([for k in global_secondary_index.value.key_schema : k.attribute_name if upper(k.key_type) == "HASH"])
      range_key       = try(one([for k in global_secondary_index.value.key_schema : k.attribute_name if upper(k.key_type) == "RANGE"]), null)
      projection_type = upper(global_secondary_index.value.projection.projection_type)
      non_key_attributes = try(global_secondary_index.value.projection.non_key_attributes, null)

      read_capacity  = upper(var.aws_dynamodb.billing_mode) == "PROVISIONED" ? try(global_secondary_index.value.provisioned_throughput.read_capacity_units, null)  : null
      write_capacity = upper(var.aws_dynamodb.billing_mode) == "PROVISIONED" ? try(global_secondary_index.value.provisioned_throughput.write_capacity_units, null) : null
    }
  }

  # ---------------------------------------------------------------------------
  #  Local Secondary Indexes (optional)
  # ---------------------------------------------------------------------------
  dynamic "local_secondary_index" {
    for_each = try(var.aws_dynamodb.local_secondary_indexes, [])
    content {
      name            = local_secondary_index.value.index_name
      # HASH key must be the same as table, so we only pull RANGE:
      range_key       = one([for k in local_secondary_index.value.key_schema : k.attribute_name if upper(k.key_type) == "RANGE"])
      projection_type = upper(local_secondary_index.value.projection.projection_type)
      non_key_attributes = try(local_secondary_index.value.projection.non_key_attributes, null)
    }
  }

  # ---------------------------------------------------------------------------
  #  Tags
  # ---------------------------------------------------------------------------
  tags = merge(
    {
      "Name" = "dynamodb/${var.aws_dynamodb.table_name}"
    },
    try(var.aws_dynamodb.tags, {})
  )
}

# -----------------------------------------------------------------------------
#  Outputs – map 1-to-1 with AwsDynamodbStackOutputs proto message.
# -----------------------------------------------------------------------------
output "table_arn" {
  description = "Fully-qualified ARN of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

output "table_name" {
  description = "Name of the DynamoDB table (possibly with suffixes)."
  value       = aws_dynamodb_table.this.name
}

output "table_id" {
  description = "AWS-assigned logical ID of the table."
  value       = aws_dynamodb_table.this.id
}

output "stream" {
  description = "Current (latest) stream identifiers, null when streams disabled."
  value = (
    aws_dynamodb_table.this.stream_enabled ? {
      stream_arn   = aws_dynamodb_table.this.stream_arn
      stream_label = aws_dynamodb_table.this.stream_label
    } : null
  )
}

output "kms_key_arn" {
  description = "ARN of the CMK used for server-side encryption, when applicable."
  value       = local.kms_key_arn
}

output "global_secondary_index_names" {
  description = "Names of the provisioned GSIs."
  value       = try([for g in aws_dynamodb_table.this.global_secondary_index : g.name], [])
}

output "local_secondary_index_names" {
  description = "Names of the provisioned LSIs."
  value       = try([for l in aws_dynamodb_table.this.local_secondary_index : l.name], [])
}
