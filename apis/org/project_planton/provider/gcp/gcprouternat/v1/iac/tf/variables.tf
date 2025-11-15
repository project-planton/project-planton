variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name    = string,
    id      = optional(string),
    org     = optional(string),
    env     = optional(string),
    labels  = optional(map(string)),
    tags    = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "Specification for the GCP Cloud Router NAT"
  type = object({
    vpc_self_link          = string
    region                 = string
    subnetwork_self_links  = optional(list(string), [])
    nat_ip_names           = optional(list(string), [])
    log_filter             = optional(string, "ERRORS_ONLY") # DISABLED, ERRORS_ONLY, or ALL
  })
}
