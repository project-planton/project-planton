resource "kubernetes_namespace_v1" "kafka_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
