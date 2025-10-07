package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/mergestringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func clickhouse(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	// https://github.com/bitnami/charts/blob/main/bitnami/clickhouse/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride":  pulumi.String(locals.KubeServiceName),
		"namespaceOverride": createdNamespace.Metadata.Name(),
		"shards":            pulumi.Int(int(locals.ClickhouseKubernetes.Spec.Container.Replicas)),
		"replicaCount":      pulumi.Int(int(locals.ClickhouseKubernetes.Spec.Container.Replicas)),
		"resources":         containerresources.ConvertToPulumiMap(locals.ClickhouseKubernetes.Spec.Container.Resources),
		"persistence": pulumi.Map{
			"enabled": pulumi.Bool(locals.ClickhouseKubernetes.Spec.Container.IsPersistenceEnabled),
			"size":    pulumi.String(locals.ClickhouseKubernetes.Spec.Container.DiskSize),
		},
		"podLabels":    pulumi.ToStringMap(locals.KubernetesLabels),
		"commonLabels": pulumi.ToStringMap(locals.KubernetesLabels),
		"auth": pulumi.Map{
			"existingSecret":    pulumi.String(locals.KubeServiceName),
			"existingSecretKey": pulumi.String(vars.ClickhousePasswordKey),
			"username":          pulumi.String(vars.DefaultUsername),
		},
		// Use bitnamilegacy registry due to Bitnami discontinuing free Docker Hub images (Sep 2025)
		// See: https://github.com/bitnami/containers/issues/83267
		// Global image registry override for all Bitnami images (including ClickHouse and ZooKeeper)
		"global": pulumi.Map{
			"imageRegistry": pulumi.String("docker.io/bitnamilegacy"),
		},
		// ClickHouse specific image - just the image name (no registry prefix)
		"image": pulumi.Map{
			"repository": pulumi.String("clickhouse"),
		},
	}

	// Configure clustering if enabled
	if locals.ClickhouseKubernetes.Spec.Cluster != nil && locals.ClickhouseKubernetes.Spec.Cluster.IsEnabled {
		helmValues["shards"] = pulumi.Int(int(locals.ClickhouseKubernetes.Spec.Cluster.ShardCount))
		helmValues["replicaCount"] = pulumi.Int(int(locals.ClickhouseKubernetes.Spec.Cluster.ReplicaCount))
		helmValues["keeper"] = pulumi.Map{
			"enabled": pulumi.Bool(true),
		}
		helmValues["zookeeper"] = pulumi.Map{
			"enabled": pulumi.Bool(true),
			// ZooKeeper image - just the image name (global.imageRegistry handles the registry)
			"image": pulumi.Map{
				"repository": pulumi.String("zookeeper"),
			},
		}
	}

	mergestringmaps.MergeMapToPulumiMap(helmValues, locals.ClickhouseKubernetes.Spec.HelmValues)

	// install helm-chart
	_, err := helmv3.NewChart(ctx,
		locals.ClickhouseKubernetes.Metadata.Name,
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
		return errors.Wrap(err, "failed to create clickhouse helm-chart")
	}
	return nil
}
