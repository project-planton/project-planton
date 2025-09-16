# Pulumi Labels Package

This package defines standardized Kubernetes labels for configuring Pulumi backend state management directly within ProjectPlanton resource manifests.

## Overview

The `pulumilabels` package provides constant definitions for labels that can be applied to any ProjectPlanton resource manifest to specify where and how Pulumi should store its state. This enables infrastructure-as-code deployments to be fully self-contained, with backend configuration embedded in the manifest itself.

## Label Constants

### Primary Label

- **`StackFqdnLabelKey`** (`pulumi.project-planton.org/stack.fqdn`)
  - Takes precedence over individual component labels
  - Format: `organization/project/stack`
  - Example: `demo-org/aws-infrastructure/production`

### Component Labels (Fallback)

When `stack.fqdn` is not specified, the following three labels must all be present:

- **`OrganizationLabelKey`** (`pulumi.project-planton.org/organization`)
  - The Pulumi organization name
  - Example: `demo-org`

- **`ProjectLabelKey`** (`pulumi.project-planton.org/project`)
  - The Pulumi project name
  - Example: `aws-infrastructure`

- **`StackNameLabelKey`** (`pulumi.project-planton.org/stack.name`)
  - The Pulumi stack name
  - Example: `production`

## Usage Examples

### Using Stack FQDN (Recommended)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsVpc
metadata:
  name: production-vpc
  labels:
    pulumi.project-planton.org/stack.fqdn: "acme-corp/network-infrastructure/prod"
spec:
  cidrBlock: "10.0.0.0/16"
```

### Using Individual Components

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpGkeCluster
metadata:
  name: app-cluster
  labels:
    pulumi.project-planton.org/organization: "acme-corp"
    pulumi.project-planton.org/project: "kubernetes-clusters"
    pulumi.project-planton.org/stack.name: "production"
spec:
  region: "us-central1"
```

## Benefits

1. **Self-Contained Manifests**: Backend configuration travels with the manifest
2. **Quick Start Ready**: Enables demo manifests that work out-of-the-box
3. **Override Flexibility**: Labels can override CLI flags when present
4. **GitOps Friendly**: Backend config is version-controlled with the resource definition

## Integration with CLI

When these labels are present in a manifest, the ProjectPlanton CLI will:
1. First check for backend configuration in manifest labels
2. Fall back to command-line flags if labels are not present
3. Use defaults if neither labels nor flags are provided

This allows for flexible deployment scenarios:
```bash
# Backend config from manifest labels
project-planton pulumi update --manifest https://example.com/manifests/vpc.yaml

# Override with CLI flags
project-planton pulumi update --manifest vpc.yaml --stack my-org/my-project/dev
```

## Best Practices

1. **Use Stack FQDN**: Prefer the single `stack.fqdn` label over individual components
2. **Consistent Naming**: Follow your organization's naming conventions for organizations, projects, and stacks
3. **Environment Separation**: Use different stacks for different environments (dev, staging, prod)
4. **Documentation**: Document your organization's stack naming strategy

## Related Packages

- `pkg/iac/pulumi/backendconfig`: Extracts and processes these labels from manifests
- `pkg/iac/pulumi/pulumistack`: Uses the extracted configuration to run Pulumi operations
