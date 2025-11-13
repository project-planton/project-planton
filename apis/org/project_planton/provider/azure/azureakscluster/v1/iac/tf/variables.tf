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
  description = "Specification for Azure AKS Cluster"
  type = object({
    # Required fields
    region         = string
    vnet_subnet_id = string

    # Optional fields
    network_plugin              = optional(string, "AZURE_CNI")
    kubernetes_version          = optional(string, "1.30")
    private_cluster_enabled     = optional(bool, false)
    authorized_ip_ranges        = optional(list(string), [])
    disable_azure_ad_rbac       = optional(bool, false)
    log_analytics_workspace_id  = optional(string, "")
  })
}