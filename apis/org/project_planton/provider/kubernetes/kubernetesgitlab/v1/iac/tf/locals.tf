###########################
# locals.tf
###########################

locals {
  # Derive a stable resource ID (fall back to name if ID is missing/empty).
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "kubernetes_gitlab"
  }

  # Organization label only if non-empty
  org_label = (
    var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if env is set
  env_label = (
    var.metadata.env != null && var.metadata.env != ""
  ) ? {
    "environment" = var.metadata.env
  } : {}

  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Namespace is the resource_id
  namespace = local.resource_id

  # GitLab service name
  gitlab_service_name = "${var.metadata.name}-gitlab"
  gitlab_service_fqdn = "${local.gitlab_service_name}.${local.namespace}.svc.cluster.local"
  gitlab_port         = 80

  # Ingress configuration
  ingress_is_enabled       = try(var.spec.ingress.enabled, false)
  ingress_external_hostname = try(var.spec.ingress.hostname, null)

  # Certificate issuer: extract domain from hostname
  ingress_cert_cluster_issuer_name = local.ingress_is_enabled && local.ingress_external_hostname != null ? (
    join(".", slice(split(".", local.ingress_external_hostname), 1, length(split(".", local.ingress_external_hostname))))
  ) : ""

  ingress_cert_secret_name = local.resource_id

  # Istio ingress configuration
  istio_ingress_namespace      = "istio-ingress"
  gateway_ingress_class_name   = "istio"
  gateway_external_lb_hostname = "ingress-external.istio-ingress.svc.cluster.local"

  # Port-forward command for debugging
  port_forward_command = "kubectl port-forward -n ${local.namespace} svc/${local.gitlab_service_name} ${local.gitlab_port}:${local.gitlab_port}"
}

