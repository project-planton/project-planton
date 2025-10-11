# PerconaPostgresqlOperator Terraform Module

## Overview

This Terraform module deploys the Percona Operator for PostgreSQL to a Kubernetes cluster. The operator enables the management of PostgreSQL database clusters within Kubernetes using custom resources.

## Key Features

### Terraform-Based Deployment
- **HCL Configuration**: Uses standard Terraform HCL for infrastructure definition
- **State Management**: Leverages Terraform state for tracking operator deployment
- **Provider Integration**: Integrates with Kubernetes and Helm providers seamlessly

### Operator Management
- **Helm Chart Deployment**: Deploys using the official Percona Helm chart
- **CRD Installation**: Automatically installs PostgreSQL Custom Resource Definitions
- **Resource Configuration**: Configurable CPU and memory resources for the operator pod
- **Namespace Isolation**: Creates dedicated namespace for the operator

### Infrastructure as Code
- **Declarative Configuration**: Define desired state using Terraform syntax
- **Version Control**: Track changes to operator configuration in Git
- **Reproducible Deployments**: Consistently deploy across environments

## Prerequisites

- Terraform >= 1.0
- Kubernetes cluster with appropriate access
- kubectl configured with cluster access
- Helm provider configured

## Module Structure

```
tf/
├── main.tf          # Main resources (namespace, Helm release)
├── provider.tf      # Provider configurations (Kubernetes, Helm)
├── variables.tf     # Input variables
└── README.md        # This file
```

## Resources Created

1. **Kubernetes Namespace**: `percona-postgresql-operator`
2. **Helm Release**: Deploys the Percona Operator for PostgreSQL
3. **CRDs**: PerconaPGCluster and related custom resources

## Input Variables

### metadata
Metadata for the resource, including name and labels.

**Type**: `object`

**Fields**:
- `name` (string, required): Name of the operator deployment
- `id` (string, optional): Unique identifier
- `org` (string, optional): Organization name
- `env` (string, optional): Environment name
- `labels` (map(string), optional): Key-value labels
- `tags` (list(string), optional): List of tags
- `version` (object, optional): Version information

### spec
Specification for the operator deployment.

**Type**: `object`

**Fields**:
- `namespace` (string, optional): Namespace to install the operator (defaults to "percona-postgresql-operator")
- `container` (object, required): Container specifications
  - `resources` (object, required): Resource allocations
    - `limits` (object): Maximum resources
      - `cpu` (string): CPU limit (e.g., "1000m")
      - `memory` (string): Memory limit (e.g., "1Gi")
    - `requests` (object): Guaranteed resources
      - `cpu` (string): CPU request (e.g., "100m")
      - `memory` (string): Memory request (e.g., "256Mi")

## Outputs

### namespace
The Kubernetes namespace where the operator is deployed.

**Type**: `string`

**Value**: `percona-postgresql-operator` (or custom namespace from spec)

## Usage

### Basic Example

```hcl
module "percona_pg_operator" {
  source = "./tf"

  metadata = {
    name = "percona-pg-operator-prod"
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

### With Custom Resources

```hcl
module "percona_pg_operator_large" {
  source = "./tf"

  metadata = {
    name = "percona-pg-operator-large"
    labels = {
      environment = "production"
      team        = "data-platform"
    }
  }

  spec = {
    namespace = "percona-postgresql-operator-prod"
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
}
```

## Deployment

### Initialize Terraform

```bash
terraform init
```

### Plan Changes

```bash
terraform plan -var-file="operator.tfvars"
```

### Apply Configuration

```bash
terraform apply -var-file="operator.tfvars"
```

### Destroy Resources

```bash
terraform destroy -var-file="operator.tfvars"
```

## Verification

After deployment, verify the operator is running:

```bash
# Check operator pod
kubectl get pods -n percona-postgresql-operator

# Check operator logs
kubectl logs -n percona-postgresql-operator -l app.kubernetes.io/name=percona-postgresql-operator

# Verify CRDs
kubectl get crds | grep percona
```

## Configuration

### Helm Chart Configuration

The module configures the Helm chart with the following defaults:

- **Chart Repository**: `https://percona.github.io/percona-helm-charts/`
- **Chart Name**: `pg-operator`
- **Chart Version**: `2.7.0`
- **Timeout**: 300 seconds
- **Atomic**: true (rollback on failure)
- **Cleanup on Fail**: true

### Resource Defaults

Default resource allocations:
- **CPU Request**: 100m
- **Memory Request**: 256Mi
- **CPU Limit**: 1000m
- **Memory Limit**: 1Gi

## Troubleshooting

### Helm Release Failed

Check Helm release status:
```bash
helm list -n percona-postgresql-operator
helm status -n percona-postgresql-operator pg-operator
```

### CRDs Not Installed

Verify CRD installation:
```bash
kubectl get crds | grep percona
```

The Percona Helm chart automatically installs CRDs.

### Resource Limits Too Low

If the operator is resource-constrained, increase the limits in your variables file:

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

## Integration with Other Modules

After deploying the operator, you can deploy PostgreSQL clusters using:
- Terraform PostgreSQL module
- PerconaPGCluster CRDs directly
- PostgreSQLKubernetes Planton resource

## State Management

### Local Backend

For development, use local state:

```hcl
terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}
```

### Remote Backend (Recommended for Production)

For production, use remote state:

```hcl
terraform {
  backend "s3" {
    bucket = "my-terraform-state"
    key    = "percona-postgresql-operator/terraform.tfstate"
    region = "us-west-2"
  }
}
```

## Contributing

Contributions are welcome! Please ensure:
- All variables are properly documented
- Examples are tested and working
- README is updated with new features

## License

This project is licensed under the [MIT License](LICENSE).

