resource "google_artifact_registry_repository" "repo" {
  project       = local.project_id
  location      = local.region
  repository_id = "${local.resource_id}-docker"
  format        = local.repo_format
  labels        = local.final_labels
}

# Grant "reader" role to either `allUsers` (if external) or to the reader service account.
resource "google_artifact_registry_repository_iam_member" "docker_repo_public_reader" {
  count    = local.enable_public_access ? 1 : 0
  project  = local.project_id
  location = local.region
  repository = google_artifact_registry_repository.repo.repository_id
  role     = "roles/artifactregistry.reader"
  member   = "allUsers"
}

resource "google_artifact_registry_repository_iam_member" "docker_repo_reader" {
  count    = local.enable_public_access ? 0 : 1
  project  = local.project_id
  location = local.region
  repository = google_artifact_registry_repository.repo.repository_id
  role     = "roles/artifactregistry.reader"
  member   = "serviceAccount:${google_service_account.reader.email}"
}

# Grant "writer" role to the writer service account.
resource "google_artifact_registry_repository_iam_member" "docker_repo_writer" {
  project  = local.project_id
  location = local.region
  repository = google_artifact_registry_repository.repo.repository_id
  role     = "roles/artifactregistry.writer"
  member   = "serviceAccount:${google_service_account.writer.email}"
}

# Grant "repoAdmin" role to the writer service account.
resource "google_artifact_registry_repository_iam_member" "docker_repo_admin" {
  project  = local.project_id
  location = local.region
  repository = google_artifact_registry_repository.repo.repository_id
  role     = "roles/artifactregistry.repoAdmin"
  member   = "serviceAccount:${google_service_account.writer.email}"
}

#############################
# Outputs
#############################

output "repo_name" {
  description = "The repository ID for the repository"
  value       = google_artifact_registry_repository.repo.repository_id
}

# Example: us-west2-docker.pkg.dev
output "repo_hostname" {
  description = "The Repository hostname (region-docker.pkg.dev)"
  value       = "${google_artifact_registry_repository.repo.location}-docker.pkg.dev"
}

# Example: us-west2-docker.pkg.dev/<project>/<repo-id>
output "repo_url" {
  description = "Full Repository URL"
  value       = "${google_artifact_registry_repository.repo.location}-docker.pkg.dev/${google_artifact_registry_repository.repo.project}/${google_artifact_registry_repository.repo.repository_id}"
}
