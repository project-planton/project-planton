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

    # List of secret names to create in AWS Secrets Manager.
    # Each name corresponds to a unique secret that will be securely stored and managed.
    # Secret names must be unique within your AWS account and region.
    secret_names = optional(list(string))
  })
}
