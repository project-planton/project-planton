##############################################
# Create a Kubernetes Secret for CronJobKubernetes
##############################################

resource "kubernetes_secret" "this" {
  count = (
    can(var.spec.env.secrets)
    && length(var.spec.env.secrets) > 0
  ) ? 1 : 0

  metadata {
    # Computed name to avoid conflicts when multiple instances share a namespace
    name      = local.env_secrets_secret_name
    namespace = local.namespace_name
    labels    = local.final_labels
  }

  type = "Opaque"

  # `data` automatically converts each map value into a string,
  # then Kubernetes encodes it as base64 in the final secret.
  data = try(var.spec.env.secrets, {})
}
