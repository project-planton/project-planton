##############################################
# Create a Kubernetes Secret for CronJobKubernetes
##############################################

resource "kubernetes_secret" "this" {
  count = (
  can(var.spec.env.secrets)
  && length(var.spec.env.secrets) > 0
  ) ? 1 : 0

  metadata {
    # If you want the secret to match the CronJob version name
    # adjust as needed (e.g. var.metadata.name).
    name      = var.spec.version
    namespace = kubernetes_namespace.this.metadata[0].name
    labels    = local.final_labels
  }

  type = "Opaque"

  # `stringData` automatically converts each map value into a string,
  # then Kubernetes encodes it as base64 in the final secret.
  string_data = try(var.spec.env.secrets, {})

  depends_on = [
    kubernetes_namespace.this
  ]
}
