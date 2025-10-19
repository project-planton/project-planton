Here are a few examples for the `GcpCloudFunction` API resource, showcasing different configurations for deploying Google Cloud Functions. These examples demonstrate basic configurations, environment variables, and use of secrets.

---

### Basic Example

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: my-cloud-function
spec:
  gcpProviderConfigId: my-gcp-credentials
  gcpProjectId: my-gcp-project
```

---

### Example with Environment Variables

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: my-cloud-function-env-vars
spec:
  gcpProviderConfigId: my-gcp-credentials
  gcpProjectId: my-gcp-project
  environmentVariables:
    DATABASE_NAME: my_database
    API_KEY: abc123
```

---

### Example with Secrets from GCP Secret Manager

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: my-cloud-function-secrets
spec:
  gcpProviderConfigId: my-gcp-credentials
  gcpProjectId: my-gcp-project
  environmentSecrets:
    DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.database-password}
    API_KEY: ${gcpsm-my-org-prod-gcp-secrets.api-key}
```

---

### Example with an Empty Spec

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudFunction
metadata:
  name: empty-spec-function
spec: {}
```
