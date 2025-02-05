###############################################################################
# Create a default Security Group for the RDS instance
###############################################################################
resource "aws_security_group" "default" {
  name        = local.resource_id
  description = "Allow inbound traffic from the security groups"
  vpc_id      = var.spec.vpc_id
  tags        = local.final_labels
}

###############################################################################
# Ingress rule allowing traffic from user-supplied Security Groups
###############################################################################
resource "aws_security_group_rule" "ingress_sg" {
  for_each = toset(var.spec.security_group_ids)

  description              = "Allow inbound traffic from existing Security Groups"
  type                     = "ingress"
  from_port                = var.spec.port
  to_port                  = var.spec.port
  protocol                 = "tcp"
  source_security_group_id = each.value
  security_group_id        = aws_security_group.default.id
}

###############################################################################
# Ingress rule allowing traffic from user-supplied CIDR blocks
###############################################################################
resource "aws_security_group_rule" "ingress_cidr" {
  count = length(var.spec.allowed_cidr_blocks) > 0 ? 1 : 0

  description       = "Allow inbound traffic from CIDR blocks"
  type              = "ingress"
  from_port         = var.spec.port
  to_port           = var.spec.port
  protocol          = "tcp"
  cidr_blocks       = var.spec.allowed_cidr_blocks
  security_group_id = aws_security_group.default.id
}

###############################################################################
# Egress rule allowing all outbound traffic
###############################################################################
resource "aws_security_group_rule" "egress_all" {
  description       = "Allow all egress traffic"
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks = ["0.0.0.0/0"]
  security_group_id = aws_security_group.default.id
}
