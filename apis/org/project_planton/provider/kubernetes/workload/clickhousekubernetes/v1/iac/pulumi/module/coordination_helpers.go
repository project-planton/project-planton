package module

import (
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/workload/clickhousekubernetes/v1"
)

// shouldCreateClickHouseKeeper determines if we should auto-create ClickHouseKeeperInstallation
// Returns true only when:
// - Clustering is enabled (requires coordination)
// - coordination.type = keeper (auto-managed)
func shouldCreateClickHouseKeeper(spec *clickhousekubernetesv1.ClickHouseKubernetesSpec) bool {
	// Only create Keeper for clustered deployments
	if spec.Cluster == nil || !spec.Cluster.IsEnabled {
		return false
	}

	// Check if using new coordination field
	if spec.Coordination != nil {
		coordinationType := spec.Coordination.Type
		// Default unspecified to keeper
		if coordinationType == clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_unspecified {
			coordinationType = clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_keeper
		}
		// Only create for keeper type (not external_keeper or external_zookeeper)
		return coordinationType == clickhousekubernetesv1.ClickHouseKubernetesCoordinationConfig_keeper
	}

	// If using deprecated zookeeper field, don't auto-create Keeper
	if spec.Zookeeper != nil {
		return false
	}

	// Default behavior: clustered without explicit coordination = create Keeper
	return true
}

// getKeeperConfig extracts the keeper configuration from spec
// Returns nil if not specified (will use defaults)
func getKeeperConfig(spec *clickhousekubernetesv1.ClickHouseKubernetesSpec) *clickhousekubernetesv1.ClickHouseKubernetesKeeperConfig {
	if spec.Coordination != nil && spec.Coordination.KeeperConfig != nil {
		return spec.Coordination.KeeperConfig
	}
	return nil
}
