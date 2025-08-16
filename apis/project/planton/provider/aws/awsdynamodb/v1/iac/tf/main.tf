resource "aws_dynamodb_table" "this" {
  name = var.spec.table_name != "" ? var.spec.table_name : var.metadata.name

  # Billing
  billing_mode  = var.spec.billing_mode == "PAY_PER_REQUEST" ? "PAY_PER_REQUEST" : null
  read_capacity  = var.spec.billing_mode == "PROVISIONED" ? var.spec.read_capacity_units : null
  write_capacity = var.spec.billing_mode == "PROVISIONED" ? var.spec.write_capacity_units : null

  # Keys
  hash_key  = var.spec.partition_key_name
  range_key = var.spec.sort_key_name != "" ? var.spec.sort_key_name : null

  # Attributes
  attribute {
    name = var.spec.partition_key_name
    type = var.spec.partition_key_type
  }

  dynamic "attribute" {
    for_each = var.spec.sort_key_name != "" ? [1] : []
    content {
      name = var.spec.sort_key_name
      type = var.spec.sort_key_type
    }
  }

  # Point-in-time recovery
  point_in_time_recovery {
    enabled = try(var.spec.point_in_time_recovery_enabled, false)
  }

  # Server-side encryption (AWS owned key when enabled)
  dynamic "server_side_encryption" {
    for_each = try(var.spec.server_side_encryption_enabled, false) ? [1] : []
    content {
      enabled = true
    }
  }

  tags = {
    "planton:resource"      = "true"
    "planton:organization"  = var.metadata.org
    "planton:environment"   = var.metadata.env
    "planton:resource_kind" = "AwsDynamodb"
    "planton:resource_id"   = var.metadata.id
  }
}


