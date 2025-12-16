# Create namespace for Temporal deployment (only if create_namespace is true)
resource "kubernetes_namespace_v1" "temporal_namespace" {
  count = var.spec.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}

# Create secret for external database password (only when external database is configured)
resource "kubernetes_secret_v1" "db_password" {
  count = local.has_external_database ? 1 : 0

  metadata {
    name      = local.database_secret_name
    namespace = local.namespace
    labels    = local.final_labels
  }

  data = {
    (local.database_secret_key) = local.external_db_password
  }

  type = "Opaque"
}

# Deploy Temporal using Helm chart
resource "helm_release" "temporal" {
  name       = var.metadata.name
  repository = local.helm_chart_repository
  chart      = local.helm_chart_name
  version    = local.helm_chart_version
  namespace  = local.namespace

  # Wait for deployment to complete
  wait          = true
  wait_for_jobs = true
  timeout       = 600

  # Override default name
  set {
    name  = "fullnameOverride"
    value = var.metadata.name
  }

  # ---------------------------------------------------------------- Database Configuration
  # External database configuration (when provided)
  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "cassandra.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "mysql.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "postgresql.enabled"
      value = "false"
    }
  }

  # External database SQL configuration for default persistence
  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.driver"
      value = "sql"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.driver"
      value = local.sql_driver
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.host"
      value = local.external_db_host
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.port"
      value = local.external_db_port
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.database"
      value = local.database_name
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.user"
      value = local.external_db_username
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.existingSecret"
      value = local.database_secret_name
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.tls.enabled"
      value = "true"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.default.sql.tls.enableHostVerification"
      value = "false"
    }
  }

  # External database SQL configuration for visibility persistence
  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.driver"
      value = "sql"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.driver"
      value = local.sql_driver
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.host"
      value = local.external_db_host
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.port"
      value = local.external_db_port
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.database"
      value = local.visibility_name
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.user"
      value = local.external_db_username
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.existingSecret"
      value = local.database_secret_name
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.tls.enabled"
      value = "true"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.visibility.sql.tls.enableHostVerification"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = local.has_external_database ? [1] : []
    content {
      name  = "server.config.persistence.driver"
      value = "sql"
    }
  }

  # Embedded database configuration (when external database is not provided)
  # Cassandra
  dynamic "set" {
    for_each = !local.has_external_database && local.is_cassandra ? [1] : []
    content {
      name  = "cassandra.enabled"
      value = "true"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_cassandra ? [1] : []
    content {
      name  = "cassandra.replicaCount"
      value = var.spec.cassandra_replicas
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_cassandra ? [1] : []
    content {
      name  = "cassandra.config.dev"
      value = "true"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_cassandra ? [1] : []
    content {
      name  = "cassandra.config.cluster_size"
      value = var.spec.cassandra_replicas
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_cassandra ? [1] : []
    content {
      name  = "mysql.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_cassandra ? [1] : []
    content {
      name  = "postgresql.enabled"
      value = "false"
    }
  }

  # MySQL
  dynamic "set" {
    for_each = !local.has_external_database && local.is_mysql ? [1] : []
    content {
      name  = "cassandra.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_mysql ? [1] : []
    content {
      name  = "mysql.enabled"
      value = "true"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_mysql ? [1] : []
    content {
      name  = "postgresql.enabled"
      value = "false"
    }
  }

  # PostgreSQL
  dynamic "set" {
    for_each = !local.has_external_database && local.is_postgresql ? [1] : []
    content {
      name  = "cassandra.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_postgresql ? [1] : []
    content {
      name  = "mysql.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = !local.has_external_database && local.is_postgresql ? [1] : []
    content {
      name  = "postgresql.enabled"
      value = "true"
    }
  }

  # ---------------------------------------------------------------- Frontend Service Ports
  set {
    name  = "server.config.services.frontend.rpc.grpcPort"
    value = local.frontend_grpc_port
  }

  set {
    name  = "server.config.services.frontend.rpc.httpPort"
    value = local.frontend_http_port
  }

  # ---------------------------------------------------------------- Schema Setup
  set {
    name  = "schema.createDatabase.enabled"
    value = !local.disable_auto_schema_setup
  }

  set {
    name  = "schema.setup.enabled"
    value = "true"
  }

  set {
    name  = "schema.update.enabled"
    value = "true"
  }

  # ---------------------------------------------------------------- Web UI
  dynamic "set" {
    for_each = var.spec.disable_web_ui ? [1] : []
    content {
      name  = "web.enabled"
      value = "false"
    }
  }

  # ---------------------------------------------------------------- Monitoring Stack
  set {
    name  = "prometheus.enabled"
    value = local.enable_monitoring_stack
  }

  set {
    name  = "grafana.enabled"
    value = local.enable_monitoring_stack
  }

  set {
    name  = "kubePrometheusStack.enabled"
    value = local.enable_monitoring_stack
  }

  # ---------------------------------------------------------------- Elasticsearch Configuration
  # External Elasticsearch
  dynamic "set" {
    for_each = local.has_external_elasticsearch ? [1] : []
    content {
      name  = "elasticsearch.enabled"
      value = "false"
    }
  }

  dynamic "set" {
    for_each = local.has_external_elasticsearch ? [1] : []
    content {
      name  = "elasticsearch.host"
      value = local.external_es_host
    }
  }

  dynamic "set" {
    for_each = local.has_external_elasticsearch ? [1] : []
    content {
      name  = "elasticsearch.port"
      value = local.external_es_port
    }
  }

  dynamic "set" {
    for_each = local.has_external_elasticsearch ? [1] : []
    content {
      name  = "elasticsearch.scheme"
      value = "http"
    }
  }

  dynamic "set" {
    for_each = local.has_external_elasticsearch && local.external_es_user != "" ? [1] : []
    content {
      name  = "elasticsearch.username"
      value = local.external_es_user
    }
  }

  dynamic "set" {
    for_each = local.has_external_elasticsearch && local.external_es_password != "" ? [1] : []
    content {
      name  = "elasticsearch.password"
      value = local.external_es_password
    }
  }

  # Embedded Elasticsearch (when not using external and enabled in spec)
  dynamic "set" {
    for_each = !local.has_external_elasticsearch && !var.spec.enable_embedded_elasticsearch ? [1] : []
    content {
      name  = "elasticsearch.enabled"
      value = "false"
    }
  }

  depends_on = [
    kubernetes_secret_v1.db_password
  ]
}

# Create LoadBalancer service for frontend gRPC ingress (when enabled)
resource "kubernetes_service_v1" "frontend_grpc_lb" {
  count = local.frontend_ingress_enabled && local.frontend_grpc_hostname != "" ? 1 : 0

  metadata {
    name      = "${var.metadata.name}-frontend-grpc-lb"
    namespace = local.namespace
    labels    = local.final_labels
    annotations = {
      "external-dns.alpha.kubernetes.io/hostname" = local.frontend_grpc_hostname
    }
  }

  spec {
    type = "LoadBalancer"

    port {
      name        = "grpc"
      port        = local.frontend_grpc_port
      target_port = local.frontend_grpc_port
      protocol    = "TCP"
    }

    selector = {
      "app.kubernetes.io/name"      = "temporal"
      "app.kubernetes.io/instance"  = var.metadata.name
      "app.kubernetes.io/component" = "frontend"
    }
  }

  depends_on = [
    helm_release.temporal
  ]
}
