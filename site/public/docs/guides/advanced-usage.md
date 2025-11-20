---
title: "Advanced Usage"
description: "Advanced Project Planton techniques - runtime overrides, URL manifests, validation, module customization, and power-user workflows"
icon: "settings"
order: 5
---

# Advanced Usage Guide

Master advanced Project Planton techniques for power users and complex scenarios.

---

## Runtime Value Overrides with --set

The `--set` flag lets you override manifest values without editing files. Think of it like command-line arguments for your infrastructure.

### Basic Syntax

```bash
project-planton pulumi up \
  -f deployment.yaml \
  --set key=value
```

### Nested Fields

Use dot notation to access nested fields:

```bash
# Override nested spec fields
project-planton pulumi up \
  -f api.yaml \
  --set spec.container.replicas=5 \
  --set spec.container.image.tag=v2.0.0 \
  --set spec.container.resources.limits.cpu=2000m
```

### Multiple Overrides

Repeat the `--set` flag multiple times:

```bash
project-planton pulumi up \
  -f deployment.yaml \
  --set spec.replicas=10 \
  --set spec.container.image.tag=v1.5.0 \
  --set metadata.labels.version=v1.5.0 \
  --set metadata.labels.environment=staging
```

### When to Use --set

**✅ Good use cases**:
- **Quick testing**: "What if I used 10 replicas instead of 3?"
- **CI/CD parameterization**: Dynamic image tags from build pipeline
- **Emergency overrides**: Temporary configuration changes
- **A/B testing**: Compare different configurations

**❌ Bad use cases**:
- **Permanent configuration**: Changes that should be committed
- **Complex changes**: Better to edit manifest directly
- **Team deployments**: Others won't see the override in Git

### Examples

**Testing different replica counts**:

```bash
# Test with 1 replica (cheapest)
project-planton pulumi preview \
  -f api.yaml \
  --set spec.replicas=1

# Test with 5 replicas (more realistic)
project-planton pulumi preview \
  -f api.yaml \
  --set spec.replicas=5

# Deploy with 3 (commit to manifest for permanence)
vim api.yaml  # Set replicas: 3
project-planton pulumi up -f api.yaml
```

**Dynamic image tags in CI/CD**:

```bash
# In GitHub Actions
IMAGE_TAG="${GITHUB_SHA:0:7}"  # Short commit hash

project-planton pulumi up \
  -f deployment.yaml \
  --set spec.container.image.tag=$IMAGE_TAG \
  --yes
```

**Emergency scaling**:

```bash
# Production is slow, scale up immediately
project-planton pulumi up \
  -f prod-api.yaml \
  --set spec.replicas=10 \
  --yes

# Later: Update manifest and revert to normal
vim prod-api.yaml  # Set permanent replica count
project-planton pulumi up -f prod-api.yaml
```

### Limitations

**Cannot override**:
- Lists/arrays directly (override the whole field instead)
- Maps with specific keys (override entire map)
- Complex nested structures (edit manifest instead)

**Can override**:
- Strings
- Numbers (int, float)
- Booleans
- Nested scalar fields
- Message fields (creates if needed)

---

## Loading Manifests from URLs

Deploy infrastructure directly from URLs without downloading files manually.

### Basic Usage

```bash
# Deploy from GitHub raw URL
project-planton pulumi up \
  -f https://raw.githubusercontent.com/myorg/manifests/main/prod/database.yaml

# Deploy from any HTTPS URL
project-planton pulumi up \
  -f https://config-server.example.com/api/manifests/vpc.yaml
```

### How It Works

```
1. CLI detects manifestPath is a URL (has scheme + host)
   ↓
2. Downloads file using http.Get
   ↓
3. Saves to temp file with generated name (ULID)
   ↓
4. Processes temp file as normal manifest
   ↓
5. Deploys
   ↓
6. Cleans up temp file
```

**Temp file location**: `~/.project-planton/manifests/downloaded/<ulid>.yaml`

### Use Cases

**Centralized manifest repository**:

```bash
# Team maintains manifests in central repo
# Developers deploy from URLs

MANIFEST_REPO="https://raw.githubusercontent.com/myorg/infra-manifests/main"

project-planton pulumi up \
  -f $MANIFEST_REPO/prod/database.yaml

project-planton pulumi up \
  -f $MANIFEST_REPO/prod/cache.yaml
```

**Generated manifests from API**:

```bash
# CI/CD generates manifests via API
MANIFEST_URL=$(curl -s https://config-api.example.com/generate?env=prod&service=api)

project-planton pulumi up -f $MANIFEST_URL
```

**Version-pinned deployments**:

```bash
# Deploy from specific Git tag/commit
project-planton pulumi up \
  -f https://raw.githubusercontent.com/myorg/manifests/v1.0.0/database.yaml

# Rollback to previous version
project-planton pulumi up \
  -f https://raw.githubusercontent.com/myorg/manifests/v0.9.0/database.yaml
```

### URL Requirements

- Must be HTTPS (HTTP not allowed for security)
- Must return raw YAML (not HTML page)
- Must be publicly accessible (no authentication)

**GitHub raw URLs**:

```bash
# ✅ Correct - raw.githubusercontent.com
https://raw.githubusercontent.com/myorg/manifests/main/database.yaml

# ❌ Wrong - github.com (returns HTML)
https://github.com/myorg/manifests/blob/main/database.yaml
```

---

## Validation Command

Validate manifests before deploying to catch errors early.

### Basic Validation

```bash
# Validate a single manifest
project-planton validate -f database.yaml

# If valid, no output (exit code 0)
# If invalid, shows detailed errors
```

### What Gets Validated

**1. Schema validation**: Fields match proto definition  
**2. Type validation**: Values are correct types (string, int, bool)  
**3. Field constraints**: Proto-validate rules (min/max, patterns, etc.)  
**4. Required fields**: All required fields present  
**5. Custom rules**: CEL expressions for complex validation

### Example Validation Errors

```bash
$ project-planton validate -f bad-config.yaml

╔═══════════════════════════════════════════════════════════╗
║                 ❌  MANIFEST VALIDATION FAILED            ║
╚═══════════════════════════════════════════════════════════╝

⚠️  Validation Errors:

spec.container.replicas: value must be >= 1 and <= 10 (got: 15)
spec.container.cpu: value must match pattern "^[0-9]+m$" (got: "2cores")
spec.container.image.repo: value is required
```

### Pre-Deployment Workflow

```bash
# ✅ Good: Validate before deploying
project-planton validate -f resource.yaml
# See errors immediately (2 seconds)

project-planton pulumi up -f resource.yaml
# Deploy with confidence

# ⚠️ Risky: Deploy without validation
project-planton pulumi up -f resource.yaml
# Might fail after 5 minutes of deployment
```

### Validation in CI/CD

```yaml
# GitHub Actions
- name: Validate Manifests
  run: |
    for manifest in ops/manifests/**/*.yaml; do
      project-planton validate -f $manifest
    done

- name: Deploy
  run: project-planton pulumi up -f ops/manifests/prod.yaml --yes
```

---

## Load-Manifest Command

View the effective manifest with defaults applied and overrides resolved.

### Basic Usage

```bash
# Load and display manifest
project-planton load-manifest database.yaml
```

**Output**: YAML with all defaults filled in.

### With Overrides

```bash
# See what --set would produce
project-planton load-manifest \
  -f deployment.yaml \
  --set spec.replicas=5 \
  --set spec.container.image.tag=v2.0.0
```

### With Kustomize

```bash
# See what kustomize builds
project-planton load-manifest \
  --kustomize-dir services/api/kustomize \
  --overlay prod
```

### Use Cases

**Debug default application**:

```yaml
# Your manifest (minimal)
spec:
  namespace: custom-ns
  # version omitted

# Load to see defaults
$ project-planton load-manifest kubernetes-kubernetes-cert-manager.yaml

# Output shows:
spec:
  namespace: custom-ns
  version: v1.13.0           # ← Default applied
  helmChartVersion: 1.13.0   # ← Default applied
```

**Verify kustomize output**:

```bash
# Make sure your overlay produces expected result
project-planton load-manifest \
  --kustomize-dir services/api/kustomize \
  --overlay prod \
  | grep "replicas:"

# Should show: replicas: 3 (from prod overlay)
```

**Test --set before deploying**:

```bash
# Verify override works as expected
project-planton load-manifest \
  -f api.yaml \
  --set spec.container.image.tag=v2.0.0 \
  | grep "tag:"

# Should show: tag: v2.0.0
```

---

## Module Directory Override

Customize or test IaC modules locally before deploying.

### When to Use --module-dir

- **Local development**: Testing changes to Pulumi/OpenTofu modules
- **Custom modules**: Using forked/modified deployment modules
- **Module testing**: Validating module changes before committing

### Default Behavior

Without `--module-dir`, Project Planton uses current working directory:

```bash
cd /path/to/module
project-planton pulumi up -f resource.yaml
# Uses current directory as module directory
```

### Override Behavior

With `--module-dir`, you can deploy from anywhere:

```bash
# From any location
project-planton pulumi up \
  -f ~/manifests/database.yaml \
  --module-dir ~/projects/custom-modules/postgres-k8s
```

### Local Module Development Workflow

```bash
# 1. Clone or fork a module
git clone https://github.com/project-planton/project-planton
cd apis/org/project_planton/provider/kubernetes/postgresqk8s/v1/iac/pulumi

# 2. Make changes to module code
vim main.go

# 3. Test with your manifest
project-planton pulumi preview \
  -f ~/test-manifests/postgres.yaml \
  --module-dir .

# 4. Iterate
vim main.go
project-planton pulumi preview -f ~/test-manifests/postgres.yaml --module-dir .

# 5. Deploy when ready
project-planton pulumi up \
  -f ~/test-manifests/postgres.yaml \
  --module-dir .
```

### Custom Module Example

**Fork and customize**:

```bash
# 1. Fork the default module
cp -r apis/org/project_planton/provider/aws/awss3bucket/v1/iac/pulumi \
      ~/custom-modules/my-s3-module

# 2. Customize
cd ~/custom-modules/my-s3-module
vim main.go  # Add custom logic

# 3. Test
project-planton pulumi up \
  -f s3-bucket.yaml \
  --module-dir ~/custom-modules/my-s3-module

# 4. If it works, consider contributing back or maintaining internally
```

---

## Combining Techniques

### Kustomize + --set

```bash
# Base from kustomize, runtime override for image tag
project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay prod \
  --set spec.container.image.tag=$GIT_SHA
```

**Order of application**:
1. Kustomize builds base + overlay
2. --set overrides applied to result
3. Final manifest validated
4. Deployed

### URL Manifest + --set

```bash
# Load from URL, override specific values
project-planton pulumi up \
  -f https://manifests.example.com/database.yaml \
  --set spec.region=us-west-2 \
  --set spec.instanceSize=large
```

### Kustomize + Custom Module

```bash
# Custom overlay + custom module
project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay prod \
  --module-dir ~/custom-modules/api-module
```

---

## Pro Tips

### 1. Use Shell Variables for Readability

```bash
# ✅ Good: Clear and reusable
MANIFEST="ops/resources/database.yaml"
OVERLAY="prod"
IMAGE_TAG="v1.2.3"

project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay $OVERLAY \
  --set spec.container.image.tag=$IMAGE_TAG

# ❌ Bad: Hard to read, error-prone
project-planton pulumi up --kustomize-dir services/api/kustomize --overlay prod --set spec.container.image.tag=v1.2.3
```

### 2. Create Deployment Scripts

```bash
#!/bin/bash
# deploy-api.sh

set -e

ENVIRONMENT=${1:-dev}
IMAGE_TAG=${2:-latest}

echo "Deploying API to $ENVIRONMENT with tag $IMAGE_TAG"

project-planton pulumi preview \
  --kustomize-dir services/api/kustomize \
  --overlay $ENVIRONMENT \
  --set spec.container.image.tag=$IMAGE_TAG

read -p "Apply? (y/N) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    project-planton pulumi up \
      --kustomize-dir services/api/kustomize \
      --overlay $ENVIRONMENT \
      --set spec.container.image.tag=$IMAGE_TAG
fi
```

**Usage**:

```bash
./deploy-api.sh staging v1.2.0
```

### 3. Validate in Pre-Commit Hooks

**`.git/hooks/pre-commit`**:

```bash
#!/bin/bash

# Validate all manifests before committing
for manifest in $(git diff --cached --name-only --diff-filter=ACM | grep '\.yaml$'); do
    if project-planton validate -f $manifest; then
        echo "✓ $manifest valid"
    else
        echo "✗ $manifest invalid"
        exit 1
    fi
done
```

### 4. Use Make for Complex Workflows

**`Makefile`**:

```makefile
.PHONY: validate deploy-dev deploy-staging deploy-prod

validate:
	@for f in ops/manifests/*.yaml; do \
		project-planton validate -f $$f; \
	done

deploy-dev:
	project-planton pulumi up \
		--kustomize-dir services/api/kustomize \
		--overlay dev \
		--yes

deploy-staging:
	project-planton pulumi preview \
		--kustomize-dir services/api/kustomize \
		--overlay staging
	@read -p "Deploy to staging? (y/N) " REPLY; \
	if [ "$$REPLY" = "y" ]; then \
		project-planton pulumi up \
			--kustomize-dir services/api/kustomize \
			--overlay staging; \
	fi

deploy-prod:
	@echo "⚠️  Production deployment - review carefully"
	project-planton pulumi preview \
		--kustomize-dir services/api/kustomize \
		--overlay prod
	@echo "Deploy to PRODUCTION?"
	@read -p "Type 'yes' to confirm: " REPLY; \
	if [ "$$REPLY" = "yes" ]; then \
		project-planton pulumi up \
			--kustomize-dir services/api/kustomize \
			--overlay prod; \
	fi
```

**Usage**:

```bash
make validate
make deploy-dev
make deploy-prod
```

---

## Advanced Patterns

### Pattern 1: Environment Matrix Testing

Test configuration across all environments:

```bash
#!/bin/bash
# test-all-environments.sh

for env in dev staging prod; do
    echo "Testing $env environment..."
    
    if project-planton validate \
        --kustomize-dir services/api/kustomize \
        --overlay $env; then
        echo "✓ $env configuration valid"
    else
        echo "✗ $env configuration invalid"
        exit 1
    fi
done
```

### Pattern 2: Progressive Rollout

Deploy gradually across environments with validation:

```bash
#!/bin/bash
# progressive-rollout.sh

IMAGE_TAG=$1

# Deploy to dev
project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay dev \
  --set spec.container.image.tag=$IMAGE_TAG \
  --yes

# Wait and test
sleep 60
curl -f https://api-dev.example.com/health || exit 1

# Deploy to staging
project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay staging \
  --set spec.container.image.tag=$IMAGE_TAG \
  --yes

# Wait and test
sleep 60
curl -f https://api-staging.example.com/health || exit 1

# Deploy to prod (with manual confirmation)
project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay prod \
  --set spec.container.image.tag=$IMAGE_TAG
```

### Pattern 3: Manifest Generation

Generate manifests programmatically:

```python
#!/usr/bin/env python3
import yaml

def generate_manifest(name, replicas, memory):
    manifest = {
        'apiVersion': 'kubernetes.project-planton.org/v1',
        'kind': 'MicroserviceKubernetes',
        'metadata': {'name': name},
        'spec': {
            'container': {
                'replicas': replicas,
                'resources': {
                    'limits': {
                        'memory': memory
                    }
                }
            }
        }
    }
    
    with open(f'{name}.yaml', 'w') as f:
        yaml.dump(manifest, f)

# Generate manifests
generate_manifest('api', 3, '2Gi')
generate_manifest('worker', 5, '4Gi')
```

Then deploy:

```bash
python3 generate-manifests.py
project-planton pulumi up -f api.yaml
project-planton pulumi up -f worker.yaml
```

---

## Debugging Techniques

### Inspecting Final Manifest

See exactly what gets deployed:

```bash
# Load manifest with all transformations applied
project-planton load-manifest \
  --kustomize-dir services/api/kustomize \
  --overlay prod \
  --set spec.container.image.tag=v1.2.0 \
  > final-manifest.yaml

# Review
cat final-manifest.yaml
```

### Validating Before Deploy

```bash
# Validate with overrides
project-planton validate \
  --kustomize-dir services/api/kustomize \
  --overlay prod

# If validation passes, deploy
project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay prod
```

### Testing Module Changes

```bash
# Test modified module without deploying
cd ~/projects/custom-module

# 1. Validate manifest
project-planton validate -f ~/test-manifests/test.yaml

# 2. Preview changes
project-planton pulumi preview \
  -f ~/test-manifests/test.yaml \
  --module-dir .

# 3. If preview looks good, deploy to test environment
project-planton pulumi up \
  -f ~/test-manifests/test.yaml \
  --module-dir . \
  --stack test-org/test-project/test-stack
```

---

## Power User Workflows

### Multi-Manifest Deployment

Deploy multiple related resources:

```bash
#!/bin/bash
# deploy-stack.sh

OVERLAY=${1:-dev}

MANIFESTS=(
    "vpc.yaml"
    "database.yaml"
    "cache.yaml"
    "application.yaml"
)

for manifest in "${MANIFESTS[@]}"; do
    echo "Deploying $manifest to $OVERLAY..."
    project-planton pulumi up \
        --kustomize-dir ops/resources/kustomize \
        --overlay $OVERLAY \
        --yes
    
    # Wait between deployments
    sleep 10
done
```

### Conditional Deployment

Deploy only if validation passes:

```bash
if project-planton validate -f database.yaml; then
    echo "✓ Validation passed, deploying..."
    project-planton pulumi up -f database.yaml --yes
else
    echo "✗ Validation failed, aborting"
    exit 1
fi
```

### Parameterized Deployments

```bash
#!/bin/bash
# deploy-with-params.sh

ENVIRONMENT=$1
REPLICAS=$2
CPU=$3
MEMORY=$4

project-planton pulumi up \
  --kustomize-dir services/api/kustomize \
  --overlay $ENVIRONMENT \
  --set spec.replicas=$REPLICAS \
  --set spec.container.resources.limits.cpu=$CPU \
  --set spec.container.resources.limits.memory=$MEMORY
```

**Usage**:

```bash
./deploy-with-params.sh prod 5 2000m 4Gi
```

---

## Common Mistakes to Avoid

### ❌ Using --set for Permanent Changes

```bash
# BAD: Override in CI/CD but not in manifest
project-planton pulumi up \
  -f api.yaml \
  --set spec.replicas=10 \
  --yes

# 3 months later: "Why is prod running 10 replicas?"
# Nobody knows because it's not in the manifest!

# GOOD: Update manifest first
vim api.yaml  # Set replicas: 10
git commit -m "scale: increase API replicas to 10"
project-planton pulumi up -f api.yaml --yes
```

### ❌ Loading Non-Raw GitHub URLs

```bash
# BAD: Returns HTML page, not YAML
project-planton pulumi up \
  -f https://github.com/myorg/repo/blob/main/manifest.yaml

# GOOD: Use raw.githubusercontent.com
project-planton pulumi up \
  -f https://raw.githubusercontent.com/myorg/repo/main/manifest.yaml
```

### ❌ Complex --set Overrides

```bash
# BAD: Trying to override complex structures
project-planton pulumi up \
  -f api.yaml \
  --set spec.container.resources='{"limits":{"cpu":"2000m"}}'  # Won't work

# GOOD: Override individual fields
project-planton pulumi up \
  -f api.yaml \
  --set spec.container.resources.limits.cpu=2000m
```

---

## Related Documentation

- [Manifest Structure](/docs/guides/manifests) - Understanding manifests
- [Kustomize Integration](/docs/guides/kustomize) - Detailed kustomize guide
- [Pulumi Commands](/docs/cli/pulumi-commands) - Pulumi deployment
- [OpenTofu Commands](/docs/cli/tofu-commands) - OpenTofu deployment

---

## Next Steps

1. **Experiment with --set**: Try overriding values in a test manifest
2. **Set up Kustomize**: Create overlays for your environments
3. **Create Scripts**: Automate your common deployment workflows
4. **Test Validation**: Add validate step to your CI/CD pipeline

Master these advanced techniques to unlock Project Planton's full power for complex deployment scenarios.

