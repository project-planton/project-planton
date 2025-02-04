resource "kubernetes_service" "this" {
  metadata {
    name      = var.spec.version
    namespace = kubernetes_namespace.this.metadata[0].name
    labels    = local.final_labels
  }

  spec {
    type     = "ClusterIP"
    selector = local.final_labels

    dynamic "port" {
      for_each = try(var.spec.container.app.ports, [])
      content {
        name        = port.value.name
        protocol    = port.value.network_protocol
        port        = port.value.service_port
        target_port = port.value.container_port
        # The appProtocol is only recognized in newer
        # Kubernetes versions. If you need it, you can
        # set it via an annotation or the official field:
        app_protocol = port.value.app_protocol
      }
    }
  }

  depends_on = [
    kubernetes_deployment.this
  ]
}
