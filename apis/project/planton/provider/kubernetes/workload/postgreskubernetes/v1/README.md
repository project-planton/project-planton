# Overview

The **Postgres Kubernetes API resource** is designed to simplify the deployment and management of PostgreSQL databases within Kubernetes environments. This API resource allows users to configure PostgreSQL containers, resource allocations, and ingress settings efficiently, ensuring consistency and reliability in managing PostgreSQL instances.

## Why We Created This API Resource

Deploying and managing PostgreSQL databases in Kubernetes can be complex due to the need to handle container resources, replicas, and ingress configurations. This API resource was developed to:

- **Simplify PostgreSQL Deployment**: Provides an easy-to-use interface for deploying PostgreSQL in Kubernetes environments.
- **Ensure Consistency**: Offers a standardized approach to deploying PostgreSQL across different Kubernetes clusters and environments.
- **Optimize Resource Management**: Allows fine-tuning of CPU, memory, and storage allocations for PostgreSQL containers.
- **Streamline Ingress Management**: Facilitates ingress configuration to allow secure access to PostgreSQL instances.

## Key Features

### Environment Integration

- **Environment Info**: Automatically integrates with the Planton Cloud environment management system, ensuring that PostgreSQL is deployed in the right context.
- **Stack Job Settings**: Supports stack job settings for consistent infrastructure-as-code deployments.

### Kubernetes Credential Management

- **Kubernetes Cluster Credential ID**: Specifies the Kubernetes credentials (`kubernetes_cluster_credential_id`) for securely deploying and managing the PostgreSQL container in Kubernetes.

### PostgreSQL Container Configuration

#### Resource Management

- **Replicas**: Define the number of PostgreSQL replicas to ensure high availability and fault tolerance. The recommended default is 1 replica.

- **Container Resources**: Customize the CPU and memory resources for the PostgreSQL container to ensure optimal performance. The recommended default values are:
    - **CPU Requests**: `50m`
    - **Memory Requests**: `256Mi`
    - **CPU Limits**: `1`
    - **Memory Limits**: `1Gi`

- **Disk Size**: Configure the storage size for each PostgreSQL instance. The default value is `1Gi`, but you can specify a different size based on your requirements.

### Ingress Configuration

- **Ingress Spec**: Manage and control ingress traffic to the PostgreSQL instance by configuring ingress rules to expose the PostgreSQL service securely.

## Benefits

- **Simplified Deployment**: Reduces the complexity of deploying PostgreSQL on Kubernetes by providing an easy-to-use API resource.
- **Consistent Configuration**: Ensures PostgreSQL deployments follow a consistent configuration across different Kubernetes clusters and environments.
- **Resource Optimization**: Enables fine-tuning of CPU, memory, and storage allocations to optimize the performance of PostgreSQL instances.
- **Secure Access**: Configures ingress rules to allow secure and controlled access to PostgreSQL instances.
