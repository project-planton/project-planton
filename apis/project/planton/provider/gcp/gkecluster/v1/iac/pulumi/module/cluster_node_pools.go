package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// clusterNodePools creates node pools for the given GKE cluster based on the specifications provided.
// It iterates over each node pool specification and configures the node pool with autoscaling, management, and node settings.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - createdCluster: The GKE cluster for which the node pools are being created.
//
// Returns:
// - []pulumi.Resource: A slice of created node pool resources.
// - error: An error object if there is any issue during the node pool creation.
//
// The function performs the following steps:
//  1. Iterates over each node pool specification provided in the locals.
//  2. Creates a node pool with the specified configuration, including location, project, cluster, node count,
//     autoscaling, management, node configuration, and upgrade settings.
//  3. Adds OAuth scopes, machine type, labels, metadata, tags, and workload metadata configuration to the node config.
//  4. Sets node pool management options, such as auto-repair and auto-upgrade.
//  5. Configures upgrade settings for the node pool with max surge and max unavailable settings.
//  6. Handles errors and returns a slice of created node pool resources and any errors encountered.
func clusterNodePools(ctx *pulumi.Context,
	locals *localz.Locals,
	createdCluster *container.Cluster) ([]pulumi.Resource, error) {
	createdNodePoolResources := make([]pulumi.Resource, 0)

	for _, nodePoolSpec := range locals.GkeCluster.Spec.NodePools {
		createdNodePool, err := container.NewNodePool(ctx, nodePoolSpec.Name, &container.NodePoolArgs{
			Location:  pulumi.String(locals.GkeCluster.Spec.Zone),
			Project:   createdCluster.Project,
			Cluster:   createdCluster.Name,
			NodeCount: pulumi.Int(nodePoolSpec.MinNodeCount),
			Autoscaling: container.NodePoolAutoscalingPtrInput(&container.NodePoolAutoscalingArgs{
				MinNodeCount: pulumi.Int(nodePoolSpec.MinNodeCount),
				MaxNodeCount: pulumi.Int(nodePoolSpec.MaxNodeCount),
			}),
			Management: container.NodePoolManagementPtrInput(&container.NodePoolManagementArgs{
				AutoRepair:  pulumi.Bool(true),
				AutoUpgrade: pulumi.Bool(true),
			}),
			NodeConfig: &container.NodePoolNodeConfigArgs{
				Labels:      pulumi.ToStringMap(locals.GcpLabels),
				MachineType: pulumi.String(nodePoolSpec.MachineType),
				Metadata:    pulumi.StringMap{"disable-legacy-endpoints": pulumi.String("true")},
				OauthScopes: pulumi.StringArray{
					pulumi.String("https://www.googleapis.com/auth/monitoring"),
					pulumi.String("https://www.googleapis.com/auth/monitoring.write"),
					pulumi.String("https://www.googleapis.com/auth/devstorage.read_only"),
					pulumi.String("https://www.googleapis.com/auth/logging.write"),
				},
				Preemptible: pulumi.Bool(nodePoolSpec.IsSpotEnabled),
				Tags: pulumi.StringArray{
					pulumi.String(locals.NetworkTag),
				},
				WorkloadMetadataConfig: container.NodePoolNodeConfigWorkloadMetadataConfigPtrInput(
					&container.NodePoolNodeConfigWorkloadMetadataConfigArgs{
						Mode: pulumi.String("GKE_METADATA")}),
			},
			UpgradeSettings: container.NodePoolUpgradeSettingsPtrInput(&container.NodePoolUpgradeSettingsArgs{
				MaxSurge:       pulumi.Int(2),
				MaxUnavailable: pulumi.Int(1),
			}),
		},
			pulumi.Parent(createdCluster),
			pulumi.IgnoreChanges([]string{"nodeCount"}),
			pulumi.DeleteBeforeReplace(true),
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create node-pool")
		}

		createdNodePoolResources = append(createdNodePoolResources, createdNodePool)
	}

	return createdNodePoolResources, nil
}