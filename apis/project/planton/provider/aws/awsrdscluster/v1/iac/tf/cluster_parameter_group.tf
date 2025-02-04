resource "aws_rds_cluster_parameter_group" "this" {
  count = (
  var.spec.cluster_parameter_group_name == null
  || var.spec.cluster_parameter_group_name == ""
  ) ? 1 : 0

  name_prefix = "${local.resource_id}-"
  family      = var.spec.cluster_family
  tags        = local.final_labels

  dynamic "parameter" {
    for_each = var.spec.cluster_parameters
    content {
      name         = parameter.value.name
      value        = parameter.value.value
      apply_method = parameter.value.apply_method
    }
  }
}
