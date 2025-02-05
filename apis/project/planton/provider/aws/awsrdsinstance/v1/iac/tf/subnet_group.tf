###############################################################################
# Create a DB Subnet Group only if user has supplied subnet_ids but not an
# existing db_subnet_group_name
###############################################################################
resource "aws_db_subnet_group" "auto" {
  count = ((
  length(var.spec.subnet_ids) > 0
  && var.spec.db_subnet_group_name == ""
  ) ? 1 : 0)

  name       = local.resource_id
  subnet_ids = var.spec.subnet_ids
  tags       = local.final_labels
}

###############################################################################
# Local reference to the final DB Subnet Group name
###############################################################################
locals {
  final_db_subnet_group_name = (var.spec.db_subnet_group_name != ""
    ? var.spec.db_subnet_group_name
    : length(aws_db_subnet_group.auto) > 0
      ? aws_db_subnet_group.auto[0].name
      : null)
}
