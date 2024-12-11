variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name   = string
  })
}

variable "spec" {
  description = "Specification for the S3Bucket, including whether it's public and its AWS region"
  type = object({
    is_public  = bool
    aws_region = string
  })
  default = {
    is_public  = false
    aws_region = "us-west-2"
  }
}

variable "aws_credential" {
  description = "AWS Credential data, including account_id, access_key_id, secret_access_key, and region. Optional."
  type = object({
    account_id        = string
    access_key_id     = string
    secret_access_key = string
    region            = string
  })
  default = null
}
