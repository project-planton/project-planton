###############################################################################
# Solr Operator
#
# 1. Create the "solr-operator" namespace, labeled with final_kubernetes_labels.
# 2. Apply the CRDs from the Solr Operator manifest download URL (multiple docs).
# 3. Deploy the Solr Operator Helm chart into the namespace, referencing local.final_kubernetes_labels.
###############################################################################

##############################################
# 1. solr-operator Namespace
##############################################
resource "kubernetes_namespace_v1" "solr_operator_namespace" {
  count = var.spec.install_solr_operator ? 1 : 0

  metadata {
    name   = "solr-operator"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 2. CRD Resources from the Solr Operator manifest
##############################################
data "http" "solr_operator_crd_manifest" {
  count = var.spec.install_solr_operator ? 1 : 0
  url   = "https://solr.apache.org/operator/downloads/crds/v0.7.0/all-with-dependencies.yaml"
}

resource "kubectl_manifest" "solr_operator_crds" {
  count = var.spec.install_solr_operator ? 1 : 0

  # This provider supports multi-doc YAML out of the box,
  # so we can just feed the entire file in "yaml_body".
  yaml_body = data.http.solr_operator_crd_manifest[count.index].response_body

  # Optional: you can customize validation, patch strategies, etc.
  # skip_validation = true
  # replace_on_change = false

  depends_on = [
    kubernetes_namespace_v1.solr_operator_namespace
  ]
}

##############################################
# 3. Helm Release for the Solr Operator
##############################################
resource "helm_release" "solr_operator" {
  count            = var.spec.install_solr_operator ? 1 : 0
  name             = "solr-operator"
  repository       = "https://solr.apache.org/charts"
  chart            = "solr-operator"
  version          = "0.7.0"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.solr_operator_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait             = true

  # Provide any custom values if needed
  values = [
    yamlencode({})
  ]

  depends_on = [
    kubectl_manifest.solr_operator_crds
  ]
}
