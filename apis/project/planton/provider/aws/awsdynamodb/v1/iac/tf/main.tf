# AWS DynamoDB Table
resource "aws_dynamodb_table" "table" {
  name           = local.safe_spec.table_name
  billing_mode   = local.safe_spec.billing_mode
  hash_key       = local.safe_spec.partition_key_name
  
  # Sort key (range key) - only if specified
  range_key = local.has_sort_key ? local.safe_spec.sort_key_name : null
  
  # Read capacity units - only for PROVISIONED billing
  read_capacity = local.is_provisioned_billing ? local.safe_read_capacity_units : null
  
  # Write capacity units - only for PROVISIONED billing
  write_capacity = local.is_provisioned_billing ? local.safe_write_capacity_units : null
  
  # Point-in-time recovery
  point_in_time_recovery {
    enabled = local.safe_point_in_time_recovery_enabled
  }
  
  # Server-side encryption
  server_side_encryption {
    enabled = local.safe_server_side_encryption_enabled
  }
  
  # Attribute definitions
  attribute {
    name = local.safe_spec.partition_key_name
    type = local.safe_spec.partition_key_type
  }
  
  # Sort key attribute definition - only if specified
  dynamic "attribute" {
    for_each = local.has_sort_key ? [1] : []
    content {
      name = local.safe_spec.sort_key_name
      type = local.safe_spec.sort_key_type
    }
  }
  
  # Tags
  tags = merge(
    {
      Name = local.safe_spec.table_name
      Environment = local.safe_metadata.env
      Organization = local.safe_metadata.org
      ManagedBy = "project-planton"
    },
    {
      for tag in local.safe_metadata.tags : tag => tag
    }
  )
}
