package module

import (
	"github.com/pkg/errors"
	cronjobkubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/workload/cronjobkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the main entry point for creating all necessary Kubernetes resources
// for a CronJob, based on the CronJobKubernetes API resource definition.
func Resources(ctx *pulumi.Context, stackInput *cronjobkubernetesv1.CronJobKubernetesStackInput) error {
	// Initialize local references
	locals, err := initializeLocals(ctx, stackInput)
	if err != nil {
		return errors.Wrap(err, "failed to initialize locals")
	}

	// Create a Pulumi Kubernetes provider from the given credentials
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx,
		stackInput.ProviderConfig,
		"kubernetes",
	)
	if err != nil {
		return errors.Wrap(err, "failed to create Kubernetes provider")
	}

	// Create (or get) the namespace resource
	createdNamespace, err := corev1.NewNamespace(
		ctx,
		locals.Namespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:   pulumi.String(locals.Namespace),
				Labels: pulumi.ToStringMap(locals.Labels),
			},
		},
		pulumi.Provider(kubernetesProvider),
	)
	if err != nil {
		return errors.Wrapf(err, "failed to create namespace %s", locals.Namespace)
	}

	// Create the main secret resource
	if err := secret(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create secret")
	}

	// Create an image pull secret if Docker credentials are provided
	if locals.ImagePullSecretData != nil {
		if err := createImagePullSecret(ctx, locals, createdNamespace); err != nil {
			return errors.Wrap(err, "failed to create image pull secret")
		}
	}

	// Create the CronJob resource
	_, err = cronJob(ctx, locals, createdNamespace)
	if err != nil {
		return errors.Wrap(err, "failed to create cronjob")
	}

	return nil
}
