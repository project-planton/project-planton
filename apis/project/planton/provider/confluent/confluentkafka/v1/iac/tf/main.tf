resource "aws_s3_bucket" "my_bucket" {
  bucket = var.metadata.name

  # If you want to reference other attributes (e.g., tags) from metadata or spec:
  # tags = var.metadata.labels
}

output "bucketName" {
  value = aws_s3_bucket.my_bucket.bucket
}
