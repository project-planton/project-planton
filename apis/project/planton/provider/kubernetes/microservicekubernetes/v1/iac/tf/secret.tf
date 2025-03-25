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
    namespace = kubernetes_namespace.this.metadata[0].name
    labels    = local.final_labels
  }

  type = "Opaque"

  # Populate the secret with key-value pairs.
  # Note: 'string_data' automatically converts the map values into string form.
  string_data = try(var.spec.container.app.env.secrets, {})

  depends_on = [
    kubernetes_namespace.this
  ]
}
