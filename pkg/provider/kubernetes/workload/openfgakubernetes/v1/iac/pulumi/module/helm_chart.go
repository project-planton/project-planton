package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func helmChart(ctx *pulumi.Context,
	locals *Locals, createdNamespace *kubernetescorev1.Namespace) error {

	// https://github.com/openfga/helm-charts/blob/main/charts/openfga/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride": pulumi.String(locals.OpenFgaKubernetes.Metadata.Name),
		"replicaCount":     pulumi.Int(locals.OpenFgaKubernetes.Spec.Container.Replicas),
		"datastore": pulumi.Map{
			"engine": pulumi.String(locals.OpenFgaKubernetes.Spec.Datastore.Engine),
			"uri":    pulumi.String(locals.OpenFgaKubernetes.Spec.Datastore.Uri),
		},
		"resources": containerresources.ConvertToPulumiMap(locals.OpenFgaKubernetes.Spec.Container.Resources),
	}

	//merge extra helm values provided in the spec with base values
	//mergemaps.MergeMapToPulumiMap(helmValues, locals.OpenFgaKubernetes.Spec.HelmValues)

	//install openfga helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.OpenFgaKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values:    helmValues,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create helm chart")
	}

	return nil
}
