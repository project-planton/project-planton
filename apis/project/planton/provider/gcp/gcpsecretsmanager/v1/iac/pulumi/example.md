Here are a few examples for the `GcpSecretsManager` API resource, modeled in a similar way to how you created examples for the `MicroserviceKubernetes` API.

### Example 1: Basic Google Cloud Secrets Manager Setup

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpSecretsManager
metadata:
  name: prod-secrets
spec:
  gcp_credential_id: my-gcp-credentials
  project_id: my-gcp-project
  secret_names:
    - database-password
    - api-key
```

### Example 2: Google Cloud Secrets Manager with Multiple Secrets

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpSecretsManager
metadata:
  name: dev-secrets
spec:
  gcp_credential_id: my-dev-gcp-credentials
  project_id: dev-gcp-project
  secret_names:
    - jwt-secret
    - database-url
    - oauth-token
```

### Example 3: Google Cloud Secrets Manager with Empty Spec

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpSecretsManager
metadata:
  name: minimal-secrets
spec:
  gcp_credential_id: my-gcp-credentials
  project_id: my-gcp-project
```
