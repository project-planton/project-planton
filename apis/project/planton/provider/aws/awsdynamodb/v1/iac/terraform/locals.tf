###############################################################################
#  Derived helpers for the AwsDynamodb module                                 #
#                                                                              #
#  This file converts the incoming protobuf-shaped "spec" variable into        #
#  the exact structures and literal strings expected by the Terraform          #
#  aws_dynamodb_table resource and the surrounding plumbing.                   #
###############################################################################

locals {
  ###########################################################################
  #  Enum value look-up tables                                              #
  ###########################################################################
  billing_mode_map = {
    1 = "PROVISIONED"
    2 = "PAY_PER_REQUEST"
  }

  attribute_type_map = {
    1 = "S" # STRING
    2 = "N" # NUMBER
    3 = "B" # BINARY
  }

  key_type_map = {
    1 = "HASH"
    2 = "RANGE"
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

  ###########################################################################
  #  Convenience flags & simple derivations                                 #
  ###########################################################################
  billing_mode   = lookup(local.billing_mode_map, var.spec.billing_mode, null)
  is_provisioned = local.billing_mode == "PROVISIONED"

  stream_enabled = try(var.spec.stream_specification.stream_enabled, false)
  ttl_enabled    = try(var.spec.ttl_specification.ttl_enabled, false)
  sse_enabled    = try(var.spec.sse_specification.enabled, false)

  ###########################################################################
  #  Simple scalar conversions                                              #
  ###########################################################################
  stream_view_type = local.stream_enabled ?
    lookup(local.stream_view_type_map, var.spec.stream_specification.stream_view_type, null)
    : null

  ttl_attribute_name = local.ttl_enabled ? var.spec.ttl_specification.attribute_name : null

  sse_type          = local.sse_enabled ? lookup(local.sse_type_map, var.spec.sse_specification.sse_type, null) : null
  kms_master_key_id = (local.sse_enabled && var.spec.sse_specification.sse_type == 2) ? var.spec.sse_specification.kms_master_key_id : null

  point_in_time_recovery_enabled = var.spec.point_in_time_recovery_enabled

  ###########################################################################
  #  Attribute definitions & key schema                                     #
  ###########################################################################
  attribute_definitions = [
    for a in var.spec.attribute_definitions : {
      name = a.attribute_name
      type = lookup(local.attribute_type_map, a.attribute_type, null)
    }
  ]

  key_schema = [
    for k in var.spec.key_schema : {
      attribute_name = k.attribute_name
      key_type       = lookup(local.key_type_map, k.key_type, null)
    }
  ]

  ###########################################################################
  #  Provisioned capacity (table level)                                     #
  ###########################################################################
  provisioned_throughput = local.is_provisioned ? {
    read_capacity  = var.spec.provisioned_throughput.read_capacity_units
    write_capacity = var.spec.provisioned_throughput.write_capacity_units
  } : null

  ###########################################################################
  #  Global secondary indexes                                               #
  ###########################################################################
  global_secondary_indexes = [
    for g in var.spec.global_secondary_indexes : {
      name     = g.index_name
      hash_key = one([for ks in g.key_schema : ks.attribute_name if ks.key_type == 1])
      range_key = try(
        one([for ks in g.key_schema : ks.attribute_name if ks.key_type == 2]),
        null
      )

      projection_type   = lookup(local.projection_type_map, g.projection.projection_type, null)
      non_key_attributes = g.projection.projection_type == 3 ? g.projection.non_key_attributes : null

      read_capacity  = (local.is_provisioned && g.provisioned_throughput != null) ? g.provisioned_throughput.read_capacity_units  : null
      write_capacity = (local.is_provisioned && g.provisioned_throughput != null) ? g.provisioned_throughput.write_capacity_units : null
    }
  ]

  ###########################################################################
  #  Local secondary indexes                                                #
  ###########################################################################
  local_secondary_indexes = [
    for l in var.spec.local_secondary_indexes : {
      name       = l.index_name
      range_key  = one([for ks in l.key_schema : ks.attribute_name if ks.key_type == 2])

      projection_type   = lookup(local.projection_type_map, l.projection.projection_type, null)
      non_key_attributes = l.projection.projection_type == 3 ? l.projection.non_key_attributes : null
    }
  ]

  ###########################################################################
  #  Consolidated tags                                                      #
  ###########################################################################
  tags = merge(
    try(var.default_tags, {}),  # module / account defaults
    try(var.spec.tags, {}),     # user-supplied in the spec
    {
      "Name" = var.spec.table_name
    }
  )
}
