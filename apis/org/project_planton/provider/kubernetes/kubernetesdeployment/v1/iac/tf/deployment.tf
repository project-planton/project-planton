# 1) Create a ServiceAccount
resource "kubernetes_service_account" "this" {
  metadata {
    name      = local.resource_id
    namespace = local.namespace
  }
}

# 3) Create the Deployment
resource "kubernetes_deployment" "this" {
  metadata {
    name      = var.metadata.name
    namespace = local.namespace
    labels    = local.final_labels
    annotations = {
      # Example annotation (remove or modify as needed)
      "example.annotation" = "true"
    }
  }

  spec {
    # If not provided, default to 1 replica
    replicas = try(var.spec.availability.min_replicas, 1)

    selector {
      match_labels = local.final_labels
    }

    template {
      metadata {
        labels = local.final_labels
      }

      spec {
        service_account_name             = kubernetes_service_account.this.metadata[0].name
        termination_grace_period_seconds = 60

        container {
          name  = "microservice"
          image = "${var.spec.container.app.image.repo}:${var.spec.container.app.image.tag}"

          # Container ports
          dynamic "port" {
            for_each = try(var.spec.container.app.ports, [])
            content {
              name           = port.value.name
              container_port = port.value.container_port
            }
          }

          # Add built-in environment variables
          env {
            name = "HOSTNAME"
            value_from {
              field_ref {
                field_path = "status.podIP"
              }
            }
          }

          env {
            name = "K8S_POD_ID"
            value_from {
              field_ref {
                field_path  = "metadata.name"
                api_version = "v1"
              }
            }
          }

          # Add env variables from var.spec.container.app.env.variables
          dynamic "env" {
            for_each = try(var.spec.container.app.env.variables, {})
            content {
              name  = env.key
              value = env.value
            }
          }

          # Add env variables from secrets (referenced by var.spec.version)
          dynamic "env" {
            for_each = try(var.spec.container.app.env.secrets, {})
            content {
              name = env.key
              value_from {
                secret_key_ref {
                  name = var.spec.version
                  key  = env.key
                }
              }
            }
          }

          # Resource requests/limits
          resources {
            limits = {
              cpu = try(var.spec.container.app.resources.limits.cpu, null)
              memory = try(var.spec.container.app.resources.limits.memory, null)
            }
            requests = {
              cpu = try(var.spec.container.app.resources.requests.cpu, null)
              memory = try(var.spec.container.app.resources.requests.memory, null)
            }
          }

          # Lifecycle pre-stop hook (sleep 60 seconds)
          lifecycle {
            pre_stop {
              exec {
                command = ["/bin/sleep", "60"]
              }
            }
          }
        }
      }
    }
  }

  depends_on = [
    kubernetes_namespace.this
  ]
}
