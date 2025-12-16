##############################################
# Create a Kubernetes Secret for CronJobKubernetes
##############################################

resource "kubernetes_secret" "this" {
  count = (
    can(var.spec.env.secrets)
    && length(var.spec.env.secrets) > 0
  ) ? 1 : 0

  metadata {
    # Secret name is "main" - referenced by cron_job.tf
    name      = "main"
    namespace = local.namespace_name
    labels    = local.final_labels
  }

  type = "Opaque"

  # `data` automatically converts each map value into a string,
  # then Kubernetes encodes it as base64 in the final secret.
  data = try(var.spec.env.secrets, {})
}
