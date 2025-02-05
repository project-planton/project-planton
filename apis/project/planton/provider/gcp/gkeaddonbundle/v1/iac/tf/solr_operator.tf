###############################################################################
# Solr Operator
#
# 1. Create the "solr-operator" namespace, labeled with final_kubernetes_labels.
# 2. Apply the CRDs from the Solr Operator manifest download URL.
# 3. Deploy the Solr Operator Helm chart into the namespace, referencing local.final_kubernetes_labels.
###############################################################################

##############################################
# 1. solr-operator Namespace
##############################################
resource "kubernetes_namespace_v1" "solr_operator_namespace" {
  metadata {
    name   = "solr-operator"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 2. CRD Resources from the Solr Operator manifest
##############################################
data "http" "solr_operator_crd_manifest" {
  url = "https://solr.apache.org/operator/downloads/crds/v0.7.0/all-with-dependencies.yaml"
}

resource "kubernetes_manifest" "solr_operator_crds" {
  manifest = yamldecode(data.http.solr_operator_crd_manifest.response_body)

  depends_on = [
    kubernetes_namespace_v1.solr_operator_namespace
  ]
}

##############################################
# 3. Helm Release for the Solr Operator
##############################################
resource "helm_release" "solr_operator" {
  name             = "solr-operator"
  repository       = "https://solr.apache.org/charts"
  chart            = "solr-operator"
  version          = "0.7.0"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.solr_operator_namespace.metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Provide any custom values if needed
  values = [
    yamlencode({})
  ]

  lifecycle {
    ignore_changes = [
      status,
      description
    ]
  }

  depends_on = [
    kubernetes_manifest.solr_operator_crds
  ]
}
