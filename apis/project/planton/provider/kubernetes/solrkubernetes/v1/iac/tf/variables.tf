variable "metadata" {
  description = "metadata"
  type = object({

    # name of the resource
    name = string

    # id of the resource
    id = string

    # id of the organization to which the api-resource belongs to
    org = string

    # environment to which the resource belongs to
    env = object({

      # name of the environment
      name = string

      # id of the environment
      id = string
    })

    # labels for the resource
    labels = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # tags for the resource
    tags = list(string)
  })
}

variable "spec" {
  description = "spec"
  type = object({

    # The specifications for the Solr container deployment.
    solr_container = object({

      # The number of Solr pods in the Solr Kubernetes deployment.
      replicas = number

      # The CPU and memory resources allocated to the Solr container.
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

      # The size of the persistent volume attached to each Solr pod (e.g., "1Gi").
      disk_size = string

      # The container image for the Solr deployment.
      # Example repository: "solr", example tag: "8.7.0".
      image = object({

        # The repository of the image (e.g., "gcr.io/project/image").
        repo = string

        # The tag of the image (e.g., "latest" or "1.0.0").
        tag = string

        # The name of the image pull secret for private image repositories.
        pull_secret_name = string
      })
    })

    # The Solr-specific configuration options.
    config = object({

      # JVM memory settings for Solr.
      java_mem = string

      # Custom Solr options (e.g., "-Dsolr.autoSoftCommit.maxTime=10000").
      opts = string

      # Solr garbage collection tuning configuration (e.g., "-XX:SurvivorRatio=4 -XX:TargetSurvivorRatio=90 -XX:MaxTenuringThreshold=8").
      garbage_collection_tuning = string
    })

    # The specifications for the Zookeeper container deployment.
    zookeeper_container = object({

      # The number of Zookeeper pods in the Zookeeper cluster.
      replicas = number

      # The CPU and memory resources allocated to the Zookeeper container.
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

      # The size of the persistent volume attached to each Zookeeper pod (e.g., "1Gi").
      disk_size = string
    })

    # The ingress configuration for the Solr deployment.
    ingress = object({

      # A flag to enable or disable ingress.
      is_enabled = bool

      # The dns domain.
      dns_domain = string
    })
  })
}