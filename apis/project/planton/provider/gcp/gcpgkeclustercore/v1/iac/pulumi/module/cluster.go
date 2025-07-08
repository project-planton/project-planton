package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// cluster provisions the control‑plane (no node pools, networking assumed pre‑created).
func cluster(ctx *pulumi.Context, locals *Locals, gcpProvider *gcp.Provider) (*container.Cluster, error) {

	// ------ Private‑cluster settings ------------------------------------------------------------
	privateClusterCfg := &container.ClusterPrivateClusterConfigArgs{
		// Invert the enable_public_nodes flag: if public nodes requested,
		// do NOT enable private nodes.  The masters remain private either way.
		EnablePrivateNodes:    pulumi.Bool(!locals.GcpGkeClusterCore.Spec.EnablePublicNodes),
		EnablePrivateEndpoint: pulumi.Bool(false),
		MasterIpv4CidrBlock:   pulumi.String(locals.GcpGkeClusterCore.Spec.MasterIpv4CidrBlock),
	}

	// ------ Workload Identity -------------------------------------------------------------------
	var workloadIdentityCfg container.ClusterWorkloadIdentityConfigPtrInput
	if !locals.GcpGkeClusterCore.Spec.DisableWorkloadIdentity {
		workloadIdentityCfg = container.ClusterWorkloadIdentityConfigPtrInput(
			&container.ClusterWorkloadIdentityConfigArgs{
				WorkloadPool: pulumi.Sprintf("%s.svc.id.goog",
					locals.GcpGkeClusterCore.Spec.ProjectId.GetValue()),
			})
	}

	// ------ Network Policy ----------------------------------------------------------------------
	addonsCfg := container.ClusterAddonsConfigPtrInput(&container.ClusterAddonsConfigArgs{
		NetworkPolicyConfig: container.ClusterAddonsConfigNetworkPolicyConfigPtrInput(
			&container.ClusterAddonsConfigNetworkPolicyConfigArgs{
				// Disable if flag is set; otherwise enable enforcement (Calico).
				Disabled: pulumi.Bool(locals.GcpGkeClusterCore.Spec.DisableNetworkPolicy),
			}),
	})

	// ------ Cluster ElasticOperatorKubernetes --------------------------------------------------------------------
	createdCluster, err := container.NewCluster(ctx,
		"cluster",
		&container.ClusterArgs{
			Name:                  pulumi.String(locals.GcpGkeClusterCore.Metadata.Name),
			Project:               pulumi.String(locals.GcpGkeClusterCore.Spec.ProjectId.GetValue()),
			Location:              pulumi.String(locals.GcpGkeClusterCore.Spec.Location),
			Subnetwork:            pulumi.String(locals.GcpGkeClusterCore.Spec.SubnetworkSelfLink.GetValue()),
			RemoveDefaultNodePool: pulumi.Bool(true),
			DeletionProtection:    pulumi.Bool(false),
			InitialNodeCount:      pulumi.Int(1), // required by API even if we delete later
			PrivateClusterConfig:  privateClusterCfg,
			IpAllocationPolicy: container.ClusterIpAllocationPolicyPtrInput(
				&container.ClusterIpAllocationPolicyArgs{
					ClusterSecondaryRangeName:  pulumi.String(locals.GcpGkeClusterCore.Spec.ClusterSecondaryRangeName.GetValue()),
					ServicesSecondaryRangeName: pulumi.String(locals.GcpGkeClusterCore.Spec.ServicesSecondaryRangeName.GetValue()),
				}),
			ReleaseChannel: container.ClusterReleaseChannelPtrInput(
				&container.ClusterReleaseChannelArgs{
					Channel: pulumi.String(locals.ReleaseChannelStr),
				}),
			WorkloadIdentityConfig: workloadIdentityCfg,
			AddonsConfig:           addonsCfg,
		},
		pulumi.Provider(gcpProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create container cluster")
	}

	// ------ Outputs -----------------------------------------------------------------------------
	ctx.Export(OpEndpoint, createdCluster.Endpoint)
	ctx.Export(OpClusterCaCertificate, createdCluster.MasterAuth.ClusterCaCertificate())
	if workloadIdentityCfg != nil {
		ctx.Export(OpWorkloadIdentityPool,
			pulumi.Sprintf("%s.svc.id.goog", locals.GcpGkeClusterCore.Spec.ProjectId.GetValue()))
	}

	return createdCluster, nil
}
