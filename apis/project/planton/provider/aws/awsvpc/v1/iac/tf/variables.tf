variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "spec"
  type = object({

    # The CIDR (Classless Inter-Domain Routing) block for the VPC.
    # This defines the IP address range for the VPC.
    # Example: "10.0.0.0/16" allows IP addresses from 10.0.0.0 to 10.0.255.255.
    vpc_cidr = string

    # The list of availability zones where the VPC will be spanned.
    # AWS regions are divided into multiple availability zones (AZs) for high availability.
    # Example: ["us-west-2a", "us-west-2b"] indicates that resources will be spread across these two AZs.
    availability_zones = list(string)

    # The number of subnets to be created in each availability zone.
    # Subnets are segments of the VPC's IP address range where you can place groups of isolated resources.
    subnets_per_availability_zone = number

    # The number of hosts (IP addresses) in each subnet.
    # This determines the size of each subnet's CIDR block.
    subnet_size = number

    # Toggle to enable or disable a NAT (Network Address Translation) gateway for private subnets created in the VPC.
    # A NAT gateway allows instances in a private subnet to connect to the internet or other AWS services, but prevents
    # the internet from initiating a connection with those instances.
    is_nat_gateway_enabled = optional(bool, false)

    # Toggle to enable or disable DNS hostnames in the VPC.
    # When enabled, instances with public IP addresses receive corresponding public DNS hostnames.
    # See AWS documentation: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-hostnames
    is_dns_hostnames_enabled = optional(bool, false)

    # Toggle to enable or disable DNS resolution in the VPC through the Amazon-provided DNS server.
    # When enabled, the Amazon DNS server resolves DNS hostnames for your instances.
    is_dns_support_enabled = optional(bool, false)
  })
}
