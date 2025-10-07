# ClickhouseKubernetes - Example Configurations

This document provides a series of examples demonstrating various configurations of the **ClickhouseKubernetes** API resource. Each example shows a typical use case, with corresponding YAML that can be applied via `planton apply -f <filename>`.

---

## 1. Minimal Configuration

A simple example deploying ClickHouse with default settings for resource allocation and persistence enabled.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: minimal-clickhouse
spec:
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

**Key points**:
- Single replica with 8Gi persistent volume
- Default resource allocations suitable for development/testing
- No external ingress configured

---

## 2. Production Configuration with Larger Resources

Demonstrates a production-ready ClickHouse deployment with substantial resource allocations for high-performance workloads.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: production-clickhouse
spec:
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

**Key points**:
- Larger persistent volume (100Gi) for storing substantial datasets
- Higher CPU and memory allocations for better query performance
- Still single replica (standalone mode)

---

## 3. Clustered Deployment with Sharding and Replication

Example of a distributed ClickHouse cluster with sharding for horizontal scalability and replication for high availability.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: clustered-clickhouse
spec:
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

**Key points**:
- Clustering enabled with 3 shards and 2 replicas per shard
- Distributed query processing across multiple nodes
- ZooKeeper automatically deployed for cluster coordination
- Suitable for large-scale analytics workloads

---

## 4. Custom Helm Values Configuration

Demonstrates how to customize the ClickHouse deployment using Helm chart values for advanced configurations.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: custom-clickhouse
spec:
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

**Key points**:
- Custom username configuration
- Override ClickHouse server configuration parameters
- Fine-tune connection and query limits
- Leverage full power of Bitnami ClickHouse Helm chart

---

## 5. Ingress-Enabled Deployment for External Access

In this example, ingress is enabled to allow external access to the ClickHouse cluster via a LoadBalancer service.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: ingress-clickhouse
spec:
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
    dnsDomain: example.com
```

**Key points**:
- LoadBalancer service created with external DNS annotation
- External hostname: `ingress-clickhouse.example.com`
- Internal hostname also configured for in-cluster access
- Both HTTP (8123) and native (9000) ports exposed

---

## 6. Development Environment with Minimal Resources

This example shows a minimal configuration suitable for local development or testing with reduced resource allocations and no persistence.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: dev-clickhouse
spec:
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

**Key points**:
- No persistence (ephemeral storage)
- Minimal resource allocation
- Quick startup for development/testing
- Data lost on pod restart

---

## 7. High Availability Setup

Example of a highly available ClickHouse deployment with multiple replicas and adequate resources.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: ha-clickhouse
spec:
  container:
    replicas: 3
    isPersistenceEnabled: true
    diskSize: 50Gi
    resources:
      requests:
        cpu: 1000m
        memory: 4Gi
      limits:
        cpu: 4000m
        memory: 16Gi
  cluster:
    isEnabled: true
    shardCount: 2
    replicaCount: 3
```

**Key points**:
- 2 shards with 3 replicas each (total 6 nodes)
- High resource allocation for production workloads
- Data replicated across nodes for redundancy
- Query load distributed across shards

---

## 8. Using Alternative Docker Registry

Due to Bitnami's registry changes, you can override the default bitnamilegacy registry to use official ClickHouse images or other sources.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: custom-registry-clickhouse
spec:
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
  helmValues:
    image.registry: docker.io
    image.repository: clickhouse/clickhouse-server
    image.tag: "24.8"
```

**Key points**:
- Override default bitnamilegacy registry with official ClickHouse images
- Use official vendor images directly from ClickHouse
- Provides long-term stability without Bitnami dependency
- Custom image tags for version control

---

## Conclusion

These examples illustrate the breadth of **ClickhouseKubernetes** features, from basic single-node deployments to advanced distributed clusters with sharding and replication. By consolidating Kubernetes manifests behind a concise API resource definition, you can maintain consistency, reduce error-prone manual config, and accelerate delivery cycles.

> **Getting Started**
> 1. Create a YAML file for your ClickHouse cluster (e.g., `clickhouse.yaml`).
> 2. Run:
>    ```shell
>    planton apply -f clickhouse.yaml
>    ```
> 3. Verify the deployment:
>    ```shell
>    kubectl get pods -n <namespace>
>    kubectl logs -n <namespace> <pod-name>
>    ```

For additional details, see the [ClickhouseKubernetes API documentation](../README.md), or reach out to our support team.
