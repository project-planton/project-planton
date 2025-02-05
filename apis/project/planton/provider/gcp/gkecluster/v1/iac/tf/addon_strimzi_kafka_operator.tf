###############################################################################
# Strimzi Kafka Operator
#
# 1. Create the "strimzi-operator" namespace, labeled with final_kubernetes_labels.
# 2. Deploy the Strimzi Kafka Operator Helm chart into that namespace.
# 3. Provide any necessary chart values, such as watchAnyNamespace = true.
###############################################################################

##############################################
# 1. strimzi-operator Namespace
##############################################
resource "kubernetes_namespace_v1" "strimzi_operator_namespace" {
  metadata {
    name   = "strimzi-operator"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 2. Helm Release for the Strimzi Kafka Operator
##############################################
resource "helm_release" "strimzi_kafka_operator" {
  name             = "strimzi-kafka-operator"
  repository       = "https://strimzi.io/charts/"
  chart            = "strimzi-kafka-operator"
  version          = "0.42.0"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.strimzi_operator_namespace.metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait             = true

  # Provide any custom values if needed
  values = [
    yamlencode({
      watchAnyNamespace = true
    })
  ]

  lifecycle {
    ignore_changes = [
      status,
      description
    ]
  }

  depends_on = [
    kubernetes_namespace_v1.strimzi_operator_namespace
  ]
}
