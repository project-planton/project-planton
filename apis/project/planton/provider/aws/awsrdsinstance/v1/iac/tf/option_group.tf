###############################################################################
# Derive major engine version if user hasn't explicitly provided one
###############################################################################
locals {
  derived_major_engine_version = (
    var.spec.major_engine_version != ""
    ? var.spec.major_engine_version
    : (
    var.spec.engine_version != "" && var.spec.engine == "postgres"
    ? split(".", var.spec.engine_version)[0]
    : (
    length(split(".", var.spec.engine_version)) >= 2
    ? join(".", slice(split(".", var.spec.engine_version), 0, 2))
    : ""
  )))
  final_option_group_name = (var.spec.option_group_name != ""
    ? var.spec.option_group_name
    : (aws_db_option_group.auto[0].name))
}

###############################################################################
# Create an RDS Option Group only if user has not provided an existing one
###############################################################################
resource "aws_db_option_group" "auto" {
  count = var.spec.option_group_name == "" ? 1 : 0

  name_prefix          = "${local.resource_id}-"
  engine_name          = var.spec.engine
  major_engine_version = local.derived_major_engine_version
  tags                 = local.final_labels

  dynamic "option" {
    for_each = var.spec.options
    content {
      option_name = option.value.option_name
      port        = option.value.port
      version = option.value.version

      # Attach DB Security Groups
      db_security_group_memberships = option.value.db_security_group_memberships
      # Attach VPC Security Groups
      vpc_security_group_memberships = option.value.vpc_security_group_memberships

      dynamic "option_settings" {
        for_each = option.value.option_settings
        content {
          name  = option_settings.value.name
          value = option_settings.value.value
        }
      }
    }
  }
}
