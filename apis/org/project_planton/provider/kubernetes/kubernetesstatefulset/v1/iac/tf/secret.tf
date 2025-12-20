# Create a secret for environment secrets if any are defined
resource "kubernetes_secret" "env_secrets" {
  count = length(try(var.spec.container.app.env.secrets, {})) > 0 ? 1 : 0

  metadata {
    name      = local.env_secret_name
    namespace = local.namespace
    labels    = local.final_labels
  }

  type = "Opaque"

  data = try(var.spec.container.app.env.secrets, {})

  depends_on = [
    kubernetes_namespace.this
  ]
}
