package module

import (
	"github.com/pkg/errors"
	kubernetesneo4jv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/workload/kubernetesneo4j/v1"
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

	// Create a namespace for the Neo4j deployment. (Namespace name is derived from the resource metadata)
	createdNamespace, err := kubernetescorev1.NewNamespace(
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

	// Install the Neo4j Helm chart in the newly created namespace, applying user-specified config.
	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to deploy neo4j helm chart")
	}

	return nil
}
