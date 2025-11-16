# Civo Load Balancer Terraform Module
#
# This module creates a Civo Load Balancer with forwarding rules, health checks,
# and backend instance configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    civo = {
      source  = "civo/civo"
      version = "~> 1.1"
    }
  }
}

# Civo Load Balancer Resource
resource "civo_loadbalancer" "this" {
  name       = var.spec.load_balancer_name
  region     = lower(var.spec.region)
  network_id = local.network_id
  algorithm  = local.algorithm

  # Reserved IP (optional)
  reserved_ip_id = local.reserved_ip_id != "" ? local.reserved_ip_id : null

  # Backend configurations
  dynamic "backend" {
    for_each = local.backends
    content {
      instance_id  = backend.value.instance_id
      protocol     = backend.value.protocol
      source_port  = backend.value.source_port
      target_port  = backend.value.target_port
    }
  }

  # Health check path (for HTTP/HTTPS health checks)
  health_check_path = local.health_check_path
}

