resource "google_artifact_registry_repository" "maven" {
  project       = local.project_id
  location      = local.region
  repository_id = "${local.resource_id}-maven"
  format        = "MAVEN"
  labels        = local.final_labels
}

# Grant "reader" role for the reader service account
resource "google_artifact_registry_repository_iam_member" "maven_reader" {
  project    = local.project_id
  location   = local.region
  repository = google_artifact_registry_repository.maven.repository_id
  role       = "roles/artifactregistry.reader"
  member     = "serviceAccount:${google_service_account.reader.email}"
}

# Grant "writer" role for the writer service account
resource "google_artifact_registry_repository_iam_member" "maven_writer" {
  project    = local.project_id
  location   = local.region
  repository = google_artifact_registry_repository.maven.repository_id
  role       = "roles/artifactregistry.writer"
  member     = "serviceAccount:${google_service_account.writer.email}"
}

# Grant "repoAdmin" role for the writer service account
resource "google_artifact_registry_repository_iam_member" "maven_admin" {
  project    = local.project_id
  location   = local.region
  repository = google_artifact_registry_repository.maven.repository_id
  role       = "roles/artifactregistry.repoAdmin"
  member     = "serviceAccount:${google_service_account.writer.email}"
}

#############################
# Outputs
#############################

output "maven_repo_name" {
  description = "The repository ID for the Maven repository"
  value       = google_artifact_registry_repository.maven.repository_id
}

output "maven_repo_url" {
  description = "A reference that can be used to address the Maven repository"
  value       = google_artifact_registry_repository.maven.id
}
