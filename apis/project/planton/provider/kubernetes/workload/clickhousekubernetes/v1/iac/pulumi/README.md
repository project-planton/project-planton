# ClickHouse Kubernetes Pulumi Module

## Key Features

- **Standardized API Resource Model**  
  Provides a unified way to define and deploy ClickHouse on Kubernetes. By describing resource allocations, persistence settings, and optional clustering in a simple API resource, you ensure consistency across environments.

- **Automated Kubernetes Resource Creation**  
  Automatically creates Namespaces, Helm chart deployments, Services, and optional Load Balancer resources based on the provided specifications. Eliminates the need for hand-maintained YAML files.

- **ClickHouse Configuration**  
  Supports detailed specifications for ClickHouse containers, including resource limits/requests, persistence volumes, and clustering configurations for sharding and replication.

- **Clustering Support**  
  Enable distributed ClickHouse deployments with configurable sharding and replication for horizontal scalability and high availability.

- **Persistence Management**  
  Toggle data persistence with configurable disk sizes. When enabled, data is stored in persistent volumes, allowing data to survive pod restarts.

- **Password Management**  
  Automatically generates secure random passwords stored in Kubernetes Secrets, ensuring credentials stay out of version control and container images.

- **Ingress Integration**  
  When enabled, the module sets up Load Balancer services with external DNS annotations, allowing for external access to your ClickHouse cluster.

- **Output Exports**  
  Exports useful values such as namespace, service name, internal service FQDN, port-forward commands, and credentials. These can be leveraged for further automation or debugging.

## Usage

See [examples.md](examples.md) for usage details and step-by-step examples. In general:

1. Define a YAML resource describing your ClickHouse cluster using the **ClickhouseKubernetes** API.
2. Run:
   ```bash
   planton pulumi up --stack-input <your-clickhouse-file.yaml>
   ```

to apply the resource on your cluster.

## Important: Docker Image Registry

**⚠️ Bitnami Registry Changes (September 2025)**

Due to Bitnami's transition to a paid model, this module now uses `docker.io/bitnamilegacy` registry for ClickHouse and ZooKeeper images. The legacy images receive no updates or security patches but provide a temporary migration solution.

**Long-term Alternatives:**
- Subscribe to Bitnami Secure Images ($50k-$72k/year)
- Use official ClickHouse images from clickhouse.com
- Build custom images from open-source Bitnami code (Apache 2.0)

For more details, see: https://github.com/bitnami/containers/issues/83267

## Getting Started

1. **Craft Your Specification**  
   Include container resource info, persistence settings, and optionally clustering preferences. If you need external access, enable ingress.

2. **Apply via CLI**  
   Execute `planton pulumi up --stack-input <clickhouse-spec.yaml>` (or your organization's standard CLI command). The Pulumi module automatically compiles your specification into Kubernetes resources.

3. **Validate & Observe**  
   Check the logs of your ClickHouse deployment, confirm the Namespace, StatefulSet, and Service are created, and if ingress is enabled, verify external access.

## Module Structure

1. **Initialization**  
   Reads your `ClickhouseKubernetesStackInput` (containing cluster creds, resource definitions), sets up local variables, and merges labels.

2. **Provider Setup**  
   Establishes a Pulumi Kubernetes Provider for your target cluster.

3. **Namespace Creation**  
   Creates (or identifies) a namespace to house all your ClickHouse resources.

4. **Secret Management**  
   Generates a secure random password and stores it in a Kubernetes Secret for ClickHouse authentication.

5. **Helm Chart Deployment**  
   Deploys the Bitnami ClickHouse Helm chart with configured values for resources, persistence, clustering, and authentication.

6. **Service Configuration (Optional)**  
   If ingress is enabled, creates a LoadBalancer Service for external access with DNS annotations.

7. **Output Exports**  
   Publishes final references (e.g., namespace, service name, cluster endpoints, credentials), which can aid in post-deployment automation.

## Benefits

- **Simplified Deployment**  
  Focus on high-level configuration rather than writing raw Kubernetes manifests. Consistent patterns reduce the risk of misconfiguration.

- **Security & Compliance**  
  Minimizes exposure of credentials by auto-generating passwords and storing them securely in Kubernetes Secrets.

- **Scalability**  
  Easily configure standalone or clustered deployments with sharding and replication for handling large-scale analytical workloads.

- **Extensibility**  
  The module is built on Pulumi's Kubernetes provider. You can augment or override resources if your team needs advanced configurations through helm_values.

## Contributing

Contributions are always welcome! Please open an issue or submit a pull request in the main repository if you want to add features, fix bugs, or improve documentation.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to adapt it for your internal workflows.
