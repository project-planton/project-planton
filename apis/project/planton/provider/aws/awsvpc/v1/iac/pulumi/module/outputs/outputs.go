package outputs

import (
	"fmt"
)

const (
	VpcId             = "vpc-id"
	InternetGatewayId = "internet-gateway-id"
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
