# ClickhouseKubernetes - Example Configurations

This document provides examples demonstrating various configurations of the **ClickhouseKubernetes** API resource using the Altinity ClickHouse Operator. Each example shows a typical use case with corresponding YAML that can be applied via `planton pulumi up`.

---

## Prerequisites

The **Altinity ClickHouse Operator** must be installed on your cluster before deploying these examples. See the main [README.md](README.md) for installation instructions.

---

## 1. Minimal Standalone Instance

A simple single-node ClickHouse deployment suitable for development and testing.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: dev-clickhouse
spec:
  clusterName: dev-cluster
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

**Key points**:
- Single replica with 20Gi persistent volume
- Suitable resource allocations for development
- Operator manages all Kubernetes resources
- No clustering - standalone mode

---

## 2. Production Standalone with Custom Cluster Name

Production-ready single-node deployment with larger resources and explicit cluster naming.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: prod-clickhouse
  org: my-org
  env: production
spec:
  clusterName: production-analytics
  version: "24.3"
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

**Key points**:
- Custom cluster name: `production-analytics`
- Pinned ClickHouse version for stability
- Large persistent volume (200Gi)
- High resource allocation for production workloads
- Organization and environment metadata for tracking

---

## 3. Distributed Cluster with Sharding

Horizontally scaled cluster with multiple shards for parallel query processing.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: distributed-analytics
spec:
  clusterName: analytics-cluster
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

**Key points**:
- 4 shards for horizontal scaling
- 1 replica per shard (4 total nodes)
- ZooKeeper automatically managed by operator
- Distributed query execution across shards
- Each shard processes subset of data

---

## 4. High Availability Cluster with Replication

Production cluster with replication for high availability and fault tolerance.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: ha-clickhouse
spec:
  clusterName: ha-analytics
  version: "24.3"
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
- Data automatically replicated across replicas
- Query load balanced across replicas
- Operator manages ZooKeeper for coordination

---

## 5. Cluster with External ZooKeeper

Advanced configuration using external ZooKeeper for shared coordination across multiple ClickHouse clusters.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: enterprise-clickhouse
spec:
  clusterName: enterprise-cluster
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

**Key points**:
- External ZooKeeper for production-grade coordination
- 6 shards with 2 replicas (12 total nodes)
- Enterprise-scale resources
- Shared ZooKeeper can coordinate multiple ClickHouse clusters
- Better isolation and control over ZooKeeper

---

## 6. Development Cluster with Ephemeral Storage

Lightweight cluster for testing and development with no persistence.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: test-clickhouse
spec:
  clusterName: test-cluster
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
- Cost-effective for temporary workloads

---

## 7. Ingress-Enabled Cluster with External Access

Cluster exposed to external clients via LoadBalancer with DNS.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: public-analytics
spec:
  clusterName: public-cluster
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
- External DNS: `public-analytics.example.com`
- Both HTTP (8123) and native (9000) ports exposed
- Suitable for external analytics tools and dashboards
- Consider security implications (firewall, authentication)

---

## 8. Multi-Environment with Custom Versions

Example showing version pinning for different environments.

```yaml
# Staging Environment
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: staging-clickhouse
  env: staging
spec:
  clusterName: staging-analytics
  version: "24.4"  # Testing newer version
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 50Gi
    resources:
      requests:
        cpu: 500m
        memory: 2Gi
      limits:
        cpu: 2000m
        memory: 8Gi

---

# Production Environment
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: production-clickhouse
  env: production
spec:
  clusterName: production-analytics
  version: "24.3"  # Stable version
  container:
    isPersistenceEnabled: true
    diskSize: 200Gi
    resources:
      requests:
        cpu: 2000m
        memory: 8Gi
      limits:
        cpu: 8000m
        memory: 32Gi
  cluster:
    isEnabled: true
    shardCount: 4
    replicaCount: 2
```

**Key points**:
- Different versions for testing vs production
- Staging validates new versions before production
- Environment-specific resource allocations
- Consistent configuration structure across environments

---

## Deployment Instructions

1. **Ensure Operator is Installed**:
   ```bash
   kubectl get pods -n clickhouse-operator
   ```

2. **Apply Configuration**:
   ```bash
   planton pulumi up --stack-input clickhouse.yaml \
     --module-dir apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1/iac/pulumi
   ```

3. **Verify Deployment**:
   ```bash
   # Check ClickHouseInstallation
   kubectl get clickhouseinstallations -n <namespace>
   
   # Check pods
   kubectl get pods -n <namespace>
   
   # Check services
   kubectl get svc -n <namespace>
   ```

4. **Access ClickHouse**:
   ```bash
   # Use port-forward command from stack outputs
   kubectl port-forward -n <namespace> service/<cluster-name> 8123:8123
   
   # Connect with clickhouse-client
   clickhouse-client --host localhost --port 8123
   ```

---

## Best Practices

### Resource Sizing
- **Development**: 500m CPU, 1-2Gi memory
- **Production standalone**: 2-4 CPU cores, 8-16Gi memory  
- **Production clustered**: 4-8 CPU cores, 16-48Gi memory per node

### Storage Planning
- **Small datasets** (<100GB): 20-50Gi per node
- **Medium datasets** (100GB-1TB): 100-500Gi per node
- **Large datasets** (>1TB): 500Gi-2Ti per node, consider sharding

### Clustering Strategy
- **Development/Testing**: Standalone (1 node)
- **Small production**: 2-3 shards, 2 replicas (4-6 nodes)
- **Large production**: 4-8 shards, 2-3 replicas (8-24 nodes)

### Version Management
- Pin specific versions in production
- Test new versions in staging first
- Follow Altinity/ClickHouse release notes for breaking changes

---

## Troubleshooting

### Check Operator Logs
```bash
kubectl logs -n clickhouse-operator deployment/clickhouse-operator
```

### Check ClickHouseInstallation Status
```bash
kubectl describe clickhouseinstallation <cluster-name> -n <namespace>
```

### Check ClickHouse Logs
```bash
kubectl logs -n <namespace> <pod-name>
```

### Verify ZooKeeper Connectivity (Clustered Mode)
```bash
kubectl exec -n <namespace> <pod-name> -- clickhouse-client --query "SELECT * FROM system.zookeeper WHERE path='/'"
```

---

For additional details and advanced configurations, see the [ClickhouseKubernetes API documentation](../README.md) and the [Altinity Operator documentation](https://github.com/Altinity/clickhouse-operator).
