resource "kubernetes_namespace" "ingress_nginx" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}

resource "helm_release" "ingress_nginx" {
  name       = local.release_name
  namespace  = kubernetes_namespace.ingress_nginx.metadata[0].name
  repository = local.helm_chart_repo
  chart      = local.helm_chart_name
  version    = local.chart_version

  create_namespace = false
  atomic           = true
  cleanup_on_fail  = true
  wait_for_jobs    = true
  timeout          = 180

  # Controller configuration
  set {
    name  = "controller.service.type"
    value = local.service_type
  }

  set {
    name  = "controller.ingressClassResource.default"
    value = "true"
  }

  set {
    name  = "controller.watchIngressWithoutClass"
    value = "true"
  }

  # Apply service annotations dynamically
  dynamic "set" {
    for_each = local.service_annotations
    content {
      name  = "controller.service.annotations.${replace(set.key, "/", "\\.")}"
      value = set.value
    }
  }
}

