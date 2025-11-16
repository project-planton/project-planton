# Pulumi Module for KubernetesSolrOperator

This Pulumi module deploys the Apache Solr Operator to a Kubernetes cluster, enabling declarative management of SolrCloud instances through Kubernetes Custom Resource Definitions.

## Overview

The Apache Solr Operator is the official Kubernetes operator for managing Apache SolrCloud deployments. This Pulumi module handles:

- Installing Solr Operator Custom Resource Definitions (CRDs)
- Deploying the operator via Helm chart
- Configuring operator namespace and resources
- Exporting relevant stack outputs

## Prerequisites

- **Pulumi CLI**: Version 3.0 or later
- **Kubernetes Cluster**: Access to a cluster with kubectl configured
- **Kubernetes Provider Config**: Valid Project Planton Kubernetes credential

## Project Structure

```
iac/pulumi/
├── main.go              # Entry point for Pulumi program
├── Pulumi.yaml          # Pulumi project configuration
├── Makefile            # Common commands (install, preview, deploy)
├── README.md           # This file
├── overview.md         # Architecture and design decisions
├── debug.sh            # Debug script for local testing
└── module/
    ├── main.go         # Core resource creation logic
    ├── locals.go       # Computed values and derived config
    ├── vars.go         # Constants and version defaults
    └── outputs.go      # Output constant definitions
```

## Quick Start

### 1. Configure Stack

Create a Pulumi stack configuration file (e.g., `Pulumi.dev.yaml`):

```yaml
config:
  provider-config: k8s-cluster-credential-id
```

Or set via CLI:

```bash
pulumi config set provider-config k8s-cluster-credential-id
```

### 2. Preview Changes

```bash
pulumi preview
```

### 3. Deploy

```bash
pulumi up
```

### 4. View Outputs

```bash
pulumi stack output namespace
```

## Configuration

### Required

- **`provider-config`**: Kubernetes cluster credential ID (string)

### Optional

No additional configuration is currently required. The operator uses default resource limits and chart versions.

## Helm Chart Details

The module deploys the official Apache Solr Operator Helm chart:

- **Chart Repository**: https://solr.apache.org/charts
- **Chart Name**: `solr-operator`
- **Default Version**: `0.7.0` (configurable in `module/vars.go`)
- **CRD Version**: `v0.7.0`

### Installed CRDs

The module installs these Custom Resource Definitions:

1. **SolrCloud**: Defines a SolrCloud cluster
2. **SolrBackup**: Manages backup operations
3. **SolrPrometheusExporter**: Deploys Prometheus metrics exporters

## Deployment Process

1. **Setup Kubernetes Provider**: Connects to target cluster using provided credentials
2. **Create Namespace**: Creates `solr-operator` namespace
3. **Install CRDs**: Applies CRDs from upstream manifest
4. **Deploy Helm Chart**: Installs operator with dependency on CRDs
5. **Export Outputs**: Makes namespace available as stack output

## Outputs

After deployment, the stack exports:

| Output | Type | Description |
|--------|------|-------------|
| `namespace` | string | Kubernetes namespace containing the operator |

Access outputs:

```bash
# All outputs
pulumi stack output

# Specific output
pulumi stack output namespace
```

## Managing the Operator

### Checking Operator Status

```bash
# Get operator pod
kubectl get pods -n solr-operator

# View logs
kubectl logs -n solr-operator <pod-name>

# Check CRDs
kubectl get crds | grep solr
```

### Creating SolrCloud Clusters

After the operator is deployed, create SolrCloud clusters:

```yaml
apiVersion: solr.apache.org/v1beta1
kind: SolrCloud
metadata:
  name: example
  namespace: solr-production
spec:
  replicas: 3
  solrImage:
    repository: solr
    tag: 8.11.3
  zookeeperRef:
    provided:
      replicas: 3
```

Apply with:

```bash
kubectl apply -f solrcloud.yaml
```

## Upgrading

### Operator Version

Update the chart version in `module/vars.go`:

```go
DefaultStableVersion: "0.8.0",  // Update this
```

Then redeploy:

```bash
pulumi up
```

### CRDs

CRD upgrades require updating the `CrdManifestDownloadURL` in `module/vars.go`:

```go
CrdManifestDownloadURL: "https://solr.apache.org/operator/downloads/crds/v0.8.0/all-with-dependencies.yaml",
```

**Note**: CRD upgrades may require destroying and recreating SolrCloud resources depending on API changes.

## Troubleshooting

### CRDs Not Installing

If CRDs fail to install:

```bash
# Check CRD manifest URL is accessible
curl https://solr.apache.org/operator/downloads/crds/v0.7.0/all-with-dependencies.yaml

# Manually apply CRDs
kubectl apply -f https://solr.apache.org/operator/downloads/crds/v0.7.0/all-with-dependencies.yaml
```

### Helm Chart Fails to Deploy

```bash
# Check Helm repository
helm repo add solr https://solr.apache.org/charts
helm repo update

# Verify chart exists
helm search repo solr/solr-operator
```

### Pulumi State Issues

```bash
# Refresh state
pulumi refresh

# View detailed logs
pulumi up --logtostderr --logflow -v=9
```

## Development

### Local Testing

Use the `debug.sh` script:

```bash
./debug.sh
```

This runs Pulumi in debug mode with verbose logging.

### Module Development

When modifying `module/` files:

1. Update Go code
2. Test with `pulumi preview`
3. Verify with `pulumi up`
4. Check operator deployment in cluster

### Adding Custom Values

To pass custom Helm values, modify `module/main.go`:

```go
Values: pulumi.Map{
	"resources": pulumi.Map{
		"limits": pulumi.Map{
			"cpu":    pulumi.String("500m"),
			"memory": pulumi.String("512Mi"),
		},
	},
},
```

## Resource Requirements

The operator pod uses these default limits (from Helm chart):

- **CPU**: 100m (request), 500m (limit)
- **Memory**: 64Mi (request), 256Mi (limit)

These are separate from the KubernetesSolrOperatorSpec resource limits which are not currently used (operator deployment uses Helm chart defaults).

## Best Practices

1. **Pin Versions**: Specify exact chart versions in `vars.go`
2. **Namespace Isolation**: Keep operator in dedicated namespace
3. **Monitor Deployments**: Use `pulumi watch` for continuous monitoring
4. **Stack Per Environment**: Use separate Pulumi stacks for dev/staging/prod
5. **State Backend**: Use remote state backend (S3, GCS, etc.) for team collaboration

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Deploy Solr Operator
on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: pulumi/actions@v4
        with:
          command: up
          stack-name: production
        env:
          PULUMI_ACCESS_TOKEN: ${{ secrets.PULUMI_ACCESS_TOKEN }}
```

## Additional Resources

- **Architecture Overview**: [overview.md](overview.md)
- **Component Documentation**: [../../README.md](../../README.md)
- **Apache Solr Operator Docs**: https://apache.github.io/solr-operator/
- **Pulumi Kubernetes Provider**: https://www.pulumi.com/docs/reference/pkg/kubernetes/

## Support

For issues related to:
- **Module bugs**: File issue on Project Planton repository
- **Operator behavior**: Check Apache Solr Operator documentation
- **Pulumi questions**: Consult Pulumi documentation

## License

This Pulumi module is part of Project Planton. The Apache Solr Operator is licensed under Apache License 2.0.

