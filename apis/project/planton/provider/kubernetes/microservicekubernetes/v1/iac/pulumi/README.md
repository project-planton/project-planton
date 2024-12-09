# Microservice Kubernetes Pulumi Module

## Key Features

- **Standardized API Resource Model**: Utilizes a unified structure for API resources, ensuring consistency across
  different deployments and simplifying the development process.

- **Automated Kubernetes Resource Creation**: Automatically creates Kubernetes namespaces, deployments, services, and
  ingress resources based on the provided specifications.

- **Container Configuration**: Supports detailed container specifications, including image details, resource limits,
  environment variables, secrets, and port configurations.

- **Ingress Management**: Handles the setup of ingress resources using the Gateway API and Cert-Manager, enabling secure
  external and internal access to microservices with automatic TLS certificate provisioning.

- **Environment and Credential Integration**: Integrates with environment-specific configurations and credentials, such
  as Kubernetes cluster credentials and Docker registry credentials, to facilitate secure deployments.

- **Scalability and Availability**: Supports configurations for replicas and horizontal pod autoscaling to ensure the
  microservice can scale based on demand.

- **Secret Management**: Utilizes external secret management for securely injecting secrets into microservices,
  integrating with secret stores like GCP Secret Manager.

- **Output Exports**: Provides outputs such as namespace, service name, endpoints, and commands for port forwarding,
  which can be used for further automation or integration.

## Usage

Refer to [example](example.md) for usage instructions.

## Getting Started

To deploy a microservice using this Pulumi module, define your microservice specifications in a YAML file following the
standardized API resource model. Use the CLI command:

```bash
platon pulumi up --stack-input <api-resource.yaml>
```

The module reads the specifications from the provided YAML file, sets up the Kubernetes provider using the specified
cluster credentials, and proceeds to create and configure all necessary Kubernetes resources.

## Module Structure

- **Initialization**: Reads the API resource specifications and initializes local variables and labels used throughout
  the deployment process.

- **Provider Setup**: Creates a Kubernetes provider instance using the provided cluster credentials for subsequent
  resource creation.

- **Namespace Creation**: Generates a unique namespace for the microservice deployment to encapsulate all related
  resources.

- **Image Pull Secrets**: Creates Kubernetes secrets for pulling images from private Docker registries, based on the
  provided Docker credentials.

- **Deployment Configuration**: Sets up the Kubernetes Deployment resource, including containers, environment variables,
  secrets, ports, and resource requests and limits.

- **Service Configuration**: Creates a Kubernetes Service to expose the microservice within the cluster, configuring it
  according to the specified ports and protocols.

- **Ingress Setup**: If ingress is enabled, sets up ingress resources using the Gateway API and Cert-Manager, handling
  both external and internal access with HTTPS termination and automatic certificate provisioning.

- **Secret Management**: Integrates with external secret stores to securely inject secrets into the microservice's
  environment.

- **Output Exports**: Exports key information such as namespace, service name, endpoints, and commands, which can be
  used for accessing the microservice or for further automation.

## Benefits

- **Simplified Deployment Process**: Reduces complexity by automating resource creation and configuration.

- **Consistency and Standardization**: Ensures all deployments adhere to a standardized structure, making it easier to
  manage multiple microservices across different environments.

- **Security and Compliance**: Incorporates best practices for handling secrets and credentials, ensuring sensitive
  information is managed securely.

- **Scalable Architecture**: Supports scaling configurations based on resource utilization, ensuring high availability
  and performance.

- **Extensibility**: Can be extended or customized to accommodate additional requirements or integrate with other tools
  and services.

## Contributing

Contributions are welcome. Please submit issues or pull requests to the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).
