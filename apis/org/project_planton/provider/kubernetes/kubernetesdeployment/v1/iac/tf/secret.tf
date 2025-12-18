##############################################
# Create a Kubernetes Secret instead of ExternalSecret
# Uses computed name to avoid conflicts when multiple deployments share a namespace
##############################################
resource "kubernetes_secret" "this" {
  count = (
  can(var.spec.container.app.env.secrets)
  && length(var.spec.container.app.env.secrets) > 0
  ) ? 1 : 0

  metadata {
    name      = local.env_secret_name
    namespace = local.namespace
    labels    = local.final_labels
  }

  type = "Opaque"

  # Populate the secret with key-value pairs.
  data = { for k, v in try(var.spec.container.app.env.secrets, {}) : k => base64encode(v) }
}
