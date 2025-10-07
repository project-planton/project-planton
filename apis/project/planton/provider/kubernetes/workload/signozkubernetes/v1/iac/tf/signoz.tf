resource "helm_release" "signoz" {
  name       = local.resource_id
  repository = "https://charts.signoz.io"
  chart      = "signoz"
  version    = "0.52.0"
  namespace  = kubernetes_namespace_v1.signoz_namespace.metadata[0].name

  values = [
    yamlencode(
      merge(
        {
          fullnameOverride = var.metadata.name
          podLabels        = local.final_labels
          commonLabels     = local.final_labels

          # Use bitnamilegacy registry due to Bitnami discontinuing free Docker Hub images (Sep 2025)
          # See: https://github.com/bitnami/containers/issues/83267
          # Global image registry override for all Bitnami images (ClickHouse and ZooKeeper dependencies)
          global = {
            imageRegistry = "docker.io/bitnamilegacy"
          }

          # Configure SigNoz container (main binary with UI, API, Ruler, Alertmanager)
          signoz = merge(
            {
              replicaCount = var.spec.signoz_container.replicas
              resources = {
                limits = {
                  cpu    = var.spec.signoz_container.resources.limits.cpu
                  memory = var.spec.signoz_container.resources.limits.memory
                }
                requests = {
                  cpu    = var.spec.signoz_container.resources.requests.cpu
                  memory = var.spec.signoz_container.resources.requests.memory
                }
              }
            },
            # Add image configuration if provided
            var.spec.signoz_container.image != null ? {
              image = {
                repository = var.spec.signoz_container.image.repo
                tag        = var.spec.signoz_container.image.tag
              }
            } : {}
          )

          # Configure OpenTelemetry Collector
          otelCollector = merge(
            {
              replicaCount = var.spec.otel_collector_container.replicas
              resources = {
                limits = {
                  cpu    = var.spec.otel_collector_container.resources.limits.cpu
                  memory = var.spec.otel_collector_container.resources.limits.memory
                }
                requests = {
                  cpu    = var.spec.otel_collector_container.resources.requests.cpu
                  memory = var.spec.otel_collector_container.resources.requests.memory
                }
              }
            },
            # Add image configuration if provided
            var.spec.otel_collector_container.image != null ? {
              image = {
                repository = var.spec.otel_collector_container.image.repo
                tag        = var.spec.otel_collector_container.image.tag
              }
            } : {}
          )

          # Configure database (ClickHouse)
          # External ClickHouse configuration
          clickhouse = !var.spec.database.is_external ? merge(
            {
              enabled      = true
              replicaCount = var.spec.database.managed_database.container.replicas
              resources = {
                limits = {
                  cpu    = var.spec.database.managed_database.container.resources.limits.cpu
                  memory = var.spec.database.managed_database.container.resources.limits.memory
                }
                requests = {
                  cpu    = var.spec.database.managed_database.container.resources.requests.cpu
                  memory = var.spec.database.managed_database.container.resources.requests.memory
                }
              }
              persistence = {
                enabled = var.spec.database.managed_database.container.is_persistence_enabled
                size    = var.spec.database.managed_database.container.disk_size
              }
            },
            # Add image configuration if provided
            var.spec.database.managed_database.container.image != null ? {
              image = {
                repository = var.spec.database.managed_database.container.image.repo
                tag        = var.spec.database.managed_database.container.image.tag
              }
            } : {},
            # Add clustering configuration if enabled
            local.cluster_is_enabled ? {
              layout = {
                shardsCount   = local.shard_count
                replicasCount = local.replica_count
              }
            } : {}
            ) : {
            enabled = false
          }

          # External ClickHouse configuration
          externalClickhouse = var.spec.database.is_external && var.spec.database.external_database != null ? {
            host     = var.spec.database.external_database.host
            httpPort = var.spec.database.external_database.http_port
            tcpPort  = var.spec.database.external_database.tcp_port
            cluster  = var.spec.database.external_database.cluster_name
            secure   = var.spec.database.external_database.is_secure
            user     = var.spec.database.external_database.username
            password = var.spec.database.external_database.password
          } : null

          # Zookeeper configuration (required for distributed ClickHouse)
          zookeeper = !var.spec.database.is_external && local.zookeeper_is_enabled && var.spec.database.managed_database.zookeeper.container != null ? merge(
            {
              enabled      = true
              replicaCount = var.spec.database.managed_database.zookeeper.container.replicas
              resources = {
                limits = {
                  cpu    = var.spec.database.managed_database.zookeeper.container.resources.limits.cpu
                  memory = var.spec.database.managed_database.zookeeper.container.resources.limits.memory
                }
                requests = {
                  cpu    = var.spec.database.managed_database.zookeeper.container.resources.requests.cpu
                  memory = var.spec.database.managed_database.zookeeper.container.resources.requests.memory
                }
              }
              persistence = {
                size = var.spec.database.managed_database.zookeeper.container.disk_size
              }
            },
            # Add image configuration if provided
            var.spec.database.managed_database.zookeeper.container.image != null ? {
              image = {
                repository = var.spec.database.managed_database.zookeeper.container.image.repo
                tag        = var.spec.database.managed_database.zookeeper.container.image.tag
              }
            } : {}
            ) : {
            enabled = false
          }
        },
        # Merge any user-provided helm_values
        var.spec.helm_values != null ? var.spec.helm_values : {}
      )
    )
  ]

  depends_on = [
    kubernetes_namespace_v1.signoz_namespace
  ]
}

