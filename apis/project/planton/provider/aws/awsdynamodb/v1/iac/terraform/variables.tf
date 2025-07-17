############################################
# Amazon DynamoDB – module input variables  #
############################################
#
# The variables declared in this file mirror the fields exposed by the
# AwsDynamodbSpec protobuf message.  Wherever practical, protobuf
# constraints (min/max length, enumerations, cross-field requirements,
# …) are implemented through native Terraform type definitions and the
# `validation` blocks that were introduced in Terraform 0.13.
#
# NOTE:  Object attributes are written in lower-snake-case (matching
# Terraform conventions) whereas the protobuf message uses lower-camel.
#
#-------------------------------------------------------------------------------

############################
# Basic table information  #
############################

variable "table_name" {
  description = "Name of the DynamoDB table (3-255 characters).  A runtime-generated suffix may be appended by the stack implementation."
  type        = string

  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "`table_name` must contain between 3 and 255 characters."
  }
}

###################################
# Attribute & key-schema settings #
###################################

variable "attribute_definitions" {
  description = <<-EOT
    Definitions for every attribute referenced by the table or any index.
    Each item must contain:
      * attribute_name – 1-255 characters
      * attribute_type – one of "S", "N", "B" (case-insensitive)
  EOT

  type = list(object({
    attribute_name = string
    attribute_type = string # "S", "N" or "B"
  }))

  validation {
    condition     = length(var.attribute_definitions) >= 1
    error_message = "At least one `attribute_definition` is required."
  }

  validation {
    condition = alltrue([
      for a in var.attribute_definitions :
      length(a.attribute_name) >= 1 &&
      length(a.attribute_name) <= 255 &&
      contains(["S", "N", "B"], upper(a.attribute_type))
    ])
    error_message = "Every attribute definition must have a 1-255 character `attribute_name` and an `attribute_type` of `S`, `N` or `B`."
  }
}

variable "key_schema" {
  description = <<-EOT
    Primary key schema for the table.  Must contain 1 or 2 elements:
      * Exactly one element with key_type = "HASH" (the partition key)
      * (Optional) one element with key_type = "RANGE" (the sort key)
  EOT

  type = list(object({
    attribute_name = string
    key_type       = string # "HASH" or "RANGE"
  }))

  validation {
    condition     = length(var.key_schema) >= 1 && length(var.key_schema) <= 2
    error_message = "`key_schema` must contain 1 or 2 elements."
  }

  validation {
    condition = (
      length([for ks in var.key_schema : ks if upper(ks.key_type) == "HASH"]) == 1 &&
      length([for ks in var.key_schema : ks if upper(ks.key_type) == "RANGE"]) <= 1 &&
      alltrue([
        for ks in var.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], upper(ks.key_type))
      ])
    )
    error_message = "`key_schema` must include exactly one HASH key and at most one RANGE key, with valid attribute names (1-255 chars)."
  }
}

######################
# Billing & capacity #
######################

variable "billing_mode" {
  description = "How the table is billed.  Valid values: \"PROVISIONED\" or \"PAY_PER_REQUEST\"."
  type        = string

  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], upper(var.billing_mode))
    error_message = "`billing_mode` must be either `PROVISIONED` or `PAY_PER_REQUEST`."
  }
}

variable "provisioned_throughput" {
  description = <<-EOT
    Provisioned capacity settings – required when `billing_mode` is PROVISIONED and **must** be omitted when billing is on-demand.

    Example:
      provisioned_throughput = {
        read_capacity_units  = 5
        write_capacity_units = 2
      }
  EOT

  type    = any   # kept flexible so that `null` is allowed when PAY_PER_REQUEST
  default = null

  validation {
    condition = (
      upper(var.billing_mode) == "PROVISIONED" ?
      (
        var.provisioned_throughput != null &&
        try(var.provisioned_throughput.read_capacity_units, 0) > 0 &&
        try(var.provisioned_throughput.write_capacity_units, 0) > 0
      ) : var.provisioned_throughput == null
    )

    error_message = "When `billing_mode` is PROVISIONED, a non-null `provisioned_throughput` object with positive read & write capacity units is required; otherwise it must be null."
  }
}

##################################
# Secondary index configurations #
##################################

# Helper types used for both GSI and LSI
locals {
  projection_type_allowed = ["ALL", "KEYS_ONLY", "INCLUDE"]
  stream_view_type_allowed = ["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "STREAM_KEYS_ONLY"]
}

variable "global_secondary_indexes" {
  description = "Definitions for global secondary indexes (GSIs).  An empty list (default) means no GSIs are created."

  type = list(object({
    index_name = string

    key_schema = list(object({
      attribute_name = string
      key_type       = string # "HASH" or "RANGE"
    }))

    projection = object({
      projection_type   = string           # \"ALL\", \"KEYS_ONLY\" or \"INCLUDE\"
      non_key_attributes = list(string)    # <= 20 items, each 1-255 chars
    })

    # Optional per-GSI provisioned throughput.  Must follow the same rule
    # as the table: required when billing_mode == PROVISIONED, forbidden otherwise.
    provisioned_throughput = any
  }))

  default = []

  # Basic structural validation (name lengths, allowed enums, etc.).
  # Full cross-field validation is left to the provider/API.
  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes :
      length(g.index_name) >= 3 && length(g.index_name) <= 255 &&
      # key_schema 1-2 items & valid key_type
      length(g.key_schema) >= 1 && length(g.key_schema) <= 2 &&
      alltrue([
        for ks in g.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], upper(ks.key_type))
      ]) &&
      # projection rules
      contains(local.projection_type_allowed, upper(g.projection.projection_type)) &&
      (
        upper(g.projection.projection_type) == "INCLUDE" ?
        length(g.projection.non_key_attributes) > 0 && length(g.projection.non_key_attributes) <= 20 &&
        alltrue([for a in g.projection.non_key_attributes : length(a) >=1 && length(a) <= 255])
        : length(g.projection.non_key_attributes) == 0
      )
    ])

    error_message = "Every GSI must have a valid name (3-255 chars), key_schema (1-2 entries of HASH/RANGE), and a valid projection (non_key_attributes only when projection_type == INCLUDE)."
  }

  # Capacity rule mirroring the CEL constraint in the protobuf spec.
  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes :
      (
        upper(var.billing_mode) == "PROVISIONED" ?
        g.provisioned_throughput != null &&
        try(g.provisioned_throughput.read_capacity_units, 0) > 0 &&
        try(g.provisioned_throughput.write_capacity_units, 0) > 0
        : g.provisioned_throughput == null
      )
    ])

    error_message = "When `billing_mode` is PROVISIONED, each GSI requires a non-null `provisioned_throughput` object with positive RCU/WCU; otherwise it must be null."
  }
}

variable "local_secondary_indexes" {
  description = "Definitions for local secondary indexes (LSIs).  An empty list (default) means no LSIs are created."

  type = list(object({
    index_name = string

    # Exactly two elements – table HASH key + alternate RANGE key
    key_schema = list(object({
      attribute_name = string
      key_type       = string # HASH or RANGE
    }))

    projection = object({
      projection_type   = string
      non_key_attributes = list(string)
    })
  }))

  default = []

  validation {
    condition = alltrue([
      for l in var.local_secondary_indexes :
      length(l.index_name) >= 3 && length(l.index_name) <= 255 &&
      length(l.key_schema) == 2 &&
      # must contain exactly one HASH and one RANGE key
      length([for ks in l.key_schema : ks if upper(ks.key_type) == "HASH"]) == 1 &&
      length([for ks in l.key_schema : ks if upper(ks.key_type) == "RANGE"]) == 1 &&
      alltrue([
        for ks in l.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], upper(ks.key_type))
      ]) &&
      # projection rules
      contains(local.projection_type_allowed, upper(l.projection.projection_type)) &&
      (
        upper(l.projection.projection_type) == "INCLUDE" ?
        length(l.projection.non_key_attributes) > 0 && length(l.projection.non_key_attributes) <= 20 &&
        alltrue([for a in l.projection.non_key_attributes : length(a) >=1 && length(a) <= 255])
        : length(l.projection.non_key_attributes) == 0
      )
    ])

    error_message = "Each LSI must have a valid name (3-255 chars), exactly two key_schema elements (one HASH, one RANGE), and a valid projection configuration."
  }
}

#######################
# Streams & TTL setup #
#######################

variable "stream_specification" {
  description = <<-EOT
    DynamoDB Streams configuration.
    Example to enable NEW_AND_OLD_IMAGES:
      stream_specification = {
        stream_enabled   = true
        stream_view_type = "NEW_AND_OLD_IMAGES"
      }
  EOT

  type = object({
    stream_enabled   = bool
    stream_view_type = string # "NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "STREAM_KEYS_ONLY" – required when enabled
  })

  default = {
    stream_enabled   = false
    stream_view_type = "" # ignored when disabled
  }

  validation {
    condition = (
      var.stream_specification.stream_enabled == false ||
      contains(local.stream_view_type_allowed, upper(var.stream_specification.stream_view_type))
    )

    error_message = "When `stream_enabled` is true, `stream_view_type` must be one of NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES, STREAM_KEYS_ONLY."
  }

  validation {
    condition = (
      var.stream_specification.stream_enabled || var.stream_specification.stream_view_type == ""
    )
    error_message = "`stream_view_type` must be empty when streams are disabled."
  }
}

variable "ttl_specification" {
  description = <<-EOT
    Time-to-Live (TTL) configuration for automatic item expiry.
    Example (enable):
      ttl_specification = {
        ttl_enabled    = true
        attribute_name = "expires_at_epoch"
      }
  EOT

  type = object({
    ttl_enabled    = bool
    attribute_name = string
  })

  default = {
    ttl_enabled    = false
    attribute_name = ""
  }

  validation {
    condition = (
      var.ttl_specification.ttl_enabled == false || length(var.ttl_specification.attribute_name) > 0
    )
    error_message = "`attribute_name` must be provided (non-empty) when `ttl_enabled` is true."
  }

  validation {
    condition = length(var.ttl_specification.attribute_name) <= 255
    error_message = "TTL `attribute_name` must be 255 characters or fewer."
  }
}

#############################
# Server-side encryption (SSE)
#############################

variable "sse_specification" {
  description = <<-EOT
    Server-side encryption configuration.
    Example using AWS-managed CMK:
      sse_specification = {
        enabled           = true
        sse_type          = "KMS"
        kms_master_key_id = "arn:aws:kms:us-east-1:111122223333:key/abcd-1234"
      }
  EOT

  type = object({
    enabled           = bool
    sse_type          = string  # "AES256" or "KMS" – required when enabled
    kms_master_key_id = string  # required only when sse_type == "KMS"
  })

  default = {
    enabled           = false
    sse_type          = ""
    kms_master_key_id = ""
  }

  validation {
    condition = (
      var.sse_specification.enabled == false && var.sse_specification.sse_type == "" && var.sse_specification.kms_master_key_id == "" ||
      (
        var.sse_specification.enabled == true &&
        contains(["AES256", "KMS"], upper(var.sse_specification.sse_type)) &&
        (
          upper(var.sse_specification.sse_type) == "KMS" ? length(var.sse_specification.kms_master_key_id) > 0 && length(var.sse_specification.kms_master_key_id) <= 2048 : var.sse_specification.kms_master_key_id == ""
        )
      )
    )

    error_message = "When SSE is enabled, `sse_type` must be AES256 or KMS and `kms_master_key_id` is required only for KMS; when disabled the related fields must be empty."
  }
}

###################################
# Miscellaneous / convenience vars #
###################################

variable "point_in_time_recovery_enabled" {
  description = "Enable point-in-time recovery (continuous backups).  Defaults to `false`."
  type        = bool
  default     = false
}

variable "tags" {
  description = "Key/value tags to apply to the DynamoDB table.  Both keys and values must be non-empty strings."
  type        = map(string)
  default     = {}

  validation {
    condition = alltrue([
      for key, value in var.tags :
      length(key) > 0 && length(value) > 0
    ])
    error_message = "Tag keys and values must be non-empty strings."
  }
}
