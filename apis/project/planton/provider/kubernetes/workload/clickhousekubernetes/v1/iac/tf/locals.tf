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
    "resource_kind" = "clickhouse_kubernetes"
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

  # Use resource_id as the namespace name
  namespace = local.resource_id

  # The service name for ClickHouse (Helm chart overrides "fullname")
  kube_service_name = var.metadata.name

  # Construct labels for the ClickHouse pods
  clickhouse_pod_selector_labels = {
    "app.kubernetes.io/component" = "clickhouse"
    "app.kubernetes.io/instance"  = local.resource_id
    "app.kubernetes.io/name"      = "clickhouse"
  }

  # Safely handle optional ingress values
  ingress_is_enabled = try(var.spec.ingress.is_enabled, false)
  ingress_dns_domain = try(var.spec.ingress.dns_domain, "")

  # External LB hostnames
  ingress_external_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}.${local.ingress_dns_domain}" : null

  ingress_internal_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}-internal.${local.ingress_dns_domain}" : null

  # Cluster configuration
  cluster_is_enabled = try(var.spec.cluster.is_enabled, false)
  shard_count        = local.cluster_is_enabled ? try(var.spec.cluster.shard_count, 1) : 1
  replica_count      = local.cluster_is_enabled ? try(var.spec.cluster.replica_count, 1) : 1
}
