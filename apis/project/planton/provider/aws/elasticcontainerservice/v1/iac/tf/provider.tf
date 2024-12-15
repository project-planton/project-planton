variable "aws_credential" {
  description = "AWS Credential data, including account_id, access_key_id, secret_access_key, and region. Optional."
  type = object({
    account_id        = string
    access_key_id     = string
    secret_access_key = string
    region            = string
  })
}

provider "aws" {
  # If aws_credential is provided, use it; otherwise, fall back to spec.aws_region only.
  region     = var.aws_credential == null ? var.spec.aws_region : var.aws_credential.region
  access_key = var.aws_credential == null ? null : var.aws_credential.access_key_id
  secret_key = var.aws_credential == null ? null : var.aws_credential.secret_access_key
}
