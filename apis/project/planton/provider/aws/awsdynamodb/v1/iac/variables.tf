# -----------------------------------------------------------------------------
# Input variables for the aws_dynamodb module
# -----------------------------------------------------------------------------

###############################################################################
# Core table arguments
###############################################################################

variable "table_name" {
  type        = string
  description = "Name of the DynamoDB table (3–255 characters)."

  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "table_name must be between 3 and 255 characters."
  }
}

variable "attribute_definitions" {
  description = "Definitions for every attribute referenced by the table or any index."
  type = list(object({
    attribute_name = string
    attribute_type = string # STRING | NUMBER | BINARY
  }))

  validation {
    condition = length(var.attribute_definitions) > 0 &&
      alltrue([
        for a in var.attribute_definitions :
        length(a.attribute_name) >= 1 &&
        length(a.attribute_name) <= 255 &&
        contains(["STRING", "NUMBER", "BINARY"], upper(a.attribute_type))
      ])
    error_message = "Each attribute_definition must have attribute_name 1-255 chars and attribute_type one of STRING, NUMBER or BINARY."
  }
}

variable "key_schema" {
  description = "Primary key schema for the table (partition and optional sort key)."
  type = list(object({
    attribute_name = string
    key_type       = string # HASH | RANGE
  }))

  validation {
    condition = length(var.key_schema) >= 1 && length(var.key_schema) <= 2 &&
      alltrue([
        for k in var.key_schema :
        length(k.attribute_name) >= 1 &&
        length(k.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], upper(k.key_type))
      ])
    error_message = "Provide 1–2 key_schema elements with key_type HASH or RANGE and attribute names 1-255 chars."
  }
}

variable "billing_mode" {
  description = "How the table is billed. Valid values: PROVISIONED or PAY_PER_REQUEST."
  type        = string
  default     = "PAY_PER_REQUEST"

  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], upper(var.billing_mode))
    error_message = "billing_mode must be either PROVISIONED or PAY_PER_REQUEST."
  }
}

variable "provisioned_throughput" {
  description = "Provisioned capacity settings (required when billing_mode is PROVISIONED)."
  type = object({
    read_capacity_units  = number
    write_capacity_units = number
  })
  default = null

  validation {
    condition     = var.provisioned_throughput == null || (var.provisioned_throughput.read_capacity_units > 0 && var.provisioned_throughput.write_capacity_units > 0)
    error_message = "When provided, read_capacity_units and write_capacity_units must be greater than zero."
  }
}

###############################################################################
# Secondary indexes
###############################################################################

variable "global_secondary_indexes" {
  description = "Definitions for global secondary indexes (GSIs)."
  default     = []
  type = list(object({
    index_name = string

    key_schema = list(object({
      attribute_name = string
      key_type       = string # HASH | RANGE
    }))

    projection = object({
      projection_type    = string # ALL | KEYS_ONLY | INCLUDE
      non_key_attributes = list(string)
    })

    # Optional when billing_mode == PAY_PER_REQUEST
    provisioned_throughput = object({
      read_capacity_units  = number
      write_capacity_units = number
    })
  }))

  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes :
      length(g.index_name) >= 3 && length(g.index_name) <= 255 &&
      length(g.key_schema) >= 1 && length(g.key_schema) <= 2 &&
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], upper(g.projection.projection_type)) &&
      (
        upper(g.projection.projection_type) == "INCLUDE" ? length(g.projection.non_key_attributes) > 0 : length(g.projection.non_key_attributes) == 0
      )
    ])
    error_message = "Each GSI must meet name length requirements, have 1-2 key_schema elements, a valid projection_type, and non_key_attributes only when projection_type is INCLUDE."
  }
}

variable "local_secondary_indexes" {
  description = "Definitions for local secondary indexes (LSIs)."
  default     = []
  type = list(object({
    index_name = string

    key_schema = list(object({
      attribute_name = string
      key_type       = string # HASH | RANGE
    }))

    projection = object({
      projection_type    = string # ALL | KEYS_ONLY | INCLUDE
      non_key_attributes = list(string)
    })
  }))

  validation {
    condition = alltrue([
      for l in var.local_secondary_indexes :
      length(l.index_name) >= 3 && length(l.index_name) <= 255 &&
      length(l.key_schema) == 2 &&
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], upper(l.projection.projection_type)) &&
      (
        upper(l.projection.projection_type) == "INCLUDE" ? length(l.projection.non_key_attributes) > 0 : length(l.projection.non_key_attributes) == 0
      )
    ])
    error_message = "Each LSI must meet name length requirements, have exactly 2 key_schema elements, a valid projection_type, and non_key_attributes only when projection_type is INCLUDE."
  }
}

###############################################################################
# Table-level features
###############################################################################

variable "stream_specification" {
  description = "DynamoDB Streams configuration."
  type = object({
    stream_enabled   = bool
    stream_view_type = string # NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | STREAM_KEYS_ONLY
  })
  default = {
    stream_enabled   = false
    stream_view_type = ""
  }

  validation {
    condition     = !var.stream_specification.stream_enabled || contains(["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "STREAM_KEYS_ONLY", "KEYS_ONLY"], upper(var.stream_specification.stream_view_type))
    error_message = "When stream_enabled is true, stream_view_type must be set to a valid value."
  }
}

variable "ttl_specification" {
  description = "Time-to-live (TTL) configuration."
  type = object({
    ttl_enabled    = bool
    attribute_name = string
  })
  default = {
    ttl_enabled    = false
    attribute_name = ""
  }

  validation {
    condition = (!var.ttl_specification.ttl_enabled && var.ttl_specification.attribute_name == "") ||
      (var.ttl_specification.ttl_enabled && length(var.ttl_specification.attribute_name) > 0 && length(var.ttl_specification.attribute_name) <= 255)
    error_message = "When ttl_enabled is true, attribute_name must be provided (1-255 chars). When false it must be empty."
  }
}

variable "sse_specification" {
  description = "Server-side encryption (SSE) settings."
  type = object({
    enabled           = bool
    sse_type          = string # AES256 | KMS (required when enabled)
    kms_master_key_id = string # Required only when sse_type == KMS
  })
  default = {
    enabled           = false
    sse_type          = ""
    kms_master_key_id = ""
  }

  validation {
    condition = (!var.sse_specification.enabled && var.sse_specification.sse_type == "" && var.sse_specification.kms_master_key_id == "") ||
      (
        var.sse_specification.enabled &&
        contains(["AES256", "KMS"], upper(var.sse_specification.sse_type)) &&
        (
          (upper(var.sse_specification.sse_type) == "AES256" && var.sse_specification.kms_master_key_id == "") ||
          (upper(var.sse_specification.sse_type) == "KMS" && length(var.sse_specification.kms_master_key_id) > 0 && length(var.sse_specification.kms_master_key_id) <= 2048)
        )
      )
    error_message = "When SSE is enabled, sse_type must be AES256 or KMS. kms_master_key_id is required only when sse_type is KMS. When disabled all fields must be unset/empty."
  }
}

variable "point_in_time_recovery_enabled" {
  description = "Enables point-in-time recovery (continuous backups)."
  type        = bool
  default     = false
}

###############################################################################
# Miscellaneous
###############################################################################

variable "tags" {
  description = "Key/value tags applied to the table."
  type        = map(string)
  default     = {}

  validation {
    condition     = alltrue([for k, v in var.tags : length(k) > 0 && length(v) > 0])
    error_message = "Tag keys and values must be non-empty strings."
  }
}
