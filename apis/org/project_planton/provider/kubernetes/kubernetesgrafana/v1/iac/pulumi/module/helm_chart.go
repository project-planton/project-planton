package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	grafanaHelmChartName    = "grafana"
	grafanaHelmChartVersion = "8.7.0"
	grafanaHelmChartRepoUrl = "https://grafana.github.io/helm-charts"
)

func helmChart(ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	// https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride": pulumi.String(locals.KubernetesGrafana.Metadata.Name),
		"resources":        containerresources.ConvertToPulumiMap(locals.KubernetesGrafana.Spec.Container.Resources),
		"service": pulumi.Map{
			"type": pulumi.String("ClusterIP"),
		},
		"adminUser":     pulumi.String("admin"),
		"adminPassword": pulumi.String("admin"),
		"persistence": pulumi.Map{
			"enabled": pulumi.Bool(false),
		},
	}

	//install grafana helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.KubernetesGrafana.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(grafanaHelmChartName),
			Version:   pulumi.String(grafanaHelmChartVersion),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values:    helmValues,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(grafanaHelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create helm chart")
	}

	return nil
}
