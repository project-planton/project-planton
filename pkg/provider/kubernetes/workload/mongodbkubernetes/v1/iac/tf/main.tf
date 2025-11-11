resource "kubernetes_namespace_v1" "mongodb_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
