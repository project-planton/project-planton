variable "metadata" {
  description = "metadata"
  type = object({

    # name of the resource
    name = string

    # id of the resource
    id = string

    # id of the organization to which the api-resource belongs to
    org = string

    # environment to which the resource belongs to
    env = object({

      # name of the environment
      name = string

      # id of the environment
      id = string
    })

    # labels for the resource
    labels = object({

      # Description for key
      key = string

      # Description for value
      value = string
    })

    # tags for the resource
    tags = list(string)
  })
}

variable "spec" {
  description = "spec"
  type = object({

    # The GCP project ID in which the GKE cluster will be created.
    cluster_project_id = string

    # Required.** The GCP region where the GKE cluster will be created.
    # **Warning:** The GKE cluster will be recreated if this value is updated.
    # Refer to: https://cloud.google.com/compute/docs/regions-zones
    region = string

    # Required.** The GCP zone where the GKE cluster will be created.
    # Refer to: https://cloud.google.com/compute/docs/regions-zones
    zone = string

    # GkeClusterSharedVpcConfig** specifies the shared VPC network settings for GKE clusters.
    # This message includes the project ID for the shared VPC network where the GKE cluster is created.
    # For more details, visit: https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-shared-vpc
    shared_vpc_config = object({

      # A flag indicating whether the cluster should be created in a shared VPC network.
      # **Warning:** The GKE cluster will be recreated if this is updated.
      is_enabled = bool

      # Description for vpc_project_id
      vpc_project_id = string
    })

    # A flag to toggle workload logs for the GKE cluster environment.
    # When enabled, logs from Kubernetes pods will be sent to Google Cloud Logging.
    # **Warning:** Enabling log forwarding may increase cloud bills depending on the log volume.
    is_workload_logs_enabled = bool

    # Configuration for cluster autoscaling.
    cluster_autoscaling_config = object({

      # A flag to enable or disable autoscaling of Kubernetes worker nodes.
      # When set to true, the cluster will automatically scale up or down based on resource requirements.
      is_enabled = bool

      # The minimum number of CPU cores the cluster can scale down to when autoscaling is enabled.
      # This is the total number of CPU cores across all nodes in the cluster.
      cpu_min_cores = number

      # The maximum number of CPU cores the cluster can scale up to when autoscaling is enabled.
      # This is the total number of CPU cores across all nodes in the cluster.
      cpu_max_cores = number

      # The minimum amount of memory in gigabytes (GB) the cluster can scale down to when autoscaling is enabled.
      # This is the total memory across all nodes in the cluster.
      memory_min_gb = number

      # The maximum amount of memory in gigabytes (GB) the cluster can scale up to when autoscaling is enabled.
      # This is the total memory across all nodes in the cluster.
      memory_max_gb = number
    })

    # A list of node pools for the GKE cluster.
    node_pools = list(object({

      # The name of the node pool.
      # This name is added as a label to the node pool and can be used to schedule workloads.
      name = string

      # Required.** The machine type for the node pool (e.g., 'n2-custom-8-16234').
      machine_type = string

      # The minimum number of nodes in the node pool. Defaults to 1.
      min_node_count = number

      # The maximum number of nodes in the node pool. Defaults to 1.
      max_node_count = number

      # A flag to enable spot instances on the node pool. Defaults to false.
      is_spot_enabled = bool
    }))

    # Specifications for Kubernetes addons in the GKE cluster.
    kubernetes_addons = object({

      # A flag to control the installation of the PostgreSQL operator.
      is_install_postgres_operator = bool

      # A flag to control the installation of the Kafka operator.
      is_install_kafka_operator = bool

      # A flag to control the installation of the Solr operator.
      is_install_solr_operator = bool

      # A flag to control the installation of Kubecost.
      is_install_kubecost = bool

      # A flag to control the installation of Ingress NGINX.
      is_install_ingress_nginx = bool

      # A flag to control the installation of Istio.
      is_install_istio = bool

      # A flag to control the installation of Cert-Manager.
      is_install_cert_manager = bool

      # A flag to control the installation of External DNS.
      is_install_external_dns = bool

      # A flag to control the installation of External Secrets.
      is_install_external_secrets = bool

      # A flag to control the installation of the Elastic operator.
      is_install_elastic_operator = bool

      # A flag to control the installation of the Keycloak operator.
      is_install_keycloak_operator = bool
    })

    # Ingress DNS domains to be configured in the GKE cluster.
    ingress_dns_domains = list(object({

      # A unique identifier for the ingress DNS domain.
      id = string

      # Required.** The DNS domain name (e.g., 'example.com').
      name = string

      # A flag to enable TLS for the endpoint domain. Defaults to false.
      # **Important:** Certificates are not created for endpoints that do not require TLS.
      # Also, ingress DNS domains without TLS enabled cannot be used for creating endpoints for microservice instances,
      # PostgreSQL clusters, Kafka clusters, Redis clusters, or Solr clouds.
      is_tls_enabled = bool

      # The GCP project ID containing the DNS zone for the endpoint domain.
      # This value is retrieved from the DNS domains in the organization's DNS data.
      # It is required for configuring the certificate issuer to perform DNS validations.
      dns_zone_gcp_project_id = string
    }))
  })
}
