# Terraform module for Percona Operator for MongoDB

resource "kubernetes_namespace" "percona_operator" {
  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "helm_release" "percona_operator" {
  name       = local.helm_chart_name
  repository = local.helm_chart_repo
  chart      = local.helm_chart_name
  version    = local.helm_chart_version
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

  set {
    name  = "watchAllNamespaces"
    value = "true"
  }

  timeout         = 300
  atomic          = true
  cleanup_on_fail = true
  wait            = true
}

