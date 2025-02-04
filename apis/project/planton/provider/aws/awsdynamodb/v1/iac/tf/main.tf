########################################
# main.tf
########################################

# ----------------------------
# 1) "Safe" locals definitions
# ----------------------------
locals {
  # ------------------------------------------------------------------------
  # If var.spec.auto_scale is null, we create a "dummy" object so that we
  # never attempt .is_enabled on a null. Feel free to set default min_capacity,
  # etc. as you prefer. The key point: is_enabled must exist and be false
  # if no scaling is desired.
  # ------------------------------------------------------------------------
  safe_auto_scale = (
    var.spec.auto_scale == null
    ? {
    is_enabled      = false
    read_capacity   = {
      min_capacity         = 0
      max_capacity         = 0
      target_utilization   = 0
    }
    write_capacity  = {
      min_capacity         = 0
      max_capacity         = 0
      target_utilization   = 0
    }
  }
    : var.spec.auto_scale
  )

  # Auto-scaling is enabled only if:
  # 1) The safe_auto_scale.is_enabled is true
  # 2) billing_mode == "PROVISIONED"
  auto_scaling_is_enabled = local.safe_auto_scale.is_enabled == true&& var.spec.billing_mode == "PROVISIONED"

  # ------------------------------------------------------------------------
  # If var.spec.range_key is null, use a dummy object (with empty name/type),
  # so referencing .name or .type never fails.
  # ------------------------------------------------------------------------
  safe_range_key = (
    var.spec.range_key == null
    ? {
    name = ""
    type = ""
  }
    : var.spec.range_key
  )

  # We'll consider a range key provided only if name != ""
  range_key_is_provided = local.safe_range_key.name != ""

  # ------------------------------------------------------------------------
  # If var.spec.attributes is null, treat it like an empty list
  # ------------------------------------------------------------------------
  safe_attributes = (
    var.spec.attributes == null
    ? []
    : var.spec.attributes
  )

  # ------------------------------------------------------------------------
  # Build a combined attribute list:
  #   1) Always include the hash_key
  #   2) If range_key_is_provided, add it
  #   3) Append safe_attributes
  # ------------------------------------------------------------------------
  all_attributes = concat(
    [
      {
        name = var.spec.hash_key.name
        type = var.spec.hash_key.type
      }
    ],
    (
      local.range_key_is_provided
      ? [
      {
        name = local.safe_range_key.name
        type = local.safe_range_key.type
      }
    ]
      : []
    ),
    [
      for a in local.safe_attributes : {
      name = a.name
      type = a.type
    }
    ]
  )

  # ------------------------------------------------------------------------
  # If var.spec.replica_region_names is null, treat it like an empty list
  # ------------------------------------------------------------------------
  safe_replica_region_names = var.spec.replica_region_names == null ? [] : var.spec.replica_region_names

  # ------------------------------------------------------------------------
  # If var.spec.enable_streams is null, default to false
  # If replica list is non-empty, we must enable streams no matter what
  # ------------------------------------------------------------------------
  safe_enable_streams = coalesce(var.spec.enable_streams, false)
  stream_is_enabled   = length(local.safe_replica_region_names) > 0 || local.safe_enable_streams

  # ------------------------------------------------------------------------
  # stream_view_type is only used if streams are enabled; otherwise null
  # ------------------------------------------------------------------------
  stream_view_type = local.stream_is_enabled ? coalesce(var.spec.stream_view_type, "") : null

  # ------------------------------------------------------------------------
  # If var.spec.global_secondary_indexes is null, treat it like an empty list
  # ------------------------------------------------------------------------
  safe_global_secondary_indexes = var.spec.global_secondary_indexes == null ? [] : var.spec.global_secondary_indexes

  # ------------------------------------------------------------------------
  # If var.spec.local_secondary_indexes is null, treat it like an empty list
  # ------------------------------------------------------------------------
  safe_local_secondary_indexes = var.spec.local_secondary_indexes == null ? [] : var.spec.local_secondary_indexes

  # ------------------------------------------------------------------------
  # If var.spec.import_table is null, we'll skip it (with an empty list)
  # ------------------------------------------------------------------------
  safe_import_table = var.spec.import_table == null ? [] : [var.spec.import_table]
}

# -----------------------------
# 2) The aws_dynamodb_table Resource
# -----------------------------
resource "aws_dynamodb_table" "this" {
  # Basic configuration
  name         = var.spec.table_name
  billing_mode = var.spec.billing_mode

  # If PROVISIONED, set read/write from safe_auto_scale or fallback 5
  read_capacity = (
    var.spec.billing_mode == "PROVISIONED"
    ? coalesce(try(local.safe_auto_scale.read_capacity.min_capacity, null), 5)
    : null
  )

  write_capacity = (
    var.spec.billing_mode == "PROVISIONED"
    ? coalesce(try(local.safe_auto_scale.write_capacity.min_capacity, null), 5)
    : null
  )

  # Keys
  hash_key = var.spec.hash_key.name
  range_key = (
    local.range_key_is_provided
    ? local.safe_range_key.name
    : null
  )

  # Streams
  stream_enabled   = local.stream_is_enabled
  stream_view_type = local.stream_view_type

  table_class                 = "STANDARD"
  deletion_protection_enabled = false

  # Server-side encryption
  server_side_encryption {
    enabled     = try(var.spec.server_side_encryption.is_enabled, false)
    kms_key_arn = try(var.spec.server_side_encryption.kms_key_arn, null)
  }

  # PITR
  point_in_time_recovery {
    enabled = try(var.spec.point_in_time_recovery.is_enabled, false)
  }

  # TTL
  dynamic "ttl" {
    for_each = var.spec.ttl != null ? [var.spec.ttl] : []
    content {
      enabled        = try(ttl.value.is_enabled, false)
      attribute_name = try(ttl.value.attribute_name, "")
    }
  }

  # Final attribute list
  dynamic "attribute" {
    for_each = distinct(local.all_attributes)
    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  # Global Secondary Indexes
  dynamic "global_secondary_index" {
    for_each = local.safe_global_secondary_indexes
    content {
      name               = global_secondary_index.value.name
      hash_key           = global_secondary_index.value.hash_key
      range_key          = (
        global_secondary_index.value.range_key != null
        && global_secondary_index.value.range_key != ""
        ? global_secondary_index.value.range_key
        : null
      )
      projection_type    = global_secondary_index.value.projection_type
      non_key_attributes = global_secondary_index.value.non_key_attributes

      # For PROVISIONED, use GSI read/write
      read_capacity = (
        var.spec.billing_mode == "PROVISIONED"
        ? (
        global_secondary_index.value.read_capacity != 0
        ? global_secondary_index.value.read_capacity
        : coalesce(try(local.safe_auto_scale.read_capacity.min_capacity, null), 5)
      )
        : null
      )

      write_capacity = (
        var.spec.billing_mode == "PROVISIONED"
        ? (
        global_secondary_index.value.write_capacity != 0
        ? global_secondary_index.value.write_capacity
        : coalesce(try(local.safe_auto_scale.write_capacity.min_capacity, null), 5)
      )
        : null
      )
    }
  }

  # Local Secondary Indexes
  dynamic "local_secondary_index" {
    for_each = local.safe_local_secondary_indexes
    content {
      name               = local_secondary_index.value.name
      range_key          = local_secondary_index.value.range_key
      projection_type    = local_secondary_index.value.projection_type
      non_key_attributes = local_secondary_index.value.non_key_attributes
    }
  }

  # Replicas
  dynamic "replica" {
    for_each = local.safe_replica_region_names
    content {
      region_name            = replica.value
      kms_key_arn            = null
      point_in_time_recovery = false
      propagate_tags         = false
    }
  }

  # Import Table
  dynamic "import_table" {
    for_each = local.safe_import_table
    content {
      input_compression_type = try(import_table.value.input_compression_type, null)
      input_format           = try(import_table.value.input_format, null)

      dynamic "input_format_options" {
        for_each = import_table.value.input_format_options != null ? [import_table.value.input_format_options] : []
        content {
          csv {
            delimiter   = try(input_format_options.value.csv.delimiter, null)
            header_list = try(input_format_options.value.csv.headers, [])
          }
        }
      }

      dynamic "s3_bucket_source" {
        for_each = import_table.value.s3_bucket_source != null ? [import_table.value.s3_bucket_source] : []
        content {
          bucket       = try(s3_bucket_source.value.bucket, null)
          bucket_owner = try(s3_bucket_source.value.bucket_owner, null)
          key_prefix   = try(s3_bucket_source.value.key_prefix, null)
        }
      }
    }
  }

  # Example usage: define or reference your final labels from a separate locals file
  tags = local.final_labels
}
