Here are a few examples for the `GcpCloudSql` API resource, following a similar format as the ones you provided for `MicroserviceKubernetes`.

### Example 1: Basic Google Cloud SQL Instance

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudSql
metadata:
  name: my-sql-instance
spec:
  gcp_credential_id: my-gcp-credentials
  gcp_project_id: my-gcp-project
  instance:
    database_version: POSTGRES_13
    settings:
      tier: db-f1-micro
      backup_configuration:
        enabled: true
      activation_policy: ALWAYS
```

### Example 2: Google Cloud SQL Instance with Environment Variables

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudSql
metadata:
  name: ecommerce-db
spec:
  gcp_credential_id: my-gcp-credentials
  gcp_project_id: my-gcp-project
  instance:
    database_version: MYSQL_8_0
    settings:
      tier: db-n1-standard-1
      storage_auto_resize: true
      backup_configuration:
        enabled: true
      activation_policy: ALWAYS
  environmentInfo:
    envId: prod
  container:
    app:
      env:
        variables:
          DATABASE_NAME: ecommerce
          REGION: us-central1
```

### Example 3: Google Cloud SQL Instance with Secrets

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudSql
metadata:
  name: secure-sql-instance
spec:
  gcp_credential_id: my-gcp-credentials
  gcp_project_id: my-gcp-project
  instance:
    database_version: POSTGRES_12
    settings:
      tier: db-n1-standard-1
      backup_configuration:
        enabled: true
      activation_policy: ALWAYS
  container:
    app:
      env:
        secrets:
          DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.db-password}
        variables:
          DATABASE_NAME: secure_app
```

### Example 4: Minimal Google Cloud SQL Deployment

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcpCloudSql
metadata:
  name: minimal-sql-instance
spec:
  gcp_credential_id: my-gcp-credentials
  gcp_project_id: my-gcp-project
```
