package module

import (
	"github.com/pkg/errors"
	awsec2instancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsec2instance/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the entry‑point invoked by ProjectPlanton’s CLI.
// It wires provider credentials, initialises locals, and delegates
// to ec2Instance(...) to create the EC2 VM.
func Resources(ctx *pulumi.Context, stackInput *awsec2instancev1.AwsEc2InstanceStackInput) error {
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
			Token:     pulumi.StringPtr(awsCredential.SessionToken),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	if err := ec2Instance(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "create aws ec2 instance resource")
	}

	return nil
}
