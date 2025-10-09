package module

import (
	"github.com/pkg/errors"
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	altinityv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/kubernetes/clickhouse/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// clickhouseInstallation creates a ClickHouseInstallation custom resource using the Altinity operator
func clickhouseInstallation(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
	createdSecret *kubernetescorev1.Secret,
) error {
	spec := locals.ClickHouseKubernetes.Spec

	// Determine cluster name
	clusterName := spec.ClusterName
	if clusterName == "" {
		clusterName = locals.ClickHouseKubernetes.Metadata.Name
	}

	// Determine ClickHouse version
	version := spec.Version
	if version == "" {
		version = vars.ClickhouseVersion
	}

	// Build the ClickHouseInstallation CRD
	_, err := altinityv1.NewClickHouseInstallation(ctx,
		clusterName,
		&altinityv1.ClickHouseInstallationArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(clusterName),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: &altinityv1.ClickHouseInstallationSpecArgs{
				Configuration: buildConfiguration(clusterName, spec, createdSecret),
				Defaults:      buildDefaults(),
				Templates:     buildTemplates(spec, version),
			},
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create ClickHouseInstallation")
	}

	return nil
}

// buildConfiguration generates the configuration section of the CHI spec
func buildConfiguration(
	clusterName string,
	spec *clickhousekubernetesv1.ClickHouseKubernetesSpec,
	createdSecret *kubernetescorev1.Secret,
) *altinityv1.ClickHouseInstallationSpecConfigurationArgs {
	// Build users configuration with secret reference
	// Using modern valueFrom.secretKeyRef syntax (recommended since operator v0.23.x)
	// Note: Using 'password' field (plaintext in config, will be hashed by operator)
	// The secret contains plaintext password, operator will hash it when deploying to ClickHouse
	//
	// Network access configuration:
	// This is necessary for external applications (like SigNoz) to connect to ClickHouse
	// Allow any IPv4 and IPv6 address
	users := pulumi.Map{
		vars.DefaultUsername + "/password": pulumi.Map{
			"valueFrom": pulumi.Map{
				"secretKeyRef": pulumi.Map{
					"name": createdSecret.Metadata.Name(),
					"key":  pulumi.String(vars.ClickhousePasswordKey),
				},
			},
		},
		vars.DefaultUsername + "/networks/ip": pulumi.Array{
			pulumi.String("0.0.0.0/0"), // Allow connections from any IPv4 address
			pulumi.String("::/0"),      // Allow connections from any IPv6 address
		},
	}

	// Determine cluster layout
	isClustered := spec.Cluster != nil && spec.Cluster.IsEnabled
	shardCount := 1
	replicaCount := 1
	if isClustered {
		shardCount = int(spec.Cluster.ShardCount)
		replicaCount = int(spec.Cluster.ReplicaCount)
	} else {
		replicaCount = int(spec.Container.Replicas)
	}

	// Build cluster configuration
	cluster := &altinityv1.ClickHouseInstallationSpecConfigurationClustersArgs{
		Name: pulumi.String(clusterName),
		Layout: &altinityv1.ClickHouseInstallationSpecConfigurationClustersLayoutArgs{
			ShardsCount:   pulumi.IntPtr(shardCount),
			ReplicasCount: pulumi.IntPtr(replicaCount),
		},
		// Enable inter-cluster authentication for distributed queries
		// Required for multi-node clusters to authenticate with each other
		Secret: &altinityv1.ClickHouseInstallationSpecConfigurationClustersSecretArgs{
			Auto: pulumi.String("true"),
		},
	}

	config := &altinityv1.ClickHouseInstallationSpecConfigurationArgs{
		Users: users,
		Clusters: altinityv1.ClickHouseInstallationSpecConfigurationClustersArray{
			cluster,
		},
	}

	// Add logging configuration if specified (overrides Altinity operator's debug default)
	// Uses the enum's string representation (e.g., "information", "debug", "trace")
	if spec.Logging != nil {
		logLevel := spec.Logging.Level.String()
		config.Files = pulumi.Map{
			"config.d/logging.xml": pulumi.Sprintf(`<clickhouse>
    <logger>
        <level>%s</level>
    </logger>
</clickhouse>`, logLevel),
		}
	}

	// Add coordination configuration for clustered deployments
	if isClustered {
		config.Zookeeper = buildCoordinationConfig(spec)
	}

	return config
}

// buildDefaults generates the defaults section referencing templates
func buildDefaults() *altinityv1.ClickHouseInstallationSpecDefaultsArgs {
	return &altinityv1.ClickHouseInstallationSpecDefaultsArgs{
		Templates: &altinityv1.ClickHouseInstallationSpecDefaultsTemplatesArgs{
			PodTemplate:             pulumi.String("clickhouse-pod"),
			DataVolumeClaimTemplate: pulumi.String("data-volume"),
		},
	}
}

// buildTemplates generates the templates section (pod and volume claim templates)
func buildTemplates(
	spec *clickhousekubernetesv1.ClickHouseKubernetesSpec,
	version string,
) *altinityv1.ClickHouseInstallationSpecTemplatesArgs {
	return &altinityv1.ClickHouseInstallationSpecTemplatesArgs{
		PodTemplates:         buildPodTemplates(spec, version),
		VolumeClaimTemplates: buildVolumeClaimTemplates(spec),
	}
}

// buildPodTemplates generates the pod template with container resources
func buildPodTemplates(
	spec *clickhousekubernetesv1.ClickHouseKubernetesSpec,
	version string,
) altinityv1.ClickHouseInstallationSpecTemplatesPodTemplatesArray {
	resources := spec.Container.Resources

	podSpec := pulumi.Map{
		"containers": pulumi.Array{
			pulumi.Map{
				"name":  pulumi.String("clickhouse"),
				"image": pulumi.Sprintf("clickhouse/clickhouse-server:%s", version),
				"resources": pulumi.Map{
					"requests": pulumi.Map{
						"cpu":    pulumi.String(resources.Requests.Cpu),
						"memory": pulumi.String(resources.Requests.Memory),
					},
					"limits": pulumi.Map{
						"cpu":    pulumi.String(resources.Limits.Cpu),
						"memory": pulumi.String(resources.Limits.Memory),
					},
				},
			},
		},
	}

	return altinityv1.ClickHouseInstallationSpecTemplatesPodTemplatesArray{
		&altinityv1.ClickHouseInstallationSpecTemplatesPodTemplatesArgs{
			Name: pulumi.String("clickhouse-pod"),
			Spec: podSpec,
		},
	}
}

// buildVolumeClaimTemplates generates the persistence volume claim template
func buildVolumeClaimTemplates(
	spec *clickhousekubernetesv1.ClickHouseKubernetesSpec,
) altinityv1.ClickHouseInstallationSpecTemplatesVolumeClaimTemplatesArray {
	diskSize := "1Gi" // minimal default
	if spec.Container.IsPersistenceEnabled && spec.Container.DiskSize != "" {
		diskSize = spec.Container.DiskSize
	}

	volumeClaimSpec := pulumi.Map{
		"accessModes": pulumi.Array{
			pulumi.String("ReadWriteOnce"),
		},
		"resources": pulumi.Map{
			"requests": pulumi.Map{
				"storage": pulumi.String(diskSize),
			},
		},
	}

	return altinityv1.ClickHouseInstallationSpecTemplatesVolumeClaimTemplatesArray{
		&altinityv1.ClickHouseInstallationSpecTemplatesVolumeClaimTemplatesArgs{
			Name: pulumi.String("data-volume"),
			Spec: volumeClaimSpec,
		},
	}
}
