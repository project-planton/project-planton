resource "helm_release" "openfga_helm_chart" {
  name             = local.resource_id
  repository       = "https://openfga.github.io/helm-charts"
  chart            = "openfga"
  version          = "0.2.12"
  namespace        = local.namespace
  create_namespace = false

  values = [
    yamlencode({
      fullnameOverride = local.kube_service_name
      replicaCount     = var.spec.container.replicas
      datastore = {
        engine = var.spec.datastore.engine
        uri    = var.spec.datastore.uri
      }
      resources = {
        requests = {
          cpu    = try(var.spec.container.resources.requests.cpu, null)
          memory = try(var.spec.container.resources.requests.memory, null)
        }
        limits = {
          cpu    = try(var.spec.container.resources.limits.cpu, null)
          memory = try(var.spec.container.resources.limits.memory, null)
        }
      }
    })
  ]
}
