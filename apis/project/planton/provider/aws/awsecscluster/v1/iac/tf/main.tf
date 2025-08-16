resource "aws_ecs_cluster" "this" {
  name = local.resource_name

  setting {
    name  = "containerInsights"
    value = local.container_insights_value
  }

  dynamic "configuration" {
    for_each = local.enable_execute_command ? [1] : []
    content {
      execute_command_configuration {
        logging = "DEFAULT"
      }
    }
  }

  tags = local.tags
}

# Attach capacity providers to the cluster if specified
resource "aws_ecs_cluster_capacity_providers" "this" {
  count = local.has_capacity_providers ? 1 : 0

  cluster_name       = aws_ecs_cluster.this.name
  capacity_providers = var.spec.capacity_providers

  default_capacity_provider_strategy {
    capacity_provider = var.spec.capacity_providers[0]
    weight           = 1
    base             = 0
  }
}


