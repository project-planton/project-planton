locals {
  # If metadata.id is non-empty, use that. Otherwise use metadata.name.
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "helm_release"
  }

  # Organization label if org is provided
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label if env is provided and env.id is non-empty
  env_label = (
  var.metadata.env != null && try(var.metadata.env, "") != ""
  ) ? {
    "environment" = var.metadata.env
  } : {}

  # Merge all into final_labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # For convenience, we use metadata.id as the namespace name
  namespace_name = local.resource_id

  # The helm release fields for direct reference:
  helm_repo    = var.spec.repo
  helm_chart   = var.spec.name
  helm_version = var.spec.version
  # helm_values  = var.spec.values
}
