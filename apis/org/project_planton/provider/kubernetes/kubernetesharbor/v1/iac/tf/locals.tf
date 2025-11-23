locals {
  namespace = (
    var.harbor_kubernetes.spec.namespace != null && var.harbor_kubernetes.spec.namespace != ""
    ? var.harbor_kubernetes.spec.namespace
    : var.harbor_kubernetes.metadata.name
  )

  labels = {
    "app.kubernetes.io/name"     = "harbor"
    "app.kubernetes.io/instance" = var.harbor_kubernetes.metadata.name
  }

  core_service_name     = "${var.harbor_kubernetes.metadata.name}-harbor-core"
  portal_service_name   = "${var.harbor_kubernetes.metadata.name}-harbor-portal"
  registry_service_name = "${var.harbor_kubernetes.metadata.name}-harbor-registry"
}

