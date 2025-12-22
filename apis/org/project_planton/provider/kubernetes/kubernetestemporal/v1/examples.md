# Temporal on Kubernetes Examples

Below are practical examples demonstrating how to use the `TemporalKubernetes` component within your ProjectPlanton
deployments. These examples illustrate common configurations and scenarios to help you quickly integrate Temporal
workflows into your Kubernetes environments.

---

## Example 1: Basic Deployment with Default Cassandra Database

This example shows a simple Temporal deployment using an internal Cassandra database and exposing Temporal frontend and
web UI through ingress.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-basic
spec:
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: temporal-basic
  create_namespace: true
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend.example.com
    web_ui:
      enabled: true
      hostname: temporal-ui.example.com
```

**Use Case:**

- Ideal for quickly setting up Temporal workflows for development or testing purposes without external database
  dependencies.

---

## Example 2: Deployment with External PostgreSQL Database (Using String Value)

Deploy Temporal using an external PostgreSQL database for development/testing with a plain text password.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-postgres-dev
spec:
  target_cluster:
    cluster_name: dev-gke-cluster
  namespace:
    value: temporal-dev
  create_namespace: true
  database:
    backend: postgresql
    external_database:
      host: postgres-db.example.com
      port: 5432
      username: temporaluser
      password:
        stringValue: securepassword
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend-dev.example.com
    web_ui:
      enabled: true
      hostname: temporal-ui-dev.example.com
```

**Use Case:**

- Suitable for development and testing environments where plain text passwords are acceptable.

---

## Example 2b: Deployment with External PostgreSQL Database (Using Secret Reference - Recommended)

Deploy Temporal using an external PostgreSQL database for production with a password stored in a Kubernetes Secret.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-postgres-prod
spec:
  target_cluster:
    cluster_name: prod-gke-cluster
  namespace:
    value: temporal-prod
  create_namespace: true
  database:
    backend: postgresql
    external_database:
      host: postgres-db.example.com
      port: 5432
      username: temporaluser
      password:
        secretRef:
          name: temporal-db-credentials
          key: password
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend-prod.example.com
    web_ui:
      enabled: true
      hostname: temporal-ui-prod.example.com
```

**Use Case:**

- Recommended for production environments where credentials should be stored securely in Kubernetes Secrets.
- Enables GitOps workflows where manifests can be safely committed to version control.
- Simplifies password rotation by only updating the Kubernetes Secret.

---

## Example 3: Advanced Observability with External Elasticsearch

Deploy Temporal with advanced observability features using an existing Elasticsearch cluster to power visibility
queries.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-elasticsearch
spec:
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: temporal-observability
  create_namespace: true
  database:
    backend: cassandra
  external_elasticsearch:
    host: elasticsearch.example.com
    port: 9200
    user: elasticuser
    password:
      stringValue: elasticpassword
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend.example.com
    web_ui:
      enabled: true
      hostname: temporal-ui.example.com
```

**Use Case:**

- Recommended for scenarios demanding advanced search and visibility capabilities for workflows, typically in
  monitoring-intensive production deployments.

---

## Example 3b: External Elasticsearch with Secret Reference

Deploy Temporal with external Elasticsearch using credentials stored in Kubernetes Secrets.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-elasticsearch-prod
spec:
  target_cluster:
    cluster_name: prod-gke-cluster
  namespace:
    value: temporal-observability
  create_namespace: true
  database:
    backend: postgresql
    external_database:
      host: postgres-db.example.com
      port: 5432
      username: temporaluser
      password:
        secretRef:
          name: temporal-db-credentials
          key: password
  external_elasticsearch:
    host: elasticsearch.example.com
    port: 9200
    user: elasticuser
    password:
      secretRef:
        name: es-credentials
        key: password
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend-prod.example.com
    web_ui:
      enabled: true
      hostname: temporal-ui-prod.example.com
```

**Use Case:**

- Production deployments with both external database and Elasticsearch using secure credential management.
- All sensitive values are stored in Kubernetes Secrets, enabling safe GitOps workflows.

---

## Example 4: Minimal Resource Deployment (Web UI Disabled)

Temporal deployment with minimal resource footprint by disabling the Temporal web UI component.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-minimal
spec:
  database:
    backend: mysql
    externalDatabase:
      host: mysql-db.example.com
      port: 3306
      user: temporaluser
      password: securepassword
  disableWebUi: true
  ingress:
    frontend:
      enabled: true
      hostname: temporal-frontend.example.com
```

**Use Case:**

- Optimal for environments where the UI is managed separately or unnecessary, reducing resource consumption.

---

## Example 5: Schema Control and Database Naming

Example demonstrating custom database naming conventions and controlled database schema initialization.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-custom-schema
spec:
  database:
    backend: postgresql
    externalDatabase:
      host: postgres-db.example.com
      port: 5432
      user: temporaluser
      password: securepassword
    databaseName: custom_temporal_db
    visibilityName: custom_temporal_visibility_db
    disableAutoSchemaSetup: true
  ingress:
    enabled: false
```

**Use Case:**

- Useful when integrating with existing database management procedures, enforcing naming conventions, or preventing
  automatic schema alterations.

---

## Namespace Management

### Example 6: Using Existing Namespace

Deploy Temporal into an existing namespace without creating it:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-existing-ns
spec:
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: existing-temporal-namespace
  create_namespace: false  # Use existing namespace
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend.example.com
```

**Use Case:**

- When namespace is pre-created with specific policies, quotas, or RBAC
- In environments with strict namespace governance
- When multiple components share the same namespace

### Example 7: Creating New Namespace (Default Behavior)

Create a new namespace for Temporal deployment:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-new-ns
spec:
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: temporal-namespace
  create_namespace: true  # Create new namespace (this is the default)
  database:
    backend: cassandra
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend.example.com
```

**Note:** If `create_namespace` is omitted, it defaults to `true` for backward compatibility. The module will create
the namespace with appropriate labels and metadata.

