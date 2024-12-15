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

    # cluster-config
    cluster_config = object({

      # The unique ID for the project to create the database user.
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#projectid_yaml
      project_id = string

      # Specifies the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.
      # Accepted values include:
      # REPLICASET Replica set
      # SHARDED Sharded cluster
      # GEOSHARDED Global Cluster
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#clustertype_yaml
      cluster_type = string

      # Number of electable nodes for Atlas to deploy to the region. Electable nodes can become the primary and can facilitate local reads.
      # The total number of electableNodes across all replication spec regions must total 3, 5, or 7.
      # Specify 0 if you do not want any electable nodes in the region.
      # You cannot create electable nodes in a region if priority is 0.
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#electablenodes_yaml
      electable_nodes = number

      # Election priority of the region. For regions with only read-only nodes, set this value to 0.
      # For regions where electable_nodes is at least 1, each region must have a priority of exactly one (1) less than the previous region. The first region must have a priority of 7. The lowest possible priority is 1.
      # The priority 7 region identifies the Preferred Region of the cluster. Atlas places the primary node in the Preferred Region. Priorities 1 through 7 are exclusive - no more than one region per cluster can be assigned a given priority.
      # Example: If you have three regions, their priorities would be 7, 6, and 5 respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be 4 and 3 respectively.
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#priority_yaml
      priority = number

      # Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#readonlynodes_yaml
      read_only_nodes = number

      # enable or disable cloud backup
      cloud_backup = bool

      # auto scaling disk db enabled
      auto_scaling_disk_gb_enabled = bool

      # Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: 4.4, 5.0, 6.0 or 7.0.
      # If omitted, Atlas deploys a cluster that runs MongoDB 7.0.
      # If provider_instance_size_name: M0, M2 or M5, Atlas deploys MongoDB 5.0.
      # Atlas always deploys the cluster with the latest stable release of the specified version
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#mongodbmajorversion_yaml
      mongo_db_major_version = string

      # Cloud service provider on which the servers are provisioned.
      #
      # The possible values are:
      #
      # AWS - Amazon AWS
      # GCP - Google Cloud Platform
      # AZURE - Microsoft Azure
      # TENANT - A multi-tenant deployment on one of the supported cloud service providers. Only valid when providerSettings.instanceSizeName is either M2 or M5.
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#providername_yaml
      provider_name = string

      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#providerinstancesizename_yaml
      # Atlas provides different instance sizes, each with a default storage capacity and RAM size.
      # The instance size you select is used for all the data-bearing servers in your cluster.
      # https://www.pulumi.com/registry/packages/mongodbatlas/api-docs/cluster/#providerinstancesizename_yaml
      provider_instance_size_name = string
    })
  })
}
