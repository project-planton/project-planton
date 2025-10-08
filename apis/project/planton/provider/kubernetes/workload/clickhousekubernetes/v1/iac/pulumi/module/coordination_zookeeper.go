package module

import (
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	altinityv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/kubernetes/clickhouse/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// buildExternalZookeeperReference creates configuration for external ZooKeeper
// Connects to existing ZooKeeper infrastructure (legacy systems, shared with Kafka, etc.)
func buildExternalZookeeperReference(coordination *clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig) *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	if coordination.ExternalConfig == nil || len(coordination.ExternalConfig.Nodes) == 0 {
		// No external nodes specified, fallback to default Keeper
		return buildDefaultKeeperReference()
	}

	// Build node references from external config
	nodes := make(altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArray, len(coordination.ExternalConfig.Nodes))
	for i, node := range coordination.ExternalConfig.Nodes {
		nodes[i] = &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArgs{
			Host: pulumi.String(node),
		}
	}

	return &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs{
		Nodes: nodes,
	}
}

// buildCoordinationFromDeprecatedZookeeperField handles the deprecated zookeeper field
// Maintained for backward compatibility with existing manifests
//
// DEPRECATED: Use coordination field instead
// This function will be removed in v2 when the zookeeper field is removed from spec
func buildCoordinationFromDeprecatedZookeeperField(zookeeper *clickhousekubernetesv1.ClickHouseKubernetesZookeeperConfig) *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	// If external ZooKeeper is configured
	if zookeeper.UseExternal && len(zookeeper.Nodes) > 0 {
		nodes := make(altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArray, len(zookeeper.Nodes))
		for i, node := range zookeeper.Nodes {
			nodes[i] = &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArgs{
				Host: pulumi.String(node),
			}
		}
		return &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs{
			Nodes: nodes,
		}
	}

	// Use operator-managed ZooKeeper (legacy behavior)
	// This references a ZooKeeper service that should exist at "zookeeper:2181"
	return &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs{
		Nodes: altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArray{
			&altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArgs{
				Host: pulumi.String("zookeeper"),
				Port: pulumi.IntPtr(vars.ZookeeperPort),
			},
		},
	}
}
