###############################################################################
# Create an RDS Parameter Group only if user has not provided an existing one
###############################################################################
resource "aws_db_parameter_group" "auto" {
  count = var.spec.parameter_group_name == "" ? 1 : 0

  name_prefix = "${local.resource_id}-"
  family      = var.spec.db_parameter_group
  tags        = local.final_labels

  dynamic "parameter" {
    for_each = var.spec.parameters
    content {
      apply_method = parameter.value.apply_method
      name         = parameter.value.name
      value        = parameter.value.value
    }
  }
}

###############################################################################
# Local reference to Parameter Group name (either new or user-provided)
###############################################################################
locals {
  final_parameter_group_name = (var.spec.parameter_group_name != ""
    ? var.spec.parameter_group_name
    : (aws_db_parameter_group.auto[0].name))
}
