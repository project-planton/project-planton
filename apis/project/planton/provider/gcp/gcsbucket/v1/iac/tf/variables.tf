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

    # Required.** The ID of the GCP project where the storage bucket will be created.
    gcp_project_id = string

    # Required.** The GCP region where the storage bucket will be created.
    gcp_region = string

    # A flag indicating whether the GCS bucket should have external (public) access.
    # Defaults to `false`, meaning the bucket is private by default.
    is_public = bool
  })
}
