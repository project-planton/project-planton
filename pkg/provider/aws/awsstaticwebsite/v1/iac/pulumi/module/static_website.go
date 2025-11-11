package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticWebsiteResult struct {
	BucketId pulumi.IDOutput
}

func staticWebsite(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*StaticWebsiteResult, error) {
	// Minimal viable implementation: create content S3 bucket if target is provided.
	if locals.AwsStaticWebsite == nil {
		return nil, errors.New("target AwsStaticWebsite is nil")
	}

	bucket, err := s3.NewBucket(ctx, "content-bucket", &s3.BucketArgs{
		BucketName: pulumi.String(locals.AwsStaticWebsite.Metadata.Name),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, errors.Wrap(err, "create S3 bucket")
	}

	// TODO: implement website configuration, CloudFront + Route53 conditionally per spec

	return &StaticWebsiteResult{
		BucketId: bucket.ID(),
	}, nil
}
