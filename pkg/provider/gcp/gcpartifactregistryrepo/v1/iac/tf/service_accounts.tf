resource "random_string" "service_account_suffix" {
  length  = 6
  lower   = true
  upper   = false
  numeric = true
  special = false
}

# Reader service account
resource "google_service_account" "reader" {
  project      = local.project_id
  account_id   = "${var.metadata.name}-${random_string.service_account_suffix.result}-ro"
  display_name = "${var.metadata.name}-${random_string.service_account_suffix.result}-ro"
}

resource "google_service_account_key" "reader_key" {
  service_account_id = google_service_account.reader.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}

# Writer service account
resource "google_service_account" "writer" {
  project      = local.project_id
  account_id   = "${var.metadata.name}-${random_string.service_account_suffix.result}-rw"
  display_name = "${var.metadata.name}-${random_string.service_account_suffix.result}-rw"
}

resource "google_service_account_key" "writer_key" {
  service_account_id = google_service_account.writer.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}

#############################
# Outputs
#############################

output "reader_service_account_email" {
  description = "Email address of the reader service account"
  value       = google_service_account.reader.email
}

output "reader_service_account_key_base64" {
  description = "Base64-encoded private key of the reader service account"
  value       = google_service_account_key.reader_key.private_key
  sensitive   = true
}

output "writer_service_account_email" {
  description = "Email address of the writer service account"
  value       = google_service_account.writer.email
}

output "writer_service_account_key_base64" {
  description = "Base64-encoded private key of the writer service account"
  value       = google_service_account_key.writer_key.private_key
  sensitive   = true
}
