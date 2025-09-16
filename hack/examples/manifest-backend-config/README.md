# Manifest Backend Configuration Examples

This directory contains example manifests demonstrating how to embed backend configuration directly in manifest labels, eliminating the need for CLI flags.

## Features

- **Self-contained manifests**: Deploy directly from URLs without additional configuration
- **Pulumi backend support**: Configure Pulumi stack information via labels
- **Terraform/Tofu backend support**: Configure remote state storage via labels
- **Backward compatible**: CLI flags still work and can override manifest labels

## Pulumi Backend Configuration

### Using Stack FQDN (Recommended)

```yaml
metadata:
  labels:
    pulumi.project-planton.org/stack.fqdn: "myorg/project/stack"
```

### Using Individual Components

```yaml
metadata:
  labels:
    pulumi.project-planton.org/organization: "myorg"
    pulumi.project-planton.org/project: "my-project"
    pulumi.project-planton.org/stack.name: "production"
```

## Terraform/Tofu Backend Configuration

### S3 Backend (AWS)

```yaml
metadata:
  labels:
    terraform.project-planton.org/backend.type: "s3"
    terraform.project-planton.org/backend.object: "bucket-name/path/to/state.tfstate"
```

### GCS Backend (Google Cloud)

```yaml
metadata:
  labels:
    terraform.project-planton.org/backend.type: "gcs"
    terraform.project-planton.org/backend.object: "bucket-name/prefix/path"
```

### Azure Blob Storage Backend

```yaml
metadata:
  labels:
    terraform.project-planton.org/backend.type: "azurerm"
    terraform.project-planton.org/backend.object: "container-name/path/to/state"
```

## Usage Examples

### Deploy with Pulumi (No --stack flag needed)

```bash
# Deploy using manifest with embedded Pulumi backend config
project-planton pulumi update --manifest pulumi-backend-example.yaml

# Deploy from URL
project-planton pulumi update --manifest https://raw.githubusercontent.com/example/repo/main/manifest.yaml
```

### Deploy with Tofu/Terraform (Backend auto-configured)

```bash
# Apply using manifest with embedded Terraform backend config
project-planton tofu apply --manifest tofu-s3-backend-example.yaml

# Plan changes
project-planton tofu plan --manifest tofu-gcs-backend-example.yaml
```

## Priority Order

When both manifest labels and CLI flags are present:

1. **Pulumi**: Manifest labels take precedence over CLI flags
2. **Tofu/Terraform**: Manifest labels are used if present, otherwise defaults to local backend

## Important Notes

- Backend configuration via labels is optional
- If no backend configuration is found in labels, the system falls back to CLI flags or defaults
- This feature is ideal for:
  - Quick-start demos
  - CI/CD pipelines
  - Sharing deployable manifests
  - Production use with proper credential management

## Security Considerations

- Backend configuration in labels does NOT include credentials
- Credentials must still be provided via:
  - Environment variables
  - CLI flags (`--aws-credential-id`, `--gcp-credential-id`, etc.)
  - Default credential chains (AWS IAM roles, GCP service accounts, etc.)

## Debugging

Enable debug logging to see which backend configuration is being used:

```bash
# See backend configuration details
LOG_LEVEL=debug project-planton pulumi update --manifest example.yaml
```
