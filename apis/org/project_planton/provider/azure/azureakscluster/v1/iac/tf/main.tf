# Create Resource Group
resource "azurerm_resource_group" "aks" {
  name     = local.resource_group_name
  location = var.spec.region
  tags     = local.tags
}

# Create Azure Kubernetes Service (AKS) Cluster
resource "azurerm_kubernetes_cluster" "aks" {
  name                = local.cluster_name
  location            = azurerm_resource_group.aks.location
  resource_group_name = azurerm_resource_group.aks.name
  dns_prefix          = local.dns_prefix
  kubernetes_version  = var.spec.kubernetes_version

  # Default system node pool
  default_node_pool {
    name                = "system"
    node_count          = 3
    vm_size             = "Standard_D2s_v3"
    vnet_subnet_id      = var.spec.vnet_subnet_id
    enable_auto_scaling = true
    min_count           = 3
    max_count           = 5
    
    # System node pool should be tainted for system workloads only
    node_labels = {
      "node-role" = "system"
    }
  }

  # System-Assigned Managed Identity
  identity {
    type = "SystemAssigned"
  }

  # Network Profile
  network_profile {
    network_plugin    = local.network_plugin
    load_balancer_sku = "standard"
    service_cidr      = "10.0.0.0/16"
    dns_service_ip    = "10.0.0.10"
  }

  # API Server Access Profile
  api_server_access_profile {
    authorized_ip_ranges = var.spec.authorized_ip_ranges
  }

  # Private cluster configuration
  private_cluster_enabled = var.spec.private_cluster_enabled

  # Azure AD RBAC Integration
  dynamic "azure_active_directory_role_based_access_control" {
    for_each = local.azure_ad_rbac_enabled ? [1] : []
    content {
      managed            = true
      azure_rbac_enabled = true
    }
  }

  # OMS Agent (Azure Monitor) addon
  dynamic "oms_agent" {
    for_each = var.spec.log_analytics_workspace_id != "" ? [1] : []
    content {
      log_analytics_workspace_id = var.spec.log_analytics_workspace_id
    }
  }

  tags = local.tags

  lifecycle {
    ignore_changes = [
      # Ignore changes to tags added by Azure or other processes
      tags["LastModified"],
    ]
  }
}

