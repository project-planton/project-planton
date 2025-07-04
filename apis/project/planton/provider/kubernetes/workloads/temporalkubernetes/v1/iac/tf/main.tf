resource "kubernetes_namespace" "redis_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
