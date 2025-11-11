variable "metadata" {
  description = "metadata captures identifying information (name, org, version, etc.)\nand must pass standard validations for resource naming."
  type = object({

    # name of the resource
    name = string

    # id of the resource
    id = string

    # id of the organization to which the api-resource belongs to
    org = string

    # environment to which the resource belongs to
    env = string

    # labels for the resource
    labels = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # annotations for the resource
    annotations = object({

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
  description = "spec holds the core configuration data defining how the ECS service is deployed."
  type = object({

    # enable_container_insights determines whether to enable CloudWatch
    # Container Insights for this cluster. This is highly recommended
    # for production monitoring, though it incurs CloudWatch costs.
    # If omitted, it is recommended to be "true".
    enable_container_insights = bool

    # capacity_providers is a list of capacity providers attached
    # to this cluster. For a Fargate-only cluster, typically ["FARGATE"]
    # or ["FARGATE", "FARGATE_SPOT"] for optional Spot usage.
    capacity_providers = list(string)

    # enable_execute_command controls whether ECS Exec is allowed on
    # tasks in this cluster, letting you exec into running containers
    # for debugging or operational tasks. Defaults to false.
    enable_execute_command = bool
  })
}