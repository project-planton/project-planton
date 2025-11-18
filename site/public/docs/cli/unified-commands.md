---
title: "Unified Apply/Destroy Commands"
description: "kubectl-style unified commands that automatically detect your provisioner - simplify your workflow with apply and destroy"
icon: "rocket"
order: 2
---

# Unified Apply/Destroy Commands

The unified `apply` and `destroy` commands provide a kubectl-like experience by automatically detecting the IaC provisioner from your manifest and routing to the appropriate tool (Pulumi, Tofu, or Terraform).

---

## Why Unified Commands?

### The Problem

Previously, you had to remember different commands for different provisioners:

```bash
# Pulumi
project-planton pulumi up --manifest app.yaml --stack org/project/env

# OpenTofu
project-planton tofu apply --manifest app.yaml

# Different commands, different flags, cognitive overhead!
```

### The Solution

Now, use the same commands regardless of provisioner:

```bash
# Works for Pulumi, Tofu, or Terraform
project-planton apply -f app.yaml
project-planton destroy -f app.yaml
```

The CLI automatically:
1. Reads the `project-planton.org/provisioner` label from your manifest
2. Routes to the appropriate provisioner
3. Passes through all relevant flags

---

## The Provisioner Label

Add this label to your manifest metadata:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: my-database
  labels:
    project-planton.org/provisioner: pulumi  # or tofu or terraform
spec:
  # ... your spec
```

**Supported values** (case-insensitive):
- `pulumi` - Use Pulumi for deployment
- `tofu` - Use OpenTofu for deployment
- `terraform` - Use Terraform for deployment

---

## Commands

### apply

Deploy infrastructure changes.

**Usage**:

```bash
project-planton apply -f <file> [flags]
project-planton apply --manifest <file> [flags]
```

**Examples**:

```bash
# Basic usage
project-planton apply -f database.yaml

# With kustomize
project-planton apply --kustomize-dir services/api --overlay prod

# With field overrides
project-planton apply -f app.yaml --set spec.replicas=5

# Auto-approve (Pulumi)
project-planton apply -f app.yaml --yes

# Auto-approve (Tofu/Terraform)
project-planton apply -f app.yaml --auto-approve
```

**What it does**:
1. Loads and validates your manifest
2. Detects provisioner from `project-planton.org/provisioner` label
3. If label missing, prompts interactively (defaults to Pulumi)
4. Routes to appropriate provisioner:
   - **Pulumi**: Runs `pulumi update`
   - **Tofu**: Runs `tofu apply`
   - **Terraform**: Runs `terraform apply`

---

### destroy

Teardown infrastructure.

**Usage**:

```bash
project-planton destroy -f <file> [flags]
project-planton delete -f <file> [flags]  # kubectl-style alias
```

**Examples**:

```bash
# Basic usage
project-planton destroy -f database.yaml

# kubectl-style delete
project-planton delete -f database.yaml

# With kustomize
project-planton destroy --kustomize-dir services/api --overlay staging

# Auto-approve
project-planton destroy -f app.yaml --auto-approve
```

**What it does**:
1. Loads and validates your manifest
2. Detects provisioner from label
3. Routes to appropriate destroy operation:
   - **Pulumi**: Runs `pulumi destroy`
   - **Tofu**: Runs `tofu destroy`
   - **Terraform**: Runs `terraform destroy`

---

## Interactive Provisioner Selection

If your manifest doesn't have the `project-planton.org/provisioner` label, the CLI prompts you:

```bash
$ project-planton apply -f database.yaml

✓ Manifest loaded
✓ Manifest validated
• Detecting provisioner...
ℹ Provisioner not specified in manifest
Select provisioner [Pulumi]/tofu/terraform: 
```

Simply press **Enter** to use Pulumi (the default), or type your choice.

**Tips**:
- Input is case-insensitive (`Pulumi`, `pulumi`, `PULUMI` all work)
- The prompt only appears when the label is missing
- Add the label to your manifests for a fully automated workflow

---

## Supported Flags

Both commands support all flags from their respective provisioners.

### Common Flags

| Flag | Description |
|------|-------------|
| `-f, --manifest <path>` | Path to manifest file (kubectl-style `-f` shorthand) |
| `--kustomize-dir <dir>` | Kustomize base directory |
| `--overlay <name>` | Kustomize overlay (prod, dev, staging, etc.) |
| `--set key=value` | Override manifest fields (repeatable) |
| `--module-dir <path>` | Override IaC module directory |

### Pulumi-Specific Flags

| Flag | Description |
|------|-------------|
| `--stack <org>/<project>/<stack>` | Override stack FQDN (or use manifest label) |
| `--yes` | Auto-approve without confirmation |
| `--diff` | Show detailed resource diffs |

### Tofu/Terraform-Specific Flags

| Flag | Description |
|------|-------------|
| `--auto-approve` | Skip interactive approval |

### Provider Credentials

| Flag | Description |
|------|-------------|
| `--aws-provider-config <file>` | AWS credentials |
| `--azure-provider-config <file>` | Azure credentials |
| `--gcp-provider-config <file>` | GCP credentials |
| `--kubernetes-provider-config <file>` | Kubernetes config |
| `--cloudflare-provider-config <file>` | Cloudflare credentials |
| `--confluent-provider-config <file>` | Confluent credentials |
| `--atlas-provider-config <file>` | MongoDB Atlas credentials |
| `--snowflake-provider-config <file>` | Snowflake credentials |

---

## Complete Examples

### Example 1: Pulumi PostgreSQL Database

**Manifest** (`postgres.yaml`):

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: app-database
  labels:
    project-planton.org/provisioner: pulumi
    pulumi.project-planton.org/stack.name: prod.PostgresKubernetes.app-database
spec:
  container:
    replicas: 1
    resources:
      limits:
        cpu: 1000m
        memory: 2Gi
```

**Commands**:

```bash
# Deploy
project-planton apply -f postgres.yaml

# Update with more replicas
project-planton apply -f postgres.yaml --set spec.container.replicas=3

# Destroy
project-planton destroy -f postgres.yaml
```

---

### Example 2: Tofu AWS VPC

**Manifest** (`vpc.yaml`):

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsVpc
metadata:
  name: production-vpc
  labels:
    project-planton.org/provisioner: tofu
    terraform.project-planton.org/backend.type: s3
    terraform.project-planton.org/backend.object: terraform-state/vpc/prod.tfstate
spec:
  cidrBlock: 10.0.0.0/16
  region: us-west-2
```

**Commands**:

```bash
# Deploy
project-planton apply -f vpc.yaml --auto-approve

# Destroy
project-planton delete -f vpc.yaml --auto-approve
```

---

### Example 3: Multi-Environment with Kustomize

**Directory structure**:

```
services/api/
├── base/
│   └── kustomization.yaml
└── overlays/
    ├── dev/
    │   └── kustomization.yaml
    ├── staging/
    │   └── kustomization.yaml
    └── prod/
        └── kustomization.yaml
```

**Deploy to all environments**:

```bash
for env in dev staging prod; do
    echo "Deploying to $env..."
    project-planton apply \
        --kustomize-dir services/api \
        --overlay $env \
        --yes
done
```

---

### Example 4: CI/CD Pipeline

```bash
#!/bin/bash
# deploy.sh - CI/CD deployment script

set -e

# Variables from CI environment
IMAGE_TAG="${CI_COMMIT_SHA}"
ENVIRONMENT="${CI_ENVIRONMENT_NAME}"

# Deploy with dynamic values
project-planton apply \
    -f deployment.yaml \
    --set spec.container.image.tag="$IMAGE_TAG" \
    --set metadata.labels.environment="$ENVIRONMENT" \
    --yes

echo "Deployed version $IMAGE_TAG to $ENVIRONMENT"
```

---

## Migration from Provisioner-Specific Commands

The unified commands are **fully backward compatible**. You can migrate gradually:

### Before (Provisioner-Specific)

```bash
# Pulumi
project-planton pulumi up --manifest app.yaml --stack org/project/dev

# OpenTofu
project-planton tofu apply --manifest app.yaml
```

### After (Unified)

```bash
# Works for both!
project-planton apply -f app.yaml
```

### Migration Steps

1. **Add provisioner label** to your manifests:
   ```yaml
   metadata:
     labels:
       project-planton.org/provisioner: pulumi  # or tofu
   ```

2. **Replace commands** in your scripts:
   - `pulumi up` → `apply`
   - `pulumi destroy` → `destroy`
   - `tofu apply` → `apply`
   - `tofu destroy` → `destroy`

3. **Update flags**:
   - `--manifest` → `-f` (both work, `-f` is shorter)
   - `--stack` → can be in manifest labels now
   - All other flags work the same

---

## Benefits

### 1. Simplified Mental Model

One command pattern for all provisioners:

```bash
project-planton apply -f <manifest>
project-planton destroy -f <manifest>
```

### 2. kubectl-like Experience

Familiar kubectl patterns:

```bash
kubectl apply -f deployment.yaml
project-planton apply -f deployment.yaml

kubectl delete -f deployment.yaml
project-planton delete -f deployment.yaml
```

### 3. Easier Automation

Write scripts that work regardless of provisioner:

```bash
for manifest in manifests/*.yaml; do
    project-planton apply -f "$manifest"
done
```

### 4. Lower Barrier to Entry

New team members don't need to learn multiple command patterns.

### 5. Gradual Migration

All existing commands still work - migrate at your own pace.

---

## Troubleshooting

### "Invalid provisioner value"

**Error**:
```
Invalid provisioner in manifest: invalid provisioner value 'pulum': must be one of 'pulumi', 'tofu', or 'terraform'
```

**Solution**: Check your provisioner label for typos. Valid values are: `pulumi`, `tofu`, or `terraform` (case-insensitive).

---

### "Provisioner not specified in manifest"

**Behavior**: CLI prompts you to select a provisioner.

**Solutions**:
1. **Add the label** to your manifest (recommended):
   ```yaml
   metadata:
     labels:
       project-planton.org/provisioner: pulumi
   ```

2. **Select interactively**: Type your choice when prompted

3. **Use provisioner-specific commands** if you prefer:
   ```bash
   project-planton pulumi up --manifest app.yaml
   ```

---

### Pulumi Backend Not Configured

**Error** (when using Pulumi provisioner):
```
error: no Pulumi backend configured
```

**Solution**: Set up Pulumi backend:

```bash
# Local backend (for testing)
pulumi login --local

# Or cloud backend
pulumi login
```

---

### Tofu/Terraform Backend Not Configured

**Behavior**: Uses local backend by default.

**Solution**: Configure backend via labels:

```yaml
metadata:
  labels:
    terraform.project-planton.org/backend.type: s3
    terraform.project-planton.org/backend.object: bucket/path/state.tfstate
```

---

## Best Practices

### 1. Always Add Provisioner Label

Include the provisioner label in all manifests:

```yaml
metadata:
  labels:
    project-planton.org/provisioner: pulumi
```

This enables fully automated workflows.

---

### 2. Use -f Flag for Consistency

Prefer `-f` over `--manifest` to match kubectl style:

```bash
# Good (kubectl-style)
project-planton apply -f app.yaml

# Also works
project-planton apply --manifest app.yaml
```

---

### 3. Combine with Validation

Always validate before applying in production:

```bash
project-planton validate -f app.yaml && \
project-planton apply -f app.yaml --yes
```

---

### 4. Use Kustomize for Multi-Environment

Organize environments with kustomize:

```bash
project-planton apply --kustomize-dir services/api --overlay prod
```

---

### 5. Override Values for CI/CD

Use `--set` for dynamic values in pipelines:

```bash
project-planton apply -f app.yaml \
    --set spec.image.tag="$CI_COMMIT_SHA" \
    --set metadata.labels.build="$CI_BUILD_ID"
```

---

## Related Documentation

- [CLI Reference](/docs/cli/cli-reference) - Complete command reference
- [Pulumi Commands](/docs/cli/pulumi-commands) - Pulumi-specific details
- [OpenTofu Commands](/docs/cli/tofu-commands) - OpenTofu-specific details
- [Manifest Structure](/docs/guides/manifests) - Writing manifests
- [Kustomize Integration](/docs/guides/kustomize) - Multi-environment setup

---

## Feedback

Found an issue or have a suggestion? [Open an issue](https://github.com/project-planton/project-planton/issues) on GitHub.

