package module

import (
	"github.com/pkg/errors"
	kuberneteskeycloakv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kuberneteskeycloak/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *kuberneteskeycloakv1.KubernetesKeycloakStackInput) error {
	//initialize locals
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup kubernetes provider")
	}

	//conditionally create namespace resource based on create_namespace flag
	var createdNamespace *kubernetescorev1.Namespace
	if stackInput.Target.Spec.CreateNamespace {
		createdNamespace, err = kubernetescorev1.NewNamespace(ctx,
			locals.Namespace,
			&kubernetescorev1.NamespaceArgs{
				Metadata: kubernetesmetav1.ObjectMetaPtrInput(
					&kubernetesmetav1.ObjectMetaArgs{
						Name:   pulumi.String(locals.Namespace),
						Labels: pulumi.ToStringMap(locals.Labels),
					}),
			}, pulumi.Provider(kubernetesProvider))
		if err != nil {
			return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
		}
	}

	// TODO: Future Keycloak Helm chart deployment would depend on createdNamespace
	// When implementing the Helm chart, resources should use:
	// pulumi.DependsOn([]pulumi.Resource{createdNamespace}) if namespace was created
	_ = createdNamespace // Suppress unused variable warning until Helm chart is implemented

	return nil
}
