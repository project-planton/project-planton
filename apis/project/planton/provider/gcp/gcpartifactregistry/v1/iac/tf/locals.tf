locals {
  resource_id = (
  var.metadata.id != null && var.metadata.id != ""
  ) ? var.metadata.id : var.metadata.name

  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "gcp_artifact_registry"
  }

  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env.id, "") != ""
  ) ? {
    "environment" = var.metadata.env.id
  } : {}

  final_labels = merge(
    local.base_labels,
    local.org_label,
    local.env_label
  )

  # Direct spec values
  project_id = var.spec.project_id
  region     = var.spec.region
  is_external = try(var.spec.is_external, false)
}
