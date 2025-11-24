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
  description = "NatsKubernetes specification"
  type = object({
    # Kubernetes namespace to install NATS
    namespace = string

    # Server container configuration
    server_container = object({
      # Number of NATS replicas
      replicas = number

      # CPU and memory resources
      resources = object({
        limits = object({
          cpu    = string
          memory = string
        })
        requests = object({
          cpu    = string
          memory = string
        })
      })

      # PVC size for JetStream file store (e.g., "10Gi")
      disk_size = string
    })

    # Disable JetStream persistence
    disable_jet_stream = optional(bool, false)

    # Authentication configuration
    auth = optional(object({
      enabled = bool
      scheme  = string # "bearer_token" or "basic_auth"
      no_auth_user = optional(object({
        enabled          = bool
        publish_subjects = list(string)
      }))
    }))

    # TLS encryption
    tls_enabled = optional(bool, false)

    # Ingress configuration for external access
    ingress = optional(object({
      enabled  = bool
      hostname = string
    }))

    # Toggle to deploy nats-box utility pod
    disable_nats_box = optional(bool, false)
  })
}
