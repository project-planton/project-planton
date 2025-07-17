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
# Input variables controlling the AWS provider configuration.
###############################################################################

variable "aws_region" {
  description = "AWS region to deploy resources into."
  type        = string

  validation {
    condition     = length(trimspace(var.aws_region)) > 0
    error_message = "A non-empty AWS region must be specified."
  }
}

variable "aws_profile" {
  description = "AWS CLI/SDK profile name to use for credentials (optional).  Leave empty to rely on the default credential chain."
  type        = string
  default     = ""
}

variable "aws_assume_role_arn" {
  description = "ARN of an IAM role that Terraform should assume before making AWS API calls (optional).  When empty, no role is assumed."
  type        = string
  default     = ""
}

variable "aws_assume_role_session_name" {
  description = "Session name to use when assuming the IAM role (ignored when no role ARN is supplied)."
  type        = string
  default     = "terraform"
}

###############################################################################
# Default AWS provider – uses direct credentials or an optional profile.
###############################################################################

provider "aws" {
  region  = var.aws_region
  profile = length(trimspace(var.aws_profile)) > 0 ? var.aws_profile : null
}

###############################################################################
# Optional AWS provider that assumes an IAM role.  Down-stream modules can
# opt-in by referencing the "aws.assume_role" alias when a role ARN is given.
###############################################################################

provider "aws" {
  alias   = "assume_role"
  region  = var.aws_region
  profile = length(trimspace(var.aws_profile)) > 0 ? var.aws_profile : null

  # This block is only meaningful when a non-empty role ARN is supplied.  Call
  # sites must ensure they reference this provider alias only in that case.
  assume_role {
    role_arn     = var.aws_assume_role_arn
    session_name = var.aws_assume_role_session_name
  }
}
