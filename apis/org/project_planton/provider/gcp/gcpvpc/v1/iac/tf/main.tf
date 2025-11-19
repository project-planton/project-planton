#############################################
# Enable Compute API
#############################################

resource "google_project_service" "compute_api" {
  project                    = var.spec.project_id.value
  service                    = "compute.googleapis.com"
  disable_dependent_services = true
}

#############################################
# GCP VPC Network
#############################################

resource "google_compute_network" "vpc" {
  name                    = var.spec.network_name
  project                 = var.spec.project_id.value
  auto_create_subnetworks = var.spec.auto_create_subnetworks
  routing_mode            = local.routing_mode

  # Ensure Compute API is enabled before creating VPC
  depends_on = [google_project_service.compute_api]
}

