resource "kubernetes_namespace" "helm_release_namespace" {
  metadata {
    name   = local.namespace_name
    labels = local.final_labels
  }
}
