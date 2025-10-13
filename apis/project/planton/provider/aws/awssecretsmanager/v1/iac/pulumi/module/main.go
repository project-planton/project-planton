package module

import (
	"fmt"

	"github.com/pkg/errors"
	awssecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecretsmanager/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	PlaceholderSecretValue = "placeholder"
)

func Resources(ctx *pulumi.Context, stackInput *awssecretsmanagerv1.AwsSecretsManagerStackInput) error {
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

	secretArnMap := pulumi.StringMap{}

	// For each secret in the input spec, create a secret in AWS Secrets Manager
	for _, secretName := range locals.AwsSecretsManager.Spec.SecretNames {
		if secretName == "" {
			continue
		}

		// Construct the secret ID to make it unique within the AWS account
		secretId := fmt.Sprintf("%s-%s", locals.AwsSecretsManager.Metadata.Id, secretName)

		createdSecret, err := createSecret(ctx, locals, provider, secretName, secretId)
		if err != nil {
			return errors.Wrapf(err, "secret %s", secretName)
		}

		secretArnMap[secretName] = createdSecret.Arn
	}

	ctx.Export(OpSecretArnMap, secretArnMap)

	return nil
}
