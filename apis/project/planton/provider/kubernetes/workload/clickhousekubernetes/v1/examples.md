# ClickHouse Kubernetes API - Example Configurations

## Prerequisites

The **Altinity ClickHouse Operator** must be installed on your Kubernetes cluster before deploying ClickHouse instances. Deploy the operator using the `ClickhouseOperatorKubernetes` module.

## Create Using CLI

Create a YAML file using one of the examples below, then apply it using:

```shell
planton apply -f <yaml-path>
```

---

## Example w/ Basic Standalone Configuration

This example demonstrates a minimal standalone ClickHouse deployment suitable for development and testing.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: dev-clickhouse
spec:
  clusterName: dev-cluster
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 20Gi
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 2000m
        memory: 4Gi
```

---

## Example w/ Production Resources

Production-ready standalone deployment with larger resources and explicit version pinning.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: prod-clickhouse
  org: my-org
  env: production
spec:
  clusterName: production-analytics
  version: "24.8"
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 200Gi
    resources:
      requests:
        cpu: 2000m
        memory: 8Gi
      limits:
        cpu: 8000m
        memory: 32Gi
```

---

## Example w/ Distributed Cluster (Sharding)

Horizontally scaled cluster with multiple shards for parallel query processing.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: distributed-clickhouse
spec:
  clusterName: analytics-cluster
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    isPersistenceEnabled: true
    diskSize: 100Gi
    resources:
      requests:
        cpu: 1000m
        memory: 4Gi
      limits:
        cpu: 4000m
        memory: 16Gi
  cluster:
    isEnabled: true
    shardCount: 4
    replicaCount: 1
```

---

## Example w/ High Availability (Replication)

Production cluster with replication for high availability and fault tolerance.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: ha-clickhouse
spec:
  clusterName: ha-analytics
  version: "24.8"
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    isPersistenceEnabled: true
    diskSize: 150Gi
    resources:
      requests:
        cpu: 2000m
        memory: 8Gi
      limits:
        cpu: 6000m
        memory: 24Gi
  cluster:
    isEnabled: true
    shardCount: 3
    replicaCount: 3
```

**Key points**:
- 3 shards with 3 replicas each (9 total nodes)
- High availability - survives 2 node failures per shard
- ZooKeeper automatically managed by operator
- Query load balanced across replicas

---

## Example w/ External ZooKeeper

Advanced configuration using external ZooKeeper for shared coordination across multiple ClickHouse clusters.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: enterprise-clickhouse
spec:
  clusterName: enterprise-cluster
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    isPersistenceEnabled: true
    diskSize: 200Gi
    resources:
      requests:
        cpu: 4000m
        memory: 16Gi
      limits:
        cpu: 12000m
        memory: 48Gi
  cluster:
    isEnabled: true
    shardCount: 6
    replicaCount: 2
  zookeeper:
    useExternal: true
    nodes:
      - "zk-0.zookeeper.default.svc.cluster.local:2181"
      - "zk-1.zookeeper.default.svc.cluster.local:2181"
      - "zk-2.zookeeper.default.svc.cluster.local:2181"
```

---

## Example w/ Ingress Enabled

Cluster exposed to external clients via LoadBalancer with DNS.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: public-clickhouse
spec:
  clusterName: public-cluster
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 100Gi
    resources:
      requests:
        cpu: 1000m
        memory: 4Gi
      limits:
        cpu: 4000m
        memory: 16Gi
  ingress:
    isEnabled: true
    dnsDomain: example.com
```

**Key points**:
- LoadBalancer service for external access
- External DNS: `public-clickhouse.example.com`
- Both HTTP (8123) and native (9000) ports exposed

---

## Example w/ Development Environment

Lightweight configuration for testing with minimal resources and ephemeral storage.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: dev-clickhouse
spec:
  clusterName: test-cluster
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: false
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 500m
        memory: 2Gi
```

**Key points**:
- No persistent storage - data lost on restart
- Minimal resources for quick testing
- Fast startup and teardown

---

## Deployment Architecture

All deployments use the **Altinity ClickHouse Operator** which:
- Creates ClickHouseInstallation custom resources
- Manages StatefulSets, Services, and ConfigMaps
- Handles ZooKeeper coordination for clusters
- Provides rolling upgrades and self-healing
- Uses official ClickHouse container images

For detailed architecture information, see [Pulumi Module Overview](iac/pulumi/overview.md).
