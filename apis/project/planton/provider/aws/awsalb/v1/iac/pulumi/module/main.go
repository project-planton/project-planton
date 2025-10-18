package module

import (
	"github.com/pkg/errors"
	awsalbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsalb/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the primary entry point for the aws_alb Pulumi module.
func Resources(ctx *pulumi.Context, stackInput *awsalbv1.AwsAlbStackInput) error {
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

	albResource, err := alb(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "failed to create aws_alb resource")
	}

	// If the user wants DNS, set up Route53 records
	if locals.AwsAlb.Spec.Dns.GetEnabled() {
		if err := dns(ctx, locals, provider, albResource); err != nil {
			return errors.Wrap(err, "failed to configure DNS")
		}
	}

	return nil
}
