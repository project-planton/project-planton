package module

import (
	"github.com/pkg/errors"
	helmreleasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/helmrelease/v1"
	helmreleaseoutputs "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/helmrelease/v1/iac/pulumi/module/outputs"
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
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create a new descriptive variable for the api-resource in the input.
	helmRelease := stackInput.Target

	//decide on the name of the namespace
	namespaceName := helmRelease.Metadata.Id
	ctx.Export(helmreleaseoutputs.Namespace, pulumi.String(namespaceName))

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		namespaceName,
		&kubernetescorev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:   pulumi.String(namespaceName),
				Labels: pulumi.ToStringMap(locals.Labels),
			}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", namespaceName)
	}

	//install helm-chart
	_, err = helmv3.NewChart(ctx,
		helmRelease.Metadata.Id,
		helmv3.ChartArgs{
			Chart:     pulumi.String(helmRelease.Spec.Name),
			Version:   pulumi.String(helmRelease.Spec.Version),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values:    convertstringmaps.ConvertGoStringMapToPulumiMap(helmRelease.Spec.Values),
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(helmRelease.Spec.Repo),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create helm-chart")
	}
	return nil
}
