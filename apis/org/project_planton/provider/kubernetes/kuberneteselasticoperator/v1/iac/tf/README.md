# Kubernetes Elastic Operator - Terraform Module

This Terraform module deploys the Elastic Cloud on Kubernetes (ECK) operator using the official Elastic Helm chart.

## Overview

The module installs ECK operator version 2.14.0 in the `elastic-system` namespace. The operator extends Kubernetes with Custom Resource Definitions (CRDs) for managing Elasticsearch, Kibana, APM Server, and other Elastic Stack components.

## Prerequisites

- Terraform 1.0+
- Kubernetes cluster with kubectl access
- Helm provider configured
- Kubernetes provider configured

## Usage

### Basic Example

```hcl
module "eck_operator" {
  source = "./path/to/kubernetes-elastic-operator/iac/tf"

  metadata = {
    name = "eck-operator"
    id   = "eck-op-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    container = {
      resources = {
        requests = {
          cpu    = "50m"
          memory = "100Mi"
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

### High-Availability Production

```hcl
module "eck_operator_ha" {
  source = "./path/to/kubernetes-elastic-operator/iac/tf"

  metadata = {
    name = "eck-operator"
    id   = "eck-op-prod"
    org  = "platform"
    env  = "production"
  }

  spec = {
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

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| metadata | Resource metadata including name, id, org, env | object | n/a | yes |
| spec | ECK operator specification including container resources | object | n/a | yes |

### Metadata Object

| Field | Description | Type | Required |
|-------|-------------|------|----------|
| name | Resource name | string | yes |
| id | Unique resource identifier | string | no |
| org | Organization name | string | no |
| env | Environment (dev/staging/prod) | string | no |
| labels | Additional labels | map(string) | no |
| tags | Resource tags | list(string) | no |

### Spec Object

| Field | Description | Type | Required |
|-------|-------------|------|----------|
| container.resources | Container resource limits and requests | object | yes |

## Outputs

| Name | Description |
|------|-------------|
| namespace | Kubernetes namespace where ECK operator is deployed |
| helm_release_name | Name of the Helm release |
| operator_version | ECK operator version deployed |

## Resources Created

- **kubernetes_namespace.elastic_system**: Dedicated namespace for ECK operator
- **helm_release.eck_operator**: Helm release for ECK operator chart

## Module Constants

The module uses these constants (defined in `locals.tf`):

- **Namespace**: `elastic-system`
- **Helm Chart**: `eck-operator`
- **Helm Repository**: `https://helm.elastic.co`
- **Chart Version**: `2.14.0`

## Terraform Commands

### Initialize

```bash
cd iac/tf
terraform init
```

### Plan

```bash
terraform plan
```

### Apply

```bash
terraform apply
```

### Destroy

```bash
terraform destroy
```

> **Warning**: Destroying the operator does NOT remove Elastic Stack resources it manages. Delete Elasticsearch, Kibana, and other custom resources manually before destroying the operator.

## Verification

After deployment, verify the ECK operator:

```bash
# Check operator pod
kubectl get pods -n elastic-system

# Verify CRDs
kubectl get crds | grep elastic

# View operator logs
kubectl logs -n elastic-system -l control-plane=elastic-operator
```

Expected CRDs:
- elasticsearch.k8s.elastic.co
- kibana.k8s.elastic.co
- apmserver.k8s.elastic.co
- enterprisesearch.k8s.elastic.co
- beat.k8s.elastic.co
- agent.k8s.elastic.co
- logstash.k8s.elastic.co

## Upgrading ECK Version

To upgrade the ECK operator:

1. Update `helm_chart_version` in `locals.tf`
2. Run `terraform plan` to review changes
3. Run `terraform apply` to perform upgrade

## Troubleshooting

### Operator Pod Not Starting

Check resource availability:
```bash
kubectl describe pod -n elastic-system -l control-plane=elastic-operator
kubectl top nodes
```

### CRDs Not Installing

Verify Helm release:
```bash
helm list -n elastic-system
helm status -n elastic-system eck-operator
```

### Permission Issues

Check RBAC resources:
```bash
kubectl get clusterrole elastic-operator
kubectl get clusterrolebinding elastic-operator
kubectl get serviceaccount -n elastic-system elastic-operator
```

## Examples

For complete usage examples, see [examples.md](examples.md).

## References

- [Terraform Helm Provider](https://registry.terraform.io/providers/hashicorp/helm/latest/docs)
- [Terraform Kubernetes Provider](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs)
- [ECK Documentation](https://www.elastic.co/guide/en/cloud-on-k8s/current/index.html)
- [ECK Helm Chart](https://github.com/elastic/cloud-on-k8s/tree/main/deploy/eck-operator)

