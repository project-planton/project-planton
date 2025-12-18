locals {
  # Use 'metadata.id' if set, otherwise fall back to 'metadata.name'.
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels for all resources
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "nats_kubernetes"
    "resource_name" = var.metadata.name
  }

  # Organization label only if non-empty
  org_label = (
    var.metadata.org != null && var.metadata.org != ""
    ? { "organization" = var.metadata.org }
    : {}
  )

  # Environment label only if non-empty
  env_label = (
    var.metadata.env != null && var.metadata.env != ""
    ? { "environment" = var.metadata.env }
    : {}
  )

  # Merge all labels
  labels = merge(local.base_labels, local.org_label, local.env_label)

  # Get namespace from spec
  namespace = var.spec.namespace

  # Computed resource names to avoid conflicts when multiple instances share a namespace
  # Format: {metadata.name}-{purpose}
  # Users can prefix metadata.name with component type if needed (e.g., "nats-my-bus")
  auth_secret_name         = "${var.metadata.name}-auth"
  no_auth_user_secret_name = "${var.metadata.name}-no-auth-user"
  tls_secret_name          = "${var.metadata.name}-tls"
  external_lb_service_name = "${var.metadata.name}-external-lb"

  # NATS service name (created by Helm chart)
  nats_service_name = "${var.metadata.name}-nats"

  # Internal FQDN for NATS service
  internal_client_url = "nats://${local.nats_service_name}.${local.namespace}.svc.cluster.local:4222"

  # Ingress configuration
  ingress_is_enabled = try(var.spec.ingress.enabled, false)
  ingress_hostname   = try(var.spec.ingress.hostname, null)
}
