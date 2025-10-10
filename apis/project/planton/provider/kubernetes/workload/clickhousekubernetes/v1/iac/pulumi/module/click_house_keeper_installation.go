package module

import (
	"github.com/pkg/errors"
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	clickhousekeeperv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/altinityoperator/kubernetes/clickhouse_keeper/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// clickhouseKeeperInstallation creates a ClickHouseKeeperInstallation CRD resource
// This is called when coordination.type = keeper (auto-managed ClickHouse Keeper)
//
// The ClickHouseKeeperInstallation CRD is provided by Altinity operator and manages
// ClickHouse Keeper pods, services, and configuration.
func clickhouseKeeperInstallation(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
	keeperConfig *clickhousekubernetesv1.ClickHouseKubernetesKeeperConfig,
) error {
	// Apply defaults if keeper_config is not specified
	replicas := int(1) // Default: 1 for development
	diskSize := "10Gi" // Default: 10Gi sufficient for metadata

	if keeperConfig != nil {
		if keeperConfig.Replicas > 0 {
			replicas = int(keeperConfig.Replicas)
		}
		if keeperConfig.DiskSize != "" {
			diskSize = keeperConfig.DiskSize
		}
	}

	// Keeper name - use "keeper" as the standard name
	// This creates a service named "keeper-keeper" following pattern "keeper-<name>"
	keeperName := "keeper"

	// Build the ClickHouseKeeperInstallation CRD
	_, err := clickhousekeeperv1.NewClickHouseKeeperInstallation(ctx,
		keeperName,
		&clickhousekeeperv1.ClickHouseKeeperInstallationArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(keeperName),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: &clickhousekeeperv1.ClickHouseKeeperInstallationSpecArgs{
				Configuration: buildKeeperConfiguration(replicas),
				// Defaults removed - not supported in Altinity operator 0.23.6
				Templates: buildKeeperTemplates(diskSize, keeperConfig),
			},
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create ClickHouseKeeperInstallation")
	}

	return nil
}

// buildKeeperConfiguration builds the configuration section for ClickHouseKeeperInstallation
func buildKeeperConfiguration(replicas int) *clickhousekeeperv1.ClickHouseKeeperInstallationSpecConfigurationArgs {
	return &clickhousekeeperv1.ClickHouseKeeperInstallationSpecConfigurationArgs{
		Clusters: clickhousekeeperv1.ClickHouseKeeperInstallationSpecConfigurationClustersArray{
			&clickhousekeeperv1.ClickHouseKeeperInstallationSpecConfigurationClustersArgs{
				Name: pulumi.String("default"),
				Layout: &clickhousekeeperv1.ClickHouseKeeperInstallationSpecConfigurationClustersLayoutArgs{
					ReplicasCount: pulumi.IntPtr(replicas),
				},
			},
		},
	}
}

// buildKeeperTemplates builds pod and volume templates for Keeper
func buildKeeperTemplates(
	diskSize string,
	keeperConfig *clickhousekubernetesv1.ClickHouseKubernetesKeeperConfig,
) *clickhousekeeperv1.ClickHouseKeeperInstallationSpecTemplatesArgs {
	templates := &clickhousekeeperv1.ClickHouseKeeperInstallationSpecTemplatesArgs{}

	// Always add volume claim template for persistence
	templates.VolumeClaimTemplates = clickhousekeeperv1.ClickHouseKeeperInstallationSpecTemplatesVolumeClaimTemplatesArray{
		&clickhousekeeperv1.ClickHouseKeeperInstallationSpecTemplatesVolumeClaimTemplatesArgs{
			Name: pulumi.String("default"),
			Spec: pulumi.Map{
				"accessModes": pulumi.Array{
					pulumi.String("ReadWriteOnce"),
				},
				"resources": pulumi.Map{
					"requests": pulumi.Map{
						"storage": pulumi.String(diskSize),
					},
				},
			},
		},
	}

	// Add pod template with resources if specified
	if keeperConfig != nil && keeperConfig.Resources != nil {
		templates.PodTemplates = clickhousekeeperv1.ClickHouseKeeperInstallationSpecTemplatesPodTemplatesArray{
			&clickhousekeeperv1.ClickHouseKeeperInstallationSpecTemplatesPodTemplatesArgs{
				Name: pulumi.String("default"),
				Spec: pulumi.Map{
					"containers": pulumi.Array{
						pulumi.Map{
							"name":  pulumi.String("clickhouse-keeper"),
							"image": pulumi.Sprintf("clickhouse/clickhouse-keeper:%s", vars.ClickhouseVersion),
							"resources": pulumi.Map{
								"requests": pulumi.Map{
									"cpu":    pulumi.String(keeperConfig.Resources.Requests.Cpu),
									"memory": pulumi.String(keeperConfig.Resources.Requests.Memory),
								},
								"limits": pulumi.Map{
									"cpu":    pulumi.String(keeperConfig.Resources.Limits.Cpu),
									"memory": pulumi.String(keeperConfig.Resources.Limits.Memory),
								},
							},
						},
					},
				},
			},
		}
	}

	return templates
}
