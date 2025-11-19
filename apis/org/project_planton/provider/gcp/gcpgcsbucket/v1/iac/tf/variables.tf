variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
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
    is_public = optional(bool, false)

    # Name of the GCS bucket to create in GCP
    bucket_name = string
  })
  
  validation {
    condition     = can(regex("^[a-z0-9]([a-z0-9-._]*[a-z0-9])?$", var.spec.bucket_name)) && length(var.spec.bucket_name) >= 3 && length(var.spec.bucket_name) <= 63
    error_message = "Bucket name must be 3-63 characters, globally unique, lowercase letters, numbers, hyphens, or dots, starting and ending with a letter or number."
  }
}
