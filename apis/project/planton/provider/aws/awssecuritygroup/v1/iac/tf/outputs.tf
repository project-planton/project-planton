# outputs.tf

output "security_group_id" {
  description = "The ID of the newly created Security Group."
  value       = aws_security_group.this.id
}

output "vpc_id" {
  description = "The VPC ID in which this Security Group is created."
  value       = aws_security_group.this.vpc_id
}

# The below fields are placeholders to keep consistency
# with the Pulumi module’s output structure.

output "internet_gateway_id" {
  description = "Placeholder output for internet gateway ID."
  value       = ""
}

output "private_subnets" {
  description = "Placeholder output for private subnets."
  value = []
}

output "public_subnets" {
  description = "Placeholder output for public subnets."
  value = []
}
