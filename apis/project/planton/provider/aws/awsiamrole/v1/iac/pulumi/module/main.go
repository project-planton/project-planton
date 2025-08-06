package module

import (
	"github.com/pkg/errors"
	iamrolev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsiamrole/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *iamrolev1.AwsIamRoleStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential

	var provider *aws.Provider
	var err error

	if awsCredential == nil {
		// default AWS provider
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		// custom credentials
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// create the IAM role
	if err := iamRole(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create iam role resource")
	}

	return nil
}
