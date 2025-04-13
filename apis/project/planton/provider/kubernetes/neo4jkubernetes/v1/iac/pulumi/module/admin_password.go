// File: admin_password.go
package module

import (
	"encoding/base64"

	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// adminPassword handles generating a random password for Neo4j and storing it in a Kubernetes secret.
func adminPassword(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
) error {
	// Create a random password for the Neo4j admin user.
	createdRandomPassword, err := random.NewRandomPassword(ctx,
		vars.Neo4jPasswordSecretName,
		&random.RandomPasswordArgs{
			Length:     pulumi.Int(12),
			MinSpecial: pulumi.Int(3),
			MinNumeric: pulumi.Int(2),
			MinUpper:   pulumi.Int(2),
			MinLower:   pulumi.Int(2),
			Special:    pulumi.Bool(true),
			Numeric:    pulumi.Bool(true),
			Upper:      pulumi.Bool(true),
			Lower:      pulumi.Bool(true),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create random password for neo4j")
	}

	// Convert the password to base64 to store in the Kubernetes secret data field.
	base64Password := createdRandomPassword.Result.ApplyT(func(p string) (string, error) {
		return base64.StdEncoding.EncodeToString([]byte(p)), nil
	}).(pulumi.StringOutput)

	// Create or update the Kubernetes secret to hold the admin password.
	createdSecret, err := kubernetescorev1.NewSecret(ctx,
		vars.Neo4jPasswordSecretName,
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(vars.Neo4jPasswordSecretName),
				Namespace: createdNamespace.Metadata.Name(),
			},
			Data: pulumi.StringMap{
				vars.Neo4jPasswordSecretKey: base64Password,
			},
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create neo4j password secret")
	}

	// Export some relevant fields for easy retrieval or debugging.
	ctx.Export(OpUsername, pulumi.String("neo4j"))
	ctx.Export(OpPasswordSecretName, createdSecret.Metadata.Name())
	ctx.Export(OpPasswordSecretKey, pulumi.String(vars.Neo4jPasswordSecretKey))

	return nil
}
