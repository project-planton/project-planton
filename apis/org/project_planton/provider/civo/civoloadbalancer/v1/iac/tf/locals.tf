# Local values for Civo Load Balancer Terraform Module

locals {
  # Resource ID (prefer metadata.id if set, otherwise use metadata.name)
  resource_id = coalesce(var.metadata.id, var.metadata.name)

  # Base labels for resource tagging
  base_labels = {
    "planton:resource"      = "true"
    "planton:resource_id"   = local.resource_id
    "planton:resource_name" = var.metadata.name
    "planton:resource_kind" = "civo_load_balancer"
  }

  # Organization label (if provided)
  org_label = var.metadata.org != null && var.metadata.org != "" ? {
    "planton:organization" = var.metadata.org
  } : {}

  # Environment label (if provided)
  env_label = var.metadata.env != null && var.metadata.env != "" ? {
    "planton:environment" = var.metadata.env
  } : {}

  # Merge all labels
  final_labels = merge(
    local.base_labels,
    local.org_label,
    local.env_label,
    var.metadata.labels != null ? var.metadata.labels : {}
  )

  # Extract network ID from literal value or reference
  network_id = try(var.spec.network.value, "")

  # Extract reserved IP ID if provided
  reserved_ip_id = try(var.spec.reserved_ip_id.value, "")

  # Extract instance IDs from the list
  instance_ids = [for inst in var.spec.instance_ids : try(inst.value, "") if try(inst.value, "") != ""]

  # Determine load balancing algorithm based on sticky sessions
  algorithm = var.spec.enable_sticky_sessions ? "ip_hash" : "round_robin"

  # Protocol mapping to lowercase for Civo API
  protocol_map = {
    "http"  = "http"
    "https" = "https"
    "tcp"   = "tcp"
  }

  # Build backends list
  # If instance_tag is specified, use tag-based attachment
  # Otherwise, use explicit instance IDs
  use_tag = var.spec.instance_tag != ""

  # Create backend configurations for each forwarding rule and instance
  backends = local.use_tag ? [
    for rule in var.spec.forwarding_rules : {
      instance_id   = "tag:${var.spec.instance_tag}"
      protocol      = lookup(local.protocol_map, rule.target_protocol, "http")
      source_port   = rule.entry_port
      target_port   = rule.target_port
    }
  ] : flatten([
    for rule in var.spec.forwarding_rules : [
      for instance_id in local.instance_ids : {
        instance_id   = instance_id
        protocol      = lookup(local.protocol_map, rule.target_protocol, "http")
        source_port   = rule.entry_port
        target_port   = rule.target_port
      }
    ]
  ])

  # Health check path (for HTTP/HTTPS health checks)
  health_check_path = try(var.spec.health_check.path, "/")
}

