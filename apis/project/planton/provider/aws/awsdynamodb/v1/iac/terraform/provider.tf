terraform {
  required_version = ">= 1.3.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# AWS region in which all resources will be created.
variable "aws_region" {
  description = "AWS region to deploy resources into"
  type        = string
}

provider "aws" {
  region = var.aws_region
}
