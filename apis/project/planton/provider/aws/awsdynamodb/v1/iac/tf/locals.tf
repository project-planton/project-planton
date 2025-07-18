############################
# Local helpers & mappings #
############################

locals {
  ###############################################################
  # Enum → provider value maps (keep in sync with proto enums)  #
  ###############################################################
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

  ####################################
  # Top-level table configuration    #
  ####################################

  # Primary (partition/sort) keys --------------------------------
  partition_key = [for ks in var.key_schema : ks.attribute_name if ks.key_type == "HASH"] [0]
  sort_key      = try([for ks in var.key_schema : ks.attribute_name if ks.key_type == "RANGE"] [0], null)

  # Attribute definitions ----------------------------------------
  attribute_definitions = [
    for a in var.attribute_definitions : {
      name = a.attribute_name
      type = local.attribute_type_map[a.attribute_type]
    }
  ]

  # Billing / capacity -------------------------------------------
  billing_mode = var.billing_mode == "PROVISIONED" ? "PROVISIONED" : "PAY_PER_REQUEST"

  provisioned_throughput = var.billing_mode == "PROVISIONED" ? {
    read_capacity  = var.provisioned_throughput.read_capacity_units
    write_capacity = var.provisioned_throughput.write_capacity_units
  } : null

  ####################################
  # Index helpers (GSI / LSI)        #
  ####################################

  global_secondary_indexes = [
    for g in var.global_secondary_indexes : {
      name            = g.index_name
      hash_key        = [for ks in g.key_schema : ks.attribute_name if ks.key_type == "HASH"] [0]
      range_key       = try([for ks in g.key_schema : ks.attribute_name if ks.key_type == "RANGE"] [0], null)
      projection_type = local.projection_type_map[g.projection.projection_type]
      non_key_attrs   = g.projection.projection_type == "INCLUDE" ? g.projection.non_key_attributes : null
      read_capacity   = var.billing_mode == "PROVISIONED" ? try(g.provisioned_throughput.read_capacity_units, null) : null
      write_capacity  = var.billing_mode == "PROVISIONED" ? try(g.provisioned_throughput.write_capacity_units, null) : null
    }
  ]

  local_secondary_indexes = [
    for l in var.local_secondary_indexes : {
      name            = l.index_name
      range_key       = [for ks in l.key_schema : ks.attribute_name if ks.key_type == "RANGE"] [0]
      projection_type = local.projection_type_map[l.projection.projection_type]
      non_key_attrs   = l.projection.projection_type == "INCLUDE" ? l.projection.non_key_attributes : null
    }
  ]

  ####################################
  # Streams, TTL, SSE, PITR helpers  #
  ####################################

  # Streams -------------------------------------------------------
  stream_enabled   = try(var.stream_specification.stream_enabled, false)
  stream_view_type = local.stream_enabled ? var.stream_specification.stream_view_type : null

  # TTL -----------------------------------------------------------
  ttl_enabled        = try(var.ttl_specification.ttl_enabled, false)
  ttl_attribute_name = local.ttl_enabled ? var.ttl_specification.attribute_name : null

  # Server-side encryption ---------------------------------------
  sse_enabled       = try(var.sse_specification.enabled, false)
  sse_type          = local.sse_enabled ? var.sse_specification.sse_type : null
  kms_master_key_id = (local.sse_enabled && local.sse_type == "KMS") ? var.sse_specification.kms_master_key_id : null

  ####################################
  # Tag helpers                   #
  ####################################
  tags = merge({
    "ManagedBy" = "terraform"
  }, var.tags)
}
