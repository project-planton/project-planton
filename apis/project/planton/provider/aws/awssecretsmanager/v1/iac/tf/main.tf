resource "aws_secretsmanager_secret" "secrets" {
  for_each = toset(var.spec.secret_names)

  name = "${var.metadata.id}-${each.value}"

  tags = local.aws_tags
}

resource "aws_secretsmanager_secret_version" "placeholder_versions" {
  for_each = aws_secretsmanager_secret.secrets

  secret_id     = each.value.id
  secret_string = "placeholder"

  # Ignore changes to secret_string to mimic Pulumi's IgnoreChanges([]string{"secretString"})
  lifecycle {
    ignore_changes = [secret_string]
  }
}
