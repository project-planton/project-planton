resource "aws_security_group" "default" {
  name        = local.resource_id
  description = "Allow inbound traffic from the security groups"
  vpc_id      = var.spec.vpc_id
  tags        = local.final_labels
}

# Ingress from existing Security Groups
resource "aws_security_group_rule" "ingress_security_groups" {
  for_each = toset(var.spec.security_group_ids)
  description           = "Allow inbound traffic from existing Security Groups"
  type                  = "ingress"
  from_port             = var.spec.database_port
  to_port               = var.spec.database_port
  protocol              = "tcp"
  source_security_group_id = each.value
  security_group_id     = aws_security_group.default.id
  depends_on            = [aws_security_group.default]
}

# Ingress from CIDR blocks
resource "aws_security_group_rule" "ingress_cidr_blocks" {
  count = length(var.spec.allowed_cidr_blocks) > 0 ? 1 : 0
  description       = "Allow inbound traffic from CIDR blocks"
  type              = "ingress"
  from_port         = var.spec.database_port
  to_port           = var.spec.database_port
  protocol          = "tcp"
  cidr_blocks       = var.spec.allowed_cidr_blocks
  security_group_id = aws_security_group.default.id
  depends_on        = [aws_security_group.default]
}

# Egress rule (allow all)
resource "aws_security_group_rule" "egress_rule" {
  description       = "Allow all egress traffic"
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.default.id
  depends_on        = [aws_security_group.default]
}
