locals {
  # Service name: use spec.service_name if provided, otherwise metadata.name
  service_name = var.spec.service_name != null ? var.spec.service_name : var.metadata.name

  # Construct full container image URI
  container_image = "${var.spec.container.image.repo}:${var.spec.container.image.tag}"

  # Format memory with Mi suffix
  memory = "${var.spec.container.memory}Mi"

  # CPU as string
  cpu = tostring(var.spec.container.cpu)

  # Container port
  port = var.spec.container.port

  # Labels: merge metadata labels with standard labels
  labels = merge(
    {
      "org"         = var.metadata.org != null ? var.metadata.org : "default"
      "env"         = var.metadata.env != null ? var.metadata.env : "default"
      "resource-id" = var.metadata.id != null ? var.metadata.id : var.metadata.name
    },
    var.metadata.labels != null ? var.metadata.labels : {}
  )

  # Environment variables: convert map to list of objects
  env_vars = var.spec.container.env != null && var.spec.container.env.variables != null ? [
    for k, v in var.spec.container.env.variables : {
      name  = k
      value = v
    }
  ] : []

  # Secret environment variables: convert map to list of objects
  env_secrets = var.spec.container.env != null && var.spec.container.env.secrets != null ? [
    for k, v in var.spec.container.env.secrets : {
      name   = k
      secret = v
    }
  ] : []

  # Convert ingress enum to GCP API value
  ingress_mapping = {
    "INGRESS_TRAFFIC_ALL"                      = "INGRESS_TRAFFIC_ALL"
    "INGRESS_TRAFFIC_INTERNAL_ONLY"           = "INGRESS_TRAFFIC_INTERNAL_ONLY"
    "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER" = "INGRESS_TRAFFIC_INTERNAL_AND_CLOUD_LOAD_BALANCING"
  }
  ingress = lookup(local.ingress_mapping, var.spec.ingress, "INGRESS_TRAFFIC_ALL")

  # Convert execution environment enum to GCP API value
  execution_environment = var.spec.execution_environment

  # Determine if VPC access is configured
  has_vpc_access = var.spec.vpc_access != null && (
    var.spec.vpc_access.network != null ||
    var.spec.vpc_access.subnet != null
  )

  # Determine if custom DNS is enabled
  has_custom_dns = var.spec.dns != null && var.spec.dns.enabled == true

  # DNS hostnames (empty list if DNS not configured)
  dns_hostnames = local.has_custom_dns ? var.spec.dns.hostnames : []

  # Service account email
  service_account_email = var.spec.service_account
}

