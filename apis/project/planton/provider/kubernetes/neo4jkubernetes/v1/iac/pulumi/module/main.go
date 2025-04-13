package module

import (
	"github.com/pkg/errors"
	neo4jkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/neo4jkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources sets up all the Pulumi resources for the Neo4jKubernetes component.
func Resources(ctx *pulumi.Context, stackInput *neo4jkubernetesv1.Neo4JKubernetesStackInput) error {
	// Initialize local variables from the stack input.
	locals := initializeLocals(ctx, stackInput)

	// Create the kubernetes provider from the credential in the stack input.
	createdKubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(
		ctx,
		stackInput.ProviderCredential,
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

	// Generate a random admin password (if desired) and store it in a Kubernetes secret.
	if err := adminPassword(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create admin password secret")
	}

	// Install the Neo4j Helm chart in the newly created namespace, applying user-specified config.
	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to deploy neo4j helm chart")
	}

	// If ingress is enabled in the spec, create load balancer services to expose Neo4j externally or internally.
	if locals.Neo4jKubernetes.Spec.Ingress.IsEnabled &&
		locals.Neo4jKubernetes.Spec.Ingress.DnsDomain != "" {
		if err := loadBalancerIngress(ctx, locals, createdNamespace); err != nil {
			return errors.Wrap(err, "failed to create load-balancer ingress resources")
		}
	}

	return nil
}
