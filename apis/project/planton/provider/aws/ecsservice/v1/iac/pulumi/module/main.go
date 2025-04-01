package module

import (
	"github.com/pkg/errors"
	ecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the main entry point for the Pulumi module.
// It reads the EcsServiceStackInput, initializes locals, configures
// an AWS provider, then calls the service(...) function.
func Resources(ctx *pulumi.Context, stackInput *ecsservicev1.EcsServiceStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential
	var provider *aws.Provider
	var err error

	// If the user didn't provide AWS credentials, create a default provider.
	// Otherwise, inject custom credentials for the region, access key, etc.
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
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

	// Call the service(...) function to create the ECS Service resource.
	if err := service(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create ecs service resource")
	}

	return nil
}
