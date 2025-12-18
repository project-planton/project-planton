package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/sortstringmap"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// secret creates a "main" Kubernetes Secret containing all secret environment variables
// from KubernetesCronJob.Spec.Env.Secrets.
func secret(ctx *pulumi.Context, locals *Locals, kubernetesProvider pulumi.ProviderResource) error {
	dataMap := make(map[string]string)

	if locals.KubernetesCronJob.Spec.Env != nil && locals.KubernetesCronJob.Spec.Env.Secrets != nil {
		sortedSecretKeys := sortstringmap.SortMap(locals.KubernetesCronJob.Spec.Env.Secrets)
		for _, key := range sortedSecretKeys {
			dataMap[key] = locals.KubernetesCronJob.Spec.Env.Secrets[key]
		}
	}

	// If there are no secrets, we don't need to create a secret resource.
	// But for consistency, let's create it regardless in case any
	// future changes rely on the env secrets secret existing.
	_, err := corev1.NewSecret(ctx,
		locals.EnvSecretsSecretName,
		&corev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.EnvSecretsSecretName),
				Namespace: pulumi.String(locals.Namespace),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Type:       pulumi.String("Opaque"),
			StringData: pulumi.ToStringMap(dataMap),
		},
		pulumi.Provider(kubernetesProvider),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create secret resource")
	}

	return nil
}
