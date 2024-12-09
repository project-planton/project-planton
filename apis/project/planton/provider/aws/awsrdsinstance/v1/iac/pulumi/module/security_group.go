package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func securityGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*ec2.SecurityGroup, error) {

	defaultSecurityGroup, err := ec2.NewSecurityGroup(ctx, "default", &ec2.SecurityGroupArgs{
		Name:        pulumi.String(locals.AwsRdsInstance.Metadata.Id),
		Description: pulumi.String("Allow inbound traffic from the security groups"),
		VpcId:       pulumi.String(locals.AwsRdsInstance.Spec.VpcId),
		Tags:        pulumi.ToStringMap(locals.Labels),
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create default security group")
	}

	for _, securityGroupId := range locals.AwsRdsInstance.Spec.SecurityGroupIds {
		_, err := ec2.NewSecurityGroupRule(ctx, "ingress security groups", &ec2.SecurityGroupRuleArgs{
			Description:           pulumi.String("Allow inbound traffic from existing Security Groups"),
			Type:                  pulumi.String("ingress"),
			FromPort:              pulumi.Int(locals.AwsRdsInstance.Spec.Port),
			ToPort:                pulumi.Int(locals.AwsRdsInstance.Spec.Port),
			Protocol:              pulumi.String("tcp"),
			SourceSecurityGroupId: pulumi.String(securityGroupId),
			SecurityGroupId:       defaultSecurityGroup.ID(),
		}, pulumi.Provider(awsProvider), pulumi.Parent(defaultSecurityGroup))
		if err != nil {
			return nil, errors.Wrap(err, "failed to create ingress security group rules")
		}
	}

	if len(locals.AwsRdsInstance.Spec.AllowedCidrBlocks) > 0 {
		_, err := ec2.NewSecurityGroupRule(ctx, "ingress cidr blocks", &ec2.SecurityGroupRuleArgs{
			Description:     pulumi.String("Allow inbound traffic from CIDR blocks"),
			Type:            pulumi.String("ingress"),
			FromPort:        pulumi.Int(locals.AwsRdsInstance.Spec.Port),
			ToPort:          pulumi.Int(locals.AwsRdsInstance.Spec.Port),
			Protocol:        pulumi.String("tcp"),
			CidrBlocks:      pulumi.ToStringArray(locals.AwsRdsInstance.Spec.AllowedCidrBlocks),
			SecurityGroupId: defaultSecurityGroup.ID(),
		}, pulumi.Provider(awsProvider), pulumi.Parent(defaultSecurityGroup))
		if err != nil {
			return nil, errors.Wrap(err, "failed to create ingress cidr blocks security group rules")
		}
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "egress security group rule", &ec2.SecurityGroupRuleArgs{
		Description:     pulumi.String("Allow all egress traffic"),
		Type:            pulumi.String("egress"),
		FromPort:        pulumi.Int(0),
		ToPort:          pulumi.Int(0),
		Protocol:        pulumi.String("-1"),                            // All protocols
		CidrBlocks:      pulumi.StringArray{pulumi.String("0.0.0.0/0")}, // Allow all egress
		SecurityGroupId: defaultSecurityGroup.ID(),
	}, pulumi.Provider(awsProvider), pulumi.Parent(defaultSecurityGroup))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create egress security group rule")
	}

	return defaultSecurityGroup, nil
}
