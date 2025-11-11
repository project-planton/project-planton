output "namespace" {
  description = "The Kubernetes namespace where NATS is deployed."
  value       = local.namespace
}

output "internal_client_url" {
  description = "The internal NATS client URL for cluster-local connections."
  value       = local.internal_client_url
}

output "external_hostname" {
  description = "The external hostname for NATS if ingress is enabled."
  value       = local.ingress_hostname
}
