variable "kubernetes_cluster_credential" {
  description = "Kubernetes Credential data"
  type = object({

    # The Kubernetes provider, specifying the type of cluster (e.g., gcp_gke, aws_eks, azure_aks).
    # For more detailed information, refer to the official documentation of each provider.
    provider = string

    # Credential specification for GCP GKE provider.
    # This field contains detailed information for connecting to a Google Kubernetes Engine (GKE) cluster.
    gcp_gke = object({

      # The cluster endpoint for the GKE cluster.
      # This is a required field and must contain the URL endpoint for the GKE cluster.
      cluster_endpoint = string

      # The cluster Certificate Authority (CA) data for the GKE cluster.
      # This is a required field and must contain the base64 encoded CA certificate.
      cluster_ca_data = string

      # The base64 encoded Google Service Account JSON key.
      # This is a required field and must be a valid base64 encoded string representing the service account key.
      service_account_key_base64 = string
    })

    # Credential specification for AWS EKS provider.
    # This field contains detailed information for connecting to a Elastic Kubernetes Engine(EKS) cluster.
    aws_eks = object({})

    # Credential specification for Azure AKS provider.
    # This field contains detailed information for connecting to a Azure Kubernetes Service(AKS) cluster.
    azure_aks = object({})
  })
}

provider "kubernetes" {
}
