package module

import (
	"github.com/pkg/errors"
	awskmskeyv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awskmskey/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awskmskeyv1.AwsKmsKeyStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsCredential := stackInput.ProviderCredential

	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.GetRegion()),
			Token:     pulumi.StringPtr(awsCredential.SessionToken),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create KMS key and optional alias
	result, err := kmsKey(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "failed to create kms key")
	}

	// Export outputs
	ctx.Export(OpKeyId, result.KeyId)
	ctx.Export(OpKeyArn, result.KeyArn)
	ctx.Export(OpAliasName, result.AliasName)
	ctx.Export(OpRotationEnabled, result.RotationEnabled)

	return nil
}
