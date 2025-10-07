resource "helm_release" "clickhouse" {
  name       = local.resource_id
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "clickhouse"
  version    = "6.2.15"
  namespace  = kubernetes_namespace_v1.clickhouse_namespace.metadata[0].name

  values = [
    yamlencode(
      merge(
        {
          fullnameOverride  = var.metadata.name
          namespaceOverride = local.namespace
          shards            = var.spec.container.replicas
          replicaCount      = var.spec.container.replicas
          
          # Use bitnamilegacy registry due to Bitnami discontinuing free Docker Hub images (Sep 2025)
          # See: https://github.com/bitnami/containers/issues/83267
          global = {
            imageRegistry = "docker.io/bitnamilegacy"
          }
          
          image = {
            repository = "clickhouse"
          }
          
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
          
          podLabels    = local.final_labels
          commonLabels = local.final_labels
          
          auth = {
            existingSecret    = kubernetes_secret_v1.clickhouse_password.metadata[0].name
            existingSecretKey = "admin-password"
            username          = "default"
          }
          
          # Configure clustering if enabled
          keeper = local.cluster_is_enabled ? {
            enabled = true
          } : null
          
          zookeeper = local.cluster_is_enabled ? {
            enabled = true
            image = {
              repository = "zookeeper"
            }
          } : null
        },
        # Merge any user-provided helm_values
        var.spec.helm_values != null ? var.spec.helm_values : {}
      )
    )
  ]

  depends_on = [
    kubernetes_namespace_v1.clickhouse_namespace,
    kubernetes_secret_v1.clickhouse_password
  ]
}
