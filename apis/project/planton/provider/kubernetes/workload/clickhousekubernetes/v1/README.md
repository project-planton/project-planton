# Overview

The **ClickHouse Kubernetes API Resource** provides a production-grade, operator-based way to deploy and manage ClickHouse clusters on Kubernetes. This API resource uses the **Altinity ClickHouse Operator** to deliver enterprise-level features including automated upgrades, scaling, backup, and recovery.

## Purpose

Deploying ClickHouse on Kubernetes requires managing complex distributed systems with sharding, replication, and coordination. The ClickHouse Kubernetes API Resource aims to:

- **Simplify Operations**: Leverage the Altinity operator to handle lifecycle management, rolling upgrades, and failure recovery automatically.
- **Standardize Deployments**: Offer an intuitive, deployment-agnostic interface following the 80/20 Pareto principle - focus on what 80% of users need.
- **Enable Production-Grade Clusters**: Support both standalone and distributed deployments with configurable sharding and replication.
- **Provide Type Safety**: Use strongly-typed Kubernetes custom resources for compile-time validation and better developer experience.

## Key Features

### Environment Configuration

- **Environment Info**: Tailor ClickHouse deployments to specific environments (development, staging, production) using environment-specific information.
- **Stack Job Settings**: Integrate with infrastructure-as-code (IaC) tools through stack job settings for automated and repeatable deployments.

### Credential Management

- **Kubernetes Cluster Credential ID**: Specify credentials required to access and configure the target Kubernetes cluster securely.

### Cluster Configuration

- **Cluster Name**: Configurable cluster identifier used for the ClickHouseInstallation resource.
- **Version**: Pin specific ClickHouse versions (e.g., "24.8") for consistency and stability.
- **Replicas**: Define the number of ClickHouse pod instances for standalone deployments.
- **Resources**: Allocate CPU and memory resources for optimal performance.
  - Production defaults: 500m CPU (requests), 2000m (limits), 1Gi memory (requests), 4Gi (limits)
- **Persistence**:
  - **Enable Persistence**: Toggle data persistence (strongly recommended for production).
  - **Disk Size**: Specify persistent volume size (e.g., `50Gi`, `100Gi`). Plan for growth.

### Distributed Clustering

- **Cluster Mode**: Enable distributed ClickHouse deployments with sharding and replication.
- **Shard Count**: Define the number of shards for horizontal data distribution and parallel query processing.
- **Replica Count**: Specify the number of replicas per shard for high availability and data redundancy.

### Coordination Configuration

ClickHouse clusters require coordination services for distributed operations (DDL execution, replication management). The API supports multiple coordination options following the 80/20 principle:

- **Auto-Managed ClickHouse Keeper** (Recommended - 80% use case):
  - Default option when coordination is not explicitly configured
  - 75% more resource-efficient than ZooKeeper (no JVM overhead)
  - Managed by the same Altinity operator
  - Creates a ClickHouseKeeperInstallation resource automatically
  - Configurable replicas: 1 (dev), 3 (production), 5 (large production)
  
- **External ClickHouse Keeper** (Advanced scenarios):
  - Connect to existing ClickHouse Keeper infrastructure
  - Useful for shared coordination across multiple clusters
  - Supports multi-node ensembles for high availability

- **External ZooKeeper** (Legacy/Integration):
  - Connect to existing ZooKeeper infrastructure
  - Required when sharing ZooKeeper with other services (Kafka, Solr)
  - Supports traditional ZooKeeper ensembles

### Networking and Ingress

- **Ingress Configuration**: Set up Kubernetes Ingress resources to manage external access to ClickHouse, including hostname and path routing.

## Benefits

- **Production-Ready**: Leverage the battle-tested Altinity operator used by enterprises worldwide for operational excellence.
- **Self-Healing**: Automatic recovery from failures, rolling upgrades, and continuous reconciliation to maintain desired state.
- **Simplified Operations**: The operator handles complex lifecycle operations - you focus on your data, not Kubernetes complexity.
- **Resource Efficient**: Default ClickHouse Keeper coordination uses 75% less CPU and memory compared to ZooKeeper.
- **Scalability and Flexibility**: Easily scale from single nodes to distributed clusters with hundreds of nodes.
- **Data Persistence**: Production-grade persistent storage with automatic volume provisioning and management.
- **High Availability**: Native support for clustering with sharding, replication, and efficient coordination services.
- **Type Safety**: Strongly-typed configuration with compile-time validation and IDE support.
- **Smart Defaults**: 80/20 principle applied - most common configurations work with minimal specification.

## Use Cases

- **Real-time Analytics**: Deploy ClickHouse as a high-performance analytics database for real-time data processing.
- **Observability Platforms**: Backend for SigNoz, Grafana, and other observability tools requiring fast OLAP queries.
- **Data Warehousing**: Use ClickHouse for OLAP workloads and complex analytical queries on large datasets.
- **Log Analytics**: Process and analyze large volumes of log data with ClickHouse's columnar storage engine.
- **Time-Series Data**: Handle time-series data efficiently with ClickHouse's optimized storage and query capabilities.
- **Microservices Architecture**: Deploy ClickHouse instances for services requiring fast analytical queries.
- **Development and Testing Environments**: Quickly spin up ClickHouse instances for development or testing purposes with environment-specific configurations.

## Configuration Examples

### Simple Cluster (80% Use Case)
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickHouseKubernetes
metadata:
  name: my-clickhouse
spec:
  clusterName: my-cluster
  cluster:
    isEnabled: true
    shardCount: 2
    replicaCount: 2
  # coordination: not specified = auto-managed ClickHouse Keeper with defaults
```

### Production with Custom Keeper Configuration
```yaml
spec:
  cluster:
    isEnabled: true
    shardCount: 2
    replicaCount: 2
  coordination:
    type: keeper
    keeperConfig:
      replicas: 3  # Production HA
      diskSize: "20Gi"
      resources:
        requests:
          cpu: "200m"
          memory: "512Mi"
```

### External Keeper (Shared Infrastructure)
```yaml
spec:
  cluster:
    isEnabled: true
  coordination:
    type: external_keeper
    externalConfig:
      nodes:
        - "keeper-shared:2181"
```

### External ZooKeeper (Legacy Integration)
```yaml
spec:
  cluster:
    isEnabled: true
  coordination:
    type: external_zookeeper
    externalConfig:
      nodes:
        - "zk-0.zk.svc:2181"
        - "zk-1.zk.svc:2181"
        - "zk-2.zk.svc:2181"
```

## Migration from ZooKeeper Field (Deprecated)

The original `zookeeper` field is deprecated in favor of the more flexible `coordination` field:

**Old (Deprecated):**
```yaml
spec:
  zookeeper:
    useExternal: true
    nodes:
      - "zk-0:2181"
```

**New (Recommended):**
```yaml
spec:
  coordination:
    type: external_zookeeper
    externalConfig:
      nodes:
        - "zk-0:2181"
```

The deprecated field will continue to work for backward compatibility but will be removed in v2.
