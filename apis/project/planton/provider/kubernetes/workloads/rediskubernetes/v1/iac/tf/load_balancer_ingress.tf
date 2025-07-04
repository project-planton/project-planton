resource "kubernetes_service" "redis_external_lb" {
  count = var.spec.ingress.is_enabled && var.spec.ingress.dns_domain != "" ? 1 : 0

  metadata {
    name      = "ingress-external-lb"
    namespace = kubernetes_namespace.redis_namespace.metadata[0].name

    labels = local.final_labels

    annotations = {
      "external-dns.alpha.kubernetes.io/hostname" = local.ingress_external_hostname
    }
  }

  spec {
    type = "LoadBalancer"

    port {
      name        = "tcp-redis"
      port        = 6379
      protocol    = "TCP"
      target_port = "redis"
    }

    selector = local.redis_pod_selector_labels
  }
}

resource "kubernetes_service" "redis_internal_lb" {
  count = var.spec.ingress.is_enabled && var.spec.ingress.dns_domain != "" ? 1 : 0

  metadata {
    name      = "ingress-internal-lb"
    namespace = kubernetes_namespace.redis_namespace.metadata[0].name

    labels = local.final_labels

    annotations = {
      "cloud.google.com/load-balancer-type"       = "Internal"
      "external-dns.alpha.kubernetes.io/hostname" = local.ingress_internal_hostname
    }
  }

  spec {
    type = "LoadBalancer"

    port {
      name        = "tcp-redis"
      port        = 6379
      protocol    = "TCP"
      target_port = "redis"
    }

    selector = local.redis_pod_selector_labels
  }
}
