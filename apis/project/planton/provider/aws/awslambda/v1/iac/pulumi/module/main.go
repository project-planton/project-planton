package module

import (
	"github.com/pkg/errors"
	awslambdav1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awslambda/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point that prepares locals, initialises the AWS provider,
// orchestrates the Lambda function creation, and exports outputs as defined in AwsLambdaStackOutputs.
func Resources(ctx *pulumi.Context, in *awslambdav1.AwsLambdaStackInput) error {
	locals := initializeLocals(ctx, in)

	// Initialize AWS provider with safe default when credentials are not provided
	var provider *aws.Provider
	var err error
	if in.ProviderCredential == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		region := in.ProviderCredential.Region
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(in.ProviderCredential.AccessKeyId),
			SecretKey: pulumi.String(in.ProviderCredential.SecretAccessKey),
			Region:    pulumi.String(region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create the Lambda function (and supporting log group)
	if _, err := lambdaFunction(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create lambda function")
	}

	return nil
}
