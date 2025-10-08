resource "random_password" "clickhouse_password" {
  length  = 16
  special = true
  numeric = true
  upper   = true
  lower   = true

  min_special = 3
  min_numeric = 2
  min_upper   = 2
  min_lower   = 2
}

resource "kubernetes_secret_v1" "clickhouse_password" {
  metadata {
    name      = var.metadata.name
    namespace = kubernetes_namespace_v1.clickhouse_namespace.metadata[0].name
  }

  data = {
    "admin-password" = base64encode(random_password.clickhouse_password.result)
  }

  type = "Opaque"
}
