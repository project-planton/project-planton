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
  description = "Specification for the Percona Operator for PostgreSQL deployment"
  type = object({
    # The Kubernetes cluster to install this operator on.
    target_cluster = object({
      # Name of the target Kubernetes cluster
      cluster_name = string
    })

    # Kubernetes namespace to install the operator.
    namespace = string

    # The container specifications for the operator deployment.
    container = object({
      # The CPU and memory resources allocated to the operator container.
      resources = object({
        # The resource limits for the container.
        # Specify the maximum amount of CPU and memory that the container can use.
        limits = object({
          # The amount of CPU allocated (e.g., "500m" for 0.5 CPU cores).
          cpu = string

          # The amount of memory allocated (e.g., "256Mi" for 256 mebibytes).
          memory = string
        })

        # The resource requests for the container.
        # Specify the minimum amount of CPU and memory that the container is guaranteed.
        requests = object({
          # The amount of CPU allocated (e.g., "500m" for 0.5 CPU cores).
          cpu = string

          # The amount of memory allocated (e.g., "256Mi" for 256 mebibytes).
          memory = string
        })
      })
    })
  })
}

