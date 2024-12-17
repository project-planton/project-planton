##############################
# VPC Creation
##############################

resource "aws_vpc" "this" {
  cidr_block           = var.spec.vpc_cidr
  enable_dns_support   = var.spec.is_dns_support_enabled
  enable_dns_hostnames = var.spec.is_dns_hostnames_enabled

  tags = merge(
    var.metadata.tags,
    { "Name" = var.metadata.name }
  )
}

##############################
# Internet Gateway
##############################

resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id

  tags = merge(
    var.metadata.tags,
    { "Name" = "${var.metadata.name}-igw" }
  )
}
