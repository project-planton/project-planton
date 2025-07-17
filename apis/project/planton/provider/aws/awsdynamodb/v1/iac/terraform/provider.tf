terraform {
  required_version = ">= 1.3.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# AWS provider configuration. Authentication is expected to be supplied
# through the standard AWS mechanisms (environment variables, shared
# credentials file, SSO, or IAM roles). The region can be passed via the
# `aws_region` variable or inherited from the environment.
provider "aws" {
  region = coalesce(
    try(var.aws_region, null),
    try(env("AWS_REGION"), null),
    try(env("AWS_DEFAULT_REGION"), null)
  )
}
