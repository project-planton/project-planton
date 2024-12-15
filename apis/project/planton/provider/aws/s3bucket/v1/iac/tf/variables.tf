variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = string,
    org = string,
    env = object({
      name = string,
      id = string }
    ),
    labels = object({
      key = string, value = string
    }),
    tags = list(string),
    version = object({ id = string, message = string })
  })
}

variable "spec" {
  description = "Specification for the S3Bucket, including whether it's public and its AWS region"
  type = object({
    is_public  = bool,
    aws_region = string
  })
}
