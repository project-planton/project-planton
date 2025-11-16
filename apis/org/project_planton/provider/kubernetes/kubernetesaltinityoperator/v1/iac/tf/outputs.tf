# Outputs for Altinity ClickHouse Operator deployment
# These outputs match the KubernetesAltinityOperatorStackOutputs protobuf message

output "namespace" {
  description = "The namespace where the Altinity operator is deployed"
  value       = kubernetes_namespace.kubernetes_altinity_operator.metadata[0].name
}

