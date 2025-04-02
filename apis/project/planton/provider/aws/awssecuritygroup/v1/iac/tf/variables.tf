# variables.tf

# Name of the Security Group
variable "name" {
  type        = string
  description = "Name of the AWS Security Group"
}

# VPC ID where the Security Group will be created
variable "vpc_id" {
  type        = string
  description = "ID of the VPC in which to create the Security Group"
}

# Description of the Security Group (<= 255 chars)
variable "description" {
  type        = string
  description = "Short explanation of the Security Group's purpose"
}

# List of ingress rules
variable "ingress" {
  type = list(object({
    protocol         = string
    from_port        = number
    to_port          = number
    ipv4_cidrs = list(string)
    ipv6_cidrs = list(string)
    source_security_group_ids = list(string)
    destination_security_group_ids = list(string)
    self_reference   = bool
    rule_description = string
  }))
  default = []
  description = "List of inbound Security Group rules"
}

# List of egress rules
variable "egress" {
  type = list(object({
    protocol         = string
    from_port        = number
    to_port          = number
    ipv4_cidrs = list(string)
    ipv6_cidrs = list(string)
    source_security_group_ids = list(string)
    destination_security_group_ids = list(string)
    self_reference   = bool
    rule_description = string
  }))
  default = []
  description = "List of outbound Security Group rules"
}

# Optional tags to apply to the Security Group
variable "tags" {
  type = map(string)
  default = {}
  description = "Additional tags to apply to the Security Group"
}
