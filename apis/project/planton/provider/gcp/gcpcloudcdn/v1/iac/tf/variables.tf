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

    # Required.** The ID of the GCP project where the Cloud CDN resources will be created.
    gcp_project_id = string
  })
}
