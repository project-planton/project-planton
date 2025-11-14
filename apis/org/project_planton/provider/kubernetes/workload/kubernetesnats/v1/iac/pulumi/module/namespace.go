package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmeta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// namespace creates the Kubernetes namespace that will hold every
// resource in the KubernetesNats deployment.  We return the created
// object so that downstream helpers can set it as Parent and inherit
// the namespace automatically.
//
// Terraform equivalent: a standalone kubernetes_namespace resource.
func namespace(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider pulumi.ProviderResource) (*kubernetescorev1.Namespace, error) {

	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: &kubernetesmeta.ObjectMetaArgs{
				Name:   pulumi.String(locals.Namespace),
				Labels: pulumi.ToStringMap(locals.Labels),
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	return createdNamespace, nil
}
