resource "kubernetes_namespace" "solr_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
