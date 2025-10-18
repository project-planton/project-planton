package module

import (
	"github.com/pkg/errors"
	awsiamuserv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsiamuser/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsiamuserv1.AwsIamUserStackInput) error {
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

	// Create IAM user and related resources
	results, err := iamUser(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "failed to create iam user")
	}

	// Export outputs
	ctx.Export(OpUserArn, results.UserArn)
	ctx.Export(OpUserName, results.UserName)
	ctx.Export(OpUserId, results.UserId)
	ctx.Export(OpConsoleUrl, results.ConsoleUrl)
	ctx.Export(OpAccessKeyId, results.AccessKeyId)
	ctx.Export(OpSecretAccessKey, results.SecretAccessKey)

	return nil
}
