output "namespace" {
  description = "The namespace in which the Gateway/HTTPRoute resources are created."
  value       = var.istio_ingress_namespace
}
