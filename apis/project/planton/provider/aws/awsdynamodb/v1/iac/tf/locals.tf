################################################################################
# Local helpers for the DynamoDB table module
# -------------------------------------------
#  * Enum-to-string look-ups so proto enum integers can be mapped to the values
#    expected by the Terraform AWS provider.
#  * Derivation of attribute/key/index structures in the shape required by the
#    aws_dynamodb_table resource’s argument schema.
#  * Unified tags map built by merging optional module-level default_tags with
#    the table-specific tags defined in the spec.
################################################################################

locals {
  ############################
  # Generic enum translations
  ############################
  billing_mode_map = {
    1 = "PROVISIONED"
    2 = "PAY_PER_REQUEST"
  }

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

  #######################
  # Consolidated tag set
  #######################
  #   – User tags (var.tags) win over default_tags when keys collide.
  #   – Either map may be absent in variables.tf, so use try() for safety.
  tags = merge(try(var.default_tags, {}), try(var.tags, {}))

  #################################
  # Table-level key configuration
  #################################
  hash_key = one([for ks in var.key_schema : ks.attribute_name if ks.key_type == 1])

  range_key = try(
    one([for ks in var.key_schema : ks.attribute_name if ks.key_type == 2]),
    null,
  )

  #######################################
  # Flatten & deduplicate attribute list
  #######################################
  #   DynamoDB requires every attribute referenced by the table or indexes to
  #   appear exactly once in the attribute_definition block.
  attribute_definitions = [
    for name, type in {
      for a in var.attribute_definitions :
      a.attribute_name => local.attribute_type_map[a.attribute_type]
    } : {
      name = name
      type = type
    }
  ]

  ###############################
  # Provisioned throughput (RCU/WCU)
  ###############################
  provisioned_throughput = (
    var.billing_mode == 1 ? {
      read_capacity  = var.provisioned_throughput.read_capacity_units
      write_capacity = var.provisioned_throughput.write_capacity_units
    } : null
  )

  ########################################
  # Global secondary index (GSI) helpers
  ########################################
  global_secondary_indexes = [
    for gsi in try(var.global_secondary_indexes, []) : {
      name        = gsi.index_name
      hash_key    = one([for ks in gsi.key_schema : ks.attribute_name if ks.key_type == 1])
      range_key   = try(one([for ks in gsi.key_schema : ks.attribute_name if ks.key_type == 2]), null)
      projection_type    = local.projection_type_map[gsi.projection.projection_type]
      non_key_attributes = (
        gsi.projection.projection_type == 3 ? gsi.projection.non_key_attributes : null
      )
      read_capacity  = (var.billing_mode == 1 ? gsi.provisioned_throughput.read_capacity_units  : null)
      write_capacity = (var.billing_mode == 1 ? gsi.provisioned_throughput.write_capacity_units : null)
    }
  ]

  ########################################
  # Local secondary index (LSI) helpers
  ########################################
  local_secondary_indexes = [
    for lsi in try(var.local_secondary_indexes, []) : {
      name      = lsi.index_name
      range_key = one([for ks in lsi.key_schema : ks.attribute_name if ks.key_type == 2])

      projection_type    = local.projection_type_map[lsi.projection.projection_type]
      non_key_attributes = (
        lsi.projection.projection_type == 3 ? lsi.projection.non_key_attributes : null
      )
    }
  ]

  ########################################
  # Stream / TTL / SSE derived settings
  ########################################
  stream_enabled   = try(var.stream_specification.stream_enabled, false)
  stream_view_type = (
    local.stream_enabled ?
    local.stream_view_type_map[var.stream_specification.stream_view_type] :
    null
  )

  ttl_specification = (
    try(var.ttl_specification.ttl_enabled, false) ? {
      enabled        = true
      attribute_name = var.ttl_specification.attribute_name
    } : null
  )

  sse_specification = (
    try(var.sse_specification.enabled, false) ? {
      enabled          = true
      sse_type         = local.sse_type_map[var.sse_specification.sse_type]
      kms_master_key_id = (
        var.sse_specification.sse_type == 2 ?
        var.sse_specification.kms_master_key_id :
        null
      )
    } : null
  )

  ########################################
  # Convenience look-ups for resource args
  ########################################
  billing_mode = local.billing_mode_map[var.billing_mode]
}
