package module

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/jenkinskubernetes/v1/iac/pulumi/module/outputs"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func adminCredentials(ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) (*kubernetescorev1.Secret, error) {

	createdRandomPassword, err := random.NewRandomPassword(ctx,
		"admin-password",
		&random.RandomPasswordArgs{
			Length:     pulumi.Int(12),
			Special:    pulumi.Bool(true),
			Numeric:    pulumi.Bool(true),
			Upper:      pulumi.Bool(true),
			Lower:      pulumi.Bool(true),
			MinSpecial: pulumi.Int(3),
			MinNumeric: pulumi.Int(2),
			MinUpper:   pulumi.Int(2),
			MinLower:   pulumi.Int(2),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get random password value")
	}

	// Encode the password in Base64
	base64Password := createdRandomPassword.Result.ApplyT(func(p string) (string, error) {
		return base64.StdEncoding.EncodeToString([]byte(p)), nil
	}).(pulumi.StringOutput)

	// Create or update the secret
	createdAdminPasswordSecret, err := kubernetescorev1.NewSecret(ctx,
		locals.JenkinsKubernetes.Metadata.Name,
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.JenkinsKubernetes.Metadata.Name),
				Namespace: pulumi.String(locals.JenkinsKubernetes.Metadata.Id),
			},
			Data: pulumi.StringMap{
				vars.JenkinsAdminPasswordSecretKey: base64Password,
			},
		}, pulumi.Parent(createdNamespace))

	if err != nil {
		return nil, errors.Wrap(err, "failed to create jenkins admin-password secret")
	}

	//export admin credentials to outputs
	ctx.Export(outputs.AdminUsername, pulumi.String(vars.JenkinsAdminUsername))
	ctx.Export(outputs.AdminPasswordSecretName, createdAdminPasswordSecret.Metadata.Name())
	ctx.Export(outputs.AdminPasswordSecretKey, pulumi.String(vars.JenkinsAdminPasswordSecretKey))

	return createdAdminPasswordSecret, nil
}
