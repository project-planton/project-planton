# Overview

The **ClickhouseKubernetes** API resource streamlines and standardizes how developers deploy ClickHouse database clusters onto Kubernetes. This Pulumi module interprets a `ClickhouseKubernetesStackInput`, which includes core details like Kubernetes credentials and the ClickHouse configuration. From this input, the module automatically creates and manages:

- **Kubernetes Namespaces**
- **Helm Chart Deployments** for ClickHouse using the Bitnami chart
- **StatefulSets** with persistent volumes for data storage
- **Services** for exposing ClickHouse internally
- **LoadBalancer Services** (optional) for external access
- **Secrets** for storing auto-generated passwords securely
- **ZooKeeper** (optional) for cluster coordination in distributed setups

Developers simply provide a declarative resource specification – focusing on resource allocations, persistence settings, clustering configuration, and ingress preferences – while the module handles the underlying Kubernetes constructs.

### Key Features

1. **Deployment Automation**  
   Eliminates the need to write manual Kubernetes definitions for namespaces, statefulsets, or services. The module compiles your specification into a robust, ready-to-run ClickHouse cluster.

2. **Flexible Clustering**  
   Deploy ClickHouse in standalone mode or as a distributed cluster with configurable sharding and replication. Horizontal scaling and high availability built-in.

3. **Persistence Management**  
   Toggle data persistence with configurable disk sizes. When enabled, data survives pod restarts and node failures.

4. **Security First**  
   Automatically generates secure random passwords and stores them in Kubernetes Secrets. Credentials never appear in version control or container images.

5. **Resource Optimization**  
   Specify CPU and memory requests/limits for containers. This ensures your ClickHouse cluster can run efficiently without jeopardizing the cluster's stability.

6. **Optional External Access**  
   Enable ingress to expose ClickHouse via LoadBalancer services with external DNS annotations. Perfect for connecting external analytics tools.

7. **Helm Chart Integration**  
   Leverages the official Bitnami ClickHouse Helm chart with extensive customization options through the `helmValues` field.

### Architecture

**Standalone Mode:**
```
┌─────────────────────────────────────┐
│  Namespace                          │
│  ┌───────────────────────────────┐  │
│  │  StatefulSet (ClickHouse)     │  │
│  │  - Persistent Volume          │  │
│  │  - Auto-generated Password    │  │
│  └───────────────────────────────┘  │
│  ┌───────────────────────────────┐  │
│  │  Service (ClusterIP)          │  │
│  └───────────────────────────────┘  │
│  ┌───────────────────────────────┐  │
│  │  LoadBalancer (Optional)      │  │
│  └───────────────────────────────┘  │
└─────────────────────────────────────┘
```

**Clustered Mode:**
```
┌────────────────────────────────────────────────────┐
│  Namespace                                         │
│  ┌──────────────────────────────────────────────┐  │
│  │  ZooKeeper (for coordination)                │  │
│  └──────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────┐  │
│  │  Shard 1                                     │  │
│  │  ├─ Replica 1 (Persistent Volume)           │  │
│  │  └─ Replica 2 (Persistent Volume)           │  │
│  └──────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────┐  │
│  │  Shard 2                                     │  │
│  │  ├─ Replica 1 (Persistent Volume)           │  │
│  │  └─ Replica 2 (Persistent Volume)           │  │
│  └──────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────┐  │
│  │  Services & LoadBalancers                    │  │
│  └──────────────────────────────────────────────┘  │
└────────────────────────────────────────────────────┘
```

### Use Cases

- **Real-time Analytics**: Deploy ClickHouse for high-performance OLAP queries on streaming data
- **Data Warehousing**: Build scalable data warehouses with ClickHouse's columnar storage
- **Log Analytics**: Process and analyze massive log volumes efficiently
- **Time-Series Data**: Handle time-series workloads with ClickHouse's optimized storage engine
- **Business Intelligence**: Power BI dashboards with fast analytical queries
- **Event Analytics**: Track and analyze user events at scale

Overall, **ClickhouseKubernetes** helps you focus on your data and analytics workflows. By delegating the Kubernetes resource orchestration to this module, you gain a cleaner, more consistent deployment experience across development, staging, and production environments.
