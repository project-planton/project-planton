---
title: "CLI Reference"
description: "Complete command-line reference for project-planton CLI - all commands, flags, and options"
icon: "terminal"
order: 1
---

# CLI Reference

Complete command-line reference for the `project-planton` CLI.

---

## Command Tree

```
project-planton
├── pulumi              Manage infrastructure with Pulumi
│   ├── init           Initialize Pulumi stack
│   ├── preview        Preview infrastructure changes
│   ├── up             Deploy infrastructure (or 'update')
│   ├── refresh        Sync state with cloud reality
│   ├── destroy        Teardown infrastructure
│   ├── delete         Remove stack metadata (or 'rm')
│   └── cancel         Cancel ongoing Pulumi operation
├── tofu                Manage infrastructure with OpenTofu
│   ├── init           Initialize backend and providers
│   ├── plan           Preview infrastructure changes
│   ├── apply          Deploy infrastructure
│   ├── refresh        Sync state with cloud reality
│   └── destroy        Teardown infrastructure
├── validate            Validate manifest against schema
├── load-manifest       Load and display manifest with defaults
└── version             Show CLI version
```

---

## Top-Level Commands

### pulumi

Manage infrastructure using Pulumi as the IaC engine.

**Subcommands**: `init`, `preview`, `up`/`update`, `refresh`, `destroy`, `delete`/`rm`, `cancel`

**Documentation**: See [Pulumi Commands Reference](/docs/cli/pulumi-commands)

**Example**:

```bash
project-planton pulumi up --manifest database.yaml
```

### tofu

Manage infrastructure using OpenTofu/Terraform as the IaC engine.

**Subcommands**: `init`, `plan`, `apply`, `refresh`, `destroy`

**Documentation**: See [OpenTofu Commands Reference](/docs/cli/tofu-commands)

**Example**:

```bash
project-planton tofu apply --manifest database.yaml
```

### validate

Validate a manifest against its Protocol Buffer schema without deploying.

**Usage**:

```bash
project-planton validate --manifest <file> [flags]
```

**Example**:

```bash
# Validate single manifest
project-planton validate --manifest ops/resources/database.yaml

# With kustomize
project-planton validate \
  --kustomize-dir services/api/kustomize \
  --overlay prod

# If valid: exits with code 0, no output
# If invalid: shows detailed errors, exits with code 1
```

**Flags**:
- `--manifest <file>`: Path to manifest file
- `--kustomize-dir <dir>`: Kustomize base directory
- `--overlay <name>`: Kustomize overlay name

### load-manifest

Load a manifest and display it with defaults applied and overrides resolved.

**Usage**:

```bash
project-planton load-manifest --manifest <file> [flags]
```

**Example**:

```bash
# Load manifest and see defaults
project-planton load-manifest --manifest database.yaml

# Load with overrides
project-planton load-manifest \
  --manifest api.yaml \
  --set spec.replicas=5

# Load kustomize-built manifest
project-planton load-manifest \
  --kustomize-dir services/api/kustomize \
  --overlay prod
```

**Flags**: Same as `validate`

**Output**: YAML manifest with defaults filled in and overrides applied

### version

Show Project Planton CLI version information.

**Usage**:

```bash
project-planton version
```

**Example Output**:

```
project-planton version: v0.1.0
git commit: a1b2c3d
built: 2025-11-11T10:30:00Z
```

---

## Common Flags

These flags are available across multiple commands:

### Manifest Input

**`--manifest <path>`**  
Path to manifest YAML file (local or URL).

```bash
# Local file
--manifest ops/resources/database.yaml

# URL
--manifest https://raw.githubusercontent.com/myorg/manifests/main/db.yaml
```

**`--kustomize-dir <directory>`**  
Base directory containing kustomize structure.

```bash
--kustomize-dir services/api/kustomize
```

**`--overlay <name>`**  
Kustomize overlay environment to build (must be used with `--kustomize-dir`).

```bash
--overlay prod
```

**Priority**: `--manifest` > `--kustomize-dir` + `--overlay`

### Execution Control

**`--module-dir <path>`**  
Override IaC module directory (defaults to current directory).

```bash
--module-dir ~/projects/custom-modules/my-module
```

**`--set <key>=<value>`**  
Override manifest field values at runtime (repeatable).

```bash
--set spec.replicas=5 \
--set spec.container.image.tag=v2.0.0
```

### Pulumi-Specific Flags

**`--stack <org>/<project>/<stack>`**  
Override stack FQDN (instead of using manifest label).

```bash
--stack my-org/my-project/dev-stack
```

**`--yes`**  
Auto-approve operations without confirmation (Pulumi commands).

```bash
--yes
```

**`--force`**  
Force stack removal even if resources exist (`delete`/`rm` only).

```bash
--force
```

### OpenTofu-Specific Flags

**`--auto-approve`**  
Skip interactive approval (`apply` and `destroy` commands).

```bash
--auto-approve
```

**`--destroy`**  
Create destroy plan (`plan` command).

```bash
--destroy
```

### Credential Flags

Provider credential file paths:

```bash
--aws-credential <file>
--azure-credential <file>
--gcp-credential <file>
--kubernetes-cluster <file>
--cloudflare-credential <file>
--confluent-credential <file>
--mongodb-atlas-credential <file>
--snowflake-credential <file>
```

---

## Environment Variables

### Respected by CLI

**`TF_LOG`** / **`PULUMI_LOG_LEVEL`**  
Enable verbose logging for debugging.

```bash
export TF_LOG=DEBUG
export PULUMI_LOG_LEVEL=3
```

### Provider Credentials

See [Credentials Guide](/docs/guides/credentials) for complete list of provider-specific environment variables.

---

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error (validation failed, deployment failed, etc.) |

---

## Configuration Files

### Manifest Download Directory

Downloaded URL manifests are cached in:

```
~/.project-planton/manifests/downloaded/
```

### Module Cache

Cloned IaC modules are cached in:

```
~/.project-planton/modules/
```

---

## Examples by Use Case

### First Deployment

```bash
# Using Pulumi
project-planton validate --manifest database.yaml
project-planton pulumi up --manifest database.yaml

# Using OpenTofu
project-planton validate --manifest database.yaml
project-planton tofu init --manifest database.yaml
project-planton tofu plan --manifest database.yaml
project-planton tofu apply --manifest database.yaml
```

### Multi-Environment Deployment

```bash
# Deploy across environments
for env in dev staging prod; do
    project-planton pulumi up \
        --kustomize-dir services/api/kustomize \
        --overlay $env \
        --yes
done
```

### CI/CD Deployment

```bash
# Non-interactive with dynamic values
project-planton pulumi up \
  --manifest deployment.yaml \
  --set spec.container.image.tag=$CI_COMMIT_SHA \
  --yes
```

### Testing Local Module Changes

```bash
# Point to local module during development
project-planton pulumi preview \
  --manifest test.yaml \
  --module-dir ~/dev/my-module
```

---

## Related Documentation

- [Pulumi Commands](/docs/cli/pulumi-commands) - Detailed Pulumi command guide
- [OpenTofu Commands](/docs/cli/tofu-commands) - Detailed OpenTofu command guide
- [Manifest Structure](/docs/guides/manifests) - Understanding manifests
- [Credentials Guide](/docs/guides/credentials) - Setting up cloud credentials
- [Advanced Usage](/docs/guides/advanced-usage) - Power user techniques

---

## Getting Help

**Command help**:

```bash
# General help
project-planton --help

# Command-specific help
project-planton pulumi --help
project-planton tofu apply --help
```

**Found an issue?** [Open an issue](https://github.com/project-planton/project-planton/issues)

**Need support?** Check existing issues or discussions

