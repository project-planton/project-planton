package module

import (
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	altinityv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/altinityoperator/kubernetes/clickhouse/v1"
)

// buildCoordinationConfig generates the coordination (ZooKeeper/Keeper) configuration
// Handles both new 'coordination' field and deprecated 'zookeeper' field for backward compatibility
//
// Priority Order:
//  1. coordination field (new API)
//  2. zookeeper field (deprecated, backward compat)
//  3. default (auto-managed ClickHouse Keeper)
func buildCoordinationConfig(spec *clickhousekubernetesv1.ClickHouseKubernetesSpec) *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	// Priority: coordination field takes precedence over deprecated zookeeper field
	if spec.Coordination != nil {
		return buildCoordinationFromNewField(spec.Coordination)
	}

	// Backward compatibility: fall back to deprecated zookeeper field
	if spec.Zookeeper != nil {
		return buildCoordinationFromDeprecatedZookeeperField(spec.Zookeeper)
	}

	// Default: auto-managed ClickHouse Keeper with single replica
	// Note: ClickHouse Keeper must be deployed separately via ClickHouseKeeperInstallation
	// This references the Keeper service that should exist
	return buildDefaultKeeperReference()
}

// buildCoordinationFromNewField handles the new coordination configuration
// Routes to appropriate builder based on coordination type
func buildCoordinationFromNewField(coordination *clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig) *altinityv1.ClickHouseInstallationSpecConfigurationZookeeperArgs {
	coordinationType := coordination.Type

	// Default unspecified to keeper
	if coordinationType == clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_unspecified {
		coordinationType = clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_keeper
	}

	switch coordinationType {
	case clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_keeper:
		return buildAutoManagedKeeperReference(coordination)

	case clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_external_keeper:
		return buildExternalKeeperReference(coordination)

	case clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_external_zookeeper:
		return buildExternalZookeeperReference(coordination)
	}

	// Fallback to default
	return buildDefaultKeeperReference()
}
