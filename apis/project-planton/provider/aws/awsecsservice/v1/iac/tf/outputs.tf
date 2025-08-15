output "aws_ecs_service_name" {
  description = "The ECS service name."
  value       = aws_ecs_service.this.name
}

output "ecs_cluster_name" {
  description = "The name component derived from the target cluster ARN."
  value       = local.cluster_name
}

output "load_balancer_dns_name" {
  description = "DNS name of the ALB if attached."
  value       = null
}

output "service_url" {
  description = "External endpoint if hostname was configured."
  value       = null
}

output "service_discovery_name" {
  description = "Service discovery internal DNS name if used."
  value       = null
}

output "cloudwatch_log_group_name" {
  description = "CloudWatch log group name."
  value       = local.logging_enabled ? aws_cloudwatch_log_group.this[0].name : null
}

output "cloudwatch_log_group_arn" {
  description = "CloudWatch log group ARN."
  value       = local.logging_enabled ? aws_cloudwatch_log_group.this[0].arn : null
}

output "service_arn" {
  description = "ECS service ARN."
  value       = aws_ecs_service.this.arn
}

output "target_group_arn" {
  description = "Target group ARN if ALB is enabled."
  value       = local.alb_enabled && local.container_port != null ? aws_lb_target_group.this[0].arn : null
}


