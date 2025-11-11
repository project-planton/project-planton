package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/sortstringmap"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// secret creates a "main" Kubernetes Secret containing all secret environment variables
// from CronJobKubernetes.Spec.Env.Secrets.
func secret(ctx *pulumi.Context, locals *Locals, createdNamespace *corev1.Namespace) error {
	dataMap := make(map[string]string)

	if locals.CronJobKubernetes.Spec.Env != nil && locals.CronJobKubernetes.Spec.Env.Secrets != nil {
		sortedSecretKeys := sortstringmap.SortMap(locals.CronJobKubernetes.Spec.Env.Secrets)
		for _, key := range sortedSecretKeys {
			dataMap[key] = locals.CronJobKubernetes.Spec.Env.Secrets[key]
		}
	}

	// If there are no secrets, we don't need to create a secret resource.
	// But for consistency, let's create it regardless in case any
	// future changes rely on the "main" secret existing.
	_, err := corev1.NewSecret(ctx,
		"main",
		&corev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String("main"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Type:       pulumi.String("Opaque"),
			StringData: pulumi.ToStringMap(dataMap),
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create secret resource")
	}

	return nil
}
