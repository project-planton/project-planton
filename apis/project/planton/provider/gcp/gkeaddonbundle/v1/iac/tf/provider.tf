terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.35"
    }
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.19.0"
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

provider "kubernetes" {
  config_raw = file("/Users/swarup/Desktop/deleteme/kcon/project-planton-cluster.config")
}

provider "kubectl" {
  config_path = "/Users/swarup/Desktop/deleteme/kcon/project-planton-cluster.config"
}

provider "helm" {
  kubernetes {
    config_path = "/Users/swarup/Desktop/deleteme/kcon/project-planton-cluster.config"
  }
}
