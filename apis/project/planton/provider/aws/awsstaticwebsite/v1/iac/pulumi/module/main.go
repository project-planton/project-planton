package module

import (
	"github.com/pkg/errors"
	awsstaticwebsitev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsstaticwebsite/v1"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsstaticwebsitev1.AwsStaticWebsiteStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential

	// initialize aws-native provider (fallback to default when credentials are not provided)
	var provider *aws.Provider
	var err error
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx, "native-provider", &aws.ProviderArgs{})
	} else {
		provider, err = aws.NewProvider(ctx,
			"native-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
	}
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	created, err := staticWebsite(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "failed to create static website")
	}

	ctx.Export(OpBucketId, created.BucketId)

	return nil
}
