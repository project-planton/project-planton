locals {
  # Helm chart configuration
  helm_chart_name    = "psmdb-operator"
  helm_chart_repo    = "https://percona.github.io/percona-helm-charts/"
  helm_chart_version = "1.20.1"

  # Namespace - use from spec or default to "percona-operator"
  namespace = var.spec.namespace != "" ? var.spec.namespace : "percona-operator"

  # Metadata labels
  labels = merge(
    var.metadata.labels != null ? var.metadata.labels : {},
    {
      "app.kubernetes.io/name"       = "percona-mongo-operator"
      "app.kubernetes.io/managed-by" = "project-planton"
      "app.kubernetes.io/component"  = "database-operator"
    }
  )
}

