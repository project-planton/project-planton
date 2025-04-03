package module

import (
	"github.com/pkg/errors"
	awscertv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscertmanagercert/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the main entry point for the aws_cert_manager_cert Pulumi module.
// It prepares context, configures the AWS provider, and calls certManagerCert().
func Resources(ctx *pulumi.Context, stackInput *awscertv1.AwsCertManagerCertStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential
	var provider *aws.Provider
	var err error

	// Use default provider if credentials are not explicitly provided.
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx, "default", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "custom", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Call the core logic for ACM certificate + DNS validation setup.
	if err := certManagerCert(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create aws cert manager cert resource")
	}

	return nil
}
