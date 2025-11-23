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

## Example 2: Deployment with External PostgreSQL Database

Deploy Temporal using an external PostgreSQL database for production-grade scalability and reliability.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesTemporal
metadata:
  name: temporal-postgres
spec:
  target_cluster:
    cluster_name: prod-gke-cluster
  namespace:
    value: temporal-prod
  database:
    backend: postgresql
    external_database:
      host: postgres-db.example.com
      port: 5432
      username: temporaluser
      password: securepassword
  ingress:
    frontend:
      enabled: true
      grpc_hostname: temporal-frontend-prod.example.com
    web_ui:
      enabled: true
      hostname: temporal-ui-prod.example.com
```

**Use Case:**

- Suited for production environments requiring robust database management, backup strategies, and operational isolation.

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
  database:
    backend: cassandra
  external_elasticsearch:
    host: elasticsearch.example.com
    port: 9200
    user: elasticuser
    password: elasticpassword
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

