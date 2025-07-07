# SignozKubernetes Pulumi Module

## Key Features

### Standardized API Resource Structure
- **apiVersion & kind**: Aligns with Kubernetes standards, ensuring ease of integration and familiarity for developers.
- **metadata**: Utilizes standard Kubernetes metadata fields for resource identification and management.
- **spec**: Defines the necessary parameters for deploying Signoz, including Kubernetes cluster credentials.
- **status**: Captures real-time updates and outputs from the deployed infrastructure, enhancing monitoring and visibility.

### Seamless Cloud Integration
- **Multi-Cloud Support**: Leverages Pulumi's cloud-agnostic capabilities to deploy Signoz across various cloud providers.
- **Kubernetes Credentials Management**: Utilizes Kubernetes cluster credentials to securely provision resources, adhering to best security practices.
- **Automated Outputs Handling**: Captures and manages Pulumi outputs within the API resource status, providing essential deployment information.

### Developer-Friendly CLI
- **Unified Deployment Command**: Employs the `planton pulumi up --stack-input <api-resource.yaml>` command to simplify the deployment process.
- **Default Module Configuration**: Automatically configures stack inputs using default Pulumi modules, reducing setup complexity for developers.
- **Git Integration**: Allows developers to specify custom Pulumi modules via Git repository details, enabling flexible and customized deployments.

### Robust Documentation and Validation
- **Buf.Build Documentation**: Provides comprehensive and accessible documentation through buf.build, ensuring developers have the necessary guidance.
- **Validation Rules**: Implements validation rules within the API resource specifications to enforce correct configurations and prevent deployment errors.

## Installation

To install the SignozKubernetes Pulumi module, follow the standard Pulumi module installation procedures. Ensure that you have Pulumi and the necessary dependencies installed on your development machine.

## Usage

Refer to the example section for usage instructions.

## API Reference

### SignozKubernetesSpec
Defines the configuration for deploying Signoz within a Kubernetes cluster.

- **kubernetes_cluster_credential_id**: Identifier for the Kubernetes cluster credentials to be used for provisioning resources.

### SignozKubernetesStackInputs
Specifies the inputs required for the SignozKubernetes stack.

- **pulumi**: Pulumi-specific input configurations.
- **target**: The target `SignozKubernetes` API resource.
- **kubernetes_cluster**: Specifications for the Kubernetes cluster credentials.

### SignozKubernetesStackOutputs
Provides outputs from the deployed Signoz infrastructure.

- **namespace**: Kubernetes namespace where Signoz is deployed.

## Contributing

Contributions are welcome! Please refer to the contributing guidelines for more information on how to get involved.

## License

This project is licensed under the [MIT License](LICENSE).