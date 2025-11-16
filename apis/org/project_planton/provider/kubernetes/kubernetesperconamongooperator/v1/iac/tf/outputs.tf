output "namespace" {
  description = "The Kubernetes namespace where the Percona Operator for MongoDB is installed"
  value       = kubernetes_namespace.percona_operator.metadata[0].name
}

output "operator_version" {
  description = "The version of the Percona Operator for MongoDB Helm chart deployed"
  value       = local.helm_chart_version
}

output "operator_name" {
  description = "The name of the Percona Operator for MongoDB Helm release"
  value       = helm_release.percona_operator.name
}

output "helm_status" {
  description = "The status of the Helm release deployment"
  value       = helm_release.percona_operator.status
}

