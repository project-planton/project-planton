locals {
  # basic naming and tags
  resource_name = coalesce(try(var.metadata.name, null), "aws-ecs-service")
  tags          = merge({
    "Name" = local.resource_name
  }, try(var.metadata.labels, {}))

  # cluster
  cluster_arn = coalesce(
    try(var.spec.cluster_arn.value, null),
    try(var.spec.cluster_arn.value_from.name, null)
  )

  cluster_name = length(split("/", local.cluster_arn)) > 1 ? element(split("/", local.cluster_arn), 1) : ""

  # container settings
  container_port   = try(var.spec.container.port, null)
  desired_count    = coalesce(try(var.spec.container.replicas, null), 1)
  cpu              = var.spec.container.cpu
  memory           = var.spec.container.memory
  image_repo       = try(var.spec.container.image.repo, null)
  image_tag        = try(var.spec.container.image.tag, null)
  logging_enabled  = try(var.spec.container.logging.enabled, true)

  # networking
  subnet_ids = [for s in try(var.spec.network.subnets, []) : coalesce(try(s.value, null), try(s.value_from.name, null))]
  security_group_ids = [for sg in try(var.spec.network.security_groups, []) : coalesce(try(sg.value, null), try(sg.value_from.name, null))]

  # alb
  alb_enabled          = try(var.spec.alb.enabled, false)
  alb_arn              = coalesce(try(var.spec.alb.arn.value, null), try(var.spec.alb.arn.value_from.name, null))
  alb_listener_port    = try(var.spec.alb.listener_port, 80)
  alb_listener_priority = try(var.spec.alb.listener_priority, 100)
  alb_routing_type     = upper(coalesce(try(var.spec.alb.routing_type, null), ""))
  alb_path             = try(var.spec.alb.path, null)
  alb_hostname         = try(var.spec.alb.hostname, null)
}


