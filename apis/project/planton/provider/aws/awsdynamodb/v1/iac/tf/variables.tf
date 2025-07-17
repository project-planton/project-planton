# --------------------------------------------------------------------------------
# Variables matching project.planton.provider.aws.awsdynamodb.v1.AwsDynamodbSpec
# --------------------------------------------------------------------------------

###############################################################################
# Table basics
###############################################################################

variable "table_name" {
  type        = string
  description = "Name of the DynamoDB table."

  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "table_name must contain 3–255 characters."
  }
}

###############################################################################
# Attribute / key-schema definitions
###############################################################################

variable "attribute_definitions" {
  description = "Definitions for every attribute referenced by the table or any index."
  type = list(object({
    attribute_name = string
    attribute_type = string   # Allowed: STRING|NUMBER|BINARY or S|N|B
  }))

  validation {
    condition     = length(var.attribute_definitions) >= 1
    error_message = "At least one attribute definition is required."
  }

  validation {
    condition = alltrue([
      for a in var.attribute_definitions :
      length(a.attribute_name) >= 1 && length(a.attribute_name) <= 255 &&
      contains(["STRING", "NUMBER", "BINARY", "S", "N", "B"], upper(a.attribute_type))
    ])
    error_message = "Each attribute definition must use an attribute_name 1-255 chars long and attribute_type of STRING, NUMBER, BINARY (or the short forms S, N, B)."
  }
}

variable "key_schema" {
  description = "Primary key schema — partition key and optional sort key."
  type = list(object({
    attribute_name = string
    key_type       = string   # Allowed: HASH | RANGE
  }))

  validation {
    condition     = length(var.key_schema) >= 1 && length(var.key_schema) <= 2
    error_message = "key_schema must contain 1 or 2 elements (partition key and optional sort key)."
  }

  validation {
    condition = alltrue([
      for ks in var.key_schema :
      length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
      contains(["HASH", "RANGE"], upper(ks.key_type))
    ])
    error_message = "Each key_schema element must have attribute_name 1-255 chars long and key_type of HASH or RANGE."
  }
}

###############################################################################
# Capacity / billing
###############################################################################

variable "billing_mode" {
  type        = string
  description = "How the table is billed. Allowed: PROVISIONED or PAY_PER_REQUEST."
  default     = "PAY_PER_REQUEST"

  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], upper(var.billing_mode))
    error_message = "billing_mode must be PROVISIONED or PAY_PER_REQUEST."
  }

  # Cross-field rule: provisioned_throughput must be present when PROVISIONED and
  # must be absent everywhere when PAY_PER_REQUEST.
  validation {
    condition = upper(var.billing_mode) == "PROVISIONED" ? (
      var.provisioned_throughput != null &&
      var.provisioned_throughput.read_capacity_units  > 0 &&
      var.provisioned_throughput.write_capacity_units > 0 &&
      alltrue([
        for g in var.global_secondary_indexes :
        g.provisioned_throughput != null &&
        g.provisioned_throughput.read_capacity_units  > 0 &&
        g.provisioned_throughput.write_capacity_units > 0
      ])
    ) : (
      var.provisioned_throughput == null &&
      alltrue([
        for g in var.global_secondary_indexes : g.provisioned_throughput == null
      ])
    )
    error_message = "When billing_mode is PROVISIONED, provisioned_throughput must be configured for the table and every GSI; when PAY_PER_REQUEST it must be unset everywhere."
  }
}

variable "provisioned_throughput" {
  description = "Provisioned capacity settings (required when billing_mode == PROVISIONED)."
  type = object({
    read_capacity_units  = number
    write_capacity_units = number
  })
  default = null

  validation {
    condition = var.provisioned_throughput == null || (
      var.provisioned_throughput.read_capacity_units  > 0 &&
      var.provisioned_throughput.write_capacity_units > 0
    )
    error_message = "read_capacity_units and write_capacity_units must both be > 0."
  }
}

###############################################################################
# Global secondary indexes (GSIs)
###############################################################################

variable "global_secondary_indexes" {
  description = "Definitions for global secondary indexes (GSIs)."
  type = list(object({
    index_name = string
    key_schema = list(object({
      attribute_name = string
      key_type       = string   # HASH | RANGE
    }))
    projection = object({
      projection_type    = string              # ALL | KEYS_ONLY | INCLUDE
      non_key_attributes = optional(list(string))
    })
    provisioned_throughput = optional(object({
      read_capacity_units  = number
      write_capacity_units = number
    }))
  }))
  default = []

  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes :
      # index_name length
      length(g.index_name) >= 3 && length(g.index_name) <= 255 &&

      # key_schema size & contents
      length(g.key_schema) >= 1 && length(g.key_schema) <= 2 &&
      alltrue([
        for ks in g.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], upper(ks.key_type))
      ]) &&

      # projection rules
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], upper(g.projection.projection_type)) &&
      (
        upper(g.projection.projection_type) == "INCLUDE" ? (
          length(coalesce(g.projection.non_key_attributes, [])) > 0
        ) : (
          length(coalesce(g.projection.non_key_attributes, [])) == 0
        )
      ) &&
      length(coalesce(g.projection.non_key_attributes, [])) <= 20 &&
      alltrue([
        for n in coalesce(g.projection.non_key_attributes, []) :
        length(n) >= 1 && length(n) <= 255
      ]) &&

      # per-GSI throughput rules (validated again together with billing_mode)
      (g.provisioned_throughput == null || (
        g.provisioned_throughput.read_capacity_units  > 0 &&
        g.provisioned_throughput.write_capacity_units > 0
      ))
    ])
    error_message = "Each global_secondary_indexes block is invalid — check names, key_schema, projection and throughput rules."
  }
}

###############################################################################
# Local secondary indexes (LSIs)
###############################################################################

variable "local_secondary_indexes" {
  description = "Definitions for local secondary indexes (LSIs)."
  type = list(object({
    index_name = string
    key_schema = list(object({
      attribute_name = string
      key_type       = string
    }))
    projection = object({
      projection_type    = string
      non_key_attributes = optional(list(string))
    })
  }))
  default = []

  validation {
    condition = alltrue([
      for l in var.local_secondary_indexes :
      length(l.index_name) >= 3 && length(l.index_name) <= 255 &&
      # key_schema must be exactly 2 elements (HASH shared with table, plus RANGE)
      length(l.key_schema) == 2 &&
      alltrue([
        for ks in l.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], upper(ks.key_type))
      ]) &&
      # projection rules identical to GSI
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], upper(l.projection.projection_type)) &&
      (
        upper(l.projection.projection_type) == "INCLUDE" ? (
          length(coalesce(l.projection.non_key_attributes, [])) > 0
        ) : (
          length(coalesce(l.projection.non_key_attributes, [])) == 0
        )
      ) &&
      length(coalesce(l.projection.non_key_attributes, [])) <= 20 &&
      alltrue([
        for n in coalesce(l.projection.non_key_attributes, []) :
        length(n) >= 1 && length(n) <= 255
      ])
    ])
    error_message = "Each local_secondary_indexes block is invalid — check names, key_schema and projection rules."
  }
}

###############################################################################
# Streams
###############################################################################

variable "stream_specification" {
  description = "DynamoDB Streams configuration."
  type = object({
    stream_enabled   = bool
    stream_view_type = string   # NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | KEYS_ONLY
  })
  default = null

  validation {
    condition = var.stream_specification == null ? true : (
      !var.stream_specification.stream_enabled || length(var.stream_specification.stream_view_type) > 0
    )
    error_message = "stream_view_type must be specified when streams are enabled."
  }

  validation {
    condition = var.stream_specification == null ? true : (
      var.stream_specification.stream_view_type == "" ||
      contains(["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "KEYS_ONLY"], upper(var.stream_specification.stream_view_type))
    )
    error_message = "stream_view_type must be NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES or KEYS_ONLY."
  }
}

###############################################################################
# Time-to-live (TTL)
###############################################################################

variable "ttl_specification" {
  description = "Time-to-live (TTL) configuration."
  type = object({
    ttl_enabled    = bool
    attribute_name = string
  })
  default = null

  validation {
    condition = var.ttl_specification == null ? true : (
      !var.ttl_specification.ttl_enabled || length(var.ttl_specification.attribute_name) > 0
    )
    error_message = "attribute_name must be provided when TTL is enabled."
  }

  validation {
    condition = var.ttl_specification == null ? true : (
      var.ttl_specification.attribute_name == "" || (
        length(var.ttl_specification.attribute_name) >= 1 &&
        length(var.ttl_specification.attribute_name) <= 255
      )
    )
    error_message = "attribute_name must be between 1 and 255 characters."
  }
}

###############################################################################
# Server-side encryption (SSE)
###############################################################################

variable "sse_specification" {
  description = "Server-side encryption configuration."
  type = object({
    enabled            = bool
    sse_type           = string   # AES256 | KMS
    kms_master_key_id  = string
  })
  default = null

  validation {
    condition = var.sse_specification == null ? true : (
      (
        # Disabled: all other fields empty
        !var.sse_specification.enabled &&
        var.sse_specification.sse_type == "" &&
        var.sse_specification.kms_master_key_id == ""
      ) || (
        # Enabled
        var.sse_specification.enabled &&
        contains(["AES256", "KMS"], upper(var.sse_specification.sse_type)) &&
        (
          upper(var.sse_specification.sse_type) == "KMS" ?
          length(var.sse_specification.kms_master_key_id) > 0 :
          var.sse_specification.kms_master_key_id == ""
        )
      )
    )
    error_message = "When SSE is enabled, sse_type must be AES256 or KMS and kms_master_key_id is required only for KMS; when disabled they must be unset."
  }

  validation {
    condition = var.sse_specification == null ? true : (
      var.sse_specification.kms_master_key_id == "" || (
        length(var.sse_specification.kms_master_key_id) >= 1 &&
        length(var.sse_specification.kms_master_key_id) <= 2048
      )
    )
    error_message = "kms_master_key_id must be 1-2048 characters when provided."
  }
}

###############################################################################
# Miscellaneous
###############################################################################

variable "point_in_time_recovery_enabled" {
  type        = bool
  description = "Enable point-in-time recovery (continuous backups)."
  default     = false
}

variable "tags" {
  type        = map(string)
  description = "Key/value tags applied to the table."
  default     = {}

  validation {
    condition = alltrue([
      for k, v in var.tags : length(k) > 0 && length(v) > 0
    ])
    error_message = "Tag keys and values must be non-empty strings."
  }
}
