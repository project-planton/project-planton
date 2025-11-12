package module

import (
	"github.com/pkg/errors"
	awss3bucketv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/aws/awss3bucket/v1"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awss3bucketv1.AwsS3BucketStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsProviderConfig := stackInput.ProviderConfig

	if awsProviderConfig == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsProviderConfig.AccessKeyId),
			SecretKey: pulumi.String(awsProviderConfig.SecretAccessKey),
			Region:    pulumi.String(awsProviderConfig.GetRegion()),
			Token:     pulumi.StringPtr(awsProviderConfig.SessionToken),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create an S3 bucket
	createdBucket, err := s3.NewBucket(ctx, "bucket",
		&s3.BucketArgs{
			BucketName: pulumi.String(locals.AwsS3Bucket.Metadata.Name),
		}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create S3 bucket")
	}

	// Export the bucket id
	ctx.Export(OpBucketId, createdBucket.ID())
	return nil
}
