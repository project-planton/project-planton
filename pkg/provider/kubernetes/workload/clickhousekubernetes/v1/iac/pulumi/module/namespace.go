package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createNamespace creates a dedicated Kubernetes namespace for the ClickHouse deployment
// The namespace is labeled with resource metadata for tracking and organization
//
// All ClickHouse resources (pods, services, secrets, ConfigMaps) are created within this namespace
// providing isolation and making it easy to manage the entire deployment as a unit
func createNamespace(
	ctx *pulumi.Context,
	locals *Locals,
	kubernetesProvider pulumi.ProviderResource,
) (*kubernetescorev1.Namespace, error) {
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: kubernetesmetav1.ObjectMetaPtrInput(
				&kubernetesmetav1.ObjectMetaArgs{
					Name:   pulumi.String(locals.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		}, pulumi.Provider(kubernetesProvider))

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	return createdNamespace, nil
}
