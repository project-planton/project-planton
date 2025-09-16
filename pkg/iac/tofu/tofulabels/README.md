# Tofu Labels Package

This package defines standardized Kubernetes labels for configuring Terraform/OpenTofu backend state management directly within ProjectPlanton resource manifests.

## Overview

The `tofulabels` package provides constant definitions for labels that can be applied to any ProjectPlanton resource manifest to specify backend configuration for Terraform/OpenTofu operations. This enables infrastructure deployments to be fully portable with backend configuration embedded in the manifest.

## Label Constants

### Backend Type Label

- **`BackendTypeLabelKey`** (`terraform.project-planton.org/backend.type`)
  - Specifies the type of backend to use
  - Supported values: `s3`, `gcs`, `azurerm`, `local`
  - Example: `s3`

### Backend Object Label

- **`BackendObjectLabelKey`** (`terraform.project-planton.org/backend.object`)
  - Specifies the backend-specific object path or identifier
  - Format varies by backend type:
    - **S3**: `bucket-name/path/to/state`
    - **GCS**: `bucket-name/path/to/state`
    - **Azure**: `container-name/path/to/state`
    - **Local**: `/absolute/path/to/state.tfstate`

## Usage Examples

### AWS S3 Backend

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: app-database
  labels:
    terraform.project-planton.org/backend.type: "s3"
    terraform.project-planton.org/backend.object: "terraform-states-bucket/rds/production/app-db"
spec:
  engine: "postgres"
  instanceClass: "db.t3.medium"
```

### Google Cloud Storage Backend

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudRun
metadata:
  name: api-service
  labels:
    terraform.project-planton.org/backend.type: "gcs"
    terraform.project-planton.org/backend.object: "my-tfstate-bucket/cloud-run/prod/api"
spec:
  region: "us-central1"
  image: "gcr.io/project/api:latest"
```

### Azure Storage Backend

```yaml
apiVersion: azure.project-planton.org/v1
kind: AzureAksCluster
metadata:
  name: main-cluster
  labels:
    terraform.project-planton.org/backend.type: "azurerm"
    terraform.project-planton.org/backend.object: "tfstate-container/aks/production"
spec:
  location: "eastus"
  nodeCount: 3
```

### Local Backend (Development)

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: test-service
  labels:
    terraform.project-planton.org/backend.type: "local"
    terraform.project-planton.org/backend.object: "/tmp/test-service.tfstate"
spec:
  replicas: 1
```

## Backend Type Details

### S3 Backend
- Used for AWS deployments
- Supports state locking via DynamoDB
- Object format: `bucket/prefix/resource-name`
- Additional configuration (region, encryption) handled by CLI/environment

### GCS Backend
- Used for Google Cloud deployments
- Automatic state locking
- Object format: `bucket/prefix/resource-name`
- Credentials via environment or service account

### Azure Storage Backend
- Used for Azure deployments
- Supports state locking via blob leases
- Object format: `container/prefix/resource-name`
- Requires storage account configuration

### Local Backend
- Development and testing only
- No locking support
- Full file path required
- Not recommended for production

## Label Requirements

1. **Both or Neither**: If one backend label is specified, both must be specified
2. **Non-Empty Values**: Neither label can have an empty value
3. **Valid Backend Type**: Must be one of the supported backend types
4. **Optional**: Both labels are optional - if not specified, CLI flags or defaults are used

## Benefits

1. **Portable Manifests**: Backend configuration travels with the resource definition
2. **Environment Isolation**: Different backends for dev/staging/prod
3. **Team Collaboration**: Shared state configuration in version control
4. **Disaster Recovery**: Backend location documented in manifest

## Integration with CLI

The ProjectPlanton CLI processes these labels to configure Terraform/OpenTofu backends:

```bash
# Uses backend config from manifest labels
project-planton tofu apply --manifest vpc.yaml

# Or from a URL with embedded backend config
project-planton tofu apply --manifest https://example.com/manifests/vpc.yaml

# CLI flags can override if needed
project-planton tofu apply --manifest vpc.yaml --backend-config s3://other-bucket/state
```

## Best Practices

1. **Consistent Paths**: Use a clear naming hierarchy (e.g., `environment/resource-type/resource-name`)
2. **Separate by Environment**: Use different buckets/containers for different environments
3. **Enable Versioning**: Always enable versioning on your state storage
4. **Access Control**: Restrict access to state files appropriately
5. **Backup Strategy**: Implement regular backups of state files

## Security Considerations

- Never include credentials in labels
- Use IAM roles or service accounts for backend authentication
- Enable encryption at rest for state files
- Implement proper access controls on state storage
- Consider using separate backends for sensitive resources

## Related Packages

- `pkg/iac/tofu/backendconfig`: Extracts and validates these labels from manifests
- `pkg/iac/tofu/runner`: Uses the extracted configuration to initialize backends
