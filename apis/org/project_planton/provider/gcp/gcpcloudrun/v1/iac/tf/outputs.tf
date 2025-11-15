# Public or internal URL of the Cloud Run service
output "url" {
  description = "Public or internal URL of the Cloud Run service"
  value       = google_cloud_run_v2_service.main.uri
}

# Name of the Cloud Run service
output "service_name" {
  description = "Name of the Cloud Run service"
  value       = google_cloud_run_v2_service.main.name
}

# Name of the deployed revision
output "revision" {
  description = "Name of the latest deployed revision"
  value       = google_cloud_run_v2_service.main.latest_ready_revision
}

# Additional outputs for convenience
output "project_id" {
  description = "GCP project ID where the service is deployed"
  value       = google_cloud_run_v2_service.main.project
}

output "location" {
  description = "GCP region where the service is deployed"
  value       = google_cloud_run_v2_service.main.location
}

output "service_id" {
  description = "Fully qualified service ID"
  value       = google_cloud_run_v2_service.main.id
}

