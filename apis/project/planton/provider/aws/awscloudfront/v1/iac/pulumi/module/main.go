package module

import (
	"github.com/pkg/errors"
	awscloudfrontv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscloudfront/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the primary entry point for the aws_cloudfront Pulumi module.
func Resources(ctx *pulumi.Context, stackInput *awscloudfrontv1.AwsCloudFrontStackInput) error {
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
			Region:    pulumi.String(awsCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create CloudFront distribution
	dist, err := distribution(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "cloudfront distribution")
	}

	// Optional DNS alias records
	if locals.Spec.GetDns().GetEnabled() {
		if err := dns(ctx, locals, provider, dist); err != nil {
			return errors.Wrap(err, "dns")
		}
	}

	// Export outputs
	ctx.Export(OpDistributionId, dist.Distribution.ID())
	ctx.Export(OpDomainName, dist.DomainName)
	ctx.Export(OpHostedZoneId, dist.HostedZoneId)

	return nil
}
