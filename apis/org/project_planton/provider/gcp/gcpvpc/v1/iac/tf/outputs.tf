output "network_self_link" {
  description = "The full self-link URL of the created VPC network (useful for connecting subnets or other resources)"
  value       = google_compute_network.vpc.self_link
}

