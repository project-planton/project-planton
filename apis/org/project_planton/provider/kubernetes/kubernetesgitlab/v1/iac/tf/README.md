# KubernetesGitlab Terraform Module

Terraform module for deploying GitLab on Kubernetes clusters.

## Overview

This Terraform module deploys GitLab to Kubernetes clusters. It creates the necessary namespace, service, and optionally an ingress resource for external access.

**Note:** This is a simplified implementation. For production GitLab deployments, we recommend using the official [GitLab Helm Chart](https://docs.gitlab.com/charts/) which provides comprehensive features including PostgreSQL, Redis, object storage integration, and GitLab Runner.

## Prerequisites

### Terraform Providers

This module requires the following providers:

```hcl
terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.23.0"
    }
  }
}
```

### Kubernetes Cluster Access

Configure the Kubernetes provider with access to your cluster:

```hcl
provider "kubernetes" {
  config_path = "~/.kube/config"
}
```

## Usage

### Basic Example

```hcl
module "gitlab" {
  source = "path/to/module"

  metadata = {
    name = "my-gitlab"
  }

  spec = {
    container = {
      resources = {
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
        requests = {
          cpu    = "500m"
          memory = "1Gi"
        }
      }
    }
  }
}
```

### With Ingress

```hcl
module "gitlab_with_ingress" {
  source = "path/to/module"

  metadata = {
    name = "gitlab-prod"
    env  = "production"
  }

  spec = {
    container = {
      resources = {
        limits = {
          cpu    = "4000m"
          memory = "8Gi"
        }
        requests = {
          cpu    = "1000m"
          memory = "2Gi"
        }
      }
    }
    ingress = {
      enabled  = true
      hostname = "gitlab.example.com"
    }
  }
}
```

## Variables

### `metadata` (Required)

```hcl
metadata = {
  name = string           # Required: Resource name
  id = string            # Optional: Resource ID (defaults to name)
  org = string           # Optional: Organization
  env = string           # Optional: Environment
  labels = map(string)   # Optional: Additional labels
  tags = list(string)    # Optional: Tags
  version = object({     # Optional: Version info
    id = string
    message = string
  })
}
```

### `spec` (Required)

```hcl
spec = {
  container = {           # Required
    resources = {
      limits = {
        cpu    = string   # e.g., "2000m"
        memory = string   # e.g., "4Gi"
      }
      requests = {
        cpu    = string   # e.g., "500m"
        memory = string   # e.g., "1Gi"
      }
    }
  }
  
  ingress = {            # Optional
    enabled  = bool      # Enable external access
    hostname = string    # External hostname
  }
}
```

## Outputs

| Name | Description |
|------|-------------|
| `namespace` | Namespace where GitLab is deployed |
| `service_name` | Kubernetes service name |
| `service_fqdn` | Fully qualified domain name of the service |
| `port_forward_command` | Command to port-forward to GitLab |
| `ingress_hostname` | External hostname (if ingress enabled) |
| `internal_endpoint` | Internal Kubernetes endpoint |
| `external_endpoint` | External HTTPS endpoint (if ingress enabled) |

## Example Output Usage

```hcl
output "gitlab_endpoint" {
  value = module.gitlab.external_endpoint != null ? module.gitlab.external_endpoint : module.gitlab.internal_endpoint
}

output "port_forward" {
  value = module.gitlab.port_forward_command
}
```

## Production Deployment

For production GitLab deployments, consider using the official Helm chart with this module as a base:

```hcl
resource "helm_release" "gitlab" {
  name       = var.metadata.name
  namespace  = module.gitlab.namespace
  repository = "https://charts.gitlab.io/"
  chart      = "gitlab"
  version    = "7.0.0"

  values = [
    templatefile("${path.module}/gitlab-values.yaml", {
      hostname  = var.spec.ingress.hostname
      resources = var.spec.container.resources
    })
  ]

  depends_on = [module.gitlab]
}
```

## Resources Created

This module creates the following Kubernetes resources:

- **Namespace**: Isolated namespace for GitLab
- **Service**: ClusterIP service for internal access
- **Ingress** (optional): External access with TLS

## Networking

### Internal Access

GitLab is accessible within the cluster at:

```
http://<service-name>.<namespace>.svc.cluster.local:80
```

### External Access

When ingress is enabled, GitLab is accessible at:

```
https://<ingress-hostname>
```

TLS certificates are automatically provisioned using cert-manager.

## Port Forwarding

For local development and debugging:

```bash
kubectl port-forward -n <namespace> svc/<service-name> 80:80
```

Then access at `http://localhost:80`

## Troubleshooting

### Namespace Issues

List namespaces:

```bash
kubectl get namespaces
```

### Service Issues

Check service status:

```bash
kubectl get svc -n <namespace>
kubectl describe svc <service-name> -n <namespace>
```

### Ingress Issues

Check ingress configuration:

```bash
kubectl get ingress -n <namespace>
kubectl describe ingress <ingress-name> -n <namespace>
```

Verify cert-manager certificate:

```bash
kubectl get certificate -n <namespace>
```

## Limitations

This is a simplified implementation focused on basic deployment. For production use:

- Use the official GitLab Helm chart for complete features
- Configure external PostgreSQL and Redis
- Set up object storage (S3, GCS, or Azure Blob)
- Configure GitLab Runner for CI/CD
- Enable monitoring and logging
- Configure backups and disaster recovery

## Reference

- [GitLab Charts Documentation](https://docs.gitlab.com/charts/)
- [GitLab Kubernetes Deployment Guide](https://docs.gitlab.com/ee/install/kubernetes/)
- [Terraform Kubernetes Provider](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs)

