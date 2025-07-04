# Overview

The Elasticsearch Kubernetes API resource provides a consistent and streamlined interface for deploying and managing Elasticsearch clusters within Kubernetes environments as part of our cloud infrastructure. By abstracting the complexities of Elasticsearch and Kubernetes configurations, this resource allows you to set up scalable search and analytics engines effortlessly while ensuring consistency and compliance across different environments.

## Why We Created This API Resource

Deploying Elasticsearch clusters on Kubernetes can be complex due to the intricacies involved in configuring stateful applications, managing resources, and ensuring data persistence. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:

- **Simplify Deployment**: Easily configure and deploy Elasticsearch clusters without dealing with low-level Kubernetes and Elasticsearch configurations.
- **Ensure Consistency**: Maintain uniform Elasticsearch deployments across different environments and applications.
- **Enhance Productivity**: Reduce the time and effort required to set up Elasticsearch clusters, allowing you to focus on application development and data analysis.
- **Optimize Resource Utilization**: Efficiently manage resources and enable data persistence to ensure high availability and data durability.

## Key Features

### Environment Integration

- **Environment Info**: Seamlessly integrates with our environment management system to deploy Elasticsearch clusters within specific environments.
- **Stack Job Settings**: Supports custom stack job settings for infrastructure-as-code deployments.

### Kubernetes Credential Management

- **Kubernetes Cluster Credential ID**: Utilizes specified Kubernetes credentials to ensure secure and authorized operations within Kubernetes clusters.

### Customizable Elasticsearch Deployment

#### Elasticsearch Container

- **Replicas**: Define the number of Elasticsearch pods to deploy, allowing you to scale the cluster according to your needs. The recommended default is 1.
- **Container Resources**: Specify CPU and memory resources for the Elasticsearch containers to optimize performance. Recommended defaults are:
    - CPU Requests: 50m
    - Memory Requests: 256Mi
    - CPU Limits: 1
    - Memory Limits: 1Gi
- **Persistence**: Enable or disable data persistence using `is_persistence_enabled`. When enabled, Elasticsearch in-memory data will be persisted to storage volumes, ensuring data durability across pod restarts.
- **Disk Size**: Specify the size of the persistent volume attached to each Elasticsearch pod (e.g., "10Gi"). This is required when persistence is enabled. Note that this value cannot be modified after creation due to Kubernetes limitations.

#### Kibana Container

- **Kibana Enablement**: Control whether Kibana is deployed alongside Elasticsearch using `is_enabled`. By default, this is set to `false`.
- **Replicas**: Define the number of Kibana pods to deploy.
- **Container Resources**: Specify CPU and memory resources for the Kibana containers. Recommended defaults are:
    - CPU Requests: 50m
    - Memory Requests: 256Mi
    - CPU Limits: 1
    - Memory Limits: 1Gi

### Ingress Configuration

- **Ingress Spec**: Configure ingress settings to expose Elasticsearch and Kibana services outside the cluster, including hostname, TLS settings, and ingress annotations.

## Benefits

- **Simplified Deployment**: Abstracts the complexities of deploying stateful Elasticsearch clusters on Kubernetes into an easy-to-use API.
- **Consistency**: Ensures all Elasticsearch deployments adhere to organizational standards for security, performance, and scalability.
- **Scalability**: Allows for easy scaling of Elasticsearch and Kibana pods to handle varying workloads.
- **Data Persistence**: Provides options for data persistence to ensure data is retained across pod restarts and failures.
- **Resource Optimization**: Enables precise control over resource allocation for containers, optimizing performance and cost.
- **Flexibility**: Customize Elasticsearch and Kibana configurations to meet specific application requirements without compromising best practices.
