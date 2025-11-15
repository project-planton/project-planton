variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name    = string
    id      = optional(string)
    org     = optional(string)
    env     = optional(string)
    labels  = optional(map(string))
    tags    = optional(list(string))
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "GCP Project specification"
  type = object({
    parent_type             = string                 # "organization" or "folder"
    parent_id               = string                 # Organization ID or Folder ID (numeric string)
    billing_account_id      = string                 # Format: "XXXXXX-XXXXXX-XXXXXX"
    labels                  = optional(map(string))  # Key-value metadata labels
    disable_default_network = optional(bool)         # If true, delete auto-created default VPC (default: true)
    enabled_apis            = optional(list(string)) # List of APIs to enable (e.g., "compute.googleapis.com")
    owner_member            = optional(string)       # IAM member to grant Owner role (e.g., "user:alice@example.com", "group:admins@example.com")
  })
}
