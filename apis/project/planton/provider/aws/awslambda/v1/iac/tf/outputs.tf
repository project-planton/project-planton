output "iam_role_name" {
  description = "Name of the IAM Role created for the Lambda Function"
  value       = local.create_iam_role ? aws_iam_role.lambda[0].name : null
}

output "cloudwatch_log_group_name" {
  description = "Name of the CloudWatch Log Group for the Lambda Function"
  value       = local.create_cloudwatch_log_group ? aws_cloudwatch_log_group.lambda[0].name : null
}

output "lambda_function_arn" {
  description = "ARN of the created Lambda Function"
  value       = aws_lambda_function.this.arn
}

output "lambda_function_name" {
  description = "Name of the created Lambda Function"
  value       = aws_lambda_function.this.function_name
}
