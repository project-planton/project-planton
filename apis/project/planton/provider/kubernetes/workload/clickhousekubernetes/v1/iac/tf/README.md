# ClickHouse Kubernetes Terraform Module

## Overview

This Terraform module deploys production-grade ClickHouse clusters on Kubernetes using the **Altinity ClickHouse Operator**. The operator provides enterprise-level features including automated upgrades, scaling, backup, and recovery.

## Key Features

- **Operator-Based Deployment**: Uses the Altinity ClickHouse Operator for production-grade cluster management
- **CRD-Native**: Deploys ClickHouseInstallation custom resources that the operator reconciles
- **Namespace Management**: Automatically creates and manages a dedicated Kubernetes namespace
- **Persistence**: Configurable persistent volumes for data storage with flexible sizing
- **Clustering**: Optional sharding and replication for distributed deployments
- **ZooKeeper Integration**: Auto-managed ZooKeeper for cluster coordination, or configure external ZooKeeper
- **Security**: Auto-generated passwords stored in Kubernetes Secrets
- **Ingress**: Optional LoadBalancer service for external access
- **Version Control**: Pin specific ClickHouse versions for stability
- **Production-Ready**: Leverages official ClickHouse images from clickhouse.com

## Prerequisites

The **Altinity ClickHouse Operator** must be installed on your Kubernetes cluster before using this module. Install the operator using the `ClickhouseOperatorKubernetes` module.

## Usage

### Basic Standalone Example

```hcl
module "clickhouse" {
  source = "./path/to/this/module"

  metadata = {
    name = "my-clickhouse"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    cluster_name = "production-analytics"
    version      = "24.8"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "50Gi"
      resources = {
        requests = {
          cpu    = "500m"
          memory = "1Gi"
        }
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }
    
    ingress = {
      is_enabled = true
      dns_domain = "example.com"
    }
  }
}
```

### With Distributed Clustering

```hcl
module "clickhouse_cluster" {
  source = "./path/to/this/module"

  metadata = {
    name = "clickhouse-cluster"
  }

  spec = {
    container = {
      replicas               = 3
      is_persistence_enabled = true
      disk_size              = "100Gi"
      resources = {
        requests = {
          cpu    = "1000m"
          memory = "2Gi"
        }
        limits = {
          cpu    = "4000m"
          memory = "8Gi"
        }
      }
    }
    cluster = {
      is_enabled    = true
      shard_count   = 3
      replica_count = 2
    }
  }
}
```

## Requirements

- Terraform >= 1.0
- Kubernetes cluster with kubectl access
- **Altinity ClickHouse Operator** installed on the cluster

## Providers

| Name | Version |
|------|---------|
| kubernetes | ~> 2.35 |
| random | ~> 3.5 |

## Resources Created

- `kubernetes_namespace_v1.clickhouse_namespace` - Dedicated namespace
- `random_password.clickhouse_password` - Auto-generated password
- `kubernetes_secret_v1.clickhouse_password` - Password secret
- `kubernetes_manifest.clickhouse_installation` - ClickHouseInstallation CRD
- `kubernetes_service_v1.ingress_external_lb` - LoadBalancer (if ingress enabled)

The operator then reconciles the ClickHouseInstallation to create:
- StatefulSets for ClickHouse pods
- Services for cluster communication
- ConfigMaps with ClickHouse configuration
- ZooKeeper (if clustering is enabled)

## Outputs

| Name | Description |
|------|-------------|
| namespace | The namespace in which ClickHouse is deployed |
| service | The base name of the ClickHouse service |
| kube_endpoint | Internal DNS name of the ClickHouse service |
| port_forward_command | Command to port-forward to ClickHouse |
| external_hostname | External hostname (if ingress enabled) |
| internal_hostname | Internal hostname (if ingress enabled) |
| username | ClickHouse username |
| password_secret_name | Name of the password Secret |
| password_secret_key | Key within the password Secret |

## Testing

```bash
# Initialize
terraform init

# Plan
terraform plan -var-file=hack/terraform.tfvars

# Apply
terraform apply -var-file=hack/terraform.tfvars -auto-approve

# Destroy
terraform destroy -var-file=hack/terraform.tfvars -auto-approve
```

## License

Apache 2.0
