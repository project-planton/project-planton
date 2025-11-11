---
title: "CLI Documentation"
description: "Complete command-line interface documentation for Project Planton - Pulumi commands, OpenTofu commands, and CLI reference"
icon: "terminal"
order: 10
---

# CLI Documentation

Everything you need to know about the `project-planton` command-line interface.

---

## Overview

The Project Planton CLI is your gateway to deploying infrastructure across any cloud provider with a consistent, manifest-driven workflow. Whether you prefer Pulumi or OpenTofu as your IaC engine, the CLI experience is identical.

---

## Getting Started

### Installation

```bash
# Install via Homebrew
brew install project-planton/tap/project-planton

# Verify installation
project-planton version
```

### Quick Example

```bash
# Validate your manifest
project-planton validate --manifest database.yaml

# Deploy with Pulumi
project-planton pulumi up --manifest database.yaml

# Or deploy with OpenTofu
project-planton tofu apply --manifest database.yaml
```

---

## Documentation Sections

### [CLI Reference](/docs/cli/cli-reference)

Complete reference for all CLI commands and flags.

**What you'll find**:
- Command tree structure
- All available commands
- Common flags and options
- Examples by use case
- Exit codes

**When to read**: Quick lookup for command syntax and flags.

---

### [Pulumi Commands](/docs/cli/pulumi-commands)

Comprehensive guide to managing infrastructure with Pulumi.

**What you'll find**:
- Infrastructure lifecycle (init → preview → up → refresh → destroy → delete)
- Detailed command reference with examples
- Common workflows (first deployment, updates, rollback, CI/CD)
- Troubleshooting Pulumi-specific issues
- Best practices and tips

**When to read**: If you're using Pulumi as your IaC engine.

**Key commands**:
- `pulumi init` - Initialize stack
- `pulumi preview` - Preview changes
- `pulumi up` - Deploy infrastructure
- `pulumi refresh` - Sync state
- `pulumi destroy` - Teardown infrastructure
- `pulumi delete` - Remove stack

---

### [OpenTofu Commands](/docs/cli/tofu-commands)

Comprehensive guide to managing infrastructure with OpenTofu/Terraform.

**What you'll find**:
- Infrastructure lifecycle (init → plan → apply → refresh → destroy)
- Detailed command reference with examples
- Common workflows (first deployment, updates, CI/CD)
- State management
- Troubleshooting OpenTofu-specific issues
- Best practices and tips

**When to read**: If you're using OpenTofu/Terraform as your IaC engine.

**Key commands**:
- `tofu init` - Initialize backend
- `tofu plan` - Preview changes
- `tofu apply` - Deploy infrastructure
- `tofu refresh` - Sync state
- `tofu destroy` - Teardown infrastructure

---

## Choose Your IaC Engine

### Pulumi

**Best for**:
- Teams preferring programming languages over HCL
- Complex logic and control flow
- Type safety and IDE autocomplete
- Real-time outputs during deployment

**Trade-offs**:
- Requires Pulumi backend (Pulumi Cloud, S3, GCS, etc.)
- Smaller community than Terraform
- Modules written in Go (for Project Planton)

### OpenTofu

**Best for**:
- Teams with existing Terraform experience
- Declarative infrastructure-as-code preference
- Larger ecosystem and community
- HashiCorp Configuration Language (HCL)

**Trade-offs**:
- Less flexible for complex logic
- State management is more manual
- No real-time output streaming

**The good news**: Project Planton supports both! You can switch between them based on your team's preference. The manifest format is identical—only the deployment command changes.

---

## Common Workflows

### First-Time Setup

```bash
# 1. Install CLI
brew install project-planton/tap/project-planton

# 2. Install IaC engine
brew install pulumi        # For Pulumi
# OR
brew install opentofu      # For OpenTofu

# 3. Configure credentials (see Credentials Guide)
export AWS_ACCESS_KEY_ID="..."
export AWS_SECRET_ACCESS_KEY="..."

# 4. Create your first manifest
cat > database.yaml <<EOF
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: my-database
spec:
  engine: postgres
  instanceClass: db.t3.medium
EOF

# 5. Deploy
project-planton pulumi up --manifest database.yaml
```

### Daily Development

```bash
# Morning: pull latest manifests
git pull

# Edit manifest
vim ops/resources/api-deployment.yaml

# Validate changes
project-planton validate --manifest ops/resources/api-deployment.yaml

# Preview changes
project-planton pulumi preview --manifest ops/resources/api-deployment.yaml

# Deploy
project-planton pulumi up --manifest ops/resources/api-deployment.yaml

# Evening: commit changes
git add ops/resources/api-deployment.yaml
git commit -m "scale: increase API replicas"
git push
```

---

## Related Documentation

- [Manifest Structure Guide](/docs/guides/manifests) - Learn how to write manifests
- [Credentials Guide](/docs/guides/credentials) - Set up cloud provider credentials
- [Kustomize Integration](/docs/guides/kustomize) - Multi-environment deployments
- [Advanced Usage](/docs/guides/advanced-usage) - Power user techniques
- [Troubleshooting](/docs/troubleshooting) - Solutions to common problems

---

## Getting Help

**Quick help**:

```bash
project-planton --help
project-planton pulumi --help
project-planton tofu apply --help
```

**Found an issue?** [Open an issue](https://github.com/project-planton/project-planton/issues)

**Questions?** Check [GitHub Discussions](https://github.com/project-planton/project-planton/discussions)

