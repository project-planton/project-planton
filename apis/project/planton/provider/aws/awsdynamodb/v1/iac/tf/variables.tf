##############################################
#  Amazon DynamoDB – Module input variables   #
#  Generated from project-planton AwsDynamodb  #
##############################################

############################
# Table-level information  #
############################

variable "table_name" {
  description = "Name of the DynamoDB table. Must be between 3 and 255 characters (as per AWS limits)."
  type        = string

  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "table_name must contain 3–255 characters."
  }
}

############################
# Attribute definitions   #
############################

variable "attribute_definitions" {
  description = <<EOT
List of attribute definitions referenced by the table or any index.  Each
item must contain:
  • attribute_name – string (1–255 chars)
  • attribute_type – one of "STRING", "NUMBER", "BINARY" (corresponding to S, N, B)
EOT
  type = list(object({
    attribute_name = string
    attribute_type = string
  }))

  validation {
    condition = length(var.attribute_definitions) >= 1 && alltrue([
      for d in var.attribute_definitions :
      length(d.attribute_name) >= 1 && length(d.attribute_name) <= 255 &&
      contains(["STRING", "NUMBER", "BINARY"], upper(d.attribute_type))
    ])
    error_message = "attribute_definitions must contain at least one entry; each entry requires attribute_name (1-255 chars) and attribute_type of STRING, NUMBER or BINARY."
  }
}

############################
# Key-schema (primary key) #
############################

variable "key_schema" {
  description = <<EOT
List that describes the primary key schema. 1 element -> simple key, 2
elements -> composite key.  Each element requires:
  • attribute_name – string (1–255 chars)
  • key_type       – "HASH" (partition key) or "RANGE" (sort key)
EOT
  type = list(object({
    attribute_name = string,
    key_type       = string,
  }))

  validation {
    condition = length(var.key_schema) >= 1 && length(var.key_schema) <= 2 && alltrue([
      for k in var.key_schema :
      length(k.attribute_name) >= 1 && length(k.attribute_name) <= 255 &&
      contains(["HASH", "RANGE"], upper(k.key_type))
    ])
    error_message = "key_schema must contain 1 or 2 elements describing HASH / RANGE keys with valid attribute names."
  }
}

############################
# Billing & capacity       #
############################

variable "billing_mode" {
  description = "How the table is billed. Must be either \"PROVISIONED\" or \"PAY_PER_REQUEST\"."
  type        = string

  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], upper(var.billing_mode))
    error_message = "billing_mode must be PROVISIONED or PAY_PER_REQUEST."
  }
}

variable "provisioned_throughput" {
  description = <<EOT
Read/Write capacity units for PROVISIONED mode.  Ignored / must be null when
billing_mode == "PAY_PER_REQUEST".
EOT
  type = object({
    read_capacity_units  = number,
    write_capacity_units = number,
  })
  default = null

  validation {
    condition = (
      upper(var.billing_mode) == "PROVISIONED" ?
        (var.provisioned_throughput != null &&
         var.provisioned_throughput.read_capacity_units  > 0 &&
         var.provisioned_throughput.write_capacity_units > 0)
      : var.provisioned_throughput == null
    )
    error_message = "When billing_mode is PROVISIONED, provisioned_throughput.read_capacity_units and write_capacity_units must both be > 0; when PAY_PER_REQUEST it must be unset/null."
  }
}

###########################################
# Global secondary indexes (GSI)          #
###########################################

variable "global_secondary_indexes" {
  description = <<EOT
Definitions for any global secondary indexes (GSI). Each object may contain:
  • index_name  – string 3-255 chars
  • key_schema  – list(1-2) of { attribute_name, key_type }
  • projection  – { projection_type = ALL|KEYS_ONLY|INCLUDE, non_key_attributes = list(string) }
  • provisioned_throughput – same shape as the root provisioned_throughput (required when billing_mode == PROVISIONED)
EOT
  type    = list(any) # Complex – validated below
  default = []

  validation {
    condition = alltrue([
      for g in var.global_secondary_indexes :
      length(try(g.index_name, "")) >= 3 && length(try(g.index_name, "")) <= 255 &&
      length(try(g.key_schema, [])) >= 1 && length(try(g.key_schema, [])) <= 2 &&
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], upper(try(g.projection.projection_type, "UNDEF"))) &&
      (
        upper(try(g.projection.projection_type, "")) == "INCLUDE" ?
          length(try(g.projection.non_key_attributes, [])) > 0 :
          length(try(g.projection.non_key_attributes, [])) == 0
      ) && (
        upper(var.billing_mode) == "PROVISIONED" ?
          g.provisioned_throughput != null :
          g.provisioned_throughput == null
      )
    ])
    error_message = "Each GSI must meet naming, key_schema, projection rules and obey the billing_mode capacity consistency requirement."
  }
}

###########################################
# Local secondary indexes (LSI)           #
###########################################

variable "local_secondary_indexes" {
  description = <<EOT
Definitions for local secondary indexes (LSI). Each object may contain:
  • index_name  – string 3-255 chars
  • key_schema  – MUST contain exactly 2 elements (HASH + RANGE)
  • projection  – same structure as for GSI
EOT
  type    = list(any)
  default = []

  validation {
    condition = alltrue([
      for l in var.local_secondary_indexes :
      length(try(l.index_name, "")) >= 3 && length(try(l.index_name, "")) <= 255 &&
      length(try(l.key_schema, [])) == 2 &&
      contains(["ALL", "KEYS_ONLY", "INCLUDE"], upper(try(l.projection.projection_type, "UNDEF"))) &&
      (
        upper(try(l.projection.projection_type, "")) == "INCLUDE" ?
          length(try(l.projection.non_key_attributes, [])) > 0 :
          length(try(l.projection.non_key_attributes, [])) == 0
      )
    ])
    error_message = "Each LSI must meet naming, key_schema (exactly 2 elements) and projection rules."
  }
}

###########################################
# Stream settings                         #
###########################################

variable "stream_specification" {
  description = "DynamoDB Streams configuration. Object with stream_enabled (bool) and stream_view_type (NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | STREAM_KEYS_ONLY)."
  type = object({
    stream_enabled   = bool,
    stream_view_type = string,
  })
  default = {
    stream_enabled   = false
    stream_view_type = ""
  }

  validation {
    condition = (!var.stream_specification.stream_enabled) || contains([
      "NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "STREAM_KEYS_ONLY"
    ], upper(var.stream_specification.stream_view_type))
    error_message = "When stream_enabled is true, stream_view_type must be one of NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES or STREAM_KEYS_ONLY."
  }
}

###########################################
# TTL specification                       #
###########################################

variable "ttl_specification" {
  description = "Time-to-Live (TTL) settings. Object with ttl_enabled (bool) and attribute_name (string)."
  type = object({
    ttl_enabled    = bool,
    attribute_name = string,
  })
  default = {
    ttl_enabled    = false
    attribute_name = ""
  }

  validation {
    condition = (!var.ttl_specification.ttl_enabled) || (length(var.ttl_specification.attribute_name) > 0 && length(var.ttl_specification.attribute_name) <= 255)
    error_message = "When ttl_enabled is true, attribute_name must be provided (<=255 chars)."
  }
}

###########################################
# Server-side encryption                  #
###########################################

variable "sse_specification" {
  description = <<EOT
Server-side encryption configuration. Object with:
  • enabled            – bool
  • sse_type           – "AES256" | "KMS" (required when enabled=true)
  • kms_master_key_id  – ARN of CMK (required only when sse_type == "KMS")
EOT
  type = object({
    enabled           = bool,
    sse_type          = string,
    kms_master_key_id = string,
  })
  default = {
    enabled           = false
    sse_type          = ""
    kms_master_key_id = ""
  }

  validation {
    condition = (
      !var.sse_specification.enabled && var.sse_specification.sse_type == "" && var.sse_specification.kms_master_key_id == ""
    ) || (
      var.sse_specification.enabled && contains(["AES256", "KMS"], upper(var.sse_specification.sse_type)) && (
        upper(var.sse_specification.sse_type) == "KMS" ? var.sse_specification.kms_master_key_id != "" : var.sse_specification.kms_master_key_id == ""
      )
    )
    error_message = "When encryption is enabled, sse_type must be AES256 or KMS; kms_master_key_id is required only for KMS. If encryption is disabled, both sse_type and kms_master_key_id must be empty."
  }
}

###########################################
# PITR & tags                             #
###########################################

variable "point_in_time_recovery_enabled" {
  description = "Enable point-in-time recovery (continuous backups)."
  type        = bool
  default     = false
}

variable "tags" {
  description = "Key/value tags to apply to the DynamoDB table."
  type        = map(string)
  default     = {}

  validation {
    condition     = alltrue([for k, v in var.tags : length(k) > 0 && length(v) > 0])
    error_message = "Tag keys and values must be non-empty strings."
  }
}
