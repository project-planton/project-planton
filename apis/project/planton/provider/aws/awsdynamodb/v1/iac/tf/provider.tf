terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# AWS provider region
variable "aws_region" {
  description = "AWS region in which to manage resources."
  type        = string
}

# Optional IAM role to assume when interacting with AWS
variable "aws_assume_role_arn" {
  description = "ARN of the IAM role to assume (leave empty to use the default credentials)."
  type        = string
  default     = ""
}

provider "aws" {
  region = var.aws_region

  # Conditionally add an assume_role block when a role ARN is provided.
  dynamic "assume_role" {
    for_each = length(var.aws_assume_role_arn) > 0 ? [var.aws_assume_role_arn] : []
    content {
      role_arn = assume_role.value
    }
  }
}
