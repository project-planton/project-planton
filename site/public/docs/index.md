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

Browse deployment components by cloud provider in the [Catalog](/docs/catalog):

<div class="grid grid-cols-1 md:grid-cols-2 gap-4 my-6">
  <a href="/docs/catalog/aws" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/aws.svg" alt="AWS" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">AWS</div>
      <div class="text-sm text-slate-400">22 components</div>
    </div>
  </a>
  <a href="/docs/catalog/gcp" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/gcp.svg" alt="GCP" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">GCP</div>
      <div class="text-sm text-slate-400">5 components</div>
    </div>
  </a>
  <a href="/docs/catalog/azure" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/azure.svg" alt="Azure" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">Azure</div>
      <div class="text-sm text-slate-400">7 components</div>
    </div>
  </a>
  <a href="/docs/catalog/cloudflare" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/cloudflare.svg" alt="Cloudflare" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">Cloudflare</div>
      <div class="text-sm text-slate-400">7 components</div>
    </div>
  </a>
  <a href="/docs/catalog/civo" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/civo.svg" alt="Civo" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">Civo</div>
      <div class="text-sm text-slate-400">12 components</div>
    </div>
  </a>
  <a href="/docs/catalog/digitalocean" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/digital-ocean.svg" alt="DigitalOcean" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">DigitalOcean</div>
      <div class="text-sm text-slate-400">14 components</div>
    </div>
  </a>
  <a href="/docs/catalog/atlas" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/mongodb-atlas.svg" alt="MongoDB Atlas" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">Atlas</div>
      <div class="text-sm text-slate-400">1 component</div>
    </div>
  </a>
  <a href="/docs/catalog/confluent" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors">
    <img src="/images/providers/confluent.svg" alt="Confluent" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">Confluent</div>
      <div class="text-sm text-slate-400">1 component</div>
    </div>
  </a>
  <a href="/docs/catalog/kubernetes" class="flex items-center gap-3 p-4 rounded-lg border border-purple-900/30 bg-slate-900/30 hover:bg-slate-800/50 transition-colors opacity-50">
    <img src="/images/providers/kubernetes.svg" alt="Kubernetes" class="w-8 h-8 object-contain" />
    <div>
      <div class="font-semibold text-white">Kubernetes</div>
      <div class="text-sm text-slate-400">Coming soon</div>
    </div>
  </a>
</div>

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

