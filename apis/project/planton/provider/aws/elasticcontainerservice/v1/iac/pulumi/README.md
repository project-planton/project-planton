# AWS Elastic Container Service Pulumi Module

**Note:** This module is not completely implemented because the API resource specification is currently empty.

## Introduction

The AWS Elastic Container Service (ECS) Pulumi Module provides a standardized and efficient way to define and deploy containerized applications on AWS using a Kubernetes-like API resource model. By leveraging our unified APIs, developers can specify their ECS configurations in simple YAML files, which the module then uses to create and manage AWS ECS resources through Pulumi. This approach abstracts the complexity of AWS interactions and streamlines the deployment process, enabling consistent infrastructure management across multi-cloud environments.

## Key Features

- **Kubernetes-Like API Resource Model**: Utilizes a familiar structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status`, making it intuitive for developers accustomed to Kubernetes to define AWS ECS resources.

- **Unified API Structure**: Ensures consistency across different resources and cloud providers by adhering to a standardized API resource model.

- **Pulumi Integration**: Employs Pulumi for infrastructure provisioning, enabling the use of real programming languages and providing robust state management and automation capabilities.

- **Container Orchestration**: Allows for the deployment and management of containerized applications on AWS ECS, facilitating scalable and reliable application hosting.

- **Credential Management**: Securely handles AWS credentials via the `awsCredentialId` field, ensuring authenticated and authorized resource deployments without exposing sensitive information.

- **Status Reporting**: Captures and stores outputs such as ECS cluster details in `status.stackOutputs`. This facilitates easy reference and integration with other systems, such as monitoring tools or additional automation scripts.

## Architecture

The module operates by accepting an AWS Elastic Container Service API resource definition as input. It interprets the resource definition and uses Pulumi to interact with AWS, creating the specified ECS resources. The main components involved are:

- **API Resource Definition**: A YAML file that includes all necessary information to define an ECS cluster and related resources, following the standard API structure. Developers specify the desired state in this file, including details like cluster configuration and service definitions.

- **Pulumi Module**: Written in Go, the module reads the API resource and uses Pulumi's AWS SDK to provision ECS resources based on the provided specifications. It abstracts the complexity of resource creation, update, and deletion.

- **AWS Provider Initialization**: The module initializes the AWS provider within Pulumi using the credentials specified by `awsCredentialId`. This ensures that all AWS resource operations are authenticated and authorized.

- **Resource Creation**: Provisions the ECS cluster and associated resources as defined in the `spec`, including task definitions, services, and networking configurations. The module aims to simplify the deployment of containerized applications by handling the underlying infrastructure requirements.

- **Status Outputs**: Outputs from the Pulumi deployment, such as cluster ARNs and service endpoints, are captured and stored in `status.stackOutputs`. This information is crucial for accessing deployed services and integrating with other systems.

## Usage

Refer to the example section for usage instructions.

## Limitations

- **Incomplete Implementation**: As noted, the module currently lacks a complete implementation of the AWS ECS resource creation due to an empty `spec`. Future updates will include full support for defining and deploying ECS clusters and services.

- **Advanced Features**: Certain advanced features of AWS ECS, such as custom task definitions, service autoscaling, or integration with AWS Fargate, may not be supported until the module is fully implemented.

## Contributing

We welcome contributions to enhance the functionality of this module. Please submit pull requests or open issues to help improve the module and its documentation.
