#########################################
# variables.tf
#
# Defines the input variables for a CronJobKubernetes resource.
#########################################

variable "metadata" {
  description = "Metadata for the CronJob resource, including name, org, env, labels, etc."
  type = object({
    name   = string
    id     = optional(string)
    org    = optional(string)
    env    = optional(string)
    labels = optional(map(string))
    tags   = optional(list(string))
    version = optional(object({
      id      = string
      message = string
    }))
  })
}

variable "spec" {
  description = "Spec defines the configuration for the CronJobKubernetes resource."
  type = object({
    target_cluster = object({
      cluster_name = string
      cluster_kind = optional(number)
    })
    namespace                     = string
    schedule                      = string
    concurrency_policy            = optional(string)
    suspend                       = optional(bool)
    successful_jobs_history_limit = optional(number)
    failed_jobs_history_limit     = optional(number)
    backoff_limit                 = optional(number)
    starting_deadline_seconds     = optional(number)
    restart_policy                = optional(string)
    image = object({
      repo             = string
      tag              = string
      pull_secret_name = optional(string)
    })
    resources = object({
      limits = object({
        cpu    = string
        memory = string
      })
      requests = object({
        cpu    = string
        memory = string
      })
    })
    env = object({
      variables = optional(map(string))
      secrets   = optional(map(string))
    })
    command = optional(list(string))
    args    = optional(list(string))
  })
}

variable "docker_config_json" {
  description = <<EOT
Optional Docker credentials in JSON format to create
an image pull secret (type: kubernetes.io/dockerconfigjson).
Leave empty if no private repo auth is needed.
EOT
  type        = string
  default     = ""
  sensitive   = true
}
