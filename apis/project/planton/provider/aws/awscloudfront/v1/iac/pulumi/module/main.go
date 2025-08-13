package module

import (
	"github.com/pkg/errors"
	awscloudfrontv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscloudfront/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, in *awscloudfrontv1.AwsCloudFrontStackInput) error {
	locals := initializeLocals(ctx, in)

	cred := in.ProviderCredential
	var provider *aws.Provider
	var err error
	if cred != nil {
		provider, err = aws.NewProvider(ctx, "aws-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(cred.AccessKeyId),
			SecretKey: pulumi.String(cred.SecretAccessKey),
			Region:    pulumi.String(cred.Region),
		})
		if err != nil {
			return errors.Wrap(err, "create provider")
		}
	}

	dist, err := createDistribution(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "create cloudfront distribution")
	}

	// Export outputs mapped to AwsCloudFrontStackOutputs
	ctx.Export(OpDistributionId, dist.ID())
	ctx.Export(OpDomainName, dist.DomainName)
	ctx.Export(OpHostedZoneId, pulumi.String("Z2FDTNDATAQYW2"))

	return nil
}
