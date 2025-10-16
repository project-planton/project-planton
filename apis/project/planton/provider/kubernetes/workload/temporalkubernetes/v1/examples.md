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
kind: TemporalKubernetes
metadata:
  name: temporal-basic
spec:
  database:
    backend: cassandra
  ingress:
    enabled: true
    host: temporal.example.com
```

**Use Case:**

- Ideal for quickly setting up Temporal workflows for development or testing purposes without external database
  dependencies.

---

## Example 2: Deployment with External PostgreSQL Database

Deploy Temporal using an external PostgreSQL database for production-grade scalability and reliability.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-postgres
spec:
  database:
    backend: postgresql
    externalDatabase:
      host: postgres-db.example.com
      port: 5432
      user: temporaluser
      password: securepassword
  ingress:
    enabled: true
    host: temporal-prod.example.com
```

**Use Case:**

- Suited for production environments requiring robust database management, backup strategies, and operational isolation.

---

## Example 3: Advanced Observability with External Elasticsearch

Deploy Temporal with advanced observability features using an existing Elasticsearch cluster to power visibility
queries.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-elasticsearch
spec:
  database:
    backend: cassandra
  externalElasticsearch:
    host: elasticsearch.example.com
    port: 9200
    user: elasticuser
    password: elasticpassword
  ingress:
    enabled: true
    host: temporal-advanced.example.com
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
    enabled: true
    host: temporal-minimal.example.com
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

## Example 6: Deployment with Custom Search Attributes

Deploy Temporal with custom search attributes for advanced workflow filtering and querying capabilities.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: TemporalKubernetes
metadata:
  name: temporal-with-search-attrs
spec:
  database:
    backend: postgresql
    externalDatabase:
      host: postgres.example.com
      port: 5432
      username: temporal_user
      password: secure_password
  
  externalElasticsearch:
    host: elasticsearch.example.com
    port: 9200
    user: elastic
    password: elastic_pass
  
  searchAttributes:
    - name: CustomerId
      type: Keyword
    - name: Environment
      type: Keyword
    - name: Priority
      type: Int
    - name: Amount
      type: Double
    - name: IsActive
      type: Bool
    - name: DeploymentDate
      type: Datetime
    - name: Tags
      type: KeywordList
  
  ingress:
    enabled: true
    host: temporal-search.example.com
```

**Use Case:**

- Essential for production environments requiring workflow filtering by business dimensions (customer ID, environment, priority, etc.)
- Enables building custom dashboards and analytics based on workflow metadata
- Supports both SQL and Elasticsearch visibility stores (Text type requires Elasticsearch)
