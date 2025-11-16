# Terraform Module to Deploy Argo CD on Kubernetes

This Terraform module deploys Argo CD on a Kubernetes cluster using the official Helm chart from the Argo Project.

## Prerequisites

- Kubernetes cluster with sufficient resources
- Helm provider configured
- Kubernetes provider configured
- Project Planton CLI installed

## Usage

### Initialize Terraform

```shell
project-planton tofu init --manifest hack/manifest.yaml --backend-type s3 \
  --backend-config="bucket=planton-cloud-tf-state-backend" \
  --backend-config="dynamodb_table=planton-cloud-tf-state-backend-lock" \
  --backend-config="region=us-east-1" \
  --backend-config="key=kubernetes-stacks/test-argocd.tfstate"
```

### Plan Deployment

```shell
project-planton tofu plan --manifest hack/manifest.yaml
```

### Apply Deployment

```shell
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
```

### Destroy Deployment

```shell
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

## What Gets Deployed

This module deploys the following resources:

1. **Kubernetes Namespace**: A dedicated namespace for Argo CD (prefixed with `argo-`)
2. **Argo CD Helm Release**: The official argo-cd chart from https://argoproj.github.io/argo-helm
3. **Core Components**:
   - API Server (with configured resource limits)
   - Application Controller (with configured resource limits)
   - Repo Server (with configured resource limits)
   - Redis (for caching and session storage)

## Configuration

The module reads configuration from the `KubernetesArgocd` manifest and applies it to the Helm chart values:

- **Container Resources**: CPU and memory limits/requests for all components
- **Ingress**: Optional ingress configuration for external access
- **Labels**: Resource labels for organization and tracking

## Outputs

The module exports the following outputs:

- `namespace`: The Kubernetes namespace where Argo CD is deployed
- `service`: The Kubernetes service name for the Argo CD server
- `port_forward_command`: Command to set up local port-forwarding
- `kube_endpoint`: Internal cluster endpoint
- `external_hostname`: External ingress hostname (if ingress is enabled)
- `internal_hostname`: Internal ingress hostname (if ingress is enabled)

## Accessing Argo CD

### Via Port-Forward (Local Development)

Use the `port_forward_command` output:

```shell
kubectl port-forward -n <namespace> service/<service-name> 8080:80
```

Then access Argo CD at http://localhost:8080

### Via Ingress (Production)

If ingress is enabled, access Argo CD at the `external_hostname` output.

## Default Admin Credentials

The initial admin password is auto-generated. Retrieve it with:

```shell
kubectl -n <namespace> get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

**Important**: Change the admin password and configure SSO for production use.

## Notes

- The module uses Helm chart version 7.7.12 by default
- Resource requests and limits are configurable via the manifest
- The deployment is atomic - it will rollback on failure
- Redis is deployed with minimal resources; scale up for production

