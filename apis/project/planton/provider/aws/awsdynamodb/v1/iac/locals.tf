locals {
  # Reference to the incoming protobuf-shaped specification.
  spec = var.spec

  ############################
  # Enum → provider mappings #
  ############################
  attribute_type_map = {
    1 = "S" # STRING
    2 = "N" # NUMBER
    3 = "B" # BINARY
  }

  key_type_map = {
    1 = "HASH"
    2 = "RANGE"
  }

  billing_mode_map = {
    1 = "PROVISIONED"
    2 = "PAY_PER_REQUEST"
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

  sse_type_map = {
    1 = "AES256"
    2 = "KMS"
  }

  ############################
  # Table-level configuration #
  ############################
  table_name = local.spec.table_name

  # Attribute definitions – translated & de-duplicated
  attribute_definitions = distinct([
    for attr in coalesce(local.spec.attribute_definitions, []) : {
      name = attr.attribute_name
      type = local.attribute_type_map[attr.attribute_type]
    }
  ])

  # Primary key schema
  key_schema = [
    for ks in coalesce(local.spec.key_schema, []) : {
      attribute_name = ks.attribute_name
      key_type       = local.key_type_map[ks.key_type]
    }
  ]

  #################
  # Billing model #
  #################
  billing_mode = local.billing_mode_map[local.spec.billing_mode]

  provisioned_throughput = local.spec.billing_mode == 1 ? {
    read_capacity  = local.spec.provisioned_throughput.read_capacity_units
    write_capacity = local.spec.provisioned_throughput.write_capacity_units
  } : null

  ############################
  # Global secondary indexes #
  ############################
  global_secondary_indexes = [
    for gsi in coalesce(local.spec.global_secondary_indexes, []) : {
      name      = gsi.index_name
      hash_key  = [for ks in gsi.key_schema : ks.attribute_name if ks.key_type == 1][0]
      range_key = length([for ks in gsi.key_schema : ks.key_type if ks.key_type == 2]) > 0 ? [for ks in gsi.key_schema : ks.attribute_name if ks.key_type == 2][0] : null

      projection_type    = local.projection_type_map[gsi.projection.projection_type]
      non_key_attributes = gsi.projection.projection_type == 3 ? gsi.projection.non_key_attributes : null

      # Capacities are only emitted when PROVISIONED mode is selected at the table level
      read_capacity  = local.spec.billing_mode == 1 ? gsi.provisioned_throughput.read_capacity_units  : null
      write_capacity = local.spec.billing_mode == 1 ? gsi.provisioned_throughput.write_capacity_units : null
    }
  ]

  ###########################
  # Local secondary indexes #
  ###########################
  local_secondary_indexes = [
    for lsi in coalesce(local.spec.local_secondary_indexes, []) : {
      name      = lsi.index_name
      range_key = [for ks in lsi.key_schema : ks.attribute_name if ks.key_type == 2][0]

      projection_type    = local.projection_type_map[lsi.projection.projection_type]
      non_key_attributes = lsi.projection.projection_type == 3 ? lsi.projection.non_key_attributes : null
    }
  ]

  #######################
  # Stream specification #
  #######################
  streams_enabled  = try(local.spec.stream_specification.stream_enabled, false)
  stream_view_type = local.streams_enabled ? local.stream_view_type_map[local.spec.stream_specification.stream_view_type] : null

  ########################
  # TTL (time-to-live)   #
  ########################
  ttl_specification = try(local.spec.ttl_specification.ttl_enabled, false) ? {
    enabled        = true
    attribute_name = local.spec.ttl_specification.attribute_name
  } : null

  ################################
  # Server-side encryption (SSE) #
  ################################
  sse_specification = try(local.spec.sse_specification.enabled, false) ? {
    enabled           = true
    sse_type          = local.sse_type_map[local.spec.sse_specification.sse_type]
    kms_master_key_id = local.spec.sse_specification.sse_type == 2 ? local.spec.sse_specification.kms_master_key_id : null
  } : {
    enabled = false
  }

  ########################################
  # Point-in-time recovery & user  tags  #
  ########################################
  point_in_time_recovery_enabled = local.spec.point_in_time_recovery_enabled

  tags = merge(
    local.spec.tags != null ? local.spec.tags : {},
    tomap({})
  )
}
