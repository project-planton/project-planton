locals {
  # Derive a stable resource ID
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "grafana_kubernetes"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? { "organization" = var.metadata.org } : {}

  # Environment label only if var.metadata.env is non-empty
  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env, "") != ""
  ) ? { "environment" = var.metadata.env } : {}

  # Merge base, org, and environment labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Labels for Grafana pods
  grafana_pod_selector_labels = {
    "app.kubernetes.io/name"     = "grafana"
    "app.kubernetes.io/instance" = local.resource_id
  }

  # Use resource_id as the namespace name
  namespace = local.resource_id

  # Service name
  kube_service_name = "${var.metadata.name}-grafana"

  # Internal DNS name for the service
  kube_service_fqdn = "${local.kube_service_name}.${local.namespace}.svc.cluster.local"

  # Handy port-forward command
  kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${local.kube_service_name} 8080:80"

  # Safely handle optional ingress values
  ingress_is_enabled = try(var.spec.ingress.enabled, false)
  
  # Calculate external hostname if ingress is enabled
  ingress_external_hostname = local.ingress_is_enabled && try(var.spec.ingress.dns_domain, "") != "" ? "https://grafana-${var.metadata.name}.${var.spec.ingress.dns_domain}" : ""
  
  # Calculate internal hostname if ingress is enabled
  ingress_internal_hostname = local.ingress_is_enabled && try(var.spec.ingress.dns_domain, "") != "" ? "https://grafana-${var.metadata.name}-internal.${var.spec.ingress.dns_domain}" : ""
}

