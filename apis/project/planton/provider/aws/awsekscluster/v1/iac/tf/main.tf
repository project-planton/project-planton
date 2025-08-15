resource "aws_eks_cluster" "this" {
  name     = local.resource_name
  role_arn = local.cluster_role_arn

  version = try(var.spec.version, null)

  vpc_config {
    subnet_ids              = local.subnet_ids
    endpoint_public_access  = !local.disable_public_endpoint
    endpoint_private_access = local.disable_public_endpoint
    public_access_cidrs     = local.public_access_cidrs
  }

  dynamic "encryption_config" {
    for_each = local.kms_key_arn != null && local.kms_key_arn != "" ? [1] : []
    content {
      provider {
        key_arn = local.kms_key_arn
      }
      resources = ["secrets"]
    }
  }

  enabled_cluster_log_types = local.cluster_log_types

  tags = local.tags
}


