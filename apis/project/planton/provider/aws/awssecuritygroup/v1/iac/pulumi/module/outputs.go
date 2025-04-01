package module

// The following constants define output keys for the aws_security_group module.
// They reflect the fields in AwsSecurityGroupStackOutputs.
const (
	OpSecurityGroupVpcId             = "vpc_id"
	OpSecurityGroupInternetGatewayId = "internet_gateway_id"
	OpSecurityGroupPrivateSubnets    = "private_subnets"
	OpSecurityGroupPublicSubnets     = "public_subnets"
	OpSecurityGroupId                = "security_group_id"
)
