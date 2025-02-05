# GKE Cluster Pulumi Module

This Pulumi module automates the provisioning of Google Kubernetes Engine (GKE) clusters, along with the installation of essential Kubernetes addons and operators. Designed for the multi-cloud era, it leverages a unified API resource model, allowing developers to deploy complex infrastructure with minimal configuration. By providing a standardized YAML specification, users can set up GKE clusters tailored to their specific needs, including custom networking, autoscaling, and addon installations.

The module is written in Golang and integrates seamlessly with Pulumi for infrastructure as code. It takes a custom `GkeCluster` API resource as input, interprets the specifications, and orchestrates the creation of Google Cloud resources accordingly. This includes setting up projects, networks, clusters, and addons, while handling IAM roles and service accounts for secure operations. The module abstracts the complexity of cloud resource provisioning, enabling developers to focus on application development rather than infrastructure management.

## Key Features

### Unified API Resource Model

- **Standardized Structure**: Utilizes a consistent API resource structure with fields like `apiVersion`, `kind`, `metadata`, `spec`, and `status`.
- **Simplified Configuration**: Developers provide a single YAML file to configure complex infrastructure components.
- **Extensible Specifications**: Supports detailed cluster configurations, including autoscaling, node pools, and custom network settings.

### Comprehensive Cluster Provisioning

- **Multi-Project Setup**: Automatically creates and manages Google Cloud projects for cluster and network resources.
- **Shared VPC Support**: Configurable shared VPC setups for network segmentation and isolation.
- **Custom Networking**: Sets up VPC networks, subnets, firewall rules, and NAT configurations based on specifications.

### Kubernetes Addons Installation

- **Ingress Controllers**: Installs Ingress Nginx and Istio for traffic management and service mesh capabilities.
- **Cert Manager**: Automates TLS certificate management and integrates with Let's Encrypt.
- **External DNS**: Manages DNS records for Kubernetes resources using External DNS.
- **Operators Support**: Deploys operators for Postgres, Kafka, Solr, ElasticSearch, and Keycloak.
- **Workload Identity Integration**: Configures workload identities for secure and seamless access to Google Cloud services.

### Secure and Scalable

- **IAM Management**: Automates the creation and assignment of IAM roles and service accounts.
- **Cluster Autoscaling**: Configurable autoscaling settings for CPU and memory resources.
- **Node Pools**: Supports multiple node pools with specific machine types and autoscaling configurations.
- **Logging and Monitoring**: Enables workload logs and integrates with Google Cloud's operations suite.

## Installation

To use this module, ensure you have the following prerequisites:

- **Pulumi CLI**: Installed and configured with access to your Google Cloud account.
- **Golang**: Go programming language installed for module execution.
- **Credentials**: Proper Google Cloud credentials with necessary permissions.

Clone the module repository and include it in your Pulumi project.

```bash
git clone https://github.com/your-org/gke-cluster-pulumi-module.git
```

## Usage

Refer to [example](example.md) for usage instructions.

## API Resource Specification

The module relies on a `GkeCluster` API resource that defines the desired state of the GKE cluster and associated resources. Key fields include:

- **Metadata**: Includes `name` and `id` to uniquely identify the cluster.
- **Spec**: Contains specifications for environment info, billing account, region, zone, shared VPC settings, and more.
    - **Autoscaling Config**: Enables cluster-level autoscaling with min/max CPU and memory settings.
    - **Node Pools**: Allows configuration of multiple node pools with machine types, scaling settings, and spot instances.
    - **Kubernetes Addons**: Specifies which addons and operators to install on the cluster.
    - **Ingress DNS Domains**: Defines domains for ingress resources, including TLS settings.

## Addons and Operators

The module supports the installation of various Kubernetes addons and operators:

- **Ingress Nginx**: Deploys the Nginx ingress controller for HTTP load balancing.
- **Istio**: Sets up the Istio service mesh for advanced traffic management and security.
- **Cert Manager**: Automates the issuance and renewal of TLS certificates.
- **External DNS**: Keeps DNS records in sync with Kubernetes ingresses and services.
- **Operators**: Installs operators for databases and middleware like Postgres, Kafka, Solr, ElasticSearch, and Keycloak.

## Customization and Extensibility

- **Workload Logs**: Optionally enable logging for workloads to Google Cloud Logging.
- **Shared VPC Configuration**: Choose whether to deploy the cluster within a shared VPC network.
- **Custom Labels**: Apply custom labels to Google Cloud resources for better organization and billing.
- **Workload Identity**: Leverage Google Cloud's Workload Identity for secure access to cloud services from Kubernetes pods.

## Contributing

Contributions are welcome! Please read the [contribution guidelines](CONTRIBUTING.md) and submit pull requests for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

For any questions or support, please open an issue in the repository or contact the maintainers directly.
