package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ingressNginx installs the Ingress Nginx controller in the Kubernetes cluster using Helm.
// It creates a namespace for the Ingress Nginx resources and then deploys the Helm chart.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - kubernetesProvider: The Kubernetes provider for Pulumi.
//
// Returns:
// - error: An error object if there is any issue during the installation.
//
// The function performs the following steps:
// 1. Creates a namespace for the Ingress Nginx resources, applying any necessary labels from the locals.
// 2. Deploys the Ingress Nginx Helm chart into the created namespace with specific configurations for the controller service and ingress class resource.
// 3. Uses Helm chart repository and version specified in the vars package.
// 4. Handles errors and returns any errors encountered during the namespace creation or Helm release deployment.
func ingressNginx(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider *pulumikubernetes.Provider) error {
	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.IngressNginx.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.IngressNginx.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create namespace")
	}

	//create helm-release
	_, err = helm.NewRelease(ctx, "ingress-nginx",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.IngressNginx.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.IngressNginx.HelmChartName),
			Version:         pulumi.String(vars.IngressNginx.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"controller": pulumi.Map{
					"service": pulumi.StringMap{
						"type": pulumi.String("ClusterIP"),
					},
					"ingressClassResource": pulumi.Map{
						"default": pulumi.Bool(true),
					},
				},
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.IngressNginx.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}
	return nil
}
