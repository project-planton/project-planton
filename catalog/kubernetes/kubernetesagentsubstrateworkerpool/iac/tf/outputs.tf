output "worker_pool_name" {
  description = "Name of the created WorkerPool (equals metadata.name)."
  value       = var.metadata.name
}

output "namespace" {
  description = "Namespace of the created WorkerPool."
  value       = var.spec.namespace
}
