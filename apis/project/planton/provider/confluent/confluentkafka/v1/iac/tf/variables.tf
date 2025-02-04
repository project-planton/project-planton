variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(object({
      name = optional(string),
      id = optional(string),
    })),
    labels = optional(object({
      key = string, value = string
    })),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
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