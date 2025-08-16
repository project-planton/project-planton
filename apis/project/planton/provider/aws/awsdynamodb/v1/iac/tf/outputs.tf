output "table_name" {
  value = aws_dynamodb_table.this.name
}

output "table_arn" {
  value = aws_dynamodb_table.this.arn
}

output "table_id" {
  value = aws_dynamodb_table.this.id
}

output "stream_arn" {
  value = try(aws_dynamodb_table.this.stream_arn, null)
}

output "stream_label" {
  value = try(aws_dynamodb_table.this.stream_label, null)
}


