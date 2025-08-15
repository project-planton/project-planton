resource "aws_ecs_cluster" "this" {
  name = local.resource_name

  setting {
    name  = "containerInsights"
    value = local.container_insights_value
  }

  dynamic "capacity_providers" {
    for_each = local.has_capacity_providers ? [1] : []
    content {
      # terraform requires a schema structure; use top-level argument instead
    }
  }

  capacity_providers = try(var.spec.capacity_providers, null)

  tags = local.tags
}


