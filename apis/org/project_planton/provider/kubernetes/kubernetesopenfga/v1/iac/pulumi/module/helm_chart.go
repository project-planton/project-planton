package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// helmChart installs the upstream OpenFGA Helm chart and tailors it to the spec.
func helmChart(ctx *pulumi.Context, locals *Locals, kubernetesProvider pulumi.ProviderResource) error {
	// https://github.com/openfga/helm-charts/blob/main/charts/openfga/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride": pulumi.String(locals.KubernetesOpenFga.Metadata.Name),
		"replicaCount":     pulumi.Int(locals.KubernetesOpenFga.Spec.Container.Replicas),
		"datastore": pulumi.Map{
			"engine": pulumi.String(locals.KubernetesOpenFga.Spec.Datastore.Engine),
			"uri":    pulumi.String(locals.KubernetesOpenFga.Spec.Datastore.Uri),
		},
		"resources": containerresources.ConvertToPulumiMap(locals.KubernetesOpenFga.Spec.Container.Resources),
	}

	// Install openfga helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.KubernetesOpenFga.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: pulumi.String(locals.Namespace),
			Values:    helmValues,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create helm chart")
	}

	return nil
}
