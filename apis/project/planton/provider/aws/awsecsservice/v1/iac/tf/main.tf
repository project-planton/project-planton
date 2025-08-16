# CloudWatch Log Group for ECS service logging
resource "aws_cloudwatch_log_group" "this" {
  count             = local.logging_enabled ? 1 : 0
  name              = "/ecs/${local.resource_name}"
  retention_in_days = 30
  tags              = local.tags
}

# Data source to get ALB information for outputs
data "aws_lb" "selected" {
  count = local.has_alb_config ? 1 : 0
  arn   = local.safe_alb_arn
}

# Current AWS region data source
data "aws_region" "current" {}

# ECS Task Definition
resource "aws_ecs_task_definition" "this" {
  family                   = local.resource_name
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = tostring(local.cpu)
  memory                   = tostring(local.memory)

  container_definitions = jsonencode([
    {
      name      = local.resource_name
      image     = local.container_image
      essential = true
      
      # Port mappings (only if container port is specified)
      portMappings = local.container_port != null ? [{
        containerPort = local.container_port
        protocol      = "tcp"
      }] : []
      
      # Logging configuration
      logConfiguration = local.logging_enabled ? {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.this[0].name
          awslogs-region        = data.aws_region.current.name
          awslogs-stream-prefix = "ecs"
        }
      } : null
      
      # Environment variables
      environment = [
        for k, v in try(var.spec.container.env.variables, {}) : {
          name  = k
          value = v
        }
      ]
      
      # Secrets (if any)
      secrets = [
        for k, v in try(var.spec.container.env.secrets, {}) : {
          name      = k
          valueFrom = v
        }
      ]
      
      # Environment files from S3 (if any)
      environmentFiles = [
        for s3_uri in try(var.spec.container.env.s3_files, []) : {
          value = s3_uri
          type  = "s3"
        }
      ]
    }
  ])

  # IAM roles
  execution_role_arn = local.safe_task_execution_role_arn
  task_role_arn      = local.safe_task_role_arn
  
  tags = local.tags
}

# ECS Service
resource "aws_ecs_service" "this" {
  name            = local.resource_name
  cluster         = local.safe_cluster_arn
  task_definition = aws_ecs_task_definition.this.arn
  desired_count   = local.desired_count
  launch_type     = "FARGATE"

  # Network configuration
  network_configuration {
    subnets          = local.safe_subnet_ids
    security_groups  = local.safe_security_group_ids
    assign_public_ip = false
  }

  # Load balancer configuration (only if ALB is enabled and container port is specified)
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

# ALB Target Group (only if ALB is enabled and container port is specified)
resource "aws_lb_target_group" "this" {
  count    = local.alb_enabled && local.container_port != null ? 1 : 0
  name     = substr("tg-${local.resource_name}", 0, 32)
  port     = local.container_port
  protocol = "HTTP"
  vpc_id   = null

  # Health check configuration
  health_check {
    enabled             = true
    path                = coalesce(try(var.spec.alb.health_check.path, null), "/")
    interval            = coalesce(try(var.spec.alb.health_check.interval, null), 30)
    timeout             = coalesce(try(var.spec.alb.health_check.timeout, null), 5)
    healthy_threshold   = coalesce(try(var.spec.alb.health_check.healthy_threshold, null), 5)
    unhealthy_threshold = coalesce(try(var.spec.alb.health_check.unhealthy_threshold, null), 2)
    protocol            = coalesce(try(var.spec.alb.health_check.protocol, null), "HTTP")
    port                = coalesce(try(var.spec.alb.health_check.port, null), "traffic-port")
  }

  tags = local.tags
}

# ALB Listener Rule (only if ALB is enabled and container port is specified)
resource "aws_lb_listener_rule" "this" {
  count        = local.alb_enabled && local.container_port != null ? 1 : 0
  listener_arn = "${local.safe_alb_arn}:listener/${local.alb_listener_port}"
  priority     = local.alb_listener_priority

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this[0].arn
  }

  # Path-based routing condition
  dynamic "condition" {
    for_each = local.alb_routing_type == "path" && local.alb_path != null ? [1] : []
    content {
      path_pattern {
        values = [local.alb_path]
      }
    }
  }

  # Hostname-based routing condition
  dynamic "condition" {
    for_each = local.alb_routing_type == "hostname" && local.alb_hostname != null ? [1] : []
    content {
      host_header {
        values = [local.alb_hostname]
      }
    }
  }

  depends_on = [aws_lb_target_group.this]
}


