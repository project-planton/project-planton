package module

import (
	"github.com/pkg/errors"
	awsecsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecscluster/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsecsclusterv1.AwsEcsClusterStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential

	var provider *aws.Provider
	var err error

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

	if err := cluster(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create aws ecs cluster resource")
	}

	return nil
}
