package module

import (
	"fmt"
	"github.com/pkg/errors"
	awssecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecretsmanager/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	PlaceholderSecretValue = "placeholder"
)

func Resources(ctx *pulumi.Context, stackInput *awssecretsmanagerv1.AwsSecretsManagerStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.ProviderCredential

	//create aws provider using the credentials from the input
	awsProvider, err := aws.NewProvider(ctx,
		"aws-classic",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	secretArnMap := map[string]string{}

	// For each secret in the input spec, create a secret in AWS Secrets Manager
	for _, secretName := range locals.AwsSecretsManager.Spec.SecretNames {
		if secretName == "" {
			continue
		}

		// Construct the secret ID to make it unique within the AWS account
		secretId := fmt.Sprintf("%s-%s", locals.AwsSecretsManager.Metadata.Id, secretName)

		// Create the secret resource
		createdSecret, err := secretsmanager.NewSecret(ctx,
			secretName,
			&secretsmanager.SecretArgs{
				Name: pulumi.String(secretId),
				Tags: pulumi.ToStringMap(locals.AwsTags),
			}, pulumi.Provider(awsProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create secret")
		}

		// Create a secret version with a placeholder value
		_, err = secretsmanager.NewSecretVersion(ctx, secretId, &secretsmanager.SecretVersionArgs{
			SecretId:     createdSecret.ID(),
			SecretString: pulumi.String(PlaceholderSecretValue),
		}, pulumi.Parent(createdSecret), pulumi.IgnoreChanges([]string{"secretString"})) // Ignore secret value changes to avoid diffs
		if err != nil {
			return errors.Wrap(err, "failed to create placeholder secret version")
		}

		var createdSecretArn string

		createdSecret.Arn.ApplyT(func(arn string) (string, error) {
			// Here arn is a real string value, available at runtime.
			createdSecretArn = arn
			fmt.Println("The resolved ARN is:", arn)
			return arn, nil
		})

		secretArnMap[secretName] = createdSecretArn
	}

	ctx.Export(OpSecretArnMap, pulumi.ToStringMap(secretArnMap))

	return nil
}
