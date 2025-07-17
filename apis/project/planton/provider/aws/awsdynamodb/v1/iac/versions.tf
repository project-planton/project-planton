terraform {
  # Lock Terraform CLI to the 1.x series for predictable behavior.
  required_version = ">= 1.4.0, < 2.0.0"

  required_providers {
    # Use the official AWS provider and stay within the 5.x series to avoid
    # breaking changes from future major versions.
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0.0, < 6.0.0"
    }
  }
}
