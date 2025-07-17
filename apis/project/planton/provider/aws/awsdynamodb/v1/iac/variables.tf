############################################
# Module inputs for provisioning an AWS     #
# DynamoDB table that matches the behaviour #
# described by AwsDynamodbSpec.              #
############################################

###############################################################################
# CORE TABLE SETTINGS                                                          #
###############################################################################

variable "table_name" {
  description = "Name of the DynamoDB table (3–255 characters)."
  type        = string

  validation {
    condition     = length(var.table_name) >= 3 && length(var.table_name) <= 255
    error_message = "`table_name` must contain between 3 and 255 characters."
  }
}

variable "attribute_definitions" {
  description = <<-EOT
A list of attribute definitions referenced by the table or any index.
Each item must be an object with the following keys:
  * attribute_name – (string, 1–255 chars)
  * attribute_type – (string) One of "S" (string), "N" (number) or "B" (binary).
EOT
  type = list(object({
    attribute_name = string
    attribute_type = string
  }))

  validation {
    condition     = length(var.attribute_definitions) > 0
    error_message = "At least one `attribute_definitions` entry is required."
  }

  validation {
    condition = alltrue([
      for a in var.attribute_definitions :
      length(a.attribute_name) >= 1 &&
      length(a.attribute_name) <= 255 &&
      contains(["S", "N", "B"], upper(a.attribute_type))
    ])
    error_message = "Each attribute definition must have a valid name (1–255 chars) and a type of \"S\", \"N\" or \"B\"."
  }
}

variable "key_schema" {
  description = <<-EOT
The primary key schema for the table. Must contain one (HASH) or two
(HASH + RANGE) elements.  Each element is an object with:
  * attribute_name – (string, 1–255 chars)
  * key_type       – (string) "HASH" or "RANGE"
EOT
  type = list(object({
    attribute_name = string
    key_type       = string
  }))

  validation {
    condition     = length(var.key_schema) >= 1 && length(var.key_schema) <= 2
    error_message = "`key_schema` must contain one or two elements."
  }

  validation {
    condition = alltrue([
      for k in var.key_schema :
      length(k.attribute_name) >= 1 && length(k.attribute_name) <= 255 &&
      contains(["HASH", "RANGE"], upper(k.key_type))
    ])
    error_message = "Each key schema element must have a valid attribute_name and key_type of \"HASH\" or \"RANGE\"."
  }
}

variable "billing_mode" {
  description = "How the table is billed. Allowed values are \"PROVISIONED\" or \"PAY_PER_REQUEST\" (default)."
  type        = string
  default     = "PAY_PER_REQUEST"

  validation {
    condition     = contains(["PROVISIONED", "PAY_PER_REQUEST"], upper(var.billing_mode))
    error_message = "`billing_mode` must be either \"PROVISIONED\" or \"PAY_PER_REQUEST\"."
  }
}

variable "provisioned_throughput" {
  description = <<-EOT
Provisioned capacity settings for the table. **Required when**
`billing_mode` is "PROVISIONED" and **must be omitted** when it is
"PAY_PER_REQUEST".

Expected object:
  {
    read_capacity_units  = number (>0)
    write_capacity_units = number (>0)
  }
EOT
  type    = any   # kept flexible to allow null when on-demand
  default = null

  validation {
    condition = (
      upper(var.billing_mode) == "PAY_PER_REQUEST" && var.provisioned_throughput == null
    ) || (
      upper(var.billing_mode) == "PROVISIONED" &&
      var.provisioned_throughput != null &&
      try(var.provisioned_throughput.read_capacity_units, 0)  > 0 &&
      try(var.provisioned_throughput.write_capacity_units, 0) > 0
    )
    error_message = "When `billing_mode` is PROVISIONED you must supply positive `read_capacity_units` and `write_capacity_units`; when PAY_PER_REQUEST it must be null."
  }
}

###############################################################################
# SECONDARY INDEXES                                                            #
###############################################################################

variable "global_secondary_indexes" {
  description = <<-EOT
List of global secondary indexes (GSIs).  Each list item is an object with
(at minimum) the following keys.

  * index_name              – (string, 3–255 chars)
  * key_schema              – list of up to 2 key-schema objects
  * projection              – object describing the projection
  * provisioned_throughput  – same shape as table-level provisioning;
                               required when `billing_mode` is PROVISIONED.

Full validation of inner objects is performed in the module logic; this
variable definition focuses on type-safety.  Set to an empty list when no
GSIs are required.
EOT
  type    = list(any)
  default = []
}

variable "local_secondary_indexes" {
  description = <<-EOT
List of local secondary indexes (LSIs).  Each list item is an object with
keys similar to GSIs but **must** share the HASH key from the table’s
primary key.  Set to an empty list if no LSIs are needed.
EOT
  type    = list(any)
  default = []
}

###############################################################################
# STREAMS, TTL & POINT-IN-TIME RECOVERY                                        #
###############################################################################

variable "stream_specification" {
  description = <<-EOT
Configuration for DynamoDB Streams. Example:
  {
    stream_enabled   = true
    stream_view_type = "NEW_AND_OLD_IMAGES"   # One of NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | KEYS_ONLY
  }
If omitted or set to `null` streams will be disabled.
EOT
  type    = any
  default = null

  validation {
    condition = var.stream_specification == null || (
      can(var.stream_specification.stream_enabled) &&
      (
        !var.stream_specification.stream_enabled ||
        contains(["NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES", "KEYS_ONLY"], var.stream_specification.stream_view_type)
      )
    )
    error_message = "When streams are enabled you must supply a valid `stream_view_type`."
  }
}

variable "ttl_specification" {
  description = <<-EOT
Time-to-Live (TTL) settings. Example:
  {
    ttl_enabled   = true
    attribute_name = "expires_at"   # epoch time attribute
  }
If omitted or `null`, TTL will be disabled.
EOT
  type    = any
  default = null

  validation {
    condition = var.ttl_specification == null || (
      can(var.ttl_specification.ttl_enabled) &&
      (
        !var.ttl_specification.ttl_enabled ||
        (try(length(var.ttl_specification.attribute_name), 0) > 0)
      )
    )
    error_message = "When TTL is enabled you must specify `attribute_name`."
  }
}

###############################################################################
# SERVER-SIDE ENCRYPTION                                                       #
###############################################################################

variable "sse_specification" {
  description = <<-EOT
Server-Side Encryption (SSE) configuration. Example:
  {
    enabled          = true
    sse_type         = "KMS"    # "AES256" | "KMS"
    kms_master_key_id = "arn:aws:kms:..."  # required when sse_type == "KMS"
  }
If omitted or `null`, encryption defaults to provider/region settings.
EOT
  type    = any
  default = null

  validation {
    condition = var.sse_specification == null || (
      can(var.sse_specification.enabled) &&
      (
        !var.sse_specification.enabled || (
          contains(["AES256", "KMS"], var.sse_specification.sse_type) && (
            (var.sse_specification.sse_type == "KMS" && try(length(var.sse_specification.kms_master_key_id), 0) > 0) ||
            (var.sse_specification.sse_type == "AES256" && try(length(var.sse_specification.kms_master_key_id), 0) == 0)
          )
        )
      )
    )
    error_message = "When SSE is enabled you must set `sse_type`; `kms_master_key_id` is required only for KMS encryption."
  }
}

###############################################################################
# BACKUPS                                                                      #
###############################################################################

variable "point_in_time_recovery_enabled" {
  description = "Enable point-in-time recovery (continuous backups)."
  type        = bool
  default     = false
}

###############################################################################
# TAGS                                                                         #
###############################################################################

variable "tags" {
  description = "Key/value tags to apply to the DynamoDB table."
  type        = map(string)
  default     = {}

  validation {
    condition     = alltrue([for k, v in var.tags : length(trim(k)) > 0 && length(trim(v)) > 0])
    error_message = "`tags` may not contain empty keys or values."
  }
}
