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
				if managed.Zookeeper != nil && managed.Zookeeper.IsEnabled {
					zookeeperValues := pulumi.Map{
						"enabled": pulumi.Bool(true),
					}

					if managed.Zookeeper.Container != nil {
						zookeeperValues["replicaCount"] = pulumi.Int(int(managed.Zookeeper.Container.Replicas))

						if managed.Zookeeper.Container.Resources != nil {
							zookeeperValues["resources"] = containerresources.ConvertToPulumiMap(
								managed.Zookeeper.Container.Resources)
						}

						if managed.Zookeeper.Container.Image != nil {
							zookeeperValues["image"] = pulumi.Map{
								"repository": pulumi.String(managed.Zookeeper.Container.Image.Repo),
								"tag":        pulumi.String(managed.Zookeeper.Container.Image.Tag),
							}
						}

						zookeeperValues["persistence"] = pulumi.Map{
							"size": pulumi.String(managed.Zookeeper.Container.DiskSize),
						}
					}

					helmValues["zookeeper"] = zookeeperValues
				} else {
					// Explicitly disable Zookeeper if not needed
					helmValues["zookeeper"] = pulumi.Map{
						"enabled": pulumi.Bool(false),
					}
				}
			}
		}
	}

	// Configure SigNoz UI ingress
	if locals.SignozKubernetes.Spec.SignozIngress != nil && locals.SignozKubernetes.Spec.SignozIngress.Enabled {
		ingressValues := pulumi.Map{
			"enabled": pulumi.Bool(true),
			"hosts": pulumi.Array{
				pulumi.Map{
					"host": pulumi.String(locals.IngressExternalHostname),
					"paths": pulumi.Array{
						pulumi.Map{
							"path": pulumi.String("/"),
							"port": pulumi.Int(vars.SignozUIPort),
						},
					},
				},
			},
		}

		// TLS configuration would be handled through Helm values if needed

		helmValues["signoz"] = pulumi.Map{
			"ingress": ingressValues,
		}
	}

	// Configure OTel Collector ingress
	if locals.SignozKubernetes.Spec.OtelCollectorIngress != nil &&
		locals.SignozKubernetes.Spec.OtelCollectorIngress.Enabled {
		// Note: In production, separate ingress resources may be needed for gRPC and HTTP
		// due to nginx annotation requirements (nginx.ingress.kubernetes.io/backend-protocol: "GRPC")
		ingressValues := pulumi.Map{
			"enabled": pulumi.Bool(true),
			"hosts": pulumi.Array{
				pulumi.Map{
					"host": pulumi.String(locals.OtelCollectorExternalGrpcHostname),
					"paths": pulumi.Array{
						pulumi.Map{
							"path": pulumi.String("/"),
							"port": pulumi.Int(vars.OtelGrpcPort),
						},
					},
				},
				pulumi.Map{
					"host": pulumi.String(locals.OtelCollectorExternalHttpHostname),
					"paths": pulumi.Array{
						pulumi.Map{
							"path": pulumi.String("/"),
							"port": pulumi.Int(vars.OtelHttpPort),
						},
					},
				},
			},
		}

		// TLS configuration would be handled through Helm values if needed

		helmValues["otelCollector"] = pulumi.Map{
			"ingress": ingressValues,
		}
	}

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
