# Terraform module for Percona Operator for PostgreSQL
# This is a placeholder - the operator is primarily deployed via Helm/Pulumi

locals {
  namespace = var.spec.namespace != "" ? var.spec.namespace : "kubernetes-percona-postgres-operator"
}

resource "kubernetes_namespace" "kubernetes_percona_postgres_operator" {
  metadata {
    name = local.namespace
  }
}

resource "helm_release" "kubernetes_percona_postgres_operator" {
  name       = "pg-operator"
  repository = "https://percona.github.io/percona-helm-charts/"
  chart      = "pg-operator"
  version    = "2.7.0"
  namespace  = kubernetes_namespace.kubernetes_percona_postgres_operator.metadata[0].name

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
  description = "The namespace where the Percona PostgreSQL operator is deployed"
  value       = kubernetes_namespace.kubernetes_percona_postgres_operator.metadata[0].name
}

