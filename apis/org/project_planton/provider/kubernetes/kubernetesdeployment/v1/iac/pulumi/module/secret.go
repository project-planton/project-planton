package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/sortstringmap"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func secret(ctx *pulumi.Context, locals *Locals, kubernetesProvider pulumi.ProviderResource) error {
	dataMap := make(map[string]string)

	//add all secrets to data map
	if locals.KubernetesDeployment.Spec.Container.App.Env != nil {
		var secrets = locals.KubernetesDeployment.Spec.Container.App.Env.Secrets
		if secrets != nil && len(secrets) > 0 {
			// gather all provided secrets into a simple map
			sortedSecretKeys := sortstringmap.SortMap(secrets)

			for _, sortedSecretKey := range sortedSecretKeys {
				dataMap[sortedSecretKey] = secrets[sortedSecretKey]
			}
		}
	}

	// create a standard kubernetes secret with name "main"
	secretArgs := &kubernetescorev1.SecretArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("main"),
			Namespace: pulumi.String(locals.Namespace),
			Labels:    pulumi.ToStringMap(locals.Labels),
		},
		Type:       pulumi.String("Opaque"),
		StringData: pulumi.ToStringMap(dataMap),
	}

	_, err := kubernetescorev1.NewSecret(ctx,
		"main",
		secretArgs,
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create secret resource")
	}

	return nil
}
