---
title: "Documentation"
description: "Comprehensive guides for ProjectPlanton - the open-source multi-cloud infrastructure framework"
icon: "📚"
order: 1
---

# Welcome to ProjectPlanton Documentation

ProjectPlanton is an open-source multi-cloud infrastructure framework that lets you author KRM-style YAML manifests once, validate them with Protobuf + Buf, and deploy with Pulumi or OpenTofu.

## Getting Started

New to ProjectPlanton? Start here:

- Install the CLI via Homebrew
- Validate your first manifest
- Deploy to your cloud provider or Kubernetes cluster

## Deployment Components

Browse deployment components by cloud provider:

- [AWS](/docs/aws) - 22 components
- [GCP](/docs/gcp) - 5 components
- [Azure](/docs/azure) - 7 components
- [Kubernetes](/docs/kubernetes) - Coming soon
- [Cloudflare](/docs/cloudflare) - 7 components
- [Civo](/docs/civo) - 12 components
- [DigitalOcean](/docs/digitalocean) - 14 components
- [Atlas](/docs/atlas) - 1 component
- [Confluent](/docs/confluent) - 1 component

## Key Features

- **One Model, Many Clouds**: Single API structure across AWS, GCP, Azure, and Kubernetes
- **Validation First**: Buf ProtoValidate catches errors before deployment
- **Battle-Tested Modules**: Curated Pulumi and OpenTofu modules
- **CLI-First Workflow**: Developer-grade CLI for all operations
- **Security & Governance**: Provider credentials as stack inputs, consistent labeling

## Quick Example

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: RedisKubernetes
metadata:
  name: my-redis
spec:
  replicas: 3
  resources:
    limits:
      memory: 2Gi
      cpu: 1000m
```

```bash
project-planton validate redis.yaml
project-planton pulumi up --manifest redis.yaml --stack myorg/project/dev
```

## Resources

- [GitHub Repository](https://github.com/project-planton/project-planton)
- [Buf Schema Registry](https://buf.build/project-planton/apis)
- [Issue Tracker](https://github.com/project-planton/project-planton/issues)

