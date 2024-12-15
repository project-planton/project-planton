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

    # The repository URL where the Helm chart is hosted.
    # For example, "https://charts.helm.sh/stable".
    # an example for chart-repo (redis chart) can be found in https://artifacthub.io/packages/helm/bitnami/redis?modal=install
    repo = string

    # The name of the Helm chart to deploy.
    # For example, "nginx-ingress".
    name = string

    # The version of the Helm chart to deploy.
    # For example, "1.41.3".
    version = string

    # A map of key-value pairs representing custom values for the Helm chart.
    # These values override the default settings in the chart's values.yaml file.
    values = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })
  })
}