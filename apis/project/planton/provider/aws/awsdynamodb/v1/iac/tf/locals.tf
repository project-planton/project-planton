locals {
  #############################################################################
  # resource_id: either metadata.id or metadata.name
  #############################################################################
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  #############################################################################
  # Base labels
  #############################################################################
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "aws_dynamodb"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? { "organization" = var.metadata.org } : {}

  # Environment label only if var.metadata.env is non-empty
  env_label = (
  var.metadata.env != null && try(var.metadata.env, "") != ""
  ) ? { "environment" = var.metadata.env } : {}

  #############################################################################
  # Merge labels
  #############################################################################
  final_labels = merge(local.base_labels, local.org_label, local.env_label)
}
