###############################################################################
# Zalando Postgres Operator
#
# 1. Create a dedicated "postgres-operator" namespace, labeled with final_kubernetes_labels.
# 2. Deploy the Zalando Postgres Operator Helm chart into that namespace.
# 3. Pass in the "configKubernetes.inherited_labels" values to propagate base labeling.
###############################################################################

##############################################
# 1. postgres-operator Namespace
##############################################
resource "kubernetes_namespace_v1" "zalando_postgres_operator_namespace" {
  metadata {
    name   = "postgres-operator"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 2. Helm Release for Zalando Postgres Operator
##############################################
resource "helm_release" "zalando_postgres_operator" {
  name             = "postgres-operator"
  repository       = "https://opensource.zalando.com/postgres-operator/charts/postgres-operator"
  chart            = "postgres-operator"
  version          = "1.12.2"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.zalando_postgres_operator_namespace.metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Inherit and propagate base labeling
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

  lifecycle {
    ignore_changes = [
      status,
      description
    ]
  }

  depends_on = [
    kubernetes_namespace_v1.zalando_postgres_operator_namespace
  ]
}
