locals {
  # Derive a stable resource ID (prefer `metadata.id`, fallback to `metadata.name`)
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "openfga_kubernetes"
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

  # Use resource_id as the namespace name
  namespace = local.resource_id

  # Service name (using metadata.name directly)
  kube_service_name = var.metadata.name

  # Internal DNS name for the service
  kube_service_fqdn = "${local.kube_service_name}.${local.namespace}.svc.cluster.local"

  # Handy port-forward command
  kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${local.kube_service_name} 8080:8080"

  # Ingress configuration
  ingress_is_enabled        = try(var.spec.ingress.enabled, false)
  ingress_external_hostname = try(var.spec.ingress.hostname, null)

  # Extract domain from hostname for certificate issuer
  # Example: "openfga.example.com" -> "example.com"
  ingress_cert_cluster_issuer_name = local.ingress_external_hostname != null ? (
    join(".", slice(split(".", local.ingress_external_hostname), 1,
      length(split(".", local.ingress_external_hostname))))
  ) : null

  ingress_cert_secret_name = local.resource_id
}
