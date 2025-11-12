package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// zalandoPostgresOperator installs the Zalando Postgres Operator in the Kubernetes cluster using Helm.
// It creates the necessary namespace and deploys the Helm chart with specific values.
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
// 1. Creates a namespace for the Zalando Postgres Operator and labels it with metadata from locals.
// 2. Deploys the Zalando Postgres Operator Helm chart into the created namespace with specific inherited labels and other configurations.
// 3. Uses Helm chart repository and version specified in the vars package.
// 4. Handles errors and returns any errors encountered during the namespace creation or Helm release deployment.
func zalandoPostgresOperator(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider *pulumikubernetes.Provider) error {
	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.ZalandoPostgresOperator.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.ZalandoPostgresOperator.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create namespace")
	}

	//create helm-release
	_, err = helm.NewRelease(ctx, "zalando-postgres-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.ZalandoPostgresOperator.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.ZalandoPostgresOperator.HelmChartName),
			Version:         pulumi.String(vars.ZalandoPostgresOperator.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"configKubernetes": pulumi.Map{
					"inherited_labels": pulumi.ToStringArray(
						[]string{
							kuberneteslabelkeys.Resource,
							kuberneteslabelkeys.Organization,
							kuberneteslabelkeys.Environment,
							kuberneteslabelkeys.ResourceKind,
							kuberneteslabelkeys.ResourceId,
						},
					),
				},
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.ZalandoPostgresOperator.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}
	return nil
}
