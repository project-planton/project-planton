# main.tf

resource "aws_security_group" "this" {
  name        = var.name
  description = var.description
  vpc_id      = var.vpc_id

  # ingress rules
  # each item in var.ingress maps to one AWS ingress rule
  dynamic "ingress" {
    for_each = var.ingress
    content {
      description      = ingress.value.rule_description
      protocol         = ingress.value.protocol
      from_port        = ingress.value.from_port
      to_port          = ingress.value.to_port
      cidr_blocks      = ingress.value.ipv4_cidrs
      ipv6_cidr_blocks = ingress.value.ipv6_cidrs
      security_groups  = ingress.value.source_security_group_ids
      self             = ingress.value.self_reference
    }
  }

  # egress rules
  # each item in var.egress maps to one AWS egress rule
  dynamic "egress" {
    for_each = var.egress
    content {
      description      = egress.value.rule_description
      protocol         = egress.value.protocol
      from_port        = egress.value.from_port
      to_port          = egress.value.to_port
      cidr_blocks      = egress.value.ipv4_cidrs
      ipv6_cidr_blocks = egress.value.ipv6_cidrs
      security_groups  = egress.value.destination_security_group_ids
      self             = egress.value.self_reference
    }
  }

  tags = var.tags
}
