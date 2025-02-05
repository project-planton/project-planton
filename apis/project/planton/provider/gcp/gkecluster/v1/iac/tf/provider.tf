terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.19.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.35"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.9"
    }
    http = {
      source  = "hashicorp/http"
      version = "~> 3.2"
    }
  }
}

provider "google" {
}

data "google_client_config" "default" {}

provider "kubernetes" {
  host                   = "https://${google_container_cluster.gke_cluster.endpoint}"
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(google_container_cluster.gke_cluster.master_auth[0].cluster_ca_certificate)

  ignore_annotations = [
    "^autopilot\\.gke\\.io\\/.*",
    "^cloud\\.google\\.com\\/.*"
  ]
}

provider "helm" {
  kubernetes {
    host                   = google_container_cluster.gke_cluster.endpoint
    cluster_ca_certificate = base64decode(
      google_container_cluster.gke_cluster.master_auth[0].cluster_ca_certificate
    )
    token                  = data.google_client_config.default.access_token
  }
}
