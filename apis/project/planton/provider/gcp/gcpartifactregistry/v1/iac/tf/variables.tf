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
  description = "resource spec"
  type = object({

    # Required.** The ID of the GCP project where the Artifact Registry resources will be created.
    project_id = string

    # Required.** The GCP region where the Artifact Registry will be created (e.g., "us-west2").
    # Selecting a region close to your Kubernetes clusters can reduce service startup time
    # by enabling faster downloads of container images.
    region = string

    # A flag indicating whether to allow unauthenticated access to artifacts published in the repositories.
    # Enable this for publishing artifacts for open-source projects that require public access.
    is_external = bool
  })
}
