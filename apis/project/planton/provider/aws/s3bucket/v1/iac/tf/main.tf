provider "aws" {
  # If aws_credential is provided, use it; otherwise, fall back to spec.aws_region only.
  region     = var.aws_credential == null ? var.spec.aws_region : var.aws_credential.region
  access_key = var.aws_credential == null ? null : var.aws_credential.access_key_id
  secret_key = var.aws_credential == null ? null : var.aws_credential.secret_access_key
}

resource "aws_s3_bucket" "my_bucket" {
  bucket = var.metadata.name

  # If you want to reference other attributes (e.g., tags) from metadata or spec:
  # tags = var.metadata.labels
}

output "bucketName" {
  value = aws_s3_bucket.my_bucket.bucket
}
