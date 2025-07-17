terraform {
  # Require a recent Terraform CLI but stay on the 1.x track.
  required_version = ">= 1.3, < 2.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      # Stay on the current major version (v5) to avoid breaking changes
      # when v6 is released, while still receiving backward-compatible
      # updates and new features.
      version = "~> 5.0"
    }
  }
}