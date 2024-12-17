locals {
  azs     = var.spec.availability_zones
  az_count = length(local.azs)

  # Create a list of objects, each with a `key` and a `value`,
  # by nesting your loops inside a `flatten()`.
  public_subnets_list = flatten([
    for az_i, az in local.azs : [
      for i in range(var.spec.subnets_per_availability_zone) : {
        key = "public-${az}-${i}"
        value = {
          availability_zone = az
          cidr_block = cidrsubnet(
            var.spec.vpc_cidr,
            var.spec.subnet_size,
            (az_i * var.spec.subnets_per_availability_zone) + i
          )
        }
      }
    ]
  ])

  public_subnets = {
    for s in local.public_subnets_list : s.key => s.value
  }

  # Similarly for private_subnets:
  private_subnets_list = flatten([
    for az_i, az in local.azs : [
      for i in range(var.spec.subnets_per_availability_zone) : {
        key = "private-${az}-${i}"
        value = {
          availability_zone = az
          cidr_block = cidrsubnet(
            var.spec.vpc_cidr,
            var.spec.subnet_size,
            (az_i * var.spec.subnets_per_availability_zone) + i + (var.spec.subnets_per_availability_zone * local.az_count)
          )
        }
      }
    ]
  ])

  private_subnets = {
    for s in local.private_subnets_list : s.key => s.value
  }

  # NAT gateway subnets
  nat_gateway_subnets = {
    for az_i, az in local.azs :
    az => "public-${az}-0"
  }

  az_to_nat_gw = {
    for az, subnet_key in local.nat_gateway_subnets :
    az => (var.spec.is_nat_gateway_enabled ? aws_nat_gateway.this[az].id : null)
  }
}
