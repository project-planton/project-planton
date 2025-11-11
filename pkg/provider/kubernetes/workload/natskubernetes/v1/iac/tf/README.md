# Terraform Module for NATS on Kubernetes

This Terraform module provides infrastructure-as-code for deploying NATS clusters on Kubernetes. It enables simplified deployment and management of NATS messaging infrastructure with minimal configuration.

## Key Features

### 1. **Namespace Management**
Automatically creates and manages a dedicated Kubernetes namespace for the NATS cluster, ensuring isolation from other workloads.

### 2. **Resource Labeling**
Applies consistent labels to all resources for tracking, organization, and integration with monitoring systems.

### 3. **Ingress Support**
Supports external access via LoadBalancer services with DNS hostname configuration through external-dns annotations.

### 4. **Flexible Configuration**
Allows customization of:
- Server replicas
- CPU and memory resources
- Disk size for JetStream persistence
- Authentication schemes (Bearer Token, Basic Auth)
- TLS encryption
- External ingress

## Usage

### Prerequisites

* **Terraform**: Version 0.13 or higher
* **Kubernetes Cluster**: Access to a Kubernetes cluster with valid credentials
* **Planton CLI**: For executing deployment commands

### Example Commands

Initialize Terraform:
```shell
project-planton tofu init --manifest hack/manifest.yaml --backend-type s3 \
  --backend-config="bucket=planton-cloud-tf-state-backend" \
  --backend-config="dynamodb_table=planton-cloud-tf-state-backend-lock" \
  --backend-config="region=ap-south-2" \
  --backend-config="key=kubernetes-stacks/nats-basic.tfstate"
```

Plan the deployment:
```shell
project-planton tofu plan --manifest hack/manifest.yaml
```

Apply the configuration:
```shell
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve
```

## Outputs

After successful deployment, the module provides:

* **namespace**: Kubernetes namespace where NATS is deployed
* **internal_client_url**: Internal DNS URL for NATS clients (e.g., `nats://nats-basic-nats.nats-basic.svc.cluster.local:4222`)
* **external_hostname**: External hostname if ingress is enabled

## Configuration

The module accepts configuration through a YAML manifest matching the `NatsKubernetes` protobuf specification. See `hack/manifest.yaml` for a complete example.

## Benefits

* **Infrastructure-as-Code**: Version-controlled, repeatable deployments
* **Simplified Management**: Automated namespace and resource creation
* **Flexibility**: Easy customization for various deployment scenarios
* **Consistency**: Standardized labeling and naming conventions
