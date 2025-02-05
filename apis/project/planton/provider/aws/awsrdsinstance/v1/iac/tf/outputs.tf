output "rds_instance_endpoint" {
  description = "Endpoint of the RDS instance"
  value       = aws_db_instance.this.endpoint
}

output "rds_instance_id" {
  description = "ID of the RDS instance"
  value       = aws_db_instance.this.id
}

output "rds_instance_arn" {
  description = "ARN of the RDS instance"
  value       = aws_db_instance.this.arn
}

output "rds_instance_address" {
  description = "Address of the RDS instance"
  value       = aws_db_instance.this.address
}

output "rds_subnet_group" {
  description = "Name of the DB Subnet Group used by this RDS instance"
  value       = local.final_db_subnet_group_name
}

output "rds_security_group" {
  description = "Name of the default security group created for this RDS instance"
  value       = aws_security_group.default.name
}

output "rds_parameter_group" {
  description = "Name of the DB Parameter Group used by this RDS instance"
  value       = local.final_parameter_group_name
}

output "rds_options_group" {
  description = "Name of the DB Option Group used by this RDS instance"
  value       = local.final_option_group_name
}
