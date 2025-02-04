locals {
  # Use 'metadata.id' if set, otherwise fall back to 'metadata.name'.
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "solr_kubernetes"
  }

  # Organization label only if non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if env and env.id are provided
  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env.id, "") != ""
  ) ? {
    "environment" = var.metadata.env.id
  } : {}

  # Merge all labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Namespace is resource_id
  namespace = local.resource_id

  # Solr service name
  kube_service_name = "${var.metadata.name}-solrcloud-common"

  # Internal FQDN
  kube_service_fqdn = "${local.kube_service_name}.${local.namespace}.svc.cluster.local"

  # Port-forward command
  kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${var.metadata.name} 8080:8080"

  # Ingress fields using try(...) to avoid errors if 'spec.ingress' is null.
  ingress_is_enabled = try(var.spec.ingress.is_enabled, false)
  ingress_dns_domain = try(var.spec.ingress.dns_domain, "")

  # External and internal hostnames
  ingress_external_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}.${local.ingress_dns_domain}" : null

  ingress_internal_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}-internal.${local.ingress_dns_domain}" : null

  # Combine hostnames into a list for certificate usage if both are set
  ingress_hostnames = (
  local.ingress_external_hostname != null && local.ingress_internal_hostname != null
  ) ? [
    local.ingress_external_hostname,
    local.ingress_internal_hostname
  ] : []

  # Certificate references
  ingress_cert_cluster_issuer_name = local.ingress_dns_domain
  ingress_cert_secret_name         = "cert-${local.resource_id}"

  # Hardcode or customize these as needed, matching the Pulumi vars struct.
  # If you prefer, you can turn these into variables with variable "..." blocks.
  istio_ingress_namespace      = "istio-ingress"
  gateway_external_lb_hostname = "ingress-external.istio-ingress.svc.cluster.local"
  gateway_ingress_class_name   = "istio"
}
