package module

import (
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	altinityv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/altinityoperator/kubernetes/clickhouse/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// buildAutoManagedKeeperReference creates configuration for auto-managed ClickHouse Keeper
// The module creates a ClickHouseKeeperInstallation CRD named "keeper"
// which creates a service following pattern "keeper-<name>" = "keeper-keeper"
func buildAutoManagedKeeperReference(coordination *clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig) *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	// Service name pattern: keeper-<chk-name>
	// Since we create ClickHouseKeeperInstallation named "keeper"
	// The service will be named "keeper-keeper"
	keeperHost := "keeper-keeper"

	return &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs{
		Nodes: altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArray{
			&altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArgs{
				Host: pulumi.String(keeperHost),
				Port: pulumi.IntPtr(vars.ZookeeperPort),
			},
		},
	}
}

// buildExternalKeeperReference creates configuration for external ClickHouse Keeper
// Connects to existing ClickHouse Keeper infrastructure
func buildExternalKeeperReference(coordination *clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig) *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	if coordination.ExternalConfig == nil || len(coordination.ExternalConfig.Nodes) == 0 {
		// No external nodes specified, fallback to default
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

// buildDefaultKeeperReference creates a reference to default auto-managed ClickHouse Keeper
// Used when no coordination configuration is specified
//
// The module auto-creates ClickHouseKeeperInstallation named "keeper"
// which creates service "keeper-keeper" following the pattern "keeper-<name>"
func buildDefaultKeeperReference() *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	return &altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs{
		Nodes: altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArray{
			&altinityv1.ClickHouseInstallationSpecConfigurationZookeeperNodesArgs{
				Host: pulumi.String("keeper-keeper"),
				Port: pulumi.IntPtr(vars.ZookeeperPort),
			},
		},
	}
}
