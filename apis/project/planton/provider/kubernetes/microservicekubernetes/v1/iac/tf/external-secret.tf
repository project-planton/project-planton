resource "kubernetes_manifest" "external_secret" {
  # Only create this if there are any secrets defined
  count = (
  can(var.spec.container.app.env.secrets)
  && length(var.spec.container.app.env.secrets) > 0
  ) ? 1 : 0

  manifest = {
    apiVersion = "external-secrets.io/v1beta1"
    kind       = "ExternalSecret"
    metadata = {
      name      = var.spec.version
      namespace = kubernetes_namespace.this.metadata[0].name
      labels    = local.final_labels
    }
    spec = {
      refreshInterval = "1m"
      secretStoreRef = {
        kind = "ClusterSecretStore"
        # In Pulumi code, we used a variable for the
        # secret store name. For this example, we hardcode
        # "gcp-secrets-manager" or any store you use.
        name = "gcp-secrets-manager"
      }
      target = {
        # The final Kubernetes secret name that pods read from:
        name = var.spec.version
      }
      # Convert each entry in var.spec.container.app.env.secrets into
      # an array element for ExternalSecret `data`.
      data = [
        for key, remote_key in try(var.spec.container.app.env.secrets, {}) : {
          secretKey = key
          remoteRef = {
            key     = remote_key
            version = "latest"
          }
        }
      ]
    }
  }

  depends_on = [
    kubernetes_namespace.this
  ]
}
