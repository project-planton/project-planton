locals {
  # Derive a stable resource ID
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels following Project Planton conventions
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_name" = var.metadata.name
    "resource_kind" = "kubernetes_temporal"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
    var.metadata.org != null && var.metadata.org != ""
    ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if var.metadata.env is non-empty
  env_label = (
    var.metadata.env != null && var.metadata.env != ""
    ) ? {
    "environment" = var.metadata.env
  } : {}

  # Merge all labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Namespace from spec
  namespace = var.spec.namespace

  # Service names
  frontend_service_name = "${var.metadata.name}-frontend"
  ui_service_name       = "${var.metadata.name}-web"

  # Service ports
  frontend_grpc_port = 7233
  frontend_http_port = 7243
  ui_port            = 8080

  # Internal cluster endpoints
  frontend_endpoint = "${local.frontend_service_name}.${local.namespace}.svc.cluster.local:${local.frontend_grpc_port}"
  web_ui_endpoint   = "${local.ui_service_name}.${local.namespace}.svc.cluster.local:${local.ui_port}"

  # Port-forward commands for local access
  port_forward_frontend_command = "kubectl port-forward -n ${local.namespace} service/${local.frontend_service_name} 7233:7233"
  port_forward_ui_command       = "kubectl port-forward -n ${local.namespace} service/${local.ui_service_name} 8080:8080"

  # Ingress configuration
  frontend_ingress_enabled = try(var.spec.ingress.frontend.enabled, false)
  frontend_grpc_hostname   = try(var.spec.ingress.frontend.grpc_hostname, "")
  frontend_http_hostname   = try(var.spec.ingress.frontend.http_hostname, "")
  web_ui_ingress_enabled   = try(var.spec.ingress.web_ui.enabled, false)
  web_ui_hostname          = try(var.spec.ingress.web_ui.hostname, "")

  # External hostname outputs (only set if ingress is enabled)
  external_frontend_hostname = local.frontend_ingress_enabled ? local.frontend_grpc_hostname : ""
  external_ui_hostname       = local.web_ui_ingress_enabled ? local.web_ui_hostname : ""

  # Database configuration
  database_backend          = var.spec.database.backend
  has_external_database     = var.spec.database.external_database != null
  database_name             = var.spec.database.database_name
  visibility_name           = var.spec.database.visibility_name
  disable_auto_schema_setup = var.spec.database.disable_auto_schema_setup

  # Database backend booleans
  is_cassandra  = local.database_backend == "cassandra"
  is_postgresql = local.database_backend == "postgresql"
  is_mysql      = local.database_backend == "mysql"

  # SQL sub-driver mapping
  sql_driver = local.is_postgresql ? "postgres12" : (local.is_mysql ? "mysql8" : "")

  # External database details (when provided)
  external_db_host     = try(var.spec.database.external_database.host, "")
  external_db_port     = try(var.spec.database.external_database.port, 0)
  external_db_username = try(var.spec.database.external_database.username, "")
  external_db_password = try(var.spec.database.external_database.password, "")

  # Monitoring stack configuration
  has_external_elasticsearch = var.spec.external_elasticsearch != null
  enable_monitoring_stack    = var.spec.enable_monitoring_stack || local.has_external_elasticsearch

  # External Elasticsearch details
  external_es_host     = try(var.spec.external_elasticsearch.host, "")
  external_es_port     = try(var.spec.external_elasticsearch.port, 0)
  external_es_user     = try(var.spec.external_elasticsearch.user, "")
  external_es_password = try(var.spec.external_elasticsearch.password, "")

  # Helm chart configuration
  helm_chart_name       = "temporal"
  helm_chart_repository = "https://go.temporal.io/helm-charts"
  helm_chart_version    = var.spec.version

  # Computed resource names to avoid conflicts when multiple instances share a namespace
  # Format: {metadata.name}-{purpose}
  database_secret_name = "${var.metadata.name}-db-password"
  database_secret_key  = "password"
}

