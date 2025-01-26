resource "aws_s3_bucket" "my_bucket" {
  bucket = var.metadata.name

  # If you want to reference other attributes (e.g., tags) from metadata or spec:
  # tags = var.metadata.labels
}

resource "aws_route53_zone" "my_zone" {
  name = "project-planton.com"
}

resource "aws_route53_record" "my_record" {
  zone_id = aws_route53_zone.my_zone.zone_id
  name    = "www"
  type    = "A"
  ttl     = "300"
  records = ["192.0.2.44"]
}

resource "aws_dynamodb_table" "my_table" {
  name           = "example-table"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "id"
  attribute {
    name = "id"
    type = "S"
  }
}

resource "aws_s3_bucket" "another_bucket" {
  bucket = "another-example-bucket"
}

output "bucketName" {
  value = aws_s3_bucket.my_bucket.bucket
}
