resource "aws_cloudwatch_log_group" "this" {
  count = local.logging_enabled ? 1 : 0
  name  = "/ecs/${local.resource_name}"
  retention_in_days = 30
  tags  = local.tags
}

locals {
  image = local.image_repo != null && local.image_repo != "" ? (
    local.image_tag != null && local.image_tag != "" ? "${local.image_repo}:${local.image_tag}" : local.image_repo
  ) : null
}

resource "aws_ecs_task_definition" "this" {
  family                   = local.resource_name
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = tostring(local.cpu)
  memory                   = tostring(local.memory)

  container_definitions = jsonencode([
    {
      name      = local.resource_name
      image     = local.image
      essential = true
      portMappings = local.container_port != null ? [{
        containerPort = local.container_port
        protocol      = "tcp"
      }] : null
      logConfiguration = local.logging_enabled ? {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.this[0].name
          awslogs-region        = data.aws_region.current.name
          awslogs-stream-prefix = "ecs"
        }
      } : null
      environment = [for k, v in try(var.spec.container.env.variables, {}) : { name = k, value = v }]
    }
  ])
  execution_role_arn = try(var.spec.iam.task_execution_role_arn.value, null)
  task_role_arn      = try(var.spec.iam.task_role_arn.value, null)
  tags               = local.tags
}

data "aws_region" "current" {}

resource "aws_ecs_service" "this" {
  name            = local.resource_name
  cluster         = local.cluster_arn
  task_definition = aws_ecs_task_definition.this.arn
  desired_count   = local.desired_count
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = local.subnet_ids
    security_groups = local.security_group_ids
    assign_public_ip = false
  }

  dynamic "load_balancer" {
    for_each = local.alb_enabled && local.container_port != null ? [1] : []
    content {
      target_group_arn = aws_lb_target_group.this[0].arn
      container_name   = local.resource_name
      container_port   = local.container_port
    }
  }

  depends_on = [aws_ecs_task_definition.this]
  tags       = local.tags
}

resource "aws_lb_target_group" "this" {
  count    = local.alb_enabled && local.container_port != null ? 1 : 0
  name     = substr("tg-${local.resource_name}", 0, 32)
  port     = local.container_port != null ? local.container_port : 80
  protocol = "HTTP"
  vpc_id   = null

  health_check {
    enabled             = true
    path                = coalesce(try(var.spec.alb.health_check.path, null), "/")
    interval            = coalesce(try(var.spec.alb.health_check.interval, null), 30)
    timeout             = coalesce(try(var.spec.alb.health_check.timeout, null), 5)
    healthy_threshold   = coalesce(try(var.spec.alb.health_check.healthy_threshold, null), 5)
    unhealthy_threshold = coalesce(try(var.spec.alb.health_check.unhealthy_threshold, null), 2)
  }
}

resource "aws_lb_listener_rule" "this" {
  count        = local.alb_enabled && local.container_port != null ? 1 : 0
  listener_arn = "${local.alb_arn}:listener/${local.alb_listener_port}"
  priority     = local.alb_listener_priority

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this[0].arn
  }

  dynamic "condition" {
    for_each = local.alb_routing_type == "PATH" && local.alb_path != null ? [1] : []
    content {
      path_pattern {
        values = [local.alb_path]
      }
    }
  }

  dynamic "condition" {
    for_each = local.alb_routing_type == "HOSTNAME" && local.alb_hostname != null ? [1] : []
    content {
      host_header {
        values = [local.alb_hostname]
      }
    }
  }
}


