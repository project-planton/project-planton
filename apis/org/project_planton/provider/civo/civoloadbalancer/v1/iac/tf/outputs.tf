# Outputs for Civo Load Balancer Terraform Module

output "load_balancer_id" {
  description = "The ID of the created Civo load balancer"
  value       = civo_loadbalancer.this.id
}

output "public_ip" {
  description = "The public IP address of the load balancer"
  value       = civo_loadbalancer.this.public_ip
}

output "dns_name" {
  description = "The DNS hostname for the load balancer (*.lb.civo.com)"
  value       = civo_loadbalancer.this.hostname
}

output "state" {
  description = "The current state of the load balancer"
  value       = civo_loadbalancer.this.state
}

output "algorithm" {
  description = "The load balancing algorithm in use"
  value       = civo_loadbalancer.this.algorithm
}

output "backends_count" {
  description = "The number of backend configurations"
  value       = length(civo_loadbalancer.this.backend)
}

