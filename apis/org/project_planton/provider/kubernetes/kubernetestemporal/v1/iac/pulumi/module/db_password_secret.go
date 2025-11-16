package module

import (
	"encoding/base64"

	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// dbPasswordSecret creates a Kubernetes Secret containing the external DB
// password when, and only when, external_database is provided by the user.
func dbPasswordSecret(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	if locals.KubernetesTemporal.Spec.Database.ExternalDatabase == nil {
		// No external DB or no password: nothing to create.
		return nil
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(locals.KubernetesTemporal.Spec.Database.ExternalDatabase.Password))

	_, err := kubernetescorev1.NewSecret(ctx,
		vars.DatabasePasswordSecretName,
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(vars.DatabasePasswordSecretName),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Data: pulumi.StringMap{
				vars.DatabasePasswordSecretKey: pulumi.String(encoded),
			},
			Type: pulumi.String("Opaque"),
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create database password secret")
	}

	return nil
}
