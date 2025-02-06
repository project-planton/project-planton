###############################################################################
# Elastic Operator
#
# 1. Create the "elastic-system" namespace, labeled with final_kubernetes_labels.
# 2. Deploy the Elastic ECK Operator Helm chart into that namespace.
# 3. Pass configKubernetes.inherited_labels if needed.
###############################################################################

##############################################
# 1. elastic-system Namespace
##############################################
resource "kubernetes_namespace_v1" "elastic_operator_namespace" {
  # Conditionally create this namespace based on the install_elastic_operator flag
  count = var.spec.install_elastic_operator ? 1 : 0

  metadata {
    name   = "elastic-system"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 2. Helm Release for the Elastic Operator
##############################################
resource "helm_release" "elastic_operator" {
  # Conditionally create this Helm release based on the install_elastic_operator flag
  count            = var.spec.install_elastic_operator ? 1 : 0
  name             = "eck-operator"
  repository       = "https://helm.elastic.co"
  chart            = "eck-operator"
  version          = "2.14.0"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.elastic_operator_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Provide any custom values if needed, e.g., inherited labels
  values = [
    yamlencode({
      configKubernetes = {
        inherited_labels = [
          "resource",
          "organization",
          "environment",
          "resource_kind",
          "resource_id"
        ]
      }
    })
  ]

  depends_on = [
    kubernetes_namespace_v1.elastic_operator_namespace
  ]
}
