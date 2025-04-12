# AWS DynamoDB Pulumi Module

## Introduction

The AWS DynamoDB Pulumi Module provides a standardized and efficient way to define and deploy DynamoDB tables on AWS using a Kubernetes-like API resource model. By leveraging our unified APIs, developers can specify their DynamoDB configurations in simple YAML files, which the module then uses to create and manage DynamoDB resources through Pulumi. This approach abstracts the complexity of AWS interactions and streamlines the deployment process, enabling consistent infrastructure management across multi-cloud environments.

## Key Features

- **Kubernetes-Like API Resource Model**: Utilizes a familiar structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status`, making it intuitive for developers accustomed to Kubernetes to define AWS DynamoDB resources.

- **Unified API Structure**: Ensures consistency across different resources and cloud providers by adhering to a standardized API resource model.

- **Pulumi Integration**: Employs Pulumi for infrastructure provisioning, enabling the use of real programming languages and providing robust state management and automation capabilities.

- **Comprehensive DynamoDB Configuration**: Supports detailed specification of DynamoDB table attributes, including table name, billing mode, key attributes, encryption settings, autoscaling options, global and local secondary indexes, point-in-time recovery, TTL settings, and data import configurations.

- **Credential Management**: Securely handles AWS credentials via the `awsCredentialId` field, ensuring authenticated and authorized resource deployments.

- **Autoscaling Support**: Facilitates the configuration of read and write capacity autoscaling for DynamoDB tables and indexes, allowing for responsive scaling based on workload demands.

- **Status Reporting**: Captures and stores outputs such as table names, ARNs, and policy ARNs in `status.outputs` for easy reference and further automation.

## Architecture

The module operates by accepting an AWS DynamoDB API resource definition as input. It interprets the resource definition and uses Pulumi to interact with AWS, creating the specified DynamoDB resources. The main components involved are:

- **API Resource Definition**: A YAML file that includes all necessary information to define a DynamoDB table, following the standard API structure.

- **Pulumi Module**: Written in Go, the module reads the API resource and uses Pulumi's AWS SDK to provision DynamoDB resources based on the provided specifications.

- **AWS Provider Initialization**: The module initializes the AWS provider within Pulumi using the credentials specified by `awsCredentialId`.

- **Resource Creation**: Provisions the DynamoDB table and associated resources as defined in the `spec`, including attributes, indexes, encryption settings, and autoscaling policies.

- **Status Outputs**: Outputs from the Pulumi deployment, such as table names, ARNs, and policy ARNs, are captured and stored in `status.outputs` for easy access and integration with other systems.

## Usage

Refer to the example section for usage instructions.

## Limitations

- **Advanced Features**: Certain advanced features of DynamoDB that are not specified in the current API resource definition may not be supported. Future updates may include additional capabilities based on user needs.

## Contributing

We welcome contributions to enhance the functionality of this module. Please submit pull requests or open issues to help improve the module and its documentation.

