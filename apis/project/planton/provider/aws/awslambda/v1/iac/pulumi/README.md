# AWS Lambda Pulumi Module

## Introduction

The AWS Lambda Pulumi Module provides a standardized and efficient way to define and deploy AWS Lambda functions using a Kubernetes-like API resource model. By leveraging our unified APIs, developers can specify their Lambda function configurations in simple YAML files, which the module then uses to create and manage AWS Lambda resources through Pulumi. This approach abstracts the complexity of AWS interactions and streamlines the deployment process, enabling consistent infrastructure management across multi-cloud environments.

## Key Features

- **Kubernetes-Like API Resource Model**: Utilizes a familiar structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status`, making it intuitive for developers accustomed to Kubernetes to define AWS Lambda resources.

- **Unified API Structure**: Ensures consistency across different resources and cloud providers by adhering to a standardized API resource model.

- **Pulumi Integration**: Employs Pulumi for infrastructure provisioning, enabling the use of real programming languages and providing robust state management and automation capabilities.

- **Comprehensive Lambda Function Configuration**: Supports detailed specification of Lambda function attributes, including handler, runtime, memory size, environment variables, IAM roles, VPC configurations, layers, and more.

- **IAM Role Management**: Allows for the definition of custom IAM roles and policies for the Lambda function, enhancing security and fine-grained access control.

- **Environment Variables and Secrets**: Facilitates the injection of environment variables into the Lambda function, enabling dynamic configuration and secret management.

- **VPC Integration**: Supports deploying Lambda functions within a VPC, allowing access to private resources and enhancing security.

- **Event Source Permissions**: Manages invocation permissions, defining which external sources can trigger the Lambda function.

- **CloudWatch Logs Configuration**: Enables configuration of CloudWatch Log Groups for the Lambda function, including retention policies and encryption settings.

- **Credential Management**: Securely handles AWS credentials via the `awsCredentialId` field, ensuring authenticated and authorized resource deployments.

- **Status Reporting**: Captures and stores outputs such as Lambda function IDs and ARNs in `status.stackOutputs` for easy reference and further automation.

## Architecture

The module operates by accepting an AWS Lambda API resource definition as input. It interprets the resource definition and uses Pulumi to interact with AWS, creating the specified Lambda resources. The main components involved are:

- **API Resource Definition**: A YAML file that includes all necessary information to define a Lambda function, following the standard API structure.

- **Pulumi Module**: Written in Go, the module reads the API resource and uses Pulumi's AWS SDK to provision Lambda resources based on the provided specifications.

- **AWS Provider Initialization**: The module initializes the AWS provider within Pulumi using the credentials specified by `awsCredentialId`.

- **Resource Creation**: Provisions the Lambda function and associated resources as defined in the `spec`, including IAM roles, environment variables, VPC configurations, and event source permissions.

- **Status Outputs**: Outputs from the Pulumi deployment, such as Lambda function IDs and ARNs, are captured and stored in `status.stackOutputs` for easy access and integration with other systems.

## Usage

Refer to the example section for usage instructions.

## Limitations

- **Advanced Features**: Certain advanced features of AWS Lambda that are not specified in the current API resource definition may not be supported. Future updates may include additional capabilities based on user needs.

## Contributing

We welcome contributions to enhance the functionality of this module. Please submit pull requests or open issues to help improve the module and its documentation.
