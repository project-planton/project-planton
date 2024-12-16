package outputs

import (
	"fmt"
)

const (
	VPC_ID                                 = "vpc_id"
	INTERNET_GATEWAY_ID                    = "internet_gateway_id"
	PRIVATE_SUBNETS_NAME                   = "private_subnets.name"
	PRIVATE_SUBNETS_ID                     = "private_subnets.id"
	PRIVATE_SUBNETS_CIDR                   = "private_subnets.cidr"
	PRIVATE_SUBNETS_NAT_GATEWAY_ID         = "private_subnets.nat_gateway.id"
	PRIVATE_SUBNETS_NAT_GATEWAY_PRIVATE_IP = "private_subnets.nat_gateway.private_ip"
	PRIVATE_SUBNETS_NAT_GATEWAY_PUBLIC_IP  = "private_subnets.nat_gateway.public_ip"

	PUBLIC_SUBNETS_NAME                   = "public_subnets.name"
	PUBLIC_SUBNETS_ID                     = "public_subnets.id"
	PUBLIC_SUBNETS_CIDR                   = "public_subnets.cidr"
	PUBLIC_SUBNETS_NAT_GATEWAY_ID         = "public_subnets.nat_gateway.id"
	PUBLIC_SUBNETS_NAT_GATEWAY_PRIVATE_IP = "public_subnets.nat_gateway.private_ip"
	PUBLIC_SUBNETS_NAT_GATEWAY_PUBLIC_IP  = "public_subnets.nat_gateway.public_ip"
)

func SubnetIdOutputKey(subnetName string) string {
	return fmt.Sprintf("%s-id", subnetName)
}

func SubnetCidrOutputKey(subnetName string) string {
	return fmt.Sprintf("%s-cidr", subnetName)
}

func NatGatewayIdOutputKey(subnetName string) string {
	return fmt.Sprintf("%s-nat-gw-id", subnetName)
}

func NatGatewayPrivateIpOutputKey(subnetName string) string {
	return fmt.Sprintf("%s-nat-gw-private-ip", subnetName)
}

func NatGatewayPublicIpOutputKey(subnetName string) string {
	return fmt.Sprintf("%s-nat-gw-public-ip", subnetName)
}
