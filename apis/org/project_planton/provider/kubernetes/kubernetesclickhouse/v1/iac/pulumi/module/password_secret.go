package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createPasswordSecret generates a secure random password and stores it in a Kubernetes Secret
// Returns the created Secret resource for reference in ClickHouseInstallation
//
// The password is used for the default ClickHouse admin user and is automatically
// SHA256-hashed by ClickHouse when referenced via k8s_secret in the CHI configuration.
func createPasswordSecret(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
) (*kubernetescorev1.Secret, error) {
	// Generate cryptographically secure random password
	createdRandomString, err := generateRandomPassword(ctx, createdNamespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate random password")
	}

	// Create Kubernetes Secret to store the password
	createdSecret, err := createKubernetesSecret(ctx, locals, createdNamespace, createdRandomString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create password secret")
	}

	return createdSecret, nil
}

// generateRandomPassword creates a secure random password using Pulumi's random provider
// Password contains a mix of uppercase, lowercase, numbers, and safe special characters
func generateRandomPassword(
	ctx *pulumi.Context,
	createdNamespace *kubernetescorev1.Namespace,
) (*random.RandomPassword, error) {
	// Generate random password with complexity requirements
	// Use only safe special characters to avoid encoding problems in connection strings
	createdRandomString, err := random.NewRandomPassword(ctx,
		"root-password",
		&random.RandomPasswordArgs{
			Length:     pulumi.Int(20),
			Special:    pulumi.Bool(true),
			Numeric:    pulumi.Bool(true),
			Upper:      pulumi.Bool(true),
			Lower:      pulumi.Bool(true),
			MinSpecial: pulumi.Int(2),
			MinNumeric: pulumi.Int(3),
			MinUpper:   pulumi.Int(3),
			MinLower:   pulumi.Int(3),
			// Safe special characters that work in connection strings and URLs
			OverrideSpecial: pulumi.String("-_+="),
		}, pulumi.Parent(createdNamespace))

	if err != nil {
		return nil, errors.Wrap(err, "failed to generate random password value")
	}

	return createdRandomString, nil
}

// createKubernetesSecret creates a Kubernetes Secret containing the ClickHouse admin password
// Uses StringData (not Data) to avoid double base64 encoding - Kubernetes handles encoding automatically
func createKubernetesSecret(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
	randomPassword *random.RandomPassword,
) (*kubernetescorev1.Secret, error) {
	// Create secret with the generated password
	// Note: Kubernetes automatically base64 encodes secret data, so we use StringData instead
	createdSecret, err := kubernetescorev1.NewSecret(ctx,
		locals.KubernetesClickHouse.Metadata.Name,
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.KubernetesClickHouse.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
			},
			StringData: pulumi.StringMap{
				vars.ClickhousePasswordKey: randomPassword.Result,
			},
		}, pulumi.Parent(createdNamespace))

	if err != nil {
		return nil, errors.Wrap(err, "failed to create kubernetes secret")
	}

	return createdSecret, nil
}
