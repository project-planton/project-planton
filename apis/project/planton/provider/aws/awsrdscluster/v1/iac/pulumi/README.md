# AwsRdsCluster Pulumi Module

## Overview

The `AwsRdsCluster` Pulumi module allows developers to deploy and manage Amazon RDS clusters with a Kubernetes-like API resource modeling system. This module enables the creation of highly available, scalable, and secure RDS clusters, supporting engines like Aurora, Aurora PostgreSQL, and Aurora MySQL. It abstracts away much of the complexity involved in infrastructure management and empowers developers to define infrastructure in a simple YAML format. The Pulumi module captures cloud resource details in `status.stackOutputs` for monitoring and management, making it easier to track the state of the deployed infrastructure.

## Key Features

- **Support for Multiple Database Engines:** The module supports various AWS RDS engines, including `aurora`, `aurora-postgresql`, and `aurora-mysql`. Developers can specify the engine and engine version based on their needs.
  
- **Cluster Modes and Autoscaling:** The module allows users to define different engine modes such as `provisioned`, `serverless`, `global`, and more. Additionally, it supports both manual and automatic scaling policies, including autoscaling configurations like target tracking and step scaling policies.

- **Security and Encryption:** The module provides options for VPC integration with security group configurations, along with subnet associations. It also supports encryption through AWS KMS, enabling secure data storage with customer-managed keys or default AWS keys.

- **Performance Insights:** For users who want in-depth performance monitoring, the module offers integration with AWS Performance Insights. You can enable this feature and specify KMS keys for encrypting performance data.

- **Customizable Backup and Retention Policies:** The module allows users to define custom backup windows and retention periods for their RDS clusters, ensuring that data is backed up consistently and retained for the desired duration.

- **Networking Flexibility:** The RDS clusters can be made publicly accessible or restricted to private networks. You can define CIDR blocks, security groups, and subnets to control access and ensure secure communication.

- **Easy Integration with AWS IAM:** You can specify whether to enable IAM database authentication, allowing AWS IAM roles to manage database credentials.

- **Stack Outputs for Easy Monitoring:** The module captures critical information such as the RDS cluster identifier, master endpoint, and reader endpoint in the `status.stackOutputs`, simplifying the management and connection to the RDS cluster.

- **Simplified API Resource Structure:** The module follows a Kubernetes-like structure for API resources, making it intuitive for developers familiar with Kubernetes. Each resource includes fields for `apiVersion`, `kind`, `metadata`, `spec`, and `status`, mirroring the structure of Kubernetes objects.

## Pulumi Integration

The Pulumi module utilizes Pulumi’s Go SDK for AWS, allowing dynamic resource creation based on the API resource’s specification. The Pulumi outputs from the module are captured in the `status.stackOutputs`, allowing you to monitor and manage your infrastructure directly from your Pulumi stack. The module integrates seamlessly with AWS through an AWS provider, ensuring that the appropriate credentials and permissions are applied during the resource creation process.

## API-Resource Fields

### `apiVersion`
Defines the version of the API. For this module, it is typically `code2cloud.planton.cloud/v1`.

### `kind`
The kind of resource being deployed. In this case, it is always `AwsRdsCluster`.

### `metadata`
Contains metadata about the resource such as the name and labels. The `name` field is required and must be unique within the scope of the cluster.

### `spec`
The `spec` section defines the desired state of the RDS cluster. It includes fields such as:
- `awsCredentialId`: The AWS credentials to use for creating the RDS cluster.
- `engine`: The type of database engine (e.g., Aurora, Aurora-PostgreSQL).
- `engineVersion`: The version of the database engine.
- `instanceType`: The EC2 instance type for the database.
- `clusterSize`: The number of database instances in the cluster.
- `vpcId`: The VPC in which the RDS cluster will be created.
- `securityGroupIds`: The security groups to associate with the RDS cluster.
- `subnetIds`: Subnets in which to create the RDS cluster instances.
- `autoScaling`: (Optional) Autoscaling configuration, including policies, metrics, and target values.
- `storageEncrypted`: Enables encryption for storage volumes.
- `performanceInsights`: Enables AWS Performance Insights for performance monitoring.

### `status`
This section reflects the current status of the RDS cluster. It includes details like the RDS cluster identifier, master and reader endpoints, and other outputs from Pulumi. These outputs are captured in the `status.stackOutputs`, providing real-time visibility into the state of the infrastructure.

## Usage

To use this module, simply create a YAML file that defines the `AwsRdsCluster` API resource with your desired configuration. After creating the YAML file, you can deploy the resource using the following command:

```shell
planton apply -f <yaml-path>
```

Refer to the example section for usage instructions.

## Pulumi CLI Command

This module integrates with the Planton CLI for seamless deployment and management. You can use the following command to deploy the Pulumi stack:

```shell
planton pulumi up --stack-input <api-resource.yaml>
```

This command takes your YAML file as input, configures the Pulumi stack, and applies the desired state on AWS. If no specific Pulumi module is specified, the CLI automatically selects the default module for the `AwsRdsCluster` resource type.

## Documentation

The full documentation for this API resource, including detailed descriptions of each field, can be found on the [buf.build](https://buf.build) platform. Additionally, refer to the official Pulumi documentation for more details on how Pulumi works and how it can be used to manage your infrastructure.
