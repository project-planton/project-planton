resource "aws_s3_bucket" "this" {
  bucket = var.metadata.name
}

resource "aws_s3_bucket_ownership_controls" "this" {
  bucket = aws_s3_bucket.this.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "this" {
  bucket = aws_s3_bucket.this.id

  block_public_acls       =  !var.spec.is_public
  block_public_policy     =  !var.spec.is_public
  ignore_public_acls      =  !var.spec.is_public
  restrict_public_buckets =  !var.spec.is_public
}

resource "aws_s3_bucket_acl" "public_read" {
  count = var.spec.is_public ? 1 : 0
  depends_on = [
    aws_s3_bucket_ownership_controls.this,
    aws_s3_bucket_public_access_block.this,
  ]

  bucket = aws_s3_bucket.this.id
  acl    = "public-read"
}
