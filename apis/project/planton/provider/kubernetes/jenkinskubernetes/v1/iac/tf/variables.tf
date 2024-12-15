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

    # The CPU and memory resources allocated to the Jenkins container.
    container_resources = object({

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

    # A map of key-value pairs that provide additional customization options for the Helm chart used to deploy Jenkins.
    # These values allow for further refinement of the deployment, such as customizing resource limits, setting environment variables,
    # or specifying version tags. For detailed information on the available options, refer to the Helm chart documentation at:
    # https://github.com/jenkinsci/helm-charts/blob/main/charts/jenkins/values.yaml
    helm_values = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # The ingress configuration for the Jenkins deployment.
    ingress = object({

      # A flag to enable or disable ingress.
      is_enabled = bool

      # The dns domain.
      dns_domain = string
    })
  })
}