resource "kubernetes_namespace_v1" "clickhouse_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
