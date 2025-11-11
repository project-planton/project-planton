resource "kubernetes_namespace" "openfga_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
