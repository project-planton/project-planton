package module

import (
	"github.com/pkg/errors"
	awsrdsinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1/iac/pulumi/module/outputs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsrdsinstancev1.AwsRdsInstanceStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.AwsCredential

	//create aws provider using the credentials from the input
	awsProvider, err := aws.NewProvider(ctx,
		"classic-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
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
	ctx.Export(outputs.RDS_INSTANCE_ENDPOINT, createRdsInstance.Endpoint)
	ctx.Export(outputs.RDS_INSTANCE_ID, createRdsInstance.ResourceId)
	ctx.Export(outputs.RDS_INSTANCE_ARN, createRdsInstance.Arn)
	ctx.Export(outputs.RDS_INSTANCE_ADDRESS, createRdsInstance.Address)
	ctx.Export(outputs.RDS_SECURITY_GROUP, createdSecurityGroup.Name)
	ctx.Export(outputs.RDS_PARAMETER_GROUP, createRdsInstance.ParameterGroupName)
	ctx.Export(outputs.RDS_SUBNET_GROUP, createRdsInstance.DbSubnetGroupName)
	ctx.Export(outputs.RDS_OPTIONS_GROUP, createRdsInstance.OptionGroupName)
	return nil
}
