package module

import (
	"github.com/pkg/errors"
	awss3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awss3bucketv1.AwsS3BucketStackInput) error {
	locals := initializeLocals(ctx, stackInput)
	awsCredential := stackInput.ProviderCredential

	// Initialize AWS provider with safe defaults when credentials are not provided
	var nativeProvider *aws.Provider
	var err error
	if awsCredential == nil {
		var args aws.ProviderArgs
		if locals.AwsS3Bucket != nil && locals.AwsS3Bucket.Spec.AwsRegion != "" {
			args.Region = pulumi.String(locals.AwsS3Bucket.Spec.AwsRegion)
		}
		nativeProvider, err = aws.NewProvider(ctx, "native-provider", &args)
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		region := awsCredential.Region
		if region == "" && locals.AwsS3Bucket != nil && locals.AwsS3Bucket.Spec.AwsRegion != "" {
			region = locals.AwsS3Bucket.Spec.AwsRegion
		}
		nativeProvider, err = aws.NewProvider(ctx,
			"native-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create an S3 bucket
	createdBucket, err := s3.NewBucket(ctx, "bucket",
		&s3.BucketArgs{
			BucketName: pulumi.String(locals.AwsS3Bucket.Metadata.Name),
		}, pulumi.Provider(nativeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create S3 bucket")
	}

	// Export the bucket id
	ctx.Export(OpBucketId, createdBucket.ID())
	return nil
}
