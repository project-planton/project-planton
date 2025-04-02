package module

import (
	"github.com/pkg/errors"
	awsecrrepov1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecrrepo/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the entry point for the aws_ecr_repo Pulumi module.
// It initializes locals, configures a provider (default or custom), then calls ecrRepo.
func Resources(ctx *pulumi.Context, stackInput *awsecrrepov1.AwsEcrRepoStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsCredential := stackInput.ProviderCredential

	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx, "default-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "custom-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	if err := ecrRepo(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create aws_ecr_repo resource")
	}

	return nil
}
