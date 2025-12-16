package module

import (
	"github.com/pkg/errors"
	kubernetesneo4jv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesneo4j/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources sets up all the Pulumi resources for the KubernetesNeo4J component.
func Resources(ctx *pulumi.Context, stackInput *kubernetesneo4jv1.KubernetesNeo4JStackInput) error {
	// Initialize local variables from the stack input.
	locals := initializeLocals(ctx, stackInput)

	// Create the kubernetes provider from the credential in the stack input.
	createdKubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx,
		stackInput.ProviderConfig,
		"kubernetes",
	)
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	// Conditionally create namespace resource based on create_namespace flag
	var createdNamespace *kubernetescorev1.Namespace
	if stackInput.Target.Spec.CreateNamespace {
		createdNamespace, err = kubernetescorev1.NewNamespace(
			ctx,
			locals.Namespace,
			&kubernetescorev1.NamespaceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:   pulumi.String(locals.Namespace),
					Labels: pulumi.ToStringMap(locals.Labels),
				},
			},
			pulumi.Provider(createdKubernetesProvider),
		)
		if err != nil {
			return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
		}
	}

	// Install the Neo4j Helm chart, applying user-specified config.
	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to deploy neo4j helm chart")
	}

	return nil
}
