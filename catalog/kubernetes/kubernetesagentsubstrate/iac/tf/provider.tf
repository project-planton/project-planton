# Provider requirements for the KubernetesAgentSubstrate module.
#
# kubectl (alekc/kubectl) applies upstream's documents with server-side
# apply. http fetches the pinned release files and upstream's database setup
# script. kubernetes reads, at plan time, the workloads, namespaces and
# version stamps the preflight checks and the definitions the kept-CRD block
# is about to write, and writes the Secrets. tls generates the certificate
# authorities and the actor token signing key. helm is declared because
# helm_crds.tf is one generated file for chart and bundle kinds alike, and
# init resolves every data source in it; this module never renders a chart.
#
# Providers are configured by the calling workspace (the kubeconfig
# environment contract).

terraform {
  required_version = ">= 1.4"

  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.35"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 3.0"
    }
    kubectl = {
      source  = "alekc/kubectl"
      version = ">= 2.0"
    }
    http = {
      source  = "hashicorp/http"
      version = "~> 3.4"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "~> 4.0"
    }
  }
}

provider "kubernetes" {
}

provider "helm" {
}

provider "kubectl" {
}
