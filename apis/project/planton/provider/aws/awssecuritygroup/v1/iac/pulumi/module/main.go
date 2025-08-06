package module

import (
	"github.com/pkg/errors"
	awssecuritygroupv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecuritygroup/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the primary entry point for the aws_security_group Pulumi module.
// It reads the AwsSecurityGroupStackInput, sets up AWS credentials if provided,
// and delegates to the securityGroup() function to create the resource.
func Resources(ctx *pulumi.Context, stackInput *awssecuritygroupv1.AwsSecurityGroupStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsCredential := stackInput.ProviderCredential

	// If no credential is provided, use the default AWS provider
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		// Create a custom provider with explicit credentials
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create the AWS Security Group resource
	if err := securityGroup(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create aws_security_group resource")
	}

	return nil
}
