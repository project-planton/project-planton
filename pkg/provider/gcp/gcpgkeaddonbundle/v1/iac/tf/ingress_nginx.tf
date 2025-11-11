###############################################################################
# Ingress NGINX
#
# This file sets up the Ingress NGINX controller using the official Helm chart.
# 1. Creates a dedicated namespace, labeled with our final_kubernetes_labels.
# 2. Deploys the Ingress NGINX helm_release with appropriate values.
###############################################################################

resource "kubernetes_namespace_v1" "ingress_nginx_namespace" {
  # Conditionally create this namespace
  count = var.spec.install_ingress_nginx ? 1 : 0

  metadata {
    name   = "ingress-nginx"
    labels = local.final_kubernetes_labels
  }
}

resource "helm_release" "ingress_nginx" {
  # Conditionally create this helm_release
  count            = var.spec.install_ingress_nginx ? 1 : 0
  name             = "ingress-nginx"
  repository       = "https://kubernetes.github.io/ingress-nginx"
  chart            = "ingress-nginx"
  version          = "4.11.1"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.ingress_nginx_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait             = true

  # We encode the values in YAML. The "controller.service.type" is set to "ClusterIP"
  # and "controller.ingressClassResource.default" is true, making this the default ingress class.
  values = [
    yamlencode({
      controller = {
        service = {
          type = "ClusterIP"
        }
        ingressClassResource = {
          default = true
        }
      }
    })
  ]
}
