package module

import (
	"fmt"
	"github.com/pkg/errors"
	gcpsecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpsecretsmanager/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpsecretsmanager/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/secretmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	PlaceholderSecretValue = "placeholder"
)

func Resources(ctx *pulumi.Context, stackInput *gcpsecretsmanagerv1.GcpSecretsManagerStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create gcp provider using the credentials from the input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	secretIdMap := map[string]string{}

	//for each secret in the input spec, create a secret on gcp secrets-manager
	for _, secretName := range locals.GcpSecretsManager.Spec.SecretNames {
		if secretName == "" {
			continue
		}

		//construct the id of the secret to make it unique with in the google cloud project
		secretId := fmt.Sprintf("%s-%s", locals.GcpSecretsManager.Metadata.Id, secretName)

		//create the secret resource
		createdSecret, err := secretmanager.NewSecret(ctx, secretName, &secretmanager.SecretArgs{
			Labels:   pulumi.ToStringMap(locals.GcpLabels),
			Project:  pulumi.String(locals.GcpSecretsManager.Spec.ProjectId),
			SecretId: pulumi.String(secretId),
			Replication: secretmanager.SecretReplicationArgs{
				Auto: secretmanager.SecretReplicationAutoArgs{},
			},
		}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create secret")
		}

		//create secret-version with a placeholder value
		_, err = secretmanager.NewSecretVersion(ctx, secretId, &secretmanager.SecretVersionArgs{
			Enabled:    pulumi.Bool(true),
			Secret:     createdSecret.Name,
			SecretData: pulumi.String(PlaceholderSecretValue),
		},
			pulumi.Parent(createdSecret),
			pulumi.IgnoreChanges([]string{"secretData"}))
		if err != nil {
			return errors.Wrap(err, "failed to create placeholder secret version")
		}

		secretIdMap[secretName] = secretId
	}
	//export the id of the secret
	ctx.Export(outputs.SECRET_ID_MAP, pulumi.ToStringMap(secretIdMap))
	return nil
}
