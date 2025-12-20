package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/sortstringmap"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// secret creates a Kubernetes Secret containing all secret environment variables
// from KubernetesDaemonSet.Spec.Container.App.Env.Secrets.
func secret(ctx *pulumi.Context, locals *Locals, kubernetesProvider pulumi.ProviderResource) error {
	dataMap := make(map[string]string)

	if locals.KubernetesDaemonSet.Spec.Container.App.Env != nil &&
		locals.KubernetesDaemonSet.Spec.Container.App.Env.Secrets != nil {
		sortedSecretKeys := sortstringmap.SortMap(locals.KubernetesDaemonSet.Spec.Container.App.Env.Secrets)
		for _, key := range sortedSecretKeys {
			dataMap[key] = locals.KubernetesDaemonSet.Spec.Container.App.Env.Secrets[key]
		}
	}

	// Create the secret even if empty for consistency
	_, err := corev1.NewSecret(ctx,
		locals.EnvSecretName,
		&corev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.EnvSecretName),
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
