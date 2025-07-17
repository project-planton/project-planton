terraform {
  # Require at least Terraform CLI 1.3 (tested with newer 1.x releases)
  required_version = ">= 1.3.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      # Stick to AWS provider 5.x (major-version stability, allow patch/minor updates)
      version = "~> 5.0"
    }
  }
}
