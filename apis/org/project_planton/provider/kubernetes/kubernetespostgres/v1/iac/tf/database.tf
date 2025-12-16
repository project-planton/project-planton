resource "kubernetes_manifest" "database" {
  manifest = {
    apiVersion = "acid.zalan.do/v1"
    kind       = "postgresql"
    metadata = {
      # For the Zalando operator, the name must be prefixed by the teamId (which is "db")
      # followed by our stable resource ID.
      name      = "db-${local.resource_id}"
      namespace = local.namespace_name
      labels    = local.final_labels
    }
    spec = {
      # Number of PostgreSQL instances (replicas)
      numberOfInstances = var.spec.container.replicas

      # Patroni configuration (empty object to satisfy CRD schema)
      patroni = {}

      # Pod annotations
      podAnnotations = {
        "postgres-cluster-id" = local.resource_id
      }

      # PostgreSQL settings
      postgresql = {
        version    = "14"
        parameters = {
          "max_connections" = "200"
        }
      }

      # Resource allocations
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

      # Team ID is required by the Zalando operator
      teamId = "db"

      # Persistent volume configuration
      volume = {
        size = var.spec.container.disk_size
      }
    }
  }
}
