output "function_arn" {
  description = "The ARN of the Lambda function."
  value       = aws_lambda_function.this.arn
}

output "function_name" {
  description = "The name of the Lambda function."
  value       = aws_lambda_function.this.function_name
}

output "log_group_name" {
  description = "CloudWatch Logs log group name for the function."
  value       = aws_cloudwatch_log_group.lambda.name
}

output "role_arn" {
  description = "Execution role ARN used by the Lambda function."
  value       = aws_iam_role.lambda.arn
}

output "layer_arns" {
  description = "List of layer ARNs attached to the Lambda function."
  value       = aws_lambda_function.this.layers
}


