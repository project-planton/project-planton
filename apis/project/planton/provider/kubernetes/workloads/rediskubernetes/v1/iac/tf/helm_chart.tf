resource "helm_release" "redis" {
  name       = local.resource_id
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "redis"
  version    = "17.10.1"
  namespace  = kubernetes_namespace.redis_namespace.metadata[0].name

  # Convert your entire map into a single YAML string
  values = [
    yamlencode({
      fullnameOverride = var.metadata.name
      architecture     = "standalone"

      master = {
        podLabels = local.final_labels
        resources = {
          limits = {
            cpu    = var.spec.container.resources.limits.cpu
            memory = var.spec.container.resources.limits.memory
          }
          requests = {
            cpu    = var.spec.container.resources.requests.cpu
            memory = var.spec.container.resources.requests.memory
          }
        }
        persistence = {
          enabled = var.spec.container.is_persistence_enabled
          size    = var.spec.container.disk_size
        }
      }

      replica = {
        podLabels    = local.final_labels
        replicaCount = var.spec.container.replicas
        resources = {
          limits = {
            cpu    = var.spec.container.resources.limits.cpu
            memory = var.spec.container.resources.limits.memory
          }
          requests = {
            cpu    = var.spec.container.resources.requests.cpu
            memory = var.spec.container.resources.requests.memory
          }
        }
        persistence = {
          enabled = var.spec.container.is_persistence_enabled
          size    = var.spec.container.disk_size
        }
      }

      auth = {
        existingSecret            = "redis-password"
        existingSecretPasswordKey = "password"
      }
    })
  ]
}
