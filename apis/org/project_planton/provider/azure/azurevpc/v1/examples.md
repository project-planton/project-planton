# Azure VPC Examples

## Minimal Configuration

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureVpc
metadata:
  name: dev-vpc
spec:
  address_space_cidr: "10.10.0.0/16"
  nodes_subnet_cidr: "10.10.0.0/18"
```

## Development Environment

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureVpc
metadata:
  name: dev-vpc
  env: development
spec:
  address_space_cidr: "10.10.0.0/16"
  nodes_subnet_cidr: "10.10.0.0/18"
  tags:
    purpose: "development"
```

## Production with NAT Gateway

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureVpc
metadata:
  name: prod-vpc
  org: mycompany
  env: production
spec:
  address_space_cidr: "10.0.0.0/16"
  nodes_subnet_cidr: "10.0.0.0/18"
  is_nat_gateway_enabled: true
  tags:
    cost-center: "infrastructure"
    compliance: "soc2"
```

## With Private DNS Zone Links

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureVpc
metadata:
  name: prod-vpc-with-dns
  org: mycompany
  env: production
spec:
  address_space_cidr: "10.0.0.0/16"
  nodes_subnet_cidr: "10.0.0.0/18"
  is_nat_gateway_enabled: true
  dns_private_zone_links:
    - "/subscriptions/{sub-id}/resourceGroups/{rg}/providers/Microsoft.Network/privateDnsZones/privatelink.database.windows.net"
    - "/subscriptions/{sub-id}/resourceGroups/{rg}/providers/Microsoft.Network/privateDnsZones/privatelink.blob.core.windows.net"
  tags:
    environment: "production"
```

**Use Cases**:
- Basic VNet for development/testing
- Production VNet with NAT Gateway for reliable outbound connectivity
- VNet with Private DNS zones for Azure PaaS services (SQL, Storage)
- Multi-environment network isolation

**NAT Gateway Benefits**:
- Eliminates SNAT port exhaustion
- Provides predictable egress IPs for firewall allow-listing
- Scales to 1M+ ports with public IP prefixes
- Automatic precedence over other outbound methods

