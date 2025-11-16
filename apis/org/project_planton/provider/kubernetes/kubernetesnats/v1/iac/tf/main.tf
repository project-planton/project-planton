resource "kubernetes_namespace" "nats_namespace" {
  metadata {
    name   = local.namespace
    labels = local.labels
  }
}
