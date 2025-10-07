# AltinityOperatorKubernetes Terraform Module Examples

This document provides practical examples for deploying the Altinity ClickHouse Operator using Terraform.

---

## Example 1: Basic Deployment

### Description

Deploy the operator with default resource allocations.

### Configuration Files

#### `variables.tfvars`

```hcl
metadata = {
  name = "altinity-operator-basic"
}

spec = {
  namespace = "altinity-operator"  # Optional: defaults to "altinity-operator"
  
  container = {
    resources = {
      requests = {
        cpu    = "100m"
        memory = "256Mi"
      }
      limits = {
        cpu    = "1000m"
        memory = "1Gi"
      }
    }
  }
}
```

#### `main.tf`

```hcl
terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.0"
    }
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

module "altinity_operator" {
  source = "../"

  metadata = var.metadata
  spec     = var.spec
}

output "namespace" {
  value = module.altinity_operator.namespace
}
```

### Deploy

```bash
terraform init
terraform plan -var-file="variables.tfvars"
terraform apply -var-file="variables.tfvars"
```

---

## Example 2: Production Deployment with Custom Resources

### Description

Deploy the operator with increased resources for production workloads.

### Configuration

#### `production.tfvars`

```hcl
metadata = {
  name = "altinity-operator-prod"
  labels = {
    environment = "production"
    managed-by  = "terraform"
    team        = "data-platform"
  }
}

spec = {
  namespace = "altinity-operator-prod"
  
  container = {
    resources = {
      requests = {
        cpu    = "200m"
        memory = "512Mi"
      }
      limits = {
        cpu    = "2000m"
        memory = "2Gi"
      }
    }
  }
}
```

### Deploy

```bash
terraform init
terraform plan -var-file="production.tfvars"
terraform apply -var-file="production.tfvars"
```

---

## Example 3: Development Environment

### Description

Minimal resource configuration for development clusters.

### Configuration

#### `dev.tfvars`

```hcl
metadata = {
  name = "altinity-operator-dev"
  env  = "development"
}

spec = {
  namespace = "altinity-operator-dev"
  
  container = {
    resources = {
      requests = {
        cpu    = "50m"
        memory = "128Mi"
      }
      limits = {
        cpu    = "500m"
        memory = "512Mi"
      }
    }
  }
}
```

### Deploy

```bash
terraform init
terraform plan -var-file="dev.tfvars"
terraform apply -var-file="dev.tfvars"
```

---

## Example 4: Multi-Environment Setup

### Description

Deploy the operator to multiple environments using workspaces.

### Directory Structure

```
terraform/
├── main.tf
├── variables.tf
├── dev.tfvars
├── staging.tfvars
└── prod.tfvars
```

### Commands

```bash
# Initialize
terraform init

# Development
terraform workspace new dev
terraform apply -var-file="dev.tfvars"

# Staging
terraform workspace new staging
terraform apply -var-file="staging.tfvars"

# Production
terraform workspace new prod
terraform apply -var-file="prod.tfvars"
```

---

## Example 5: With Remote State

### Description

Use S3 backend for state management in production.

### Configuration

#### `backend.tf`

```hcl
terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "altinity-operator/terraform.tfstate"
    region         = "us-west-2"
    encrypt        = true
    dynamodb_table = "terraform-state-lock"
  }
}
```

#### Deploy

```bash
terraform init -backend-config="backend.tf"
terraform apply -var-file="production.tfvars"
```

---

## Example 6: With Outputs for Integration

### Description

Export operator information for use by other modules.

### Configuration

#### `outputs.tf`

```hcl
output "operator_namespace" {
  description = "Namespace where Altinity operator is deployed"
  value       = module.altinity_operator.namespace
}

output "operator_ready" {
  description = "Indicates operator deployment completion"
  value       = true
}
```

#### Using in Another Module

```hcl
module "clickhouse_cluster" {
  source = "./clickhouse-cluster"

  operator_namespace = module.altinity_operator.namespace
  
  # Wait for operator to be ready
  depends_on = [module.altinity_operator]
}
```

---

## Verification Commands

After applying any example, verify the deployment:

```bash
# Check namespace
kubectl get namespace altinity-operator

# Check operator pod
kubectl get pods -n altinity-operator

# Check Helm release
helm list -n altinity-operator

# Check operator logs
kubectl logs -n altinity-operator -l app.kubernetes.io/name=altinity-clickhouse-operator -f

# Verify CRDs
kubectl get crds | grep clickhouse
```

Expected CRDs:
- `clickhouseinstallations.clickhouse.altinity.com`
- `clickhouseoperatorconfigurations.clickhouse.altinity.com`

---

## Updating Resources

To update operator resources:

1. Modify the tfvars file
2. Run terraform plan to preview changes
3. Apply the changes

```bash
terraform plan -var-file="production.tfvars"
terraform apply -var-file="production.tfvars"
```

---

## Destroying Resources

To remove the operator:

```bash
terraform destroy -var-file="production.tfvars"
```

**Warning**: This will remove the operator but not the CRDs or any ClickHouse clusters managed by it.

---

## Advanced Configuration

### Custom Helm Values

To pass additional Helm values, modify `main.tf`:

```hcl
resource "helm_release" "altinity_operator" {
  # ... existing configuration ...

  # Additional custom values
  set {
    name  = "operator.watchNamespaces"
    value = "clickhouse-prod,clickhouse-dev"
  }

  set {
    name  = "operator.logLevel"
    value = "debug"
  }
}
```

### Using with Terragrunt

#### `terragrunt.hcl`

```hcl
terraform {
  source = "git::https://github.com/myorg/terraform-modules//altinity-operator"
}

inputs = {
  metadata = {
    name = "altinity-operator-${get_env("ENVIRONMENT")}"
  }
  
  spec = {
    container = {
      resources = {
        requests = {
          cpu    = "100m"
          memory = "256Mi"
        }
        limits = {
          cpu    = "1000m"
          memory = "1Gi"
        }
      }
    }
  }
}
```

---

## Troubleshooting

### Issue: Helm Provider Authentication Failed

**Solution**: Ensure kubectl is configured correctly:
```bash
kubectl cluster-info
kubectl auth can-i create namespace
```

### Issue: Insufficient Resources

**Solution**: Increase resource limits in tfvars:
```hcl
spec = {
  container = {
    resources = {
      limits = {
        cpu    = "2000m"
        memory = "2Gi"
      }
    }
  }
}
```

### Issue: CRDs Already Exist

**Solution**: If upgrading, the CRDs might already exist. This is normal and can be ignored.

---

## Next Steps

After successfully deploying the operator:

1. Deploy ClickHouse clusters using ClickHouseInstallation CRDs
2. Configure monitoring and alerting for the operator
3. Set up backup and disaster recovery procedures
4. Review operator logs regularly for any issues

