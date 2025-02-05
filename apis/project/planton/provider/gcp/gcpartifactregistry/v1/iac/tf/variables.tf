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
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
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
