package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/mergestringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/containerresources"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func signoz(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
	kubernetesProvider pulumi.ProviderResource) error {

	// https://github.com/SigNoz/charts/blob/main/charts/signoz/values.yaml
	helmValues := pulumi.Map{
		"fullnameOverride": pulumi.String(locals.SignozKubernetes.Metadata.Name),
		"podLabels":        pulumi.ToStringMap(locals.KubernetesLabels),
		"commonLabels":     pulumi.ToStringMap(locals.KubernetesLabels),
	}

	// Configure SigNoz container (main binary with UI, API, Ruler, Alertmanager)
	if locals.SignozKubernetes.Spec.SignozContainer != nil {
		signozValues := pulumi.Map{
			"replicaCount": pulumi.Int(int(locals.SignozKubernetes.Spec.SignozContainer.Replicas)),
		}

		if locals.SignozKubernetes.Spec.SignozContainer.Resources != nil {
			signozValues["resources"] = containerresources.ConvertToPulumiMap(
				locals.SignozKubernetes.Spec.SignozContainer.Resources)
		}

		if locals.SignozKubernetes.Spec.SignozContainer.Image != nil {
			signozValues["image"] = pulumi.Map{
				"repository": pulumi.String(locals.SignozKubernetes.Spec.SignozContainer.Image.Repo),
				"tag":        pulumi.String(locals.SignozKubernetes.Spec.SignozContainer.Image.Tag),
			}
		}

		helmValues["signoz"] = signozValues
	}

	// Configure OpenTelemetry Collector
	if locals.SignozKubernetes.Spec.OtelCollectorContainer != nil {
		otelValues := pulumi.Map{
			"replicaCount": pulumi.Int(int(locals.SignozKubernetes.Spec.OtelCollectorContainer.Replicas)),
		}

		if locals.SignozKubernetes.Spec.OtelCollectorContainer.Resources != nil {
			otelValues["resources"] = containerresources.ConvertToPulumiMap(
				locals.SignozKubernetes.Spec.OtelCollectorContainer.Resources)
		}

		if locals.SignozKubernetes.Spec.OtelCollectorContainer.Image != nil {
			otelValues["image"] = pulumi.Map{
				"repository": pulumi.String(locals.SignozKubernetes.Spec.OtelCollectorContainer.Image.Repo),
				"tag":        pulumi.String(locals.SignozKubernetes.Spec.OtelCollectorContainer.Image.Tag),
			}
		}

		helmValues["otelCollector"] = otelValues
	}

	// Configure database (ClickHouse)
	if locals.SignozKubernetes.Spec.Database != nil {
		if locals.SignozKubernetes.Spec.Database.IsExternal {
			// External ClickHouse configuration
			if locals.SignozKubernetes.Spec.Database.ExternalDatabase != nil {
				ext := locals.SignozKubernetes.Spec.Database.ExternalDatabase
				helmValues["clickhouse"] = pulumi.Map{
					"enabled": pulumi.Bool(false),
				}
				helmValues["externalClickhouse"] = pulumi.Map{
					"host":     pulumi.String(ext.Host),
					"httpPort": pulumi.Int(int(ext.HttpPort)),
					"tcpPort":  pulumi.Int(int(ext.TcpPort)),
					"cluster":  pulumi.String(ext.ClusterName),
					"secure":   pulumi.Bool(ext.IsSecure),
					"user":     pulumi.String(ext.Username),
					"password": pulumi.String(ext.Password),
				}
			}
		} else {
			// Self-managed ClickHouse configuration
			if locals.SignozKubernetes.Spec.Database.ManagedDatabase != nil {
				managed := locals.SignozKubernetes.Spec.Database.ManagedDatabase
				clickhouseValues := pulumi.Map{
					"enabled": pulumi.Bool(true),
					// Use bitnamilegacy registry due to Bitnami discontinuing free Docker Hub images (Sep 2025)
					// See: https://github.com/bitnami/containers/issues/83267
					// ClickHouse specific image override (not using global.imageRegistry to avoid affecting Altinity operator)
					"image": pulumi.Map{
						"registry":   pulumi.String("docker.io"),
						"repository": pulumi.String("bitnamilegacy/clickhouse"),
					},
					// ZooKeeper is a subchart dependency of ClickHouse in SigNoz
					// Override ZooKeeper image here under clickhouse.zookeeper
					"zookeeper": pulumi.Map{
						"image": pulumi.Map{
							"registry":   pulumi.String("docker.io"),
							"repository": pulumi.String("bitnamilegacy/zookeeper"),
						},
					},
				}

				// ClickHouse container configuration
				if managed.Container != nil {
					clickhouseValues["replicaCount"] = pulumi.Int(int(managed.Container.Replicas))

					if managed.Container.Resources != nil {
						clickhouseValues["resources"] = containerresources.ConvertToPulumiMap(
							managed.Container.Resources)
					}

					if managed.Container.Image != nil {
						clickhouseValues["image"] = pulumi.Map{
							"registry":   pulumi.String("docker.io"),
							"repository": pulumi.String(managed.Container.Image.Repo),
							"tag":        pulumi.String(managed.Container.Image.Tag),
						}
					}

					clickhouseValues["persistence"] = pulumi.Map{
						"enabled": pulumi.Bool(managed.Container.IsPersistenceEnabled),
						"size":    pulumi.String(managed.Container.DiskSize),
					}
				}

				// ClickHouse clustering configuration
				if managed.Cluster != nil && managed.Cluster.IsEnabled {
					clickhouseValues["layout"] = pulumi.Map{
						"shardsCount":   pulumi.Int(int(managed.Cluster.ShardCount)),
						"replicasCount": pulumi.Int(int(managed.Cluster.ReplicaCount)),
					}
				}

				helmValues["clickhouse"] = clickhouseValues

				// Zookeeper configuration (required for distributed ClickHouse)
				// Note: In SigNoz, ZooKeeper settings must be configured under clickhouse.zookeeper
				if managed.Zookeeper != nil && managed.Zookeeper.IsEnabled {
					// Get existing clickhouse.zookeeper map or create new one
					zkConfig, hasZk := clickhouseValues["zookeeper"].(pulumi.Map)
					if !hasZk {
						zkConfig = pulumi.Map{}
					}

					zkConfig["enabled"] = pulumi.Bool(true)

					if managed.Zookeeper.Container != nil {
						zkConfig["replicaCount"] = pulumi.Int(int(managed.Zookeeper.Container.Replicas))

						if managed.Zookeeper.Container.Resources != nil {
							zkConfig["resources"] = containerresources.ConvertToPulumiMap(
								managed.Zookeeper.Container.Resources)
						}

						if managed.Zookeeper.Container.Image != nil {
							zkConfig["image"] = pulumi.Map{
								"registry":   pulumi.String("docker.io"),
								"repository": pulumi.String(managed.Zookeeper.Container.Image.Repo),
								"tag":        pulumi.String(managed.Zookeeper.Container.Image.Tag),
							}
						}

						zkConfig["persistence"] = pulumi.Map{
							"size": pulumi.String(managed.Zookeeper.Container.DiskSize),
						}
					}

					// Update the clickhouse.zookeeper configuration
					clickhouseValues["zookeeper"] = zkConfig
				} else {
					// Get existing clickhouse.zookeeper map or create new one
					zkConfig, hasZk := clickhouseValues["zookeeper"].(pulumi.Map)
					if !hasZk {
						zkConfig = pulumi.Map{}
					}
					// Explicitly disable Zookeeper if not needed
					zkConfig["enabled"] = pulumi.Bool(false)
					clickhouseValues["zookeeper"] = zkConfig
				}
			}
		}
	}

	// Note: Ingress is NOT configured via Helm chart values
	// We use Kubernetes Gateway API for ingress (see ingress.go)
	// This provides better control and consistency with other workloads

	// Merge custom Helm values
	mergestringmaps.MergeMapToPulumiMap(helmValues, locals.SignozKubernetes.Spec.HelmValues)

	// Install SigNoz Helm chart
	_, err := helmv3.NewChart(ctx,
		locals.SignozKubernetes.Metadata.Name,
		helmv3.ChartArgs{
			Chart:     pulumi.String(vars.HelmChartName),
			Version:   pulumi.String(vars.HelmChartVersion),
			Namespace: pulumi.String(locals.Namespace),
			Values:    helmValues,
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String(vars.HelmChartRepoUrl),
			},
		}, pulumi.Parent(createdNamespace), pulumi.Provider(kubernetesProvider))

	if err != nil {
		return errors.Wrap(err, "failed to create signoz helm-chart")
	}

	return nil
}
