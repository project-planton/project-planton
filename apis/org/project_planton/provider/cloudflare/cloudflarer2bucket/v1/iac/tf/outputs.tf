output "bucket_name" {
  description = "The name of the R2 bucket"
  value       = cloudflare_r2_bucket.main.name
}

output "bucket_id" {
  description = "The ID of the R2 bucket"
  value       = cloudflare_r2_bucket.main.id
}

output "account_id" {
  description = "The Cloudflare account ID"
  value       = cloudflare_r2_bucket.main.account_id
}

output "location" {
  description = "The location hint for the R2 bucket"
  value       = cloudflare_r2_bucket.main.location
}

