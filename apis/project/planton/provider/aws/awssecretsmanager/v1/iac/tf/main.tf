resource "aws_secretsmanager_secret" "secrets" {
  for_each = {
    for name in local.secret_names :
    name => name
  }

  name = format("%s-%s", local.resource_id, each.value)
  tags = local.final_tags
}

resource "aws_secretsmanager_secret_version" "secrets" {
  for_each = {
    for name in local.secret_names :
    name => name
  }

  secret_id     = aws_secretsmanager_secret.secrets[each.key].id
  secret_string = "placeholder"

  lifecycle {
    ignore_changes = [secret_string]
  }
}
