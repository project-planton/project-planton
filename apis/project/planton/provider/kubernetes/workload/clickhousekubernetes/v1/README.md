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
- **ZooKeeper**: Automatically managed by the operator for cluster coordination, or configure external ZooKeeper for advanced scenarios.

### Networking and Ingress

- **Ingress Configuration**: Set up Kubernetes Ingress resources to manage external access to ClickHouse, including hostname and path routing.

## Benefits

- **Production-Ready**: Leverage the battle-tested Altinity operator used by enterprises worldwide for operational excellence.
- **Self-Healing**: Automatic recovery from failures, rolling upgrades, and continuous reconciliation to maintain desired state.
- **Simplified Operations**: The operator handles complex lifecycle operations - you focus on your data, not Kubernetes complexity.
- **Scalability and Flexibility**: Easily scale from single nodes to distributed clusters with hundreds of nodes.
- **Data Persistence**: Production-grade persistent storage with automatic volume provisioning and management.
- **High Availability**: Native support for clustering with sharding, replication, and ZooKeeper coordination.
- **Type Safety**: Strongly-typed configuration with compile-time validation and IDE support.

## Use Cases

- **Real-time Analytics**: Deploy ClickHouse as a high-performance analytics database for real-time data processing.
- **Data Warehousing**: Use ClickHouse for OLAP workloads and complex analytical queries on large datasets.
- **Log Analytics**: Process and analyze large volumes of log data with ClickHouse's columnar storage engine.
- **Time-Series Data**: Handle time-series data efficiently with ClickHouse's optimized storage and query capabilities.
- **Microservices Architecture**: Deploy ClickHouse instances for services requiring fast analytical queries.
- **Development and Testing Environments**: Quickly spin up ClickHouse instances for development or testing purposes with environment-specific configurations.
