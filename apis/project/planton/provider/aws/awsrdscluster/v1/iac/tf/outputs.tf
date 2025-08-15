output "cluster_identifier" {
  description = "The RDS cluster identifier."
  value       = aws_rds_cluster.this[0].cluster_identifier
}

output "master_endpoint" {
  description = "The endpoint of the primary writer instance."
  value       = aws_rds_cluster.this[0].endpoint
}

output "reader_endpoint" {
  description = "The endpoint for read-only connections to the cluster."
  value       = aws_rds_cluster.this[0].reader_endpoint
}

output "port" {
  description = "The port on which the cluster accepts connections."
  value       = var.spec.database_port > 0 ? var.spec.database_port : null
}
