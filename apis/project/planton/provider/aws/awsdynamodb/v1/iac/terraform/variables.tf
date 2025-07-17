###############################################################################
# Input variables for the aws_dynamodb module                                 #
#                                                                              #
# Every field exposed here mirrors (1-to-1) the AwsDynamodbSpec protobuf        #
# message that describes a DynamoDB table.  Where possible, the protobuf CEL    #
# rules have been re-implemented with Terraform "validation" blocks so users    #
# get early feedback when they pass invalid data.                               #
###############################################################################

############################
# Basic table information #
############################

variable "table_name" {
  description = "Name of the DynamoDB table (3–255 characters)."
  type        = string

  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "table_name must be between 3 and 255 characters long."
  }
}

#################################
# Attribute & key configuration #
#################################

variable "attribute_definitions" {
  description = "Definitions for every attribute referenced by the table or any index."
  type = list(object({
    attribute_name = string
    attribute_type = string # One of \"S\", \"N\", \"B\"
  }))

  validation {
    condition = length(var.attribute_definitions) >= 1 &&
      alltrue([
        for a in var.attribute_definitions :
        length(a.attribute_name) >= 1 &&
        length(a.attribute_name) <= 255 &&
        contains(["S", "N", "B"], a.attribute_type)
      ])
    error_message = "attribute_definitions must contain at least one element. Each element requires attribute_name (1–255 chars) and attribute_type of S, N or B."
  }
}

variable "key_schema" {
  description = "Primary key schema (partition key and optional sort key)."
  type = list(object({
    attribute_name = string
    key_type       = string # HASH or RANGE
  }))

  validation {
    condition = length(var.key_schema) >= 1 && length(var.key_schema) <= 2 &&
      alltrue([
        for k in var.key_schema :
        length(k.attribute_name) >= 1 &&
        length(k.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], k.key_type)
      ])
    error_message = "key_schema must have 1–2 elements. key_type must be HASH or RANGE and attribute_name 1–255 chars long."
  }
}

#############################
# Billing & capacity model  #
#############################

variable "billing_mode" {
  description = "How the table is billed: PROVISIONED or PAY_PER_REQUEST."
  type        = string

  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], var.billing_mode)
    error_message = "billing_mode must be either PROVISIONED or PAY_PER_REQUEST."
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
    condition = (
      var.billing_mode == "PROVISIONED" && var.provisioned_throughput != null &&
      var.provisioned_throughput.read_capacity_units  > 0 &&
      var.provisioned_throughput.write_capacity_units > 0
    ) || (
      var.billing_mode == "PAY_PER_REQUEST" && var.provisioned_throughput == null
    )
    error_message = "When billing_mode is PROVISIONED, provisioned_throughput must be provided with positive read/write capacity units; when PAY_PER_REQUEST it must be omitted (null)."
  }
}

#########################################
# Global & local secondary index inputs #
#########################################

variable "global_secondary_indexes" {
  description = <<EOT
Definitions for global secondary indexes (GSIs).
Each element supports:
  index_name             – string (3–255 chars)
  key_schema             – list of 1–2 key_schema objects
  projection             – object { projection_type, non_key_attributes }
  provisioned_throughput – optional object { read_capacity_units, write_capacity_units }
EOT
  type = list(object({
    index_name = string
    key_schema = list(object({
      attribute_name = string
      key_type       = string # HASH or RANGE
    }))
    projection = object({
      projection_type    = string # ALL, KEYS_ONLY, INCLUDE
      non_key_attributes = list(string)
    })
    # Optional – required only when table & GSI billing is PROVISIONED
    provisioned_throughput = optional(object({
      read_capacity_units  = number
      write_capacity_units = number
    }))
  }))
  default = []

  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes :
      length(g.index_name) >= 3 && length(g.index_name) <= 255 &&
      length(g.key_schema) >= 1 && length(g.key_schema) <= 2 &&
      alltrue([for ks in g.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], ks.key_type)
      ]) &&
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], g.projection.projection_type) &&
      (
        (g.projection.projection_type == "INCLUDE" && length(g.projection.non_key_attributes) > 0) ||
        (g.projection.projection_type != "INCLUDE" && length(g.projection.non_key_attributes) == 0)
      ) &&
      (
        var.billing_mode == "PROVISIONED" ? (
          (exists(g, "provisioned_throughput") &&
           g.provisioned_throughput.read_capacity_units  > 0 &&
           g.provisioned_throughput.write_capacity_units > 0)
        ) : (
          !exists(g, "provisioned_throughput")
        )
      )
    ])
    error_message = "Each GSI must follow the same validation rules as the table. When INCLUDE projection is used, non_key_attributes must be non-empty; otherwise it must be empty. Provisioned throughput is mandatory for GSIs only when table billing_mode is PROVISIONED."
  }
}

variable "local_secondary_indexes" {
  description = <<EOT
Definitions for local secondary indexes (LSIs).
Each element supports:
  index_name – string (3–255 chars)
  key_schema – list of exactly 2 key_schema objects (must share HASH key with the table)
  projection – object { projection_type, non_key_attributes }
EOT
  type = list(object({
    index_name = string
    key_schema = list(object({
      attribute_name = string
      key_type       = string # HASH or RANGE
    }))
    projection = object({
      projection_type    = string # ALL, KEYS_ONLY, INCLUDE
      non_key_attributes = list(string)
    })
  }))
  default = []

  validation {
    condition = alltrue([
      for l in var.local_secondary_indexes :
      length(l.index_name) >= 3 && length(l.index_name) <= 255 &&
      length(l.key_schema) == 2 &&
      alltrue([for ks in l.key_schema :
        length(ks.attribute_name) >= 1 && length(ks.attribute_name) <= 255 &&
        contains(["HASH", "RANGE"], ks.key_type)
      ]) &&
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], l.projection.projection_type) &&
      (
        (l.projection.projection_type == "INCLUDE" && length(l.projection.non_key_attributes) > 0) ||
        (l.projection.projection_type != "INCLUDE" && length(l.projection.non_key_attributes) == 0)
      )
    ])
    error_message = "Each LSI must have exactly 2 key_schema elements, follow projection rules, and index_name 3–255 chars long."
  }
}

#############################
# Streams, TTL and SSE      #
#############################

variable "stream_specification" {
  description = "DynamoDB Streams configuration (omit / null to disable)."
  type = object({
    stream_enabled   = bool
    stream_view_type = string # NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES, STREAM_KEYS_ONLY
  })
  default = null

  validation {
    condition = var.stream_specification == null || (
      (!var.stream_specification.stream_enabled) ||
      contains(["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "STREAM_KEYS_ONLY"], var.stream_specification.stream_view_type)
    )
    error_message = "When stream_enabled is true, a valid stream_view_type must be supplied."
  }
}

variable "ttl_specification" {
  description = "Time-to-Live (TTL) configuration (omit / null to disable)."
  type = object({
    ttl_enabled    = bool
    attribute_name = string
  })
  default = null

  validation {
    condition = var.ttl_specification == null || (
      (!var.ttl_specification.ttl_enabled) ||
      (var.ttl_specification.ttl_enabled && length(var.ttl_specification.attribute_name) > 0 && length(var.ttl_specification.attribute_name) <= 255)
    )
    error_message = "attribute_name must be provided (1–255 chars) when ttl_enabled is true."
  }
}

variable "sse_specification" {
  description = "Server-side encryption (SSE) configuration (omit / null to disable)."
  type = object({
    enabled            = bool
    sse_type           = string # AES256 or KMS
    kms_master_key_id  = string
  })
  default = null

  validation {
    condition = var.sse_specification == null || (
      (!var.sse_specification.enabled && var.sse_specification.sse_type == "" && var.sse_specification.kms_master_key_id == "") ||
      (var.sse_specification.enabled && contains(["AES256", "KMS"], var.sse_specification.sse_type) && (
        (var.sse_specification.sse_type == "KMS" && length(var.sse_specification.kms_master_key_id) > 0 && length(var.sse_specification.kms_master_key_id) <= 2048) ||
        (var.sse_specification.sse_type == "AES256" && var.sse_specification.kms_master_key_id == "")
      ))
    )
    error_message = "When SSE is enabled, sse_type must be AES256 or KMS. kms_master_key_id is mandatory only for KMS and must be unset otherwise. When disabled, both sse_type and kms_master_key_id must be empty."
  }
}

#################################
# Miscellaneous table features  #
#################################

variable "point_in_time_recovery_enabled" {
  description = "Enables point-in-time recovery (continuous backups)."
  type        = bool
  default     = false
}

variable "tags" {
  description = "Key/value tags applied to the table."
  type        = map(string)
  default     = {}

  validation {
    condition = alltrue([
      for k, v in var.tags :
      length(k) > 0 && length(v) > 0
    ])
    error_message = "All tag keys and values must be non-empty strings."
  }
}
