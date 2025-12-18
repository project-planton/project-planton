##############################################
# main.tf
##############################################

# 1) Create a ServiceAccount
resource "kubernetes_service_account" "this" {
  metadata {
    name      = local.resource_id
    namespace = local.namespace_name
  }
}

# 2) Create an optional image pull secret if Docker credentials are provided
resource "kubernetes_secret" "image_pull_secret" {
  metadata {
    # Computed name to avoid conflicts when multiple instances share a namespace
    name      = local.image_pull_secret_name
    namespace = local.namespace_name
    labels    = local.final_labels
  }
  type = "kubernetes.io/dockerconfigjson"

  data = var.docker_config_json != "" ? {
    ".dockerconfigjson" = var.docker_config_json
  } : {}
}

# 3) Create the CronJob
resource "kubernetes_cron_job" "this" {
  metadata {
    name      = var.metadata.name
    namespace = local.namespace_name
    labels    = local.final_labels
  }

  spec {
    schedule = var.spec.schedule

    # If not set, concurrencyPolicy defaults to "Forbid"
    concurrency_policy = try(var.spec.concurrency_policy, "Forbid")

    # If not set, default to false
    suspend = try(var.spec.suspend, false)

    # Defaults if not set
    successful_jobs_history_limit = try(var.spec.successful_jobs_history_limit, 3)
    failed_jobs_history_limit     = try(var.spec.failed_jobs_history_limit, 1)

    # For starting_deadline_seconds, if not set or zero, we use null so that it doesn't appear
    starting_deadline_seconds = (try(var.spec.starting_deadline_seconds, 0) != 0 ? var.spec.starting_deadline_seconds :
    null)

    job_template {
      metadata {
        labels = local.final_labels
      }
      
      spec {
        backoff_limit = try(var.spec.backoff_limit, 6)

        template {
          metadata {
            labels = local.final_labels
          }

          spec {
            # Typically "Never" is the recommended default for CronJobs
            restart_policy       = try(var.spec.restart_policy, "Never")
            service_account_name = kubernetes_service_account.this.metadata[0].name

            container {
              name  = "cronjob-container"
              image = "${var.spec.image.repo}:${var.spec.image.tag}"

              # Use the custom command and args if provided, otherwise default to empty lists
              command = try(var.spec.command, [])
              args    = try(var.spec.args, [])

              # Env variables
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

              # Add env variables from var.spec.env.variables
              dynamic "env" {
                for_each = try(var.spec.env.variables, {})
                content {
                  name  = env.key
                  value = env.value
                }
              }

              # Add env variables from secrets (stored in the env-secrets secret)
              dynamic "env" {
                for_each = try(var.spec.env.secrets, {})
                content {
                  name = env.key
                  value_from {
                    secret_key_ref {
                      name = local.env_secrets_secret_name
                      key  = env.key
                    }
                  }
                }
              }

              # Resource requests/limits
              resources {
                limits = {
                  cpu    = try(var.spec.resources.limits.cpu, null)
                  memory = try(var.spec.resources.limits.memory, null)
                }
                requests = {
                  cpu    = try(var.spec.resources.requests.cpu, null)
                  memory = try(var.spec.resources.requests.memory, null)
                }
              }
            }

            # If the image pull secret is non-empty, attach it
            image_pull_secrets {
              name = kubernetes_secret.image_pull_secret.metadata[0].name
            }
          }
        }
      }
    }
  }
}
