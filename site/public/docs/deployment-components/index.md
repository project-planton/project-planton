---
title: "Deployment Components"
description: "Explore available cloud resource deployment components"
icon: "package"
order: 4
---

# Deployment Components

Deployment components are pre-built, battle-tested modules for deploying cloud resources. Each component includes API definitions, IaC modules (both Pulumi and Terraform), and comprehensive documentation.

## What is a Deployment Component?

A deployment component is a self-contained unit that includes:

- **API Definition**: Protocol Buffer schema defining configuration options
- **Pulumi Module**: Infrastructure-as-code implementation in Go/Python/TypeScript
- **Terraform Module**: Alternative IaC implementation in HCL
- **Validation Rules**: Field-level validations in the protobuf schema
- **Documentation**: Auto-generated from protobuf definitions on Buf Schema Registry

## Provider-Specific Components

ProjectPlanton provides **provider-specific components** rather than abstract ones. This design choice preserves the unique capabilities of each cloud provider while maintaining a consistent structure and workflow.

### For Deploying Databases

**PostgreSQL:**
- `PostgresKubernetes` - Deploy to any Kubernetes cluster
- `AwsRdsInstance` - Deploy to AWS RDS as a single instance
- `AwsRdsCluster` - Deploy to AWS RDS Aurora cluster
- `GcpCloudSql` - Deploy to Google Cloud SQL

**Other Databases:**
- `RedisKubernetes` - Deploy Redis to any Kubernetes cluster
- `MongodbKubernetes` - Deploy MongoDB to any Kubernetes cluster
- `MySqlKubernetes` - Deploy MySQL to any Kubernetes cluster

### For Deploying Containerized Applications

- `MicroserviceKubernetes` - Deploy to any Kubernetes cluster
- `AwsEcsService` - Deploy to AWS ECS/Fargate
- `GcpCloudRun` - Deploy to GCP Cloud Run

### For Deploying Open-Source Software

- `KafkaKubernetes` - Deploy Kafka to any Kubernetes cluster
- `TemporalKubernetes` - Deploy Temporal to any Kubernetes cluster
- `AirflowKubernetes` - Deploy Apache Airflow to any Kubernetes cluster

## Example: Kubernetes Deployment

See the [Kubernetes guide](kubernetes) for detailed examples of deploying to Kubernetes clusters.

## Using Deployment Components

### Step 1: Find the Component

Browse the documentation or repository to find the component you need.

### Step 2: Explore the API

Visit the Buf Schema Registry to see:
- Required and optional fields
- Default values
- Validation rules
- Field-level documentation

### Step 3: Write Your Manifest

Create a YAML manifest following the Kubernetes Resource Model:

```yaml
apiVersion: <provider>.project-planton.org/<version>
kind: <ComponentType>
metadata:
  name: <resource-name>
  org: <organization>
  env: <environment>
spec:
  # Provider-specific configuration
```

### Step 4: Deploy

```bash
# With Pulumi
project-planton pulumi up --manifest config.yaml --stack org/project/env

# With Terraform
project-planton tofu apply --manifest config.yaml
```

## Customizing Components

While the default modules work out-of-the-box, platform engineers can:

1. **Fork and customize** the default modules
2. **Rewrite in another language** using auto-generated SDKs
3. **Build entirely custom implementations** while reusing the API definitions

## Contributing Components

Want to add a new deployment component? Check the contribution guidelines in the GitHub repository. The "Forge" workflow automates much of the component creation process.

## Next Steps

- [Kubernetes Deployment Guide](kubernetes) - Learn how to deploy to Kubernetes
- [Getting Started](../getting-started) - Deploy your first resource
- [Architecture](../concepts/architecture) - Understand how components work

