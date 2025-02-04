terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.21"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.9"
    }
  }
}

provider "kubernetes" {
  config_path      = "~/.kube/config"
}

provider "helm" {
  kubernetes {
    config_path      = "~/.kube/config"
  }
}
