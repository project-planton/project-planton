variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
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
    shared_vpc_config = optional(object({

      # A flag indicating whether the cluster should be created in a shared VPC network.
      # **Warning:** The GKE cluster will be recreated if this is updated.
      is_enabled = optional(bool, false)

      # Description for vpc_project_id
      vpc_project_id = optional(string, "")
    }), {})

    # A flag to toggle workload logs for the GKE cluster environment.
    # When enabled, logs from Kubernetes pods will be sent to Google Cloud Logging.
    # **Warning:** Enabling log forwarding may increase cloud bills depending on the log volume.
    is_workload_logs_enabled = optional(bool, false)

    # Configuration for cluster autoscaling.
    cluster_autoscaling_config = optional(object({

      # A flag to enable or disable autoscaling of Kubernetes worker nodes.
      # When set to true, the cluster will automatically scale up or down based on resource requirements.
      is_enabled = optional(bool, false)

      # The minimum number of CPU cores the cluster can scale down to when autoscaling is enabled.
      # This is the total number of CPU cores across all nodes in the cluster.
      cpu_min_cores = optional(number, 0)

      # The maximum number of CPU cores the cluster can scale up to when autoscaling is enabled.
      # This is the total number of CPU cores across all nodes in the cluster.
      cpu_max_cores = optional(number, 0)

      # The minimum amount of memory in gigabytes (GB) the cluster can scale down to when autoscaling is enabled.
      # This is the total memory across all nodes in the cluster.
      memory_min_gb = optional(number, 0)

      # The maximum amount of memory in gigabytes (GB) the cluster can scale up to when autoscaling is enabled.
      # This is the total memory across all nodes in the cluster.
      memory_max_gb = optional(number, 0)
    }), {})

    # A list of node pools for the GKE cluster.
    node_pools = optional(list(object({

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
      is_spot_enabled = optional(bool, false)
    })), [])
  })
}
