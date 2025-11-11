resource "google_storage_bucket" "main" {
  name                         = var.metadata.name
  project                      = var.spec.gcp_project_id
  location                     = var.spec.gcp_region
  labels                       = local.final_labels
  force_destroy                = true
  uniform_bucket_level_access  = !(var.spec.is_public)
}

resource "google_storage_bucket_acl" "public_bucket_acl" {
  count = var.spec.is_public ? 1 : 0

  bucket       = google_storage_bucket.main.name
  role_entity  = ["READER:allUsers"]

  depends_on = [
    google_storage_bucket.main
  ]
}

resource "google_storage_bucket_acl" "public_bucket_object_reader" {
  count = var.spec.is_public ? 1 : 0

  bucket       = google_storage_bucket.main.name
  role_entity  = ["READER:allUsers"]

  depends_on = [
    google_storage_bucket.main
  ]
}
