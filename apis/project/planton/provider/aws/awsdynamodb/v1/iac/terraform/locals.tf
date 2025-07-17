############################################################
# Locals for the Amazon DynamoDB table module
# ----------------------------------------------------------
# The following local values translate the strongly-typed
# protobuf-shaped "spec" variable into the exact structures
# and literal values expected by the Terraform AWS provider.
# Doing this translation once here keeps the individual
# resource blocks tidy and makes their intent easier to read.
############################################################

terraform {
  required_version = ">= 1.3.0"
}

locals {
  ###########################################################################
  # Enum → literal lookup tables
  ###########################################################################
  attribute_type_map = {
    STRING = "S"
    NUMBER = "N"
    BINARY = "B"
  }

  key_type_map = {
    HASH  = "HASH"
    RANGE = "RANGE"
  }

  projection_type_map = {
    ALL       = "ALL"
    KEYS_ONLY = "KEYS_ONLY"
    INCLUDE   = "INCLUDE"
  }

  stream_view_type_map = {
    NEW_IMAGE          = "NEW_IMAGE"
    OLD_IMAGE          = "OLD_IMAGE"
    NEW_AND_OLD_IMAGES = "NEW_AND_OLD_IMAGES"
    STREAM_KEYS_ONLY   = "KEYS_ONLY"   # AWS names this simply KEYS_ONLY.
    KEYS_ONLY          = "KEYS_ONLY"   # Accept either spelling.
  }

  sse_type_map = {
    AES256 = "AES256"
    KMS    = "KMS"
  }

  ###########################################################################
  # Shorthand/derived flags
  ###########################################################################
  is_provisioned = var.spec.billing_mode == "PROVISIONED"

  ###########################################################################
  # Top-level table settings
  ###########################################################################
  table_name   = var.spec.table_name
  billing_mode = var.spec.billing_mode

  provisioned_throughput = local.is_provisioned ? {
    read_capacity  = var.spec.provisioned_throughput.read_capacity_units
    write_capacity = var.spec.provisioned_throughput.write_capacity_units
  } : null

  attribute_definitions = [
    for a in var.spec.attribute_definitions : {
      name = a.attribute_name
      type = local.attribute_type_map[a.attribute_type]
    }
  ]

  key_schema = [
    for k in var.spec.key_schema : {
      attribute_name = k.attribute_name
      key_type       = local.key_type_map[k.key_type]
    }
  ]

  ###########################################################################
  # Global secondary indexes (GSIs)
  ###########################################################################
  global_secondary_indexes = [
    for g in var.spec.global_secondary_indexes : {
      name = g.index_name

      # Extract HASH / RANGE keys from the key_schema list
      hash_key  = [for k in g.key_schema : k.attribute_name if local.key_type_map[k.key_type] == "HASH"][0]
      range_key = try([
        for k in g.key_schema : k.attribute_name if local.key_type_map[k.key_type] == "RANGE"
      ][0], null)

      projection_type    = local.projection_type_map[g.projection.projection_type]
      non_key_attributes = g.projection.projection_type == "INCLUDE" ? g.projection.non_key_attributes : null

      # Capacity can only be set when the table uses PROVISIONED mode
      read_capacity  = local.is_provisioned ? g.provisioned_throughput.read_capacity_units  : null
      write_capacity = local.is_provisioned ? g.provisioned_throughput.write_capacity_units : null
    }
  ]

  gsi_names = [for g in local.global_secondary_indexes : g.name]

  ###########################################################################
  # Local secondary indexes (LSIs)
  ###########################################################################
  local_secondary_indexes = [
    for l in var.spec.local_secondary_indexes : {
      name      = l.index_name
      range_key = [for k in l.key_schema : k.attribute_name if local.key_type_map[k.key_type] == "RANGE"][0]

      projection_type    = local.projection_type_map[l.projection.projection_type]
      non_key_attributes = l.projection.projection_type == "INCLUDE" ? l.projection.non_key_attributes : null
    }
  ]

  lsi_names = [for l in local.local_secondary_indexes : l.name]

  ###########################################################################
  # Streams, TTL & Server-Side Encryption (SSE)
  ###########################################################################
  stream_specification = var.spec.stream_specification.stream_enabled ? {
    stream_enabled   = true
    stream_view_type = local.stream_view_type_map[var.spec.stream_specification.stream_view_type]
  } : null

  ttl_specification = var.spec.ttl_specification.ttl_enabled ? {
    enabled        = true
    attribute_name = var.spec.ttl_specification.attribute_name
  } : null

  sse_specification = var.spec.sse_specification.enabled ? {
    enabled          = true
    sse_type         = local.sse_type_map[var.spec.sse_specification.sse_type]
    kms_master_key_id = (
      var.spec.sse_specification.sse_type == "KMS" ?
      var.spec.sse_specification.kms_master_key_id : null
    )
  } : null

  kms_key_arn = (
    var.spec.sse_specification.enabled &&
    var.spec.sse_specification.sse_type == "KMS"
  ) ? var.spec.sse_specification.kms_master_key_id : null

  ###########################################################################
  # Tags (always add a ManagedBy tag so we can quickly identify resources)
  ###########################################################################
  tags = merge(
    {
      "ManagedBy" = "terraform"
    },
    var.spec.tags
  )
}
