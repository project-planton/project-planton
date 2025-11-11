resource "aws_s3_bucket" "this" {
  bucket = local.bucket_name
}

resource "aws_s3_bucket_ownership_controls" "this" {
  bucket = aws_s3_bucket.this.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "this" {
  bucket = aws_s3_bucket.this.id

  block_public_acls       = !local.is_public
  block_public_policy     = !local.is_public
  ignore_public_acls      = !local.is_public
  restrict_public_buckets = !local.is_public
}

resource "aws_s3_bucket_acl" "public_read" {
  count = local.is_public ? 1 : 0
  depends_on = [
    aws_s3_bucket_ownership_controls.this,
    aws_s3_bucket_public_access_block.this,
  ]

  bucket = aws_s3_bucket.this.id
  acl    = "public-read"
}
