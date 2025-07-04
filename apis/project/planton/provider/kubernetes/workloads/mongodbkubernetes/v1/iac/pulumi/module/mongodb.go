package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/mergestringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func mongodb(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	// https://github.com/bitnami/charts/blob/main/bitnami/mongodb/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride":  pulumi.String(locals.KubeServiceName),
		"namespaceOverride": createdNamespace.Metadata.Name(),
		"resources":         containerresources.ConvertToPulumiMap(locals.MongodbKubernetes.Spec.Container.Resources),
		// todo: hard-coding this to 1 since we are only using `standalone` architecture,
		// need to revisit this to handle `replicaSet` architecture
		"replicaCount": pulumi.Int(1),
		"persistence": pulumi.Map{
			"enabled": pulumi.Bool(locals.MongodbKubernetes.Spec.Container.IsPersistenceEnabled),
			"size":    pulumi.String(locals.MongodbKubernetes.Spec.Container.DiskSize),
		},
		"podLabels":      pulumi.ToStringMap(locals.KubernetesLabels),
		"commonLabels":   pulumi.ToStringMap(locals.KubernetesLabels),
		"useStatefulSet": pulumi.Bool(true),
		"auth": pulumi.Map{
			"existingSecret": pulumi.String(locals.KubeServiceName),
		},
	}

	mergestringmaps.MergeMapToPulumiMap(helmValues, locals.MongodbKubernetes.Spec.HelmValues)

	// install helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.MongodbKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: pulumi.String(locals.Namespace),
			Values:    helmValues,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace))

	if err != nil {
		return errors.Wrap(err, "failed to create mongodb helm-chart")
	}
	return nil
}
