resource "kubernetes_namespace" "jenkins_namespace" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}
