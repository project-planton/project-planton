terraform {
  # Terraform CLI tested against this module. Update with major releases only.
  required_version = ">= 1.3.0, < 2.0.0"

  required_providers {
    # AWS provider – used for all AWS resources, including DynamoDB.
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0" # Tested with 5.x; accept any compatible patch/minor update.
    }

    # Random provider – commonly used for unique suffixes (e.g., table names).
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5" # Tested with 3.5.x series.
    }
  }
}
