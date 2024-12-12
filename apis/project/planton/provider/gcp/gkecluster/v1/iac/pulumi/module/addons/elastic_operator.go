package addons

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/vars"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ElasticOperator(ctx *pulumi.Context, locals *localz.Locals,
	kubernetesProvider *pulumikubernetes.Provider) error {

	createdNamespace, err := kubernetescorev1.NewNamespace(ctx, vars.ElasticOperator.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.ElasticOperator.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create namespace")
	}

	//create helm-release
	_, err = helm.NewRelease(ctx, vars.ElasticOperator.HelmChartName,
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.ElasticOperator.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.ElasticOperator.HelmChartName),
			Version:         pulumi.String(vars.ElasticOperator.HelmChartVersion),
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
				Repo: pulumi.String(vars.ElasticOperator.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}

	return nil
}
