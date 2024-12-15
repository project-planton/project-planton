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

    # Required.** The ID of the GCP project where the secrets will be created.
    project_id = string

    # A list of secret names to create in Google Cloud Secrets Manager.
    # Each name represents a unique secret that can store sensitive data securely.
    secret_names = list(string)
  })
}
