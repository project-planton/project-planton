# Google Cloud Run Service v2
resource "google_cloud_run_v2_service" "main" {
  project  = var.spec.project_id
  location = var.spec.region
  name     = local.service_name

  # Ingress settings
  ingress = local.ingress

  labels = local.labels

  template {
    # Service account
    service_account = local.service_account_email

    # Timeout
    timeout = "${var.spec.timeout_seconds}s"

    # Execution environment
    execution_environment = local.execution_environment

    # Max concurrent requests per instance
    max_instance_request_concurrency = var.spec.max_concurrency

    # Scaling configuration
    scaling {
      min_instance_count = var.spec.container.replicas.min
      max_instance_count = var.spec.container.replicas.max
    }

    # VPC access configuration (if provided)
    dynamic "vpc_access" {
      for_each = local.has_vpc_access ? [1] : []
      content {
        dynamic "network_interfaces" {
          for_each = var.spec.vpc_access.network != null ? [1] : []
          content {
            network    = var.spec.vpc_access.network
            subnetwork = var.spec.vpc_access.subnet
          }
        }
        egress = var.spec.vpc_access.egress != null ? var.spec.vpc_access.egress : "PRIVATE_RANGES_ONLY"
      }
    }

    # Container configuration
    containers {
      image = local.container_image

      # Container port
      ports {
        container_port = local.port
      }

      # Resource limits
      resources {
        limits = {
          cpu    = local.cpu
          memory = local.memory
        }
      }

      # Plain environment variables
      dynamic "env" {
        for_each = local.env_vars
        content {
          name  = env.value.name
          value = env.value.value
        }
      }

      # Secret environment variables
      dynamic "env" {
        for_each = local.env_secrets
        content {
          name = env.value.name
          value_source {
            secret_key_ref {
              secret = env.value.secret
            }
          }
        }
      }
    }
  }

  # Lifecycle management: prevent service deletion before creating new revision
  lifecycle {
    create_before_destroy = true
  }
}

# IAM policy for public access (if allow_unauthenticated is true)
resource "google_cloud_run_v2_service_iam_member" "public_invoker" {
  count = var.spec.allow_unauthenticated ? 1 : 0

  project  = google_cloud_run_v2_service.main.project
  location = google_cloud_run_v2_service.main.location
  name     = google_cloud_run_v2_service.main.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

# Custom DNS domain mappings (if DNS is enabled)
resource "google_cloud_run_domain_mapping" "custom_domain" {
  for_each = local.has_custom_dns ? toset(local.dns_hostnames) : []

  project  = var.spec.project_id
  location = var.spec.region
  name     = each.value

  metadata {
    namespace = var.spec.project_id
    labels    = local.labels
  }

  spec {
    route_name = google_cloud_run_v2_service.main.name
  }

  # Wait for the service to be created before creating domain mappings
  depends_on = [google_cloud_run_v2_service.main]
}

# DNS records for domain verification (if DNS is enabled)
# Note: These need to be created after the domain mapping to retrieve verification records
# The actual implementation would use data sources to retrieve domain mapping status
# and create corresponding DNS records in Cloud DNS.
# This is commented out as it requires additional data source lookups for domain verification codes.
#
# data "google_cloud_run_domain_mapping" "verification" {
#   for_each = local.has_custom_dns ? toset(local.dns_hostnames) : []
#   
#   project  = var.spec.project_id
#   location = var.spec.region
#   name     = each.value
#   
#   depends_on = [google_cloud_run_domain_mapping.custom_domain]
# }
#
# resource "google_dns_record_set" "verification" {
#   for_each = local.has_custom_dns ? toset(local.dns_hostnames) : []
#   
#   managed_zone = var.spec.dns.managed_zone
#   name         = "${each.value}."
#   type         = "A"
#   ttl          = 300
#   
#   rrdatas = [
#     data.google_cloud_run_domain_mapping.verification[each.key].status[0].resource_records[0].rrdata
#   ]
# }

