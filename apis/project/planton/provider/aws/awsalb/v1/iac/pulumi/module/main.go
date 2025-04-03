package module

import (
	"github.com/pkg/errors"
	awsalbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsalb/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the primary entry point for the aws_alb Pulumi module.
// It reads the AwsAlbStackInput, sets up AWS credentials if provided,
// and delegates to the alb() function to create the Application Load Balancer resource.
func Resources(ctx *pulumi.Context, stackInput *awsalbv1.AwsAlbStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsCredential := stackInput.ProviderCredential

	// If no credential is provided, use the default AWS provider
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx, "default-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		// Create a custom provider with explicit credentials
		provider, err = aws.NewProvider(ctx, "custom-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create the AWS ALB resource
	if err := alb(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create aws_alb resource")
	}

	return nil
}
