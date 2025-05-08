package module

import (
	"github.com/pkg/errors"
	awss3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awss3bucketv1.AwsS3BucketStackInput) error {
	awsCredential := stackInput.ProviderCredential

	//create aws provider using the credentials from the input
	nativeProvider, err := aws.NewProvider(ctx,
		"native-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	// Create an S3 bucket
	createdBucket, err := s3.NewBucket(ctx, "my-bucket",
		&s3.BucketArgs{
			BucketName: pulumi.String(stackInput.Target.Metadata.Name),
		}, pulumi.Provider(nativeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create S3 bucket")
	}

	// Export the bucket id
	ctx.Export(OpBucketId, createdBucket.ID())
	return nil
}
