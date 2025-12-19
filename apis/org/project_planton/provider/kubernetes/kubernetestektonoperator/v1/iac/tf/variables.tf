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
  description = "Specification for KubernetesTektonOperator"
  type = object({

    # The Kubernetes cluster to install this operator on.
    target_cluster = optional(object({
      cluster_name = string
      cluster_kind = optional(number)
    }))

    # The container specifications for the Tekton operator.
    container = object({

      # The CPU and memory resources allocated to the Tekton operator container.
      resources = object({

        # The resource limits for the container.
        limits = object({
          cpu    = string
          memory = string
        })

        # The resource requests for the container.
        requests = object({
          cpu    = string
          memory = string
        })
      })
    })

    # Configuration for which Tekton components to install.
    components = object({
      # Enable Tekton Pipelines component.
      pipelines = bool

      # Enable Tekton Triggers component.
      triggers = bool

      # Enable Tekton Dashboard component.
      dashboard = bool
    })
  })
}
