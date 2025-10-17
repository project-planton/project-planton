resource "kubernetes_service" "redis_external_lb" {
  count = var.spec.ingress.enabled && var.spec.ingress.hostname != "" ? 1 : 0

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
