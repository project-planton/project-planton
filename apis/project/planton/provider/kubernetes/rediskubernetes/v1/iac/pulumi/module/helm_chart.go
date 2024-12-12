package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func helmChart(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {
	//install helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.RedisKubernetes.Metadata.Id,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: createdNamespace.Metadata.Name().Elem(),
			//https://github.com/bitnami/charts/blob/main/bitnami/redis/values.yaml
			Values: pulumi.Map{
				"fullnameOverride": pulumi.String(locals.RedisKubernetes.Metadata.Name),
				"architecture":     pulumi.String("standalone"),
				"master": pulumi.Map{
					"podLabels": convertstringmaps.ConvertGoStringMapToPulumiMap(locals.Labels),
					"resources": containerresources.ConvertToPulumiMap(locals.RedisKubernetes.Spec.Container.Resources),
					"persistence": pulumi.Map{
						"enabled": pulumi.Bool(locals.RedisKubernetes.Spec.Container.IsPersistenceEnabled),
						"size":    pulumi.String(locals.RedisKubernetes.Spec.Container.DiskSize),
					},
				},
				"replica": pulumi.Map{
					"podLabels":    convertstringmaps.ConvertGoStringMapToPulumiMap(locals.Labels),
					"replicaCount": pulumi.Int(locals.RedisKubernetes.Spec.Container.Replicas),
					"resources":    containerresources.ConvertToPulumiMap(locals.RedisKubernetes.Spec.Container.Resources),
					"persistence": pulumi.Map{
						"enabled": pulumi.Bool(locals.RedisKubernetes.Spec.Container.IsPersistenceEnabled),
						"size":    pulumi.String(locals.RedisKubernetes.Spec.Container.DiskSize),
					},
				},
				"auth": pulumi.Map{
					"existingSecret":            pulumi.String(vars.RedisPasswordSecretName),
					"existingSecretPasswordKey": pulumi.String(vars.RedisPasswordSecretKey),
				},
			},
			//if you need to add the repository, you can specify `repo url`:
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create helm chart")
	}

	return nil
}
