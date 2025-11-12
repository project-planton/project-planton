locals {
  # resource name and tags
  resource_name = coalesce(try(var.metadata.name, null), "aws-ecs-cluster")
  tags          = merge({
    "Name" = local.resource_name
  }, try(var.metadata.labels, {}))

  # safe settings with defaults
  safe_enable_container_insights = try(var.spec.enable_container_insights, true)
  safe_enable_execute_command    = try(var.spec.enable_execute_command, false)
  safe_capacity_providers        = try(var.spec.capacity_providers, [])

  # computed values
  container_insights_value = local.safe_enable_container_insights ? "enabled" : "disabled"
  enable_execute_command   = local.safe_enable_execute_command
  has_capacity_providers   = length(local.safe_capacity_providers) > 0
}


