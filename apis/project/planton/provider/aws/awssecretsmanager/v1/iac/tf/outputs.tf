output "secret_arn_map" {
  value = { for k, v in aws_secretsmanager_secret.secrets : k => v.arn }
}
