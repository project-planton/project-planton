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
  description = "Specification for the AwsS3Bucket, including whether it's public and its AWS region"
  type = object({
    is_public = optional(bool, false),
    aws_region = optional(string, "us-west-2"),
  })
}
