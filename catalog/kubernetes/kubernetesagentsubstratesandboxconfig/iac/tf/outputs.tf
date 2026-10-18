output "sandbox_config_name" {
  description = "Name of the created SandboxConfig (its metadata.name, from spec.name)."
  value       = var.spec.name
}

output "sandbox_class" {
  description = "Sandbox class of the created SandboxConfig."
  value       = var.spec.sandboxClass
}
