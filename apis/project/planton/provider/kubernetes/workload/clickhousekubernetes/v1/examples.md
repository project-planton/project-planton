# ClickHouse Kubernetes API - Example Configurations

## Example w/ Basic Configuration

### Create Using CLI

Create a YAML file using the example shown below. After the YAML is created, use the following command to apply it:

```shell
planton apply -f <yaml-path>
```

### Basic Example

This basic example demonstrates a minimal configuration for deploying a ClickHouse Kubernetes instance using the default settings, including 1 replica with persistence enabled.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: basic-clickhouse
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 8Gi
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 2Gi
```

---

## Example w/ Larger Resources

In this example, ClickHouse is configured with more substantial resource allocations suitable for production workloads with higher performance requirements.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: production-clickhouse
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 100Gi
    resources:
      requests:
        cpu: 500m
        memory: 2Gi
      limits:
        cpu: 4000m
        memory: 8Gi
```

---

## Example w/ Clustering Enabled

This example demonstrates how to deploy a distributed ClickHouse cluster with sharding and replication for high availability and horizontal scalability.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: clustered-clickhouse
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 3
    isPersistenceEnabled: true
    diskSize: 50Gi
    resources:
      requests:
        cpu: 500m
        memory: 2Gi
      limits:
        cpu: 2000m
        memory: 4Gi
  cluster:
    isEnabled: true
    shardCount: 3
    replicaCount: 2
```

---

## Example w/ Custom Helm Values

This example demonstrates how to customize the ClickHouse deployment using Helm chart values to configure specific ClickHouse settings.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: custom-clickhouse
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 2
    isPersistenceEnabled: true
    diskSize: 30Gi
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 4Gi
  helmValues:
    auth.username: clickhouse_user
    defaultConfigurationOverrides: |
      <clickhouse>
        <max_connections>100</max_connections>
        <max_concurrent_queries>50</max_concurrent_queries>
      </clickhouse>
```

---

## Example w/ Ingress Enabled

In this example, ingress is enabled to allow external access to the ClickHouse service. This is particularly useful when ClickHouse needs to be accessed by clients outside the Kubernetes cluster.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: ingress-clickhouse
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 20Gi
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 4Gi
  ingress:
    isEnabled: true
    dnsDomain: clickhouse.example.com
```

---

## Example w/ Minimal Resources for Development

This example shows a minimal configuration suitable for development or testing environments with reduced resource allocations.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: dev-clickhouse
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: false
    resources:
      requests:
        cpu: 50m
        memory: 128Mi
      limits:
        cpu: 500m
        memory: 512Mi
```
