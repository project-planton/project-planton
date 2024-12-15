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

    # cloud provider
    # https://www.pulumi.com/registry/packages/confluentcloud/api-docs/kafkacluster/#cloud_yaml
    cloud = string

    # availability
    # https://www.pulumi.com/registry/packages/confluentcloud/api-docs/kafkacluster/#availability_yaml
    availability = string

    # environment objects represent an isolated namespace for your confluent resources for organizational purposes.
    # https://www.pulumi.com/registry/packages/confluentcloud/api-docs/kafkacluster/#environment_yaml
    environment = string
  })
}