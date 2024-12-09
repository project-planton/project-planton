## Usage

### Sample YAML Configuration

Create a YAML file (`artifact-registry.yaml`) with the desired configuration:

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpArtifactRegistry
metadata:
  id: my-artifact-registry
spec:
  gcp_credential_id: your-gcp-credential-id
  project_id: your-gcp-project-id
  region: us-central1
  is_external: false
```

### Deploying with CLI

Use the provided CLI tool to deploy the Artifact Registry repositories:

```bash
platon pulumi up --stack-input artifact-registry.yaml
```

If no Pulumi module is specified, the CLI uses the default module corresponding to the API resource.
