locals {
  # resource name and tags
  resource_name = coalesce(try(var.metadata.name, null), "aws-eks-cluster")
  tags          = merge({
    "Name" = local.resource_name
  }, try(var.metadata.labels, {}))

  # flatten foreign keys
  subnet_ids = [for s in try(var.spec.subnet_ids, []) : coalesce(try(s.value, null), try(s.value_from.name, null))]
  cluster_role_arn = coalesce(
    try(var.spec.cluster_role_arn.value, null),
    try(var.spec.cluster_role_arn.value_from.name, null)
  )
  kms_key_arn = coalesce(
    try(var.spec.kms_key_arn.value, null),
    try(var.spec.kms_key_arn.value_from.name, null)
  )

  # feature flags
  disable_public_endpoint   = try(var.spec.disable_public_endpoint, false)
  enable_control_plane_logs = try(var.spec.enable_control_plane_logs, false)
  public_access_cidrs       = try(var.spec.public_access_cidrs, [])

  # logs list per AWS
  cluster_log_types = local.enable_control_plane_logs ? [
    "api", "audit", "authenticator", "controllerManager", "scheduler"
  ] : []
}


