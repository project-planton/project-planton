# Azure AKS Cluster Examples

This document provides comprehensive examples for deploying Azure Kubernetes Service (AKS) clusters using the AzureAksCluster API resource.

## Table of Contents

- [Basic Example](#basic-example)
- [Production Example](#production-example)
- [Private Cluster Example](#private-cluster-example)
- [Development/Test Example](#developmenttest-example)
- [With Monitoring Example](#with-monitoring-example)
- [Using Foreign Key Reference](#using-foreign-key-reference)

---

## Basic Example

The simplest AKS cluster with minimal required configuration.

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: basic-aks-cluster
spec:
  region: eastus
  vnetSubnetId:
    value: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-network/providers/Microsoft.Network/virtualNetworks/app-vnet/subnets/aks-subnet
```

**Deploy:**
```shell
planton apply -f basic-example.yaml
```

**What you get:**
- AKS cluster in East US region
- Azure CNI networking (default)
- Kubernetes version 1.30 (default)
- Public cluster with Azure AD RBAC enabled
- System node pool with 3 nodes (auto-scaling to 5)
- System-Assigned Managed Identity

---

## Production Example

Production-ready configuration with security and monitoring.

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: prod-app-aks
spec:
  region: eastus
  
  # VNet integration
  vnetSubnetId:
    value: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-network/providers/Microsoft.Network/virtualNetworks/prod-vnet/subnets/aks-nodes
  
  # Pin Kubernetes version for production stability
  kubernetesVersion: "1.30"
  
  # Use Azure CNI for production
  networkPlugin: AZURE_CNI
  
  # Restrict API server access to specific networks
  privateClusterEnabled: false
  authorizedIpRanges:
    - "203.0.113.0/24"   # Corporate office network
    - "198.51.100.0/24"  # CI/CD agents
    - "192.0.2.0/24"     # Operations team VPN
  
  # Enable Azure AD RBAC for centralized access control
  disableAzureAdRbac: false
  
  # Enable monitoring with Log Analytics
  logAnalyticsWorkspaceId: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-monitoring/providers/Microsoft.OperationalInsights/workspaces/prod-logs
```

**Deploy:**
```shell
planton apply -f production-example.yaml
```

**What you get:**
- Production-grade AKS cluster
- Restricted API server access (authorized IPs only)
- Azure AD RBAC integration
- Container Insights monitoring
- Pinned Kubernetes version (prevents unwanted upgrades)

---

## Private Cluster Example

Fully private AKS cluster with no public endpoint.

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: private-aks-cluster
spec:
  region: westus2
  
  # VNet subnet for nodes
  vnetSubnetId:
    value: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-network/providers/Microsoft.Network/virtualNetworks/secure-vnet/subnets/aks-private-subnet
  
  # Private cluster configuration
  privateClusterEnabled: true
  
  # No authorized IP ranges needed for private cluster (no public endpoint)
  authorizedIpRanges: []
  
  # Azure CNI recommended for private clusters
  networkPlugin: AZURE_CNI
  
  # Kubernetes version
  kubernetesVersion: "1.30"
  
  # Enable Azure AD RBAC
  disableAzureAdRbac: false
  
  # Monitoring
  logAnalyticsWorkspaceId: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-monitoring/providers/Microsoft.OperationalInsights/workspaces/private-logs
```

**Deploy:**
```shell
planton apply -f private-example.yaml
```

**What you get:**
- Private AKS cluster (no public API endpoint)
- API server only accessible from within VNet
- Ideal for highly secure environments
- Requires VPN or bastion host for kubectl access

**Access:**
```shell
# From a VM in the same VNet or via VPN
az aks get-credentials --resource-group rg-private-aks-cluster --name private-aks-cluster
kubectl get nodes
```

---

## Development/Test Example

Simplified configuration for development and testing environments.

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: dev-aks-cluster
spec:
  region: eastus
  
  # VNet subnet
  vnetSubnetId:
    value: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-dev-network/providers/Microsoft.Network/virtualNetworks/dev-vnet/subnets/aks-dev
  
  # Use default Kubernetes version
  kubernetesVersion: "1.30"
  
  # Public cluster for easier access during development
  privateClusterEnabled: false
  
  # No IP restrictions for development (not recommended for production!)
  authorizedIpRanges: []
  
  # Azure AD RBAC optional for dev
  disableAzureAdRbac: true
  
  # Skip monitoring to reduce costs in dev
  logAnalyticsWorkspaceId: ""
```

**Deploy:**
```shell
planton apply -f dev-example.yaml
```

**What you get:**
- Developer-friendly AKS cluster
- Public API access from anywhere
- No Azure AD RBAC complexity
- Cost-optimized (no monitoring addon)

**⚠️ Warning:** This configuration is for development only. Never use in production!

---

## With Monitoring Example

AKS cluster with comprehensive Azure Monitor integration.

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: monitored-aks-cluster
spec:
  region: centralus
  
  # VNet subnet
  vnetSubnetId:
    value: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-network/providers/Microsoft.Network/virtualNetworks/app-vnet/subnets/aks-subnet
  
  # Standard configuration
  kubernetesVersion: "1.30"
  networkPlugin: AZURE_CNI
  privateClusterEnabled: false
  disableAzureAdRbac: false
  
  # Restrict API access
  authorizedIpRanges:
    - "203.0.113.0/24"
  
  # Enable Container Insights monitoring
  logAnalyticsWorkspaceId: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-monitoring/providers/Microsoft.OperationalInsights/workspaces/aks-monitoring-workspace
```

**Deploy:**
```shell
planton apply -f monitoring-example.yaml
```

**What you get:**
- Container Insights enabled
- Container logs streamed to Log Analytics
- Performance metrics collection
- Kubernetes event monitoring
- Integration with Azure Monitor dashboards

**View Metrics:**
```shell
# In Azure Portal:
# Navigate to: AKS Cluster -> Monitoring -> Insights
# Or query Log Analytics with KQL:
ContainerLog | where TimeGenerated > ago(1h) | limit 100
```

---

## Using Foreign Key Reference

Reference an AzureVpc resource to automatically get the subnet ID.

```yaml
# First, create the VNet
apiVersion: azure.project-planton.org/v1
kind: AzureVpc
metadata:
  name: my-azure-vnet
spec:
  region: eastus
  cidrBlock: 10.0.0.0/16
  # ... other VNet configuration

---

# Then create AKS cluster referencing the VNet
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: aks-with-ref
spec:
  region: eastus
  
  # Reference the VNet's nodes subnet output
  vnetSubnetId:
    ref:
      kind: AzureVpc
      name: my-azure-vnet
      path: status.outputs.nodes_subnet_id
  
  kubernetesVersion: "1.30"
  networkPlugin: AZURE_CNI
  disableAzureAdRbac: false
```

**Deploy:**
```shell
# Deploy both resources
planton apply -f azure-vpc.yaml
planton apply -f aks-with-ref.yaml
```

**What you get:**
- Automatic subnet ID resolution
- Type-safe foreign key references
- Simplified configuration management
- Clear dependencies between resources

---

## Deployment Tips

### Getting Credentials

After deployment, get cluster credentials:

```shell
az aks get-credentials --resource-group rg-<cluster-name> --name <cluster-name>
```

### Verify Deployment

Check cluster status:

```shell
kubectl get nodes
kubectl get pods -n kube-system
kubectl cluster-info
```

### Common Configurations

**Recommended for Production:**
- ✅ `kubernetesVersion`: Pinned to specific version
- ✅ `networkPlugin`: `AZURE_CNI`
- ✅ `privateClusterEnabled`: `true` or `authorizedIpRanges` configured
- ✅ `disableAzureAdRbac`: `false` (enabled)
- ✅ `logAnalyticsWorkspaceId`: Configured

**Acceptable for Development:**
- ⚠️ `kubernetesVersion`: Default (latest supported)
- ⚠️ `privateClusterEnabled`: `false` with no IP restrictions
- ⚠️ `disableAzureAdRbac`: `true` (disabled)
- ⚠️ `logAnalyticsWorkspaceId`: Empty (no monitoring)

---

## Next Steps

After deploying your AKS cluster:

1. **Configure kubectl**: Get credentials and verify connectivity
2. **Install add-ons**: Deploy ingress controller, cert-manager, etc.
3. **Deploy workloads**: Start deploying your applications
4. **Set up monitoring**: Configure alerts and dashboards in Azure Monitor
5. **Implement RBAC**: Assign Azure AD users/groups to Kubernetes roles

## Support

For issues or questions:
- Check the [main README](./README.md) for component overview
- Review the [research documentation](./docs/README.md) for architecture details
- Consult the [Azure AKS Documentation](https://docs.microsoft.com/en-us/azure/aks/)

