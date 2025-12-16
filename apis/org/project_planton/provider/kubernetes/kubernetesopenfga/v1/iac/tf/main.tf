# Conditionally create namespace based on create_namespace flag
resource "kubernetes_namespace" "openfga_namespace" {
  count = var.spec.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}

# Look up existing namespace when not creating
data "kubernetes_namespace" "existing_namespace" {
  count = var.spec.create_namespace ? 0 : 1

  metadata {
    name = local.namespace
  }
}
