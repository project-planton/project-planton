# Terraform module for Percona Operator for MongoDB
# This is a placeholder - the operator is primarily deployed via Helm/Pulumi

locals {
  namespace = var.spec.namespace != "" ? var.spec.namespace : "percona-operator"
}

resource "kubernetes_namespace" "percona_operator" {
  metadata {
    name = local.namespace
  }
}

resource "helm_release" "percona_operator" {
  name       = "psmdb-operator"
  repository = "https://percona.github.io/percona-helm-charts/"
  chart      = "psmdb-operator"
  version    = "1.16.0"
  namespace  = kubernetes_namespace.percona_operator.metadata[0].name

  set {
    name  = "resources.limits.cpu"
    value = var.spec.container.resources.limits.cpu
  }

  set {
    name  = "resources.limits.memory"
    value = var.spec.container.resources.limits.memory
  }

  set {
    name  = "resources.requests.cpu"
    value = var.spec.container.resources.requests.cpu
  }

  set {
    name  = "resources.requests.memory"
    value = var.spec.container.resources.requests.memory
  }

  timeout         = 300
  atomic          = true
  cleanup_on_fail = true
  wait            = true
}

output "namespace" {
  description = "The namespace where the Percona operator is deployed"
  value       = kubernetes_namespace.percona_operator.metadata[0].name
}

