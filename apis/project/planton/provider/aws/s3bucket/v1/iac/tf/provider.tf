variable "aws_credential" {
  description = "AWS Credential data, including account_id, access_key_id, secret_access_key, and region. Optional."
  type = object({
    access_key_id     = string
    secret_access_key = string
    region            = string
  })
  default = null
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region     = var.aws_credential.region != null ? var.aws_credential.region : var.spec.aws_region
  access_key = var.aws_credential.access_key_id != null ? var.aws_credential.access_key_id : null
  secret_key = var.aws_credential.secret_access_key != null ? var.aws_credential.secret_access_key : null
}
