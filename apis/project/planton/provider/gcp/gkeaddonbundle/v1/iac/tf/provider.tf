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


provider "kubernetes" {
}

provider "helm" {
  kubernetes {
  }
}
