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

    # The container specifications for the OpenFGA deployment.
    container = object({

      # The number of OpenFGA replicas to deploy. This determines the level of concurrency and availability.
      replicas = number

      # The CPU and memory resources allocated to the OpenFGA container.
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

    # The ingress configuration for the OpenFGA deployment.
    ingress = object({

      # A flag to enable or disable ingress.
      is_enabled = bool

      # The dns domain.
      dns_domain = string
    })

    # The data store configuration for OpenFGA.
    # This specifies the backend database engine and connection details.
    datastore = object({

      # Specifies the type of data store engine to use.
      # Allowed values are "mysql" for MySQL database and "postgres" for PostgreSQL database.
      engine = string

      # Specifies the URI to connect to the selected data store engine.
      # The URI format should be appropriate for the specified engine:
      # - For MySQL: `mysql://user:password@host:port/database`
      # - For PostgreSQL: `postgres://user:password@host:port/database`
      uri = string
    })
  })
}