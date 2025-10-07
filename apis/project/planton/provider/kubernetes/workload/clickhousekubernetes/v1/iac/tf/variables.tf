variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # The container specifications for the ClickHouse deployment.
    container = object({

      # The number of ClickHouse pods to deploy.
      replicas = number

      # The CPU and memory resources allocated to the ClickHouse container.
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

      # A flag to enable or disable data persistence for ClickHouse.
      # When enabled, data is persisted to a storage volume, allowing data to survive pod restarts.
      # Defaults to `true`.
      is_persistence_enabled = bool

      # The size of the persistent volume attached to each ClickHouse pod (e.g., "10Gi").
      # If the client does not provide a value, a default value is configured.
      # This attribute is ignored when persistence is not enabled.
      # **Note:** This value cannot be modified after creation due to Kubernetes limitations on stateful sets.
      disk_size = string
    })

    # The ingress configuration for the ClickHouse deployment.
    ingress = optional(object({

      # A flag to enable or disable ingress.
      is_enabled = bool

      # The dns domain.
      dns_domain = string
    }))

    # The cluster configuration for ClickHouse sharding and replication.
    cluster = optional(object({

      # A flag to enable or disable clustering mode for ClickHouse.
      # When enabled, ClickHouse will be deployed in a distributed cluster configuration.
      # Defaults to `false`.
      is_enabled = bool

      # The number of shards in the ClickHouse cluster.
      # Sharding distributes data across multiple nodes for horizontal scaling.
      # This value is ignored if clustering is not enabled.
      shard_count = number

      # The number of replicas for each shard.
      # Replication provides data redundancy and high availability.
      # This value is ignored if clustering is not enabled.
      replica_count = number
    }))

    # A map of key-value pairs that provide additional customization options for the Helm chart used
    # to deploy ClickHouse on Kubernetes. These values allow for further refinement of the deployment,
    # such as customizing resource limits, setting environment variables, or specifying version tags.
    # For detailed information on the available options, refer to the Helm chart documentation at:
    # https://artifacthub.io/packages/helm/bitnami/clickhouse
    helm_values = optional(map(string))
  })
}
