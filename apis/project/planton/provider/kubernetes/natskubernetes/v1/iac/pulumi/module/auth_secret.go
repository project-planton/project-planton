package module

import (
	"github.com/pkg/errors"
	natskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/natskubernetes/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmeta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// authSecret provisions Kubernetes Secrets that store all
// bearer-token credentials, basic-auth admin credentials,
// and (optionally) the no-auth user credentials.
func authSecret(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	// If auth disabled, nothing to create.
	if locals.NatsKubernetes.Spec.Auth == nil || !locals.NatsKubernetes.Spec.Auth.Enabled {
		return nil
	}

	switch locals.NatsKubernetes.Spec.Auth.Scheme {

	//--------------------------------------------------------------------
	// bearer token
	//--------------------------------------------------------------------
	case natskubernetesv1.NatsKubernetesAuthScheme_bearer_token:
		createdToken, err := random.NewRandomPassword(ctx,
			"auth-token",
			&random.RandomPasswordArgs{Length: pulumi.Int(32), Special: pulumi.Bool(false)},
			pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "generate auth token")
		}

		_, err = kubernetescorev1.NewSecret(ctx,
			"auth-token-secret",
			&kubernetescorev1.SecretArgs{
				Metadata: &kubernetesmeta.ObjectMetaArgs{
					Name:      pulumi.String(vars.AdminAuthSecretName),
					Namespace: createdNamespace.Metadata.Name(),
					Labels:    pulumi.ToStringMap(locals.Labels),
				},
				StringData: pulumi.StringMap{vars.AdminAuthSecretKey: createdToken.Result},
				Type:       pulumi.String("Opaque"),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "create bearer-token secret")
		}

	//--------------------------------------------------------------------
	// basic auth
	//--------------------------------------------------------------------
	case natskubernetesv1.NatsKubernetesAuthScheme_basic_auth:

		// ---------------- admin credentials ----------------
		createdUser, err := random.NewRandomString(ctx,
			"auth-user",
			&random.RandomStringArgs{Length: pulumi.Int(8), Upper: pulumi.Bool(false)},
			pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "generate admin user")
		}

		createdPass, err := random.NewRandomPassword(ctx,
			"auth-password",
			&random.RandomPasswordArgs{Length: pulumi.Int(32), Special: pulumi.Bool(false)},
			pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "generate admin password")
		}

		_, err = kubernetescorev1.NewSecret(ctx,
			"auth-basic-secret",
			&kubernetescorev1.SecretArgs{
				Metadata: &kubernetesmeta.ObjectMetaArgs{
					Name:      pulumi.String(vars.AdminAuthSecretName),
					Namespace: createdNamespace.Metadata.Name(),
					Labels:    pulumi.ToStringMap(locals.Labels),
				},
				StringData: pulumi.StringMap{
					vars.NatsUserSecretKeyUsername: createdUser.Result,
					vars.NatsUserSecretKeyPassword: createdPass.Result,
				},
				Type: pulumi.String("Opaque"),
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "create admin secret")
		}

		// Export key pointing at password
		ctx.Export(OpAuthSecretKey, pulumi.String(vars.NatsUserSecretKeyPassword))

		// ---------------- noauth credentials ----------------
		if na := locals.NatsKubernetes.Spec.Auth.NoAuthUser; na != nil && na.Enabled {
			_, err = kubernetescorev1.NewSecret(ctx,
				vars.NoAuthUserSecretName,
				&kubernetescorev1.SecretArgs{
					Metadata: &kubernetesmeta.ObjectMetaArgs{
						Name:      pulumi.String(vars.NoAuthUserSecretName),
						Namespace: createdNamespace.Metadata.Name(),
						Labels:    pulumi.ToStringMap(locals.Labels),
					},
					StringData: pulumi.StringMap{
						vars.NatsUserSecretKeyUsername: pulumi.String(vars.NoAuthUsername),
						vars.NatsUserSecretKeyPassword: pulumi.String(vars.NoAuthPassword),
					},
					Type: pulumi.String("Opaque"),
				}, pulumi.Parent(createdNamespace))
			if err != nil {
				return errors.Wrap(err, "create noauth secret")
			}
		}
	}

	return nil
}
