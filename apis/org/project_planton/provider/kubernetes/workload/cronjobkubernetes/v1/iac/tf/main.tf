resource "kubernetes_namespace" "this" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
