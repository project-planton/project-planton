resource "aws_db_subnet_group" "default" {
  count = (length(var.spec.subnet_ids) > 0 && var.spec.db_subnet_group_name == "") ? 1 : 0

  name       = local.resource_id
  subnet_ids = var.spec.subnet_ids
  tags       = local.final_labels
}
