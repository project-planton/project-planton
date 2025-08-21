package module

import (
	"github.com/pkg/errors"
	awsrdsinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources orchestrates creation of an AWS RDS DB instance and exports outputs.
func Resources(ctx *pulumi.Context, in *awsrdsinstancev1.AwsRdsInstanceStackInput) error {
	locals := initializeLocals(ctx, in)

	// Initialize AWS provider (default when no credentials provided)
	var provider *aws.Provider
	var err error
	if in.ProviderCredential == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(in.ProviderCredential.AccessKeyId),
			SecretKey: pulumi.String(in.ProviderCredential.SecretAccessKey),
			Region:    pulumi.String(in.ProviderCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Subnet group (only when subnetIds provided and no name supplied)
	createdSubnetGroup, err := subnetGroup(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "subnet group")
	}

	// Create the RDS Instance
	instance, err := dbInstance(ctx, locals, provider, createdSubnetGroup)
	if err != nil {
		return errors.Wrap(err, "rds instance")
	}

	// Export outputs as defined in AwsRdsInstanceStackOutputs
	ctx.Export(OpRdsInstanceId, instance.ID())
	ctx.Export(OpRdsInstanceArn, instance.Arn)
	ctx.Export(OpRdsInstanceEndpoint, instance.Address)
	ctx.Export(OpRdsInstancePort, instance.Port)
	if createdSubnetGroup != nil {
		ctx.Export(OpRdsSubnetGroup, createdSubnetGroup.Name)
	} else if locals.AwsRdsInstance.Spec != nil && locals.AwsRdsInstance.Spec.DbSubnetGroupName != nil && locals.AwsRdsInstance.Spec.DbSubnetGroupName.GetValue() != "" {
		ctx.Export(OpRdsSubnetGroup, pulumi.String(locals.AwsRdsInstance.Spec.DbSubnetGroupName.GetValue()))
	}
	if locals.AwsRdsInstance.Spec != nil && locals.AwsRdsInstance.Spec.ParameterGroupName != "" {
		ctx.Export(OpRdsParameterGroup, pulumi.String(locals.AwsRdsInstance.Spec.ParameterGroupName))
	}
	if locals.AwsRdsInstance.Spec != nil && len(locals.AwsRdsInstance.Spec.SecurityGroupIds) > 0 {
		first := locals.AwsRdsInstance.Spec.SecurityGroupIds[0].GetValue()
		if first != "" {
			ctx.Export(OpRdsSecurityGroup, pulumi.String(first))
		}
	}
	return nil
}
