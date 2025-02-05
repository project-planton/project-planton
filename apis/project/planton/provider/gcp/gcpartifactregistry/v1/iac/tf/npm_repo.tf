resource "google_artifact_registry_repository" "npm" {
  project       = local.project_id
  location      = local.region
  repository_id = "${local.resource_id}-npm"
  format        = "NPM"
  labels        = local.final_labels
}

# Grant "reader" role to the reader service account
resource "google_artifact_registry_repository_iam_member" "npm_reader" {
  project    = local.project_id
  location   = local.region
  repository = google_artifact_registry_repository.npm.repository_id
  role       = "roles/artifactregistry.reader"
  member     = "serviceAccount:${google_service_account.reader.email}"
}

# Grant "writer" role to the writer service account
resource "google_artifact_registry_repository_iam_member" "npm_writer" {
  project    = local.project_id
  location   = local.region
  repository = google_artifact_registry_repository.npm.repository_id
  role       = "roles/artifactregistry.writer"
  member     = "serviceAccount:${google_service_account.writer.email}"
}

# Grant "repoAdmin" role to the writer service account
resource "google_artifact_registry_repository_iam_member" "npm_admin" {
  project    = local.project_id
  location   = local.region
  repository = google_artifact_registry_repository.npm.repository_id
  role       = "roles/artifactregistry.repoAdmin"
  member     = "serviceAccount:${google_service_account.writer.email}"
}

#############################
# Outputs
#############################

output "npm_repo_name" {
  description = "The repository ID for the NPM repository"
  value       = google_artifact_registry_repository.npm.repository_id
}
