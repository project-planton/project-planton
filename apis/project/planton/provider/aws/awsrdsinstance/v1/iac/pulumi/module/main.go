package module

import (
	"github.com/pkg/errors"
	awsrdsinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsrdsinstancev1.AwsRdsInstanceStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential

	//create aws provider using the credentials from the input (fallback to default when nil)
	var awsProvider *aws.Provider
	var err error
	if awsCredential == nil {
		awsProvider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
	} else {
		awsProvider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
	}
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	createdSecurityGroup, err := securityGroup(ctx, locals, awsProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create default security group")
	}

	// Create RDS Instance
	createRdsInstance, err := rdsInstance(ctx, locals, awsProvider, createdSecurityGroup)
	if err != nil {
		return errors.Wrap(err, "failed to create rds instance")
	}

	// Export Outputs
	ctx.Export(OpRdsInstanceEndpoint, createRdsInstance.Endpoint)
	ctx.Export(OpRdsInstanceId, createRdsInstance.ResourceId)
	ctx.Export(OpRdsInstanceArn, createRdsInstance.Arn)
	ctx.Export(OpRdsInstanceAddress, createRdsInstance.Address)
	ctx.Export(OpRdsSecurityGroup, createdSecurityGroup.Name)
	ctx.Export(OpRdsParameterGroup, createRdsInstance.ParameterGroupName)
	ctx.Export(OpRdsSubnetGroup, createRdsInstance.DbSubnetGroupName)
	ctx.Export(OpRdsOptionsGroup, createRdsInstance.OptionGroupName)
	if locals.AwsRdsInstance.Spec.Port > 0 {
		ctx.Export(OpRdsInstancePort, pulumi.Int(locals.AwsRdsInstance.Spec.Port))
	}
	return nil
}
