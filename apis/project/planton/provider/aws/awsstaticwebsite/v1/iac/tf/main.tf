locals {
  bucket_name = var.metadata.name
}

resource "aws_s3_bucket" "content" {
  bucket = local.bucket_name
}

