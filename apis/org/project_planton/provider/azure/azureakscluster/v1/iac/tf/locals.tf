locals {
  # Resource group name
  resource_group_name = "rg-${var.metadata.name}"

  # Cluster name
  cluster_name = var.metadata.name

  # Determine network plugin
  network_plugin = var.spec.network_plugin == "KUBENET" ? "kubenet" : "azure"

  # Azure AD RBAC enabled
  azure_ad_rbac_enabled = !var.spec.disable_azure_ad_rbac

  # Tags from metadata
  tags = merge(
    var.metadata.labels != null ? var.metadata.labels : {},
    {
      Name        = var.metadata.name
      Environment = var.metadata.env != null ? var.metadata.env : "default"
      ManagedBy   = "terraform"
    }
  )

  # DNS prefix
  dns_prefix = "${var.metadata.name}-dns"
}

