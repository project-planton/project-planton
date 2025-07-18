terraform {
  required_version = ">= 1.3.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  ########################################
  # Authentication & Region
  ########################################
  region      = var.aws_region
  access_key  = var.aws_access_key_id
  secret_key  = var.aws_secret_access_key

  ########################################
  # The default provider configuration is
  # kept minimal on purpose so that the
  # calling stack can decide whether to
  # inject a profile, assume-role, etc.
  ########################################
}
