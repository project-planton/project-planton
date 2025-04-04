# Overview

The GCP GCP GKE Cluster API resource provides a consistent and streamlined interface for creating and managing Google Kubernetes Engine (GKE) clusters within our cloud infrastructure. By abstracting the complexities of GKE cluster configurations, this resource allows you to deploy and manage Kubernetes clusters effortlessly while ensuring consistency and compliance across different environments.

## Why We Created This API Resource

Managing GKE clusters directly can be complex due to the numerous configuration options, networking setups, and best practices that need to be considered. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:

- **Simplify Cluster Deployment**: Easily configure and deploy GKE clusters without dealing with low-level GCP configurations.
- **Ensure Consistency**: Maintain uniform cluster configurations across different environments and projects.
- **Enhance Productivity**: Reduce the time and effort required to set up Kubernetes clusters, allowing you to focus on application development and deployment.
- **Optimize Resource Usage**: Utilize autoscaling and node pool configurations to efficiently manage resources based on workload demands.

## Key Features

### Environment Integration

- **Environment Info**: Integrates seamlessly with our environment management system to deploy clusters within specific environments.
- **Stack Job Settings**: Supports custom stack job settings for infrastructure-as-code deployments.

### GCP Credential Management

- **GCP Credential ID**: Utilizes specified GCP credentials to ensure secure and authorized operations within Google Cloud Platform.

### Customizable Cluster Specifications

#### Cluster Configuration

- **Billing Account ID**: Links GKE cluster projects to a specified GCP billing account. Planton Cloud creates one or two GCP projects per GKE cluster, which are linked to this billing account.
- **Region and Zone**: Specify the GCP region and zone where the GKE cluster will be created. Selecting the appropriate region and zone optimizes performance and complies with data residency requirements.
- **Shared VPC**: Option to create the cluster in a shared VPC network by setting `is_create_shared_vpc` to true. This allows multiple projects to share a common VPC network for better resource management.
- **Workload Logs**: Toggle workload logs for the GKE cluster environment using `is_workload_logs_enabled`. When enabled, logs from Kubernetes pods are sent to Google Cloud Logging. Note that enabling log forwarding may increase cloud costs depending on log volume.

#### Cluster Autoscaling

- **Autoscaling Configuration**: Configure cluster-level autoscaling using `cluster_autoscaling_config`.
    - **Enable Autoscaling**: Set `is_enabled` to true to allow the cluster to automatically scale up or down based on resource requirements.
    - **CPU and Memory Limits**: Define minimum and maximum CPU cores (`cpu_min_cores`, `cpu_max_cores`) and memory in GB (`memory_min_gb`, `memory_max_gb`) that the cluster can scale to.

#### Node Pools

- **Node Pool Management**: Define one or more node pools using the `node_pools` field.
    - **Name**: Assign a name to each node pool, which is added as a label and can be used for workload scheduling.
    - **Machine Type**: Specify the machine type for nodes in the pool (e.g., `n2-custom-8-16384`).
    - **Node Count**: Set minimum and maximum node counts (`min_node_count`, `max_node_count`) to control autoscaling within the node pool.
    - **Spot Instances**: Enable spot instances using `is_spot_enabled` to reduce costs for workloads that can tolerate interruptions.

#### Kubernetes Add-ons

- **Add-on Installation**: Control the installation of various Kubernetes add-ons using `kubernetes_addons`.
    - **Postgres Operator**: Install the Postgres operator (`is_install_postgres_operator`) for managing PostgreSQL databases.
    - **Kafka Operator**: Install the Kafka operator (`is_install_kafka_operator`) for managing Apache Kafka clusters.
    - **Solr Operator**: Install the Solr operator (`is_install_solr_operator`) for managing Apache Solr instances.
    - **Kubecost**: Install Kubecost (`is_install_kubecost`) for cost monitoring and optimization.
    - **Planton Cloud Kube Agent**: Configure the Planton Cloud Kube Agent add-on using `planton_cloud_kube_agent`.
        - **Install Agent**: Set `is_install` to true to install the agent.
        - **Machine Account Email**: The agent uses a machine account email (`machine_account_email`), which is created if the agent is installed.
    - **Ingress Controllers and Other Add-ons**: Enable or disable other add-ons like Ingress NGINX (`is_install_ingress_nginx`), Istio (`is_install_istio`), Cert Manager (`is_install_cert_manager`), External DNS (`is_install_external_dns`), External Secrets (`is_install_external_secrets`), Elastic Operator (`is_install_elastic_operator`), and Keycloak Operator (`is_install_keycloak_operator`).

#### Ingress DNS Domains

- **Ingress DNS Configuration**: Define ingress DNS domains using `ingress_dns_domains`.
    - **Domain Name**: Specify the DNS domain name for ingress (e.g., `example.com`).
    - **TLS Enablement**: Set `is_tls_enabled` to control TLS for the domain. Certificates are not created for domains without TLS enabled.
    - **DNS Zone Project ID**: The GCP project ID containing the DNS zone (`dns_zone_gcp_project_id`) is computed and used for DNS validation when setting up certificates.

## Benefits

- **Simplified Deployment**: Reduces the complexity involved in setting up GKE clusters with a user-friendly API.
- **Consistency**: Ensures all clusters adhere to organizational standards for security, performance, and scalability.
- **Scalability**: Leverages autoscaling features at both the cluster and node pool levels to handle varying workloads efficiently.
- **Security**: Integrates with GCP IAM and VPCs to enhance security and compliance.
- **Cost Optimization**: Utilize spot instances and autoscaling to optimize resource usage and reduce costs.
- **Flexibility**: Provides extensive customization to meet specific application requirements without compromising best practices.
- **Enhanced Observability**: Optionally enable workload logs and install monitoring add-ons like Kubecost for better visibility into cluster operations.
