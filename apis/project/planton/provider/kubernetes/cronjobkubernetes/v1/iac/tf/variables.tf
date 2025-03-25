#########################################
# variables.tf
#
# Defines the input variables for a CronJobKubernetes resource.
#########################################

variable "metadata" {
  description = "Metadata for the CronJob resource, including name, org, env, labels, etc."
  type = object({
    # The name of the CronJob resource (required).
    name = string

    # The unique identifier for this resource (optional).
    # If not provided, the resource may fall back to using `metadata.name`.
    id = optional(string)

    # The organization that owns this resource (optional).
    org = optional(string)

    # The environment in which the resource is deployed (optional).
    env = optional(string)

    # Additional labels to apply to the resource (optional).
    labels = optional(map(string))

    # An optional list of tags for indexing/categorization.
    tags = optional(list(string))

    # An optional version object to store versioning info for reference/logging.
    version = optional(object({
      id      = string
      message = string
    }))
  })
}

variable "spec" {
  description = "Spec defines the configuration for the CronJobKubernetes resource."
  type = object({

    # Required schedule in standard cron format (e.g., "0 0 * * *" for daily at midnight).
    schedule = string

    # How concurrency is handled. Common values: "Allow", "Forbid", "Replace".
    concurrency_policy = optional(string)

    # Whether to suspend subsequent runs. Default: false.
    suspend = optional(bool)

    # Maximum number of completed jobs to keep in history. Default: 3.
    successful_jobs_history_limit = optional(number)

    # Maximum number of failed jobs to keep in history. Default: 1.
    failed_jobs_history_limit = optional(number)

    # Number of retries before the job is considered failed. Default: 6.
    backoff_limit = optional(number)

    # Deadline in seconds for starting a missed job. If 0 or unset, no deadline is enforced.
    starting_deadline_seconds = optional(number)

    # Determines whether pods restart upon failure ("Always", "OnFailure", or "Never").
    # Typically "Never" is recommended for CronJobs. Default: "Never".
    restart_policy = optional(string)

    # The container image details for the CronJob.
    image = object({
      # The repository of the image (e.g., "gcr.io/project/image").
      repo = string
      # The tag of the image (e.g., "latest" or "v1.0.0").
      tag = string
      # Optional name of the image pull secret for private repos.
      pull_secret_name = optional(string)
    })

    # The resource requests and limits for the container.
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

    # Environment variables and secrets for the container.
    env = object({
      # A map of environment variable names to their values (non-sensitive).
      variables = optional(map(string))
      # A map of environment variable names to secrets.
      secrets = optional(map(string))
    })
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
