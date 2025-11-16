locals {
  # Helm chart configuration - using pg-operator for enterprise-grade PostgreSQL
  helm_chart_name    = "pg-operator"
  helm_chart_repo    = "https://percona.github.io/percona-helm-charts/"
  helm_chart_version = "2.7.0"

  # Namespace - use from spec or default to "kubernetes-percona-postgres-operator"
  namespace = var.spec.namespace != "" ? var.spec.namespace : "kubernetes-percona-postgres-operator"

  # Metadata labels
  labels = merge(
    var.metadata.labels != null ? var.metadata.labels : {},
    {
      "app.kubernetes.io/name"       = "percona-postgres-operator"
      "app.kubernetes.io/managed-by" = "project-planton"
      "app.kubernetes.io/component"  = "database-operator"
    }
  )
}

