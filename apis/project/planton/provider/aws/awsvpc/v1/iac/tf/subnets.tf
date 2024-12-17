##############################
# Public Subnets
##############################

resource "aws_subnet" "public" {
  for_each                = local.public_subnets
  vpc_id                  = aws_vpc.this.id
  cidr_block              = each.value.cidr_block
  availability_zone       = each.value.availability_zone
  map_public_ip_on_launch = true

  tags = merge(
    var.metadata.tags,
    { "Name" = each.key }
  )
}

##############################
# Public Route Table
##############################

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.this.id

  tags = merge(
    var.metadata.tags,
    { "Name" = "${var.metadata.name}-public-RT" }
  )

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }
}

##############################
# Public Route Table Associations
##############################

resource "aws_route_table_association" "public" {
  for_each      = aws_subnet.public
  route_table_id = aws_route_table.public.id
  subnet_id      = each.value.id
}

##############################
# Private Subnets
##############################

resource "aws_subnet" "private" {
  for_each          = local.private_subnets
  vpc_id            = aws_vpc.this.id
  cidr_block        = each.value.cidr_block
  availability_zone = each.value.availability_zone

  tags = merge(
    var.metadata.tags,
    { "Name" = each.key }
  )
}

#################################################
# Private Route Tables with route to NAT gateway
#################################################

# For each private subnet, we need a route table that routes to the NAT gateway in its AZ.
# Each private subnet is defined as:
# aws_subnet.private[<key>] = {
#   cidr_block        = ...
#   availability_zone = ...
# }

resource "aws_route_table" "private" {
  for_each = aws_subnet.private

  vpc_id = aws_vpc.this.id
  tags = merge(
    var.metadata.tags,
    { "Name" = "${each.key}-RT" }
  )

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = lookup(local.az_to_nat_gw, each.value.availability_zone, null)
  }
}

##############################
# Private Route Table Associations
##############################

resource "aws_route_table_association" "private" {
  for_each = aws_subnet.private

  route_table_id = aws_route_table.private[each.key].id
  subnet_id      = each.value.id
}
