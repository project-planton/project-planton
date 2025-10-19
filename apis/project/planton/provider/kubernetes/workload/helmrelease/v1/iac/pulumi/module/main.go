package module

import (
	"github.com/pkg/errors"
	helmreleasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/helmrelease/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *helmreleasev1.HelmReleaseStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:   pulumi.String(locals.Namespace),
				Labels: pulumi.ToStringMap(locals.Labels),
			}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	//install helm-chart
	_, err = helmv3.NewChart(ctx,
		locals.HelmRelease.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(locals.HelmRelease.Spec.Name),
			Version:   pulumi.String(locals.HelmRelease.Spec.Version),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values:    convertstringmaps.ConvertGoStringMapToPulumiMap(locals.HelmRelease.Spec.Values),
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(locals.HelmRelease.Spec.Repo),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create helm-chart")
	}
	return nil
}
