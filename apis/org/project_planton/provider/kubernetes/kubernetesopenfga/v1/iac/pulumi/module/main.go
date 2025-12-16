package module

import (
	"github.com/pkg/errors"
	kubernetesopenfgav1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesopenfga/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *kubernetesopenfgav1.KubernetesOpenFgaStackInput) error {
	locals := initializeLocals(ctx, stackInput)
	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	// Conditionally create or lookup namespace based on create_namespace flag
	var namespaceResource *kubernetescorev1.Namespace
	if stackInput.Target.Spec.CreateNamespace {
		// Create namespace resource
		namespaceResource, err = kubernetescorev1.NewNamespace(ctx,
			locals.Namespace,
			&kubernetescorev1.NamespaceArgs{
				Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
					Name:   pulumi.String(locals.Namespace),
					Labels: pulumi.ToStringMap(locals.Labels),
				}),
			}, pulumi.Provider(kubernetesProvider))
		if err != nil {
			return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
		}
	} else {
		// Look up existing namespace
		namespaceResource, err = kubernetescorev1.GetNamespace(ctx,
			locals.Namespace,
			pulumi.ID(locals.Namespace),
			nil,
			pulumi.Provider(kubernetesProvider))
		if err != nil {
			return errors.Wrapf(err, "failed to lookup existing %s namespace", locals.Namespace)
		}
	}

	//export name of the namespace
	ctx.Export(OpNamespace, namespaceResource.Metadata.Name())

	//install the openfga helm-chart
	if err := helmChart(ctx, locals, namespaceResource); err != nil {
		return errors.Wrap(err, "failed to create helm-chart resources")
	}

	//create istio-ingress resources if ingress is enabled.
	if locals.KubernetesOpenFga.Spec.Ingress != nil && locals.KubernetesOpenFga.Spec.Ingress.Enabled {
		if err := ingress(ctx, locals, namespaceResource, kubernetesProvider, locals.Labels); err != nil {
			return errors.Wrap(err, "failed to create ingress resources")
		}
	}

	return nil
}
