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
  description = "CivoLoadBalancer specification"
  type = object({
    # The name of the Load Balancer (required, 1-64 chars, lowercase alphanumeric and hyphens)
    load_balancer_name = string

    # The Civo region (required)
    region = string

    # Reference to the network (required)
    network = object({
      value = optional(string)
      value_from = optional(object({
        kind       = string
        env        = optional(string)
        name       = string
        field_path = string
      }))
    })

    # Forwarding rules (required, at least 1)
    forwarding_rules = list(object({
      entry_port      = number
      entry_protocol  = string
      target_port     = number
      target_protocol = string
    }))

    # Health check configuration (optional)
    health_check = optional(object({
      port     = number
      protocol = string
      path     = optional(string)
    }))

    # Instance IDs to attach (optional, mutually exclusive with instance_tag)
    instance_ids = optional(list(object({
      value = optional(string)
      value_from = optional(object({
        kind       = string
        env        = optional(string)
        name       = string
        field_path = string
      }))
    })), [])

    # Instance tag for automatic attachment (optional, mutually exclusive with instance_ids)
    instance_tag = optional(string, "")

    # Reserved IP ID (optional)
    reserved_ip_id = optional(object({
      value = optional(string)
      value_from = optional(object({
        kind       = string
        env        = optional(string)
        name       = string
        field_path = string
      }))
    }))

    # Enable sticky sessions (optional, default false)
    enable_sticky_sessions = optional(bool, false)
  })

  validation {
    condition     = length(var.spec.load_balancer_name) >= 1 && length(var.spec.load_balancer_name) <= 64
    error_message = "load_balancer_name must be between 1 and 64 characters."
  }

  validation {
    condition     = can(regex("^[a-z0-9-]+$", var.spec.load_balancer_name))
    error_message = "load_balancer_name must contain only lowercase alphanumeric characters and hyphens."
  }

  validation {
    condition     = length(var.spec.forwarding_rules) >= 1
    error_message = "At least one forwarding rule must be specified."
  }
}