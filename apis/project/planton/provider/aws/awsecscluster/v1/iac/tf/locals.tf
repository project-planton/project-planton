locals {
  # resource name and tags
  resource_name = coalesce(try(var.metadata.name, null), "aws-ecs-cluster")
  tags          = merge({
    "Name" = local.resource_name
  }, try(var.metadata.labels, {}))

  # settings
  container_insights_enabled = try(var.spec.enable_container_insights, true)
  container_insights_value   = local.container_insights_enabled ? "enabled" : "disabled"

  has_capacity_providers = length(try(var.spec.capacity_providers, [])) > 0
}


