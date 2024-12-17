##############################
# one EIP per AZ for NAT gateway IP
##############################

resource "aws_eip" "nat" {
  for_each = var.spec.is_nat_gateway_enabled ? local.nat_gateway_subnets : {}
  domain = "vpc"

  tags = merge(
    var.metadata.tags,
    { "Name" = "${each.key}-nat-eip" }
  )
}

##############################
# NAT Gateway (one per AZ)
##############################

resource "aws_nat_gateway" "this" {
  for_each = var.spec.is_nat_gateway_enabled ? local.nat_gateway_subnets : {}

  # each.key is the AZ, each.value is the chosen public subnet key for that AZ
  subnet_id    = aws_subnet.public[each.value].id
  allocation_id = aws_eip.nat[each.key].id

  tags = merge(
    var.metadata.tags,
    { "Name" = "${each.key}-nat-gw" }
  )

  depends_on = [aws_subnet.public]
}
