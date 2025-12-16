##############################################
# Create a Kubernetes Secret instead of ExternalSecret
##############################################
resource "kubernetes_secret" "this" {
  count = (
  can(var.spec.container.app.env.secrets)
  && length(var.spec.container.app.env.secrets) > 0
  ) ? 1 : 0

  metadata {
    name      = var.spec.version
    namespace = local.namespace
    labels    = local.final_labels
  }

  type = "Opaque"

  # Populate the secret with key-value pairs.
  data = { for k, v in try(var.spec.container.app.env.secrets, {}) : k => base64encode(v) }
}
