output "cluster_endpoint" {
  description = "The endpoint of the created GKE cluster"
  value       = google_container_cluster.gcp_gke_cluster.endpoint
}

output "cluster_ca_data" {
  description = "Base64-encoded public CA certificate of the cluster"
  value       = google_container_cluster.gcp_gke_cluster.master_auth[0].cluster_ca_certificate
}

output "external_nat_ip" {
  description = "External IP address for the GKE router NAT"
  value       = google_compute_address.gke_router_nat_ip.address
}

output "network_self_link" {
  description = "Self-link of the created VPC network"
  value       = google_compute_network.gke_network.self_link
}

output "sub_network_self_link" {
  description = "Self-link of the created subnetwork"
  value       = google_compute_subnetwork.gke_subnetwork.self_link
}

output "gke_webhooks_firewall_self_link" {
  description = "Self-link of the firewall rule for GKE webhooks"
  value       = google_compute_firewall.gke_webhook_firewall.self_link
}

output "router_self_link" {
  description = "Self-link of the created router"
  value       = google_compute_router.gke_router.self_link
}

output "router_nat_name" {
  description = "Name of the router NAT"
  value       = google_compute_router_nat.gke_nat.name
}

output "workload_deployer_gsa_email" {
  description = "Email of the Workload Deployer Google Service Account"
  value       = google_service_account.workload_deployer_sa.email
}

output "workload_deployer_gsa_key_base64" {
  description = "Base64-encoded private key of the Workload Deployer SA"
  value       = google_service_account_key.workload_deployer_sa_key.private_key
  sensitive   = true
}
