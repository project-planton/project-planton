###############################################################################
# Local helpers for building the aws_dynamodb_table resource from the         #
# AwsDynamodbSpec object passed in through `var.spec`.                        #
###############################################################################

locals {
  ###########################################################################
  # Convenience shortcuts & enum -> AWS SDK string look-ups                 #
  ###########################################################################
  spec = var.spec

  # Attribute types to the letter codes expected by the TF provider
  attribute_type_map = {
    "STRING" = "S"
    "NUMBER" = "N"
    "BINARY" = "B"
    1         = "S" # STRING
    2         = "N" # NUMBER
    3         = "B" # BINARY
  }

  # Key type mapping
  key_type_map = {
    "HASH"  = "HASH"
    "RANGE" = "RANGE"
    1        = "HASH"
    2        = "RANGE"
  }

  # Projection type mapping
  projection_type_map = {
    "ALL"       = "ALL"
    "KEYS_ONLY" = "KEYS_ONLY"
    "INCLUDE"   = "INCLUDE"
    1            = "ALL"
    2            = "KEYS_ONLY"
    3            = "INCLUDE"
  }

  # Stream view type mapping
  stream_view_type_map = {
    "NEW_IMAGE"          = "NEW_IMAGE"
    "OLD_IMAGE"          = "OLD_IMAGE"
    "NEW_AND_OLD_IMAGES" = "NEW_AND_OLD_IMAGES"
    "STREAM_KEYS_ONLY"   = "KEYS_ONLY"
    1                     = "NEW_IMAGE"
    2                     = "OLD_IMAGE"
    3                     = "NEW_AND_OLD_IMAGES"
    4                     = "KEYS_ONLY"
  }

  # SSE type mapping
  sse_type_map = {
    "AES256" = "AES256"
    "KMS"    = "KMS"
    1         = "AES256"
    2         = "KMS"
  }

  # Billing mode mapping
  billing_mode_map = {
    "PROVISIONED"      = "PROVISIONED"
    "PAY_PER_REQUEST"  = "PAY_PER_REQUEST"
    1                   = "PROVISIONED"
    2                   = "PAY_PER_REQUEST"
  }

  ###########################################################################
  # Core table attributes                                                    #
  ###########################################################################
  attribute_definitions = [
    for a in try(local.spec.attribute_definitions, []) : {
      name = a.attribute_name
      type = local.attribute_type_map[a.attribute_type]
    }
  ]

  key_schema = [
    for k in try(local.spec.key_schema, []) : {
      attribute_name = k.attribute_name
      key_type       = local.key_type_map[k.key_type]
    }
  ]

  # Extract top-level HASH and (optional) RANGE key names for convenience
  table_hash_key  = [for k in local.key_schema : k if k.key_type == "HASH"][0].attribute_name
  table_range_key = try(
    [for k in local.key_schema : k if k.key_type == "RANGE"][0].attribute_name,
    null,
  )

  ###########################################################################
  # Billing / capacity helpers                                               #
  ###########################################################################
  billing_mode = local.billing_mode_map[local.spec.billing_mode]

  provisioned_throughput = (
    local.billing_mode == "PROVISIONED" ? {
      read_capacity  = local.spec.provisioned_throughput.read_capacity_units
      write_capacity = local.spec.provisioned_throughput.write_capacity_units
    } : null
  )

  ###########################################################################
  # Global & local secondary index helpers                                   #
  ###########################################################################
  global_secondary_indexes = [
    for g in try(local.spec.global_secondary_indexes, []) : {
      name      = g.index_name
      hash_key  = [for ks in g.key_schema : ks if local.key_type_map[ks.key_type] == "HASH"][0].attribute_name
      range_key = try(
        [for ks in g.key_schema : ks if local.key_type_map[ks.key_type] == "RANGE"][0].attribute_name,
        null,
      )

      projection_type    = local.projection_type_map[g.projection.projection_type]
      non_key_attributes = (
        local.projection_type_map[g.projection.projection_type] == "INCLUDE" ?
        g.projection.non_key_attributes : null
      )

      read_capacity  = (
        local.billing_mode == "PROVISIONED" ? g.provisioned_throughput.read_capacity_units : null
      )
      write_capacity = (
        local.billing_mode == "PROVISIONED" ? g.provisioned_throughput.write_capacity_units : null
      )
    }
  ]

  local_secondary_indexes = [
    for l in try(local.spec.local_secondary_indexes, []) : {
      name      = l.index_name
      range_key = [for ks in l.key_schema : ks if local.key_type_map[ks.key_type] == "RANGE"][0].attribute_name

      projection_type    = local.projection_type_map[l.projection.projection_type]
      non_key_attributes = (
        local.projection_type_map[l.projection.projection_type] == "INCLUDE" ?
        l.projection.non_key_attributes : null
      )
    }
  ]

  ###########################################################################
  # Optional feature blocks                                                  #
  ###########################################################################
  stream_specification = (
    try(local.spec.stream_specification.stream_enabled, false) ? {
      stream_enabled   = true
      stream_view_type = local.stream_view_type_map[local.spec.stream_specification.stream_view_type]
    } : null
  )

  ttl_specification = (
    try(local.spec.ttl_specification.ttl_enabled, false) ? {
      enabled        = true
      attribute_name = local.spec.ttl_specification.attribute_name
    } : null
  )

  sse_specification = (
    try(local.spec.sse_specification.enabled, false) ? {
      enabled           = true
      sse_type          = local.sse_type_map[local.spec.sse_specification.sse_type]
      kms_master_key_id = (
        contains(["KMS", 2], local.spec.sse_specification.sse_type) ?
        local.spec.sse_specification.kms_master_key_id : null
      )
    } : null
  )

  ###########################################################################
  # Tags                                                                     #
  ###########################################################################
  tags = merge({}, try(local.spec.tags, {}))
}
