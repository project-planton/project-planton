locals {
  # Derive a stable resource ID
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base GCP labels
  base_gcp_labels = {
    "resource"      = "true"
    "resource_kind" = "gke-addon-bundle"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if var.metadata.env is non-empty and var.metadata.env.id is non-empty
  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env.id, "") != ""
  ) ? {
    "environment" = var.metadata.env.id
  } : {}

  # Merge base, org, environment labels and add resource_id
  final_gcp_labels = merge(
    local.base_gcp_labels,
    { "resource_id" = local.resource_id },
    local.org_label,
    local.env_label
  )

  # Base Kubernetes labels
  base_kubernetes_labels = {
    "resource"      = "true"
    "resource_kind" = "gke-addon-bundle"
  }

  # Merge base, org, environment labels and add resource_id for Kubernetes
  final_kubernetes_labels = merge(
    local.base_kubernetes_labels,
    { "resource_id" = local.resource_id },
    local.org_label,
    local.env_label
  )
}
