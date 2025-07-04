# MongoDB Kubernetes Pulumi Module

## Key Features

### Unified API Resource Model
The MongoDB Kubernetes module adheres to a standard Kubernetes resource structure (`apiVersion`, `kind`, `metadata`, `spec`, and `status`), making it familiar to Kubernetes users. The `spec` defines critical aspects such as the MongoDB container, resource limits, persistence settings, and ingress configurations, providing granular control over the deployment.

### Key Features of the API Resource:
- **Helm Customization**: The module supports passing custom Helm chart values via the `helm_values` field, allowing further customization of the MongoDB deployment. This can be used to adjust resources, configure environment variables, or set specific MongoDB version tags.
- **Persistence**: The `is_persistence_enabled` flag allows you to enable or disable persistence for MongoDB data. When enabled, a persistent volume is attached to each MongoDB pod to store data, which is restored after restarts.
- **Resource Management**: Developers can define CPU and memory limits for MongoDB pods, ensuring that the cluster is configured with optimal resource utilization.
- **Ingress Support**: The module provides the ability to configure ingress for MongoDB clusters, enabling external access to the MongoDB service. This can be particularly useful for hybrid cloud environments or services requiring external data access.
- **Replica Configuration**: The module supports the creation of multiple MongoDB replicas for high availability. By specifying the `replicas` field, developers can scale MongoDB clusters as needed to meet availability requirements.

### Key Features of the Pulumi Module:
- **Dynamic Resource Creation**: The module dynamically creates Kubernetes resources based on the `MongodbKubernetesStackInput`, including namespaces, deployments, services, and persistent volumes.
- **Kubernetes Provider Integration**: The module uses Pulumi’s Kubernetes provider to manage resources and interact with the Kubernetes cluster using provided credentials.
- **Automatic Password Generation**: A random password for MongoDB authentication is generated and securely stored in a Kubernetes secret, simplifying the process of credential management.
- **Helm Chart Deployment**: The MongoDB instance is deployed using a Helm chart, allowing for flexible customization of the MongoDB configuration and underlying infrastructure.
- **Status Outputs**: After resource creation, the module provides key outputs such as internal and external service endpoints, port forwarding commands, and the MongoDB username. These outputs are critical for accessing the MongoDB service and managing operations like backups or monitoring.

## Status and Outputs

The module provides the following outputs to simplify the operational management of the MongoDB Kubernetes cluster:
- **Namespace**: The Kubernetes namespace in which the MongoDB cluster is created, allowing for resource isolation.
- **Service Name**: The Kubernetes service name associated with the MongoDB cluster, which can be used for internal communication.
- **Port-Forwarding Command**: A Kubernetes command for setting up port-forwarding, enabling local access to the MongoDB cluster when ingress is disabled.
- **Internal and External Endpoints**: URLs for accessing MongoDB, both from within the Kubernetes cluster and externally if ingress is enabled.
- **MongoDB Credentials**: The MongoDB username and a reference to the Kubernetes secret containing the password are provided, ensuring secure access to the MongoDB service.

## Usage

To deploy and manage a MongoDB Kubernetes cluster using this module, create a YAML file representing the MongoDB Kubernetes resource. Use the CLI command `planton pulumi up --stack-input <api-resource.yaml>` to apply the configuration and provision the resources.

Refer to the example section for usage instructions.

## Requirements

- **Pulumi**: Ensure that Pulumi is installed and configured for your environment.
- **Kubernetes Provider**: The Kubernetes provider is required to manage resources on the Kubernetes cluster where MongoDB will be deployed.
- **Helm**: The module relies on Helm charts for deploying MongoDB. Ensure Helm is available and configured for your Kubernetes environment.
- **Kubernetes Cluster Credential**: This is required to authenticate and interact with the target Kubernetes cluster.
  
## Installation

You can install this Pulumi module from GitHub by cloning the repository and running the required Pulumi commands to set up the infrastructure. Ensure you have all necessary dependencies installed, including Pulumi and Kubernetes access.

1. Clone the repository:
    ```bash
    git clone https://github.com/your-repo/planton-pulumi-mongodb-kubernetes.git
    cd planton-pulumi-mongodb-kubernetes
    ```

2. Install dependencies and initialize Pulumi:
    ```bash
    pulumi stack init <stack-name>
    pulumi config set <config-parameters>
    ```

## Contributing

Contributions are welcome to enhance the functionality of this Pulumi module. If you encounter any issues or want to suggest new features, feel free to submit a pull request or open an issue on the GitHub repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
