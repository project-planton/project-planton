---
title: "Welcome"
description: "ProjectPlanton Documentation - Multi-cloud infrastructure deployment framework"
icon: "docs"
order: 1
---

# Welcome

Welcome to the ProjectPlanton documentation! ProjectPlanton is an open-source framework that provides a unified, declarative approach to deploying infrastructure and applications across cloud providers.

## Getting Started

New to ProjectPlanton? Start here:

- **[Getting Started](getting-started)** - Install the CLI and deploy your first resource
- **[Core Concepts](concepts)** - Understand the architecture and design principles
- **[Deployment Components](deployment-components)** - Explore available cloud resources

## What is ProjectPlanton?

ProjectPlanton brings Kubernetes-style consistency to infrastructure deployments across any cloud provider. It solves a fundamental problem in modern cloud-native development: **the chaos of managing deployments across different clouds, each with their own tools, APIs, and mental models**.

### The Core Promise

**One structure. One workflow. Any cloud.**

Whether you're deploying a PostgreSQL database to AWS RDS, Google Cloud SQL, or a Kubernetes cluster, ProjectPlanton provides the same consistent experience:

1. Write a YAML manifest following the Kubernetes Resource Model
2. Validate it before deployment
3. Deploy using a single CLI command
4. Get back structured outputs

## Quick Example

Deploy Redis to Kubernetes with a simple manifest:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: RedisKubernetes
metadata:
  name: session-store
  org: acme
  env: production
spec:
  container:
    replicas: 3
    resources:
      limits:
        cpu: 500m
        memory: 1Gi
      requests:
        cpu: 250m
        memory: 512Mi
    isPersistenceEnabled: true
    diskSize: 20Gi
```

Deploy with:

```bash
project-planton pulumi up --manifest redis.yaml --stack acme/platform/prod
```

## Key Features

- **Standardized Configuration**: Kubernetes Resource Model for all cloud resources
- **Language Neutrality**: Protocol Buffers enable SDKs in Go, Java, Python, TypeScript
- **Early Validation**: Catch configuration errors before deployment
- **Multi-IaC Support**: Choose between Pulumi or Terraform/OpenTofu
- **Beautiful Documentation**: Auto-generated from protobuf definitions

## Next Steps

- [Install the CLI](getting-started#installation)
- [Deploy your first resource](getting-started#your-first-deployment)
- [Explore deployment components](deployment-components)

