package module

import (
	"github.com/pkg/errors"
	natskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/natskubernetes/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmeta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// authSecret provisions a Kubernetes Secret that stores either a bearer
// token or basic-auth credentials, depending on spec.auth.scheme.
//
//   - bearer_token ➜ { token=<random> } (exported via OpAuthSecretKey)
//   - basic_auth ➜ { user=<random>, password=<random> } – we export
//     the *password* key as OpAuthSecretKey (so the single
//     KubernetesSecretKey field in stack-outputs is valid).
func authSecret(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	// If auth is disabled, nothing to create.
	if !locals.NatsKubernetes.Spec.Auth.Enabled {
		return nil
	}

	secretName := locals.AuthSecretName

	switch locals.NatsKubernetes.Spec.Auth.Scheme {
	case natskubernetesv1.NatsKubernetesAuthScheme_bearer_token:
		// ------------------------------------------------ bearer token
		createdToken, err := random.NewRandomPassword(ctx,
			"auth-token",
			&random.RandomPasswordArgs{
				Length:  pulumi.Int(32),
				Special: pulumi.Bool(false),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to generate auth token")
		}

		_, err = kubernetescorev1.NewSecret(ctx,
			"auth-token-secret",
			&kubernetescorev1.SecretArgs{
				Metadata: &kubernetesmeta.ObjectMetaArgs{
					Name:      pulumi.String(secretName),
					Namespace: createdNamespace.Metadata.Name(),
					Labels:    pulumi.ToStringMap(locals.Labels),
				},
				StringData: pulumi.StringMap{
					vars.AuthSecretKey: createdToken.Result,
				},
				Type: pulumi.String("Opaque"),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to create bearer-token secret")
		}

	case natskubernetesv1.NatsKubernetesAuthScheme_basic_auth:
		// ------------------------------------------------ basic auth
		createdUser, err := random.NewRandomString(ctx,
			"auth-user",
			&random.RandomStringArgs{
				Length: pulumi.Int(8),
				Upper:  pulumi.Bool(false),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to generate basic-auth user")
		}

		createdPass, err := random.NewRandomPassword(ctx,
			"auth-password",
			&random.RandomPasswordArgs{
				Length:  pulumi.Int(32),
				Special: pulumi.Bool(false),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to generate basic-auth password")
		}

		_, err = kubernetescorev1.NewSecret(ctx,
			"auth-basic-secret",
			&kubernetescorev1.SecretArgs{
				Metadata: &kubernetesmeta.ObjectMetaArgs{
					Name:      pulumi.String(secretName),
					Namespace: createdNamespace.Metadata.Name(),
					Labels:    pulumi.ToStringMap(locals.Labels),
				},
				StringData: pulumi.StringMap{
					vars.AuthSecretKeyUser:     createdUser.Result,
					vars.AuthSecretKeyPassword: createdPass.Result,
				},
				Type: pulumi.String("Opaque"),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to create basic-auth secret")
		}

		// Override the exported key to point at password (single-key output)
		ctx.Export(OpAuthSecretKey, pulumi.String(vars.AuthSecretKeyPassword))
	}

	return nil
}
