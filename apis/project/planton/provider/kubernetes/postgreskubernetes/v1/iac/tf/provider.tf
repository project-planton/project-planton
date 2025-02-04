terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.21"
    }
  }
}

provider "kubernetes" {
  config_path      = "~/.kube/config"
}
