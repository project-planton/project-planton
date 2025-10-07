# ClickHouse Kubernetes Terraform Module

## Overview

This Terraform module deploys ClickHouse on Kubernetes using the Bitnami Helm chart. It provides a declarative way to manage ClickHouse clusters with support for persistence, clustering, and external access.

## Key Features

- **Namespace Management**: Automatically creates and manages a dedicated Kubernetes namespace
- **Helm Chart Deployment**: Uses Bitnami ClickHouse Helm chart with customizable values
- **Persistence**: Configurable persistent volumes for data storage
- **Clustering**: Optional sharding and replication for distributed deployments
- **Security**: Auto-generated passwords stored in Kubernetes Secrets
- **Ingress**: Optional LoadBalancer service for external access
- **Bitnami Legacy Support**: Configured to use `docker.io/bitnamilegacy` registry

## Important: Docker Image Registry

**⚠️ Bitnami Registry Changes (September 2025)**

This module uses `docker.io/bitnamilegacy` registry by default due to Bitnami discontinuing free Docker Hub images. The legacy images receive no updates or security patches but provide a temporary migration solution.

See: https://github.com/bitnami/containers/issues/83267

## Usage

### Basic Example

```hcl
module "clickhouse" {
  source = "./path/to/this/module"

  metadata = {
    name = "my-clickhouse"
    org  = "my-org"
    env  = "production"
  }

  spec = {
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

### With Clustering

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
- Helm 3.x

## Providers

| Name | Version |
|------|---------|
| kubernetes | ~> 2.35 |
| helm | ~> 2.9 |
| random | ~> 3.5 |

## Resources Created

- `kubernetes_namespace_v1.clickhouse_namespace` - Dedicated namespace
- `random_password.clickhouse_password` - Auto-generated password
- `kubernetes_secret_v1.clickhouse_password` - Password secret
- `helm_release.clickhouse` - ClickHouse Helm chart
- `kubernetes_service_v1.ingress_external_lb` - LoadBalancer (if ingress enabled)

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
