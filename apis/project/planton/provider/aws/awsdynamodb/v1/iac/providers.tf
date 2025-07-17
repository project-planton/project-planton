terraform {
  required_version = ">= 1.3"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

###############################################################################
# AWS provider configuration                                                  #
###############################################################################

# Region and credential wiring is intentionally exposed through variables so
# that the same configuration can run locally (using a named AWS profile) or
# in CI/CD environments (passing static / temporary credentials via TF_VARS or
# environment variables).

variable "aws_region" {
  description = "AWS region to target for all resources."
  type        = string
}

variable "aws_profile" {
  description = "Name of the AWS shared credentials profile to use (optional)."
  type        = string
  default     = null
}

variable "aws_access_key" {
  description = "AWS access key ID (optional, used when explicit credentials are preferred over a profile)."
  type        = string
  default     = null
  sensitive   = true
}

variable "aws_secret_key" {
  description = "AWS secret access key (optional)."
  type        = string
  default     = null
  sensitive   = true
}

variable "aws_session_token" {
  description = "AWS session token when using temporary credentials (optional)."
  type        = string
  default     = null
  sensitive   = true
}

provider "aws" {
  # Required: region in which resources will be created.
  region = var.aws_region

  # If explicit access keys are supplied they take precedence over the profile.
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  token      = var.aws_session_token

  # Fallback to a shared credentials profile when keys are not provided.
  profile = var.aws_profile
}
