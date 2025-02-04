#
# 1) Generate a random root password for MongoDB
#
resource "random_password" "mongodb_root_password" {
  length     = 12
  special    = true
  numeric    = true
  upper      = true
  lower      = true
  min_special = 3
  min_numeric = 2
  min_upper   = 2
  min_lower   = 2
}

#
# 2) Store the generated password in a Kubernetes Secret
#    which the Helm chart will reference (via .auth.existingSecret)
#
resource "kubernetes_secret_v1" "mongodb_root_secret" {
  metadata {
    name      = local.kube_service_name
    namespace = kubernetes_namespace_v1.mongodb_namespace.metadata[0].name
  }

  # Base64-encode the random password before storing it
  data = {
    "mongodb-root-password" = base64encode(random_password.mongodb_root_password.result)
  }

  depends_on = [
    kubernetes_namespace_v1.mongodb_namespace
  ]
}

#
# 3) Deploy the Bitnami MongoDB Helm chart using yamlencode for the values
#    https://artifacthub.io/packages/helm/bitnami/mongodb
#
resource "helm_release" "mongodb" {
  name       = local.resource_id
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "mongodb"
  version    = "15.1.4"

  # Install the chart into the same namespace
  namespace = local.namespace

  # Merge user-provided helm_values with standard chart configuration from your proto-based spec.
  # Then wrap the entire map with yamlencode(...) so the chart sees it as one YAML block.
  values = [
    yamlencode(
      merge(
        {
          fullnameOverride  = local.kube_service_name
          namespaceOverride = local.namespace
          useStatefulSet    = true

          resources = {
            limits = {
              cpu    = try(var.spec.container.resources.limits.cpu, "1000m")
              memory = try(var.spec.container.resources.limits.memory, "1Gi")
            }
            requests = {
              cpu    = try(var.spec.container.resources.requests.cpu, "50m")
              memory = try(var.spec.container.resources.requests.memory, "100Mi")
            }
          }

          # By default, we treat this as a single replica for a standalone deployment.
          replicaCount = try(var.spec.container.replicas, 1)

          persistence = {
            enabled = try(var.spec.container.is_persistence_enabled, true)
            size    = try(var.spec.container.disk_size, "1Gi")
          }

          podLabels    = local.final_labels
          commonLabels = local.final_labels

          # Tells the MongoDB Helm chart to use an existing secret for the root password
          auth = {
            existingSecret = local.kube_service_name
          }
        },
          var.spec.helm_values != null ? var.spec.helm_values : {}
      )
    )
  ]

  depends_on = [
    kubernetes_secret_v1.mongodb_root_secret
  ]
}

#
# 4) (Optional) Create a LoadBalancer service for external access if ingress is enabled
#
resource "kubernetes_service_v1" "mongodb_ingress_lb" {
  count = (local.ingress_is_enabled && local.ingress_dns_domain != "") ? 1 : 0

  metadata {
    name      = "ingress-external-lb"
    namespace = kubernetes_namespace_v1.mongodb_namespace.metadata[0].name
    labels    = kubernetes_namespace_v1.mongodb_namespace.metadata[0].labels
    annotations = {
      "external-dns.alpha.kubernetes.io/hostname" = local.ingress_external_hostname
    }
  }

  spec {
    type = "LoadBalancer"
    port {
      name        = "tcp-mongodb"
      port        = 27017
      protocol    = "TCP"
      # The Bitnami chart exposes a named port "mongodb" on the container.
      target_port = "mongodb"
    }

    # This selector should match the labels of the pods created by the Helm chart
    selector = local.mongodb_pod_selector_labels
  }

  depends_on = [
    helm_release.mongodb
  ]
}
