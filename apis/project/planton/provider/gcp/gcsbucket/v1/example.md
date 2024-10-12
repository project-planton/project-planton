# Microservice Kubernetes Pulumi Module

## Important Note

*This module is not completely implemented. Certain features may be missing or not fully functional. Future updates will address these limitations.*

## Overview

The **Microservice Kubernetes Pulumi Module** is designed to simplify the deployment and management of microservices on Kubernetes within a multi-cloud infrastructure. Leveraging Planton Cloud's unified API framework, this module models each API resource using a Kubernetes-like structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status` fields. The `MicroserviceKubernetes` resource encapsulates the necessary specifications for provisioning Kubernetes-based microservices, enabling developers to manage their application infrastructure as code with ease and consistency.

By utilizing this Pulumi module, developers can automate the creation and configuration of Kubernetes deployments based on defined specifications such as container images, environment variables, resource allocations, and network protocols. The module seamlessly integrates with Kubernetes clusters and other cloud provider services specified in the resource definition, ensuring secure and authenticated interactions. Additionally, the outputs generated from the deployment, including service endpoints and resource identifiers, are captured in the resource's `status.stackOutputs`. This facilitates effective monitoring and management of microservices directly through the `MicroserviceKubernetes` resource, enhancing operational efficiency and infrastructure visibility.

## Key Features

### API Resource Features

- **Standardized Structure**: The `MicroserviceKubernetes` API resource adheres to a consistent schema with `apiVersion`, `kind`, `metadata`, `spec`, and `status` fields. This uniformity ensures compatibility and ease of integration within Kubernetes-like environments, promoting seamless workflow incorporation and tooling interoperability.

- **Configurable Specifications**:
    - **Environment Information**: Define the environment context (`envId`) to ensure that microservices are deployed within the correct organizational and production contexts.
    - **Versioning**: Specify the version of the microservice to manage deployments and rollbacks effectively.
    - **Container Configuration**:
        - **Image Repository and Tag**: Define the container image repository and tag for deploying the microservice.
        - **Environment Variables**: Configure environment variables to pass necessary configurations and secrets to the microservice.
        - **Resource Allocation**: Set resource requests and limits for CPU and memory to ensure optimal performance and resource utilization.
        - **Ports Configuration**: Define the ports used by the container, including ingress ports for external access.
        - **Network Protocols**: Specify the network protocols (e.g., TCP, HTTP) used by the microservice to facilitate proper communication and service discovery.

- **Validation and Compliance**: Incorporates stringent validation rules to ensure all configurations adhere to established standards and best practices, minimizing the risk of misconfigurations and enhancing overall system reliability.

### Pulumi Module Features

- **Automated Kubernetes Provider Setup**: Utilizes the provided Kubernetes credentials to automatically configure the Pulumi Kubernetes provider, enabling seamless and secure interactions with Kubernetes clusters.

- **Microservice Management**: Streamlines the creation and management of Kubernetes deployments based on the provided specifications. This includes setting up deployment parameters, service configurations, and ingress rules to align with organizational requirements and performance standards.

- **Environment Isolation**: Manages isolated environments within Kubernetes, ensuring resources are organized and segregated according to organizational needs, which aids in maintaining clear boundaries and reducing resource conflicts.

- **Exported Stack Outputs**: Captures essential outputs such as service endpoints and resource identifiers in `status.stackOutputs`. These outputs facilitate integration with other infrastructure components, enabling effective monitoring, management, and automation workflows.

- **Scalability and Flexibility**: Designed to accommodate a wide range of microservice configurations, the module supports varying levels of complexity and can be easily extended to meet evolving infrastructure demands, ensuring long-term adaptability.

- **Error Handling**: Implements robust error handling mechanisms to promptly identify and report issues encountered during deployment or configuration processes, aiding in swift troubleshooting and maintaining infrastructure integrity.

## Installation

To integrate the Microservice Kubernetes Pulumi Module into your project, retrieve it from the [GitHub repository](https://github.com/your-repo/microservice-kubernetes-pulumi-module). Ensure that you have both Pulumi and Go installed and properly configured in your development environment.

```shell
git clone https://github.com/your-repo/microservice-kubernetes-pulumi-module.git
cd microservice-kubernetes-pulumi-module
```

## Usage

Refer to the [example section](#examples) for usage instructions.

## Examples

### Basic Example

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

### Example with Environment Variables

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      env:
        variables:
          DATABASE_NAME: todo
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: rest-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

### Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      env:
        secrets:
          # value before dot 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
          # value after dot 'database-password' is one of the secrets list in 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
          DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.database-password}
        variables:
          DATABASE_NAME: todo
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: rest-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

### Example with Multiple Containers

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: multi-container-app
spec:
  environmentInfo:
    envId: dev-env
  version: v1.0.0
  container:
    app:
      image:
        repo: myapp/frontend
        tag: v1.0.0
      ports:
        - appProtocol: http
          containerPort: 80
          isIngressPort: true
          servicePort: 8080
      resources:
        requests:
          cpu: 200m
          memory: 256Mi
        limits:
          cpu: 1000m
          memory: 1Gi
    sidecar:
      image:
        repo: myapp/logging
        tag: v1.0.0
      ports:
        - appProtocol: tcp
          containerPort: 5000
          isIngressPort: false
          servicePort: 5000
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 512Mi
```

### Example with Different Resource Limits

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: high-memory-service
spec:
  environmentInfo:
    envId: staging-env
  version: beta
  container:
    app:
      image:
        repo: highmemapp/backend
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 9090
          isIngressPort: true
          servicePort: 9090
      resources:
        requests:
          cpu: 500m
          memory: 512Mi
        limits:
          cpu: 4000m
          memory: 8Gi
```

### Example with Annotations and Labels

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: annotated-service
  labels:
    app: annotated-app
    tier: backend
  annotations:
    description: "This service handles user authentication."
spec:
  environmentInfo:
    envId: production-env
  version: release
  container:
    app:
      image:
        repo: auth-service/image
        tag: release
      ports:
        - appProtocol: https
          containerPort: 8443
          isIngressPort: true
          servicePort: 443
      resources:
        requests:
          cpu: 250m
          memory: 256Mi
        limits:
          cpu: 1500m
          memory: 2Gi
```

### Example with Health Checks

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: healthcheck-service
spec:
  environmentInfo:
    envId: test-env
  version: test
  container:
    app:
      image:
        repo: healthapp/service
        tag: test
      ports:
        - appProtocol: http
          containerPort: 8000
          isIngressPort: true
          servicePort: 8000
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 1Gi
      livenessProbe:
        httpGet:
          path: /healthz
          port: 8000
        initialDelaySeconds: 30
        periodSeconds: 10
      readinessProbe:
        httpGet:
          path: /ready
          port: 8000
        initialDelaySeconds: 10
        periodSeconds: 5
```

### Example with Empty Spec

*Note: This module is not completely implemented. Certain features may be missing or not fully functional. Future updates will address these limitations.*

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: incomplete-service
spec: {}
```

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: another-incomplete-service
spec: {}
```

## Module Details

### Input Configuration

The module expects a `MicroserviceKubernetesStackInput` which includes:

- **Pulumi Input**: Configuration details required by Pulumi for managing the stack, such as stack names, project settings, and any necessary Pulumi configurations.
- **Target API Resource**: The `MicroserviceKubernetes` resource defining the desired microservice configuration, including container specifications, environment variables, and resource allocations.
- **Kubernetes Credential**: Specifications for the Kubernetes credentials used to authenticate and authorize Pulumi operations, ensuring secure interactions with Kubernetes clusters.

### Exported Outputs

Upon successful execution, the module exports the following outputs to `status.stackOutputs`:

- **Service URL**: The URL of the deployed microservice, facilitating client interactions and service accessibility.
- **Resource ID**: The unique identifier assigned to the created Kubernetes resources, which can be used for management and monitoring purposes.

These outputs facilitate integration with other infrastructure components, provide essential information for monitoring and management, and enable automation workflows to utilize the deployed microservice resources effectively.

## Contributing

We welcome contributions to enhance the Microservice Kubernetes Pulumi Module. Please refer to our [contribution guidelines](CONTRIBUTING.md) for more information on how to get involved, including coding standards, submission processes, and best practices.

## License

This project is licensed under the [MIT License](LICENSE), granting you the freedom to use, modify, and distribute the software with minimal restrictions. Please review the LICENSE file for more details.

## Support

For support, please contact our [support team](mailto:support@planton.cloud). We are here to help you with any issues, questions, or feedback you may have regarding the Microservice Kubernetes Pulumi Module.

## Acknowledgements

Special thanks to all contributors and the Planton Cloud community for their ongoing support and feedback. Your efforts and dedication are instrumental in making this module robust and effective.

## Changelog

A detailed changelog is available in the [CHANGELOG.md](CHANGELOG.md) file. It documents all significant changes, enhancements, bug fixes, and updates made to the Microservice Kubernetes Pulumi Module over time.

## Roadmap

We are continuously working to enhance the Microservice Kubernetes Pulumi Module. Upcoming features include:

- **Advanced IAM Configurations**: Implementing more granular permission controls for Kubernetes resources to enhance security and compliance.
- **Enhanced Monitoring Integrations**: Integrating with monitoring and logging tools such as Prometheus, Grafana, and the ELK stack for better observability and performance tracking.
- **Support for Additional Cloud Providers**: Extending support to more cloud platforms to increase flexibility and accommodate diverse infrastructure requirements.
- **Automated Scaling and Optimization**: Introducing automated scaling capabilities based on workload demands and performance metrics to optimize resource utilization.
- **Comprehensive Documentation and Tutorials**: Expanding documentation and providing step-by-step tutorials to assist users in effectively leveraging the module's capabilities.

Stay tuned for more updates as we continue to develop and refine the module to meet your infrastructure management needs.

## Contact

For any inquiries or feedback, please reach out to us at [contact@planton.cloud](mailto:contact@planton.cloud). We value your input and are committed to providing the support you need to effectively manage your microservice infrastructure.

## Disclaimer

*This project is maintained by Planton Cloud and is not affiliated with any third-party services unless explicitly stated. While we strive to ensure the accuracy and reliability of this module, users are encouraged to review and test configurations in their environments.*

## Security

If you discover any security vulnerabilities, please report them responsibly by contacting our security team at [security@planton.cloud](mailto:security@planton.cloud). We take security seriously and are committed to addressing any issues promptly to protect our users and their infrastructure.

## Code of Conduct

Please adhere to our [Code of Conduct](CODE_OF_CONDUCT.md) when interacting with the project. We are committed to fostering an inclusive and respectful community where all contributors feel welcome and valued.

## References

- [Pulumi Documentation](https://www.pulumi.com/docs/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Planton Cloud APIs](https://buf.build/project-planton/apis/docs)
