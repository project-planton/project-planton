locals {
  sorted_azs = sort(var.spec.availability_zones)

  az_count = length(local.sorted_azs)

  total_public_subnets  = var.spec.subnets_per_availability_zone * local.az_count
  total_private_subnets = var.spec.subnets_per_availability_zone * local.az_count

  # Public subnet mappings
  public_subnets = {
    for idx in range(local.total_public_subnets) :
    "public-${local.sorted_azs[idx % local.az_count]}-${idx}" => {
      availability_zone = local.sorted_azs[idx % local.az_count]
      cidr_block        = cidrsubnet(var.spec.vpc_cidr, var.spec.subnet_size, idx)
    }
  }

  # Private subnet mappings (start indexing after public subnets to avoid overlap)
  private_subnets = {
    for idx in range(local.total_private_subnets) :
    "private-${local.sorted_azs[idx % local.az_count]}-${idx}" => {
      availability_zone = local.sorted_azs[idx % local.az_count]
      cidr_block        = cidrsubnet(var.spec.vpc_cidr, var.spec.subnet_size, idx + local.total_public_subnets)
    }
  }

  # Now, we want to pick one public subnet per AZ for NAT gateways.
  # Step 1: Create a list of {az, key} pairs from public_subnets
  public_subnets_list = [
    for k, v in local.public_subnets : {
      key = k
      az  = v.availability_zone
    }
  ]

  # Group public subnets by AZ
  # Result: { "us-west-2a" = ["public-us-west-2a-0", "public-us-west-2a-1", ...], "us-west-2b" = [...] }
  public_subnets_by_az = {
    for az in local.sorted_azs :
    az => [
      for entry in local.public_subnets_list : entry.key
      if entry.az == az
    ]
  }

  # For each AZ, pick the first public subnet (index 0) to host the NAT gateway
  # This ensures exactly one NAT gateway per AZ.
  nat_gateway_subnets   = {
    for az in local.sorted_azs :
    az => local.public_subnets_by_az[az][0]
  }

  # Build a map from AZ to NAT gateway ID.
  az_to_nat_gw = {
    for az, subnet_key in local.nat_gateway_subnets :
    az => (var.spec.is_nat_gateway_enabled ? aws_nat_gateway.this[az].id : null)
  }

}