resource "kubernetes_namespace" "elasticsearch_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
