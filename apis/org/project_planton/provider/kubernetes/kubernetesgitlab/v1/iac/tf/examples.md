# Terraform Examples for KubernetesGitlab

Complete Terraform configurations for deploying GitLab on Kubernetes.

---

## Example 1: Basic GitLab Deployment

Minimal GitLab deployment with default settings:

```hcl
terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.23.0"
    }
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

module "gitlab_basic" {
  source = "../"

  metadata = {
    name = "gitlab-basic"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "gitlab-basic"
    container = {
      resources = {
        limits = {
          cpu    = "1000m"
          memory = "1Gi"
        }
        requests = {
          cpu    = "50m"
          memory = "100Mi"
        }
      }
    }
  }
}

output "namespace" {
  value = module.gitlab_basic.namespace
}

output "port_forward_command" {
  value = module.gitlab_basic.port_forward_command
}
```

---

## Example 2: GitLab with Ingress

GitLab deployment with external access via ingress:

```hcl
module "gitlab_with_ingress" {
  source = "../"

  metadata = {
    name = "gitlab-prod"
    env  = "production"
    org  = "my-company"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "gitlab-prod"
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
    ingress = {
      enabled  = true
      hostname = "gitlab.example.com"
    }
  }
}

output "gitlab_url" {
  value = module.gitlab_with_ingress.external_endpoint
}

output "internal_endpoint" {
  value = module.gitlab_with_ingress.internal_endpoint
}
```

---

## Example 3: High-Resource GitLab

GitLab deployment with high resource allocations for large teams:

```hcl
module "gitlab_high_resources" {
  source = "../"

  metadata = {
    name = "gitlab-enterprise"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "gitlab-enterprise"
    container = {
      resources = {
        limits = {
          cpu    = "8000m"
          memory = "16Gi"
        }
        requests = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }
    ingress = {
      enabled  = true
      hostname = "gitlab.company.com"
    }
  }
}
```

---

## Example 4: Development GitLab

Minimal resources for development/testing:

```hcl
module "gitlab_dev" {
  source = "../"

  metadata = {
    name = "gitlab-dev"
    env  = "development"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "gitlab-dev"
    container = {
      resources = {
        limits = {
          cpu    = "500m"
          memory = "512Mi"
        }
        requests = {
          cpu    = "100m"
          memory = "256Mi"
        }
      }
    }
  }
}

output "dev_port_forward" {
  value = module.gitlab_dev.port_forward_command
}
```

---

## Example 5: Multi-Environment Setup

Deploy GitLab to multiple environments:

```hcl
variable "environments" {
  description = "Map of environments"
  type = map(object({
    cpu_limit      = string
    memory_limit   = string
    cpu_request    = string
    memory_request = string
    ingress_enabled = bool
    hostname       = string
  }))
  
  default = {
    dev = {
      cpu_limit       = "1000m"
      memory_limit    = "1Gi"
      cpu_request     = "200m"
      memory_request  = "512Mi"
      ingress_enabled = false
      hostname        = ""
    }
    staging = {
      cpu_limit       = "2000m"
      memory_limit    = "2Gi"
      cpu_request     = "500m"
      memory_request  = "1Gi"
      ingress_enabled = true
      hostname        = "gitlab-staging.example.com"
    }
    prod = {
      cpu_limit       = "4000m"
      memory_limit    = "8Gi"
      cpu_request     = "1000m"
      memory_request  = "2Gi"
      ingress_enabled = true
      hostname        = "gitlab.example.com"
    }
  }
}

module "gitlab" {
  source   = "../"
  for_each = var.environments

  metadata = {
    name = "gitlab-${each.key}"
    env  = each.key
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "gitlab-${each.key}"
    container = {
      resources = {
        limits = {
          cpu    = each.value.cpu_limit
          memory = each.value.memory_limit
        }
        requests = {
          cpu    = each.value.cpu_request
          memory = each.value.memory_request
        }
      }
    }
    ingress = each.value.ingress_enabled ? {
      enabled  = true
      hostname = each.value.hostname
    } : null
  }
}

output "gitlab_endpoints" {
  value = {
    for env, instance in module.gitlab :
    env => instance.external_endpoint != null ? instance.external_endpoint : instance.internal_endpoint
  }
}
```

---

## Example 6: Using Variables

Store configuration in `terraform.tfvars`:

```hcl
# terraform.tfvars
gitlab_name     = "gitlab-main"
environment     = "production"
cpu_limit       = "4000m"
memory_limit    = "8Gi"
cpu_request     = "1000m"
memory_request  = "2Gi"
ingress_enabled = true
hostname        = "gitlab.mycompany.com"
```

Main configuration:

```hcl
# main.tf
variable "gitlab_name" {
  description = "Name for the GitLab deployment"
  type        = string
}

variable "environment" {
  description = "Environment (dev, staging, prod)"
  type        = string
}

variable "cpu_limit" {
  description = "CPU limit"
  type        = string
}

variable "memory_limit" {
  description = "Memory limit"
  type        = string
}

variable "cpu_request" {
  description = "CPU request"
  type        = string
}

variable "memory_request" {
  description = "Memory request"
  type        = string
}

variable "ingress_enabled" {
  description = "Enable ingress"
  type        = bool
}

variable "hostname" {
  description = "Ingress hostname"
  type        = string
}

module "gitlab" {
  source = "../"

  metadata = {
    name = var.gitlab_name
    env  = var.environment
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = var.gitlab_name
    container = {
      resources = {
        limits = {
          cpu    = var.cpu_limit
          memory = var.memory_limit
        }
        requests = {
          cpu    = var.cpu_request
          memory = var.memory_request
        }
      }
    }
    ingress = var.ingress_enabled ? {
      enabled  = true
      hostname = var.hostname
    } : null
  }
}
```

---

## Running the Examples

### Initialize Terraform

```bash
terraform init
```

### Plan the deployment

```bash
terraform plan
```

### Apply the configuration

```bash
terraform apply
```

### Destroy resources

```bash
terraform destroy
```

---

## Testing GitLab Deployment

After deployment, test access:

### Internal Access (Port Forward)

```bash
# Get the port-forward command from outputs
terraform output port_forward_command

# Run the command (example)
kubectl port-forward -n gitlab-prod svc/gitlab-prod-gitlab 80:80

# Access in browser
open http://localhost:80
```

### External Access (if Ingress enabled)

```bash
# Get the external endpoint
terraform output gitlab_url

# Access in browser
open https://gitlab.example.com
```

---

## Production Considerations

For production GitLab deployments, consider:

1. **Use Official Helm Chart**: The official GitLab Helm chart provides complete features
2. **External Databases**: Configure external PostgreSQL and Redis
3. **Object Storage**: Set up S3, GCS, or Azure Blob for artifacts and uploads
4. **GitLab Runner**: Deploy GitLab Runner for CI/CD pipelines
5. **High Availability**: Configure multiple replicas for critical components
6. **Backups**: Implement regular backup strategy
7. **Monitoring**: Enable Prometheus metrics and alerting
8. **SSL Certificates**: Use cert-manager for automatic certificate management

---

## Troubleshooting

### Check Deployment Status

```bash
# List all resources in namespace
kubectl get all -n <namespace>

# Check pod logs
kubectl logs -n <namespace> <pod-name>

# Describe resources
kubectl describe pod -n <namespace> <pod-name>
```

### Access Issues

```bash
# Check service endpoints
kubectl get endpoints -n <namespace>

# Check ingress status
kubectl get ingress -n <namespace>
kubectl describe ingress -n <namespace> <ingress-name>
```

---

For more information, see the [README](README.md).

