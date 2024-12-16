package outputs

const (
	VpcId                             = "vpc_id"
	InternetGatewayId                 = "internet_gateway_id"
	PrivateSubnetsName                = "private_subnets.name"
	PrivateSubnetsId                  = "private_subnets.id"
	PrivateSubnetsCidr                = "private_subnets.cidr"
	PrivateSubnetsNatGatewayId        = "private_subnets.nat_gateway.id"
	PrivateSubnetsNatGatewayPrivateIp = "private_subnets.nat_gateway.private_ip"
	PrivateSubnetsNatGatewayPublicIp  = "private_subnets.nat_gateway.public_ip"

	PublicSubnetsName                = "public_subnets.name"
	PublicSubnetsId                  = "public_subnets.id"
	PublicSubnetsCidr                = "public_subnets.cidr"
	PublicSubnetsNatGatewayId        = "public_subnets.nat_gateway.id"
	PublicSubnetsNatGatewayPrivateIp = "public_subnets.nat_gateway.private_ip"
	PublicSubnetsNatGatewayPublicIp  = "public_subnets.nat_gateway.public_ip"
)
