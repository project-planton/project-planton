package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func helmChart(ctx *pulumi.Context,
	locals *Locals, createdNamespace *kubernetescorev1.Namespace) error {

	// https://github.com/openfga/helm-charts/blob/main/charts/openfga/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride": pulumi.String(locals.OpenfgaKubernetes.Metadata.Name),
		"replicaCount":     pulumi.Int(locals.OpenfgaKubernetes.Spec.Container.Replicas),
		"datastore": pulumi.Map{
			"engine": pulumi.String(locals.OpenfgaKubernetes.Spec.Datastore.Engine),
			"uri":    pulumi.String(locals.OpenfgaKubernetes.Spec.Datastore.Uri),
		},
		"resources": containerresources.ConvertToPulumiMap(locals.OpenfgaKubernetes.Spec.Container.Resources),
	}

	//merge extra helm values provided in the spec with base values
	//mergemaps.MergeMapToPulumiMap(helmValues, locals.OpenfgaKubernetes.Spec.HelmValues)

	//install openfga helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.OpenfgaKubernetes.Metadata.Id,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values:    helmValues,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace), pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "3m", Update: "3m", Delete: "3m"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm chart")
	}

	return nil
}
