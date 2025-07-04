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
    "resource_kind" = "postgres_kubernetes"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if var.metadata.env is non-empty
  env_label = (
  var.metadata.env != null && try(var.metadata.env, "") != ""
  ) ? {
    "environment" = var.metadata.env
  } : {}

  # Merge base, org, and environment labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Labels for Postgres pods
  postgres_pod_selector_labels = {
    "planton.cloud/resource-kind" = "postgres_kubernetes"
    "planton.cloud/resource-id"   = local.resource_id
  }

  # Namespace uses the resource_id
  namespace = local.resource_id

  # Service name
  kube_service_name = "${var.metadata.name}-master"

  # Fully qualified domain name for the service
  kube_service_fqdn = "${local.kube_service_name}.${local.namespace}.svc.cluster.local"

  # Handy port-forward command
  kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${local.kube_service_name} 8080:8080"

  # Safely handle optional ingress values
  ingress_is_enabled = try(var.spec.ingress.is_enabled, false)
  ingress_dns_domain = try(var.spec.ingress.dns_domain, "")

  # External hostname (null if ingress is not enabled or domain is empty)
  ingress_external_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}.${local.ingress_dns_domain}" : null

  # Internal hostname (null if ingress is not enabled or domain is empty)
  ingress_internal_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}-internal.${local.ingress_dns_domain}" : null
}
