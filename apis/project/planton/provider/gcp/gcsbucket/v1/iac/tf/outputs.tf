output "bucket_id" {
  description = "The ID of the created GCS Bucket"
  value       = google_storage_bucket.main.id
}
