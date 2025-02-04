resource "aws_lambda_permission" "invoke" {
  count = length(local.invoke_function_permissions)

  statement_id  = "AllowInvocation-${count.index}"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.this.function_name

  principal  = local.invoke_function_permissions[count.index].principal
  source_arn = local.invoke_function_permissions[count.index].source_arn
}
