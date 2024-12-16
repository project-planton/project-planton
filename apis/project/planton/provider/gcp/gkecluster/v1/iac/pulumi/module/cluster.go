package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/vars"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// cluster creates a GKE cluster by setting up the necessary Google Cloud projects, network resources,
// and enabling required APIs. It also configures various aspects of the cluster, including autoscaling,
// network policies, and logging.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - createdFolder: The Google Cloud Folder where the projects for the GKE cluster will be grouped.
//
// Returns:
// - *container.Cluster: A pointer to the created GKE Cluster object.
// - error: An error object if there is any issue during the cluster creation.
//
// The function performs the following steps:
//  1. Generates a random suffix to ensure the cluster project ID is unique on Google Cloud.
//  2. Creates the Google Cloud project for the GKE cluster.
//  3. If shared VPC is required, creates a separate network project with a unique ID.
//  4. Enables necessary APIs for the cluster and network projects.
//  5. Creates the VPC network, subnetwork, firewall rules, and router.
//  6. Configures NAT for the router with an external IP address.
//  7. Creates shared VPC IAM resources if shared VPC is enabled.
//  8. Configures the cluster with autoscaling, network policies, logging, and other settings.
//  9. Exports important attributes of the created resources, such as network self-link, subnetwork self-link,
//     firewall self-link, router self-link, NAT IP address, and cluster name.
func cluster(ctx *pulumi.Context, locals *localz.Locals, gcpProvider *gcp.Provider) (*container.Cluster, error) {

	//keep track of all the apis enabled to add as dependencies
	createdGoogleApiResources := make([]pulumi.Resource, 0)

	//enable apis for container cluster project
	for _, api := range vars.ContainerClusterProjectApis {
		addedProjectService, err := projects.NewService(ctx,
			fmt.Sprintf("container-cluster-%s", api),
			&projects.ServiceArgs{
				Project:                  pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
				DisableDependentServices: pulumi.BoolPtr(true),
				Service:                  pulumi.String(api),
			}, pulumi.Provider(gcpProvider))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to enable %s api for container cluster project", api)
		}
		createdGoogleApiResources = append(createdGoogleApiResources, addedProjectService)
	}

	//create vpc network
	createdNetwork, err := compute.NewNetwork(ctx,
		"vpc",
		&compute.NetworkArgs{
			Project:               pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
			AutoCreateSubnetworks: pulumi.BoolPtr(false),
		}, pulumi.Provider(gcpProvider),
		pulumi.DependsOn(createdGoogleApiResources))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create network")
	}

	//export network self-link
	ctx.Export(outputs.NETWORK_SELF_LINK, createdNetwork.SelfLink)

	//create subnetwork
	createdSubNetwork, err := compute.NewSubnetwork(ctx, "sub-network", &compute.SubnetworkArgs{
		Name:                  pulumi.String(locals.GkeCluster.Metadata.Name),
		Project:               pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
		Network:               createdNetwork.ID(),
		Region:                pulumi.String(locals.GkeCluster.Spec.Region),
		IpCidrRange:           pulumi.String(vars.SubNetworkCidr),
		PrivateIpGoogleAccess: pulumi.BoolPtr(true),
		//these two ranges will be referred in the cluster input
		SecondaryIpRanges: &compute.SubnetworkSecondaryIpRangeArray{
			&compute.SubnetworkSecondaryIpRangeArgs{
				RangeName:   pulumi.String(locals.KubernetesPodSecondaryIpRangeName),
				IpCidrRange: pulumi.String(vars.KubernetesPodSecondaryIpRange),
			},
			&compute.SubnetworkSecondaryIpRangeArgs{
				RangeName:   pulumi.String(locals.KubernetesServiceSecondaryIpRangeName),
				IpCidrRange: pulumi.String(vars.KubernetesServiceSecondaryIpRange),
			},
		},
	}, pulumi.Parent(createdNetwork))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subnetwork")
	}

	//export subnetwork self-link
	ctx.Export(outputs.SUB_NETWORK_SELF_LINK, createdSubNetwork.SelfLink)

	//create firewall
	createdFirewall, err := compute.NewFirewall(ctx, "firewall", &compute.FirewallArgs{
		Name:    pulumi.Sprintf("%s-gke-webhook", locals.GkeCluster.Metadata.Name),
		Project: pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
		Network: createdNetwork.Name,
		SourceRanges: pulumi.StringArray{
			pulumi.String(vars.ApiServerIpCidr),
		},
		Allows: compute.FirewallAllowArray{
			&compute.FirewallAllowArgs{
				Protocol: pulumi.String("tcp"),
				Ports: pulumi.StringArray{
					pulumi.String(vars.ApiServerWebhookPort),
					pulumi.String(vars.IstioPilotWebhookPort),
				},
			},
		},
		TargetTags: pulumi.StringArray{
			pulumi.String(locals.NetworkTag),
		},
	}, pulumi.Parent(createdNetwork))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create firewall")
	}

	//export firewall self-link
	ctx.Export(outputs.GKE_WEBHOOKS_FIREWALL_SELF_LINK, createdFirewall.SelfLink)

	//create router
	createdRouter, err := compute.NewRouter(ctx,
		"router",
		&compute.RouterArgs{
			Name:    pulumi.String(locals.GkeCluster.Metadata.Name),
			Network: createdNetwork.SelfLink,
			Region:  pulumi.String(locals.GkeCluster.Spec.Region),
			Project: pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
		}, pulumi.Parent(createdNetwork))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create router")
	}

	//export router self-link
	ctx.Export(outputs.ROUTER_SELF_LINK, createdRouter.SelfLink)

	//create ip-address for router nat
	createdRouterNatIp, err := compute.NewAddress(ctx,
		"router-nat-ip",
		&compute.AddressArgs{
			Name:        pulumi.Sprintf("%s-router-nat", locals.GkeCluster.Metadata.Name),
			Project:     pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
			Region:      createdRouter.Region,
			AddressType: pulumi.String("EXTERNAL"),
			Labels:      pulumi.ToStringMap(locals.GcpLabels),
		}, pulumi.Parent(createdRouter))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add new compute address")
	}

	//export router nat ip
	ctx.Export(outputs.EXTERNAL_NAT_IP, createdRouterNatIp.Address)

	//create router nat
	createdRouterNat, err := compute.NewRouterNat(ctx,
		"nat-router",
		&compute.RouterNatArgs{
			Name:                          pulumi.String(locals.GkeCluster.Metadata.Name),
			Router:                        createdRouter.Name,
			Region:                        createdRouter.Region,
			Project:                       pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
			NatIpAllocateOption:           pulumi.String("MANUAL_ONLY"),
			NatIps:                        pulumi.StringArray{createdRouterNatIp.SelfLink},
			SourceSubnetworkIpRangesToNat: pulumi.String("ALL_SUBNETWORKS_ALL_IP_RANGES"),
		}, pulumi.Parent(createdRouter))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create network router nat")
	}

	//export router nat name
	ctx.Export(outputs.ROUTER_NAT_NAME, createdRouterNat.Name)

	clusterAutoscalingArgs := &container.ClusterClusterAutoscalingArgs{
		Enabled: pulumi.Bool(false),
	}

	//determine autoscaling input based on gke-cluster input spec
	if locals.GkeCluster.Spec.ClusterAutoscalingConfig != nil &&
		locals.GkeCluster.Spec.ClusterAutoscalingConfig.IsEnabled {
		clusterAutoscalingArgs = &container.ClusterClusterAutoscalingArgs{
			Enabled:            pulumi.Bool(true),
			AutoscalingProfile: pulumi.String("OPTIMIZE_UTILIZATION"),
			ResourceLimits: container.ClusterClusterAutoscalingResourceLimitArray{
				container.ClusterClusterAutoscalingResourceLimitArgs{
					ResourceType: pulumi.String("cpu"),
					Minimum:      pulumi.Int(locals.GkeCluster.Spec.ClusterAutoscalingConfig.CpuMinCores),
					Maximum:      pulumi.Int(locals.GkeCluster.Spec.ClusterAutoscalingConfig.CpuMaxCores),
				},
				container.ClusterClusterAutoscalingResourceLimitArgs{
					ResourceType: pulumi.String("memory"),
					Minimum:      pulumi.Int(locals.GkeCluster.Spec.ClusterAutoscalingConfig.MemoryMinGb),
					Maximum:      pulumi.Int(locals.GkeCluster.Spec.ClusterAutoscalingConfig.MemoryMaxGb),
				},
			},
		}
	}

	//create container cluster
	createdCluster, err := container.NewCluster(ctx,
		"cluster",
		&container.ClusterArgs{
			Name:                  pulumi.String(locals.GkeCluster.Metadata.Name),
			Project:               pulumi.String(locals.GkeCluster.Spec.ClusterProjectId),
			Location:              pulumi.String(locals.GkeCluster.Spec.Zone),
			Network:               createdNetwork.SelfLink,
			Subnetwork:            createdSubNetwork.SelfLink,
			RemoveDefaultNodePool: pulumi.Bool(true),
			DeletionProtection:    pulumi.Bool(false),
			WorkloadIdentityConfig: container.ClusterWorkloadIdentityConfigPtrInput(
				&container.ClusterWorkloadIdentityConfigArgs{
					WorkloadPool: pulumi.Sprintf("%s.svc.id.goog", locals.GkeCluster.Spec.ClusterProjectId),
				}),
			//warning: cluster is not coming into ready state with value set to 0
			InitialNodeCount: pulumi.Int(1),
			ReleaseChannel: container.ClusterReleaseChannelPtrInput(
				&container.ClusterReleaseChannelArgs{
					Channel: pulumi.String(vars.GkeReleaseChannel),
				}),
			VerticalPodAutoscaling: container.ClusterVerticalPodAutoscalingPtrInput(
				&container.ClusterVerticalPodAutoscalingArgs{Enabled: pulumi.Bool(true)}),
			AddonsConfig: container.ClusterAddonsConfigPtrInput(&container.ClusterAddonsConfigArgs{
				HorizontalPodAutoscaling: container.ClusterAddonsConfigHorizontalPodAutoscalingPtrInput(
					&container.ClusterAddonsConfigHorizontalPodAutoscalingArgs{
						Disabled: pulumi.Bool(false)}),
				HttpLoadBalancing: container.ClusterAddonsConfigHttpLoadBalancingPtrInput(
					&container.ClusterAddonsConfigHttpLoadBalancingArgs{
						Disabled: pulumi.Bool(true)}),
				IstioConfig: container.ClusterAddonsConfigIstioConfigPtrInput(
					&container.ClusterAddonsConfigIstioConfigArgs{
						Disabled: pulumi.Bool(true)}),
				NetworkPolicyConfig: container.ClusterAddonsConfigNetworkPolicyConfigPtrInput(
					&container.ClusterAddonsConfigNetworkPolicyConfigArgs{
						Disabled: pulumi.Bool(true)}),
			}),
			PrivateClusterConfig: container.ClusterPrivateClusterConfigPtrInput(&container.ClusterPrivateClusterConfigArgs{
				EnablePrivateEndpoint: pulumi.Bool(false),
				EnablePrivateNodes:    pulumi.Bool(true),
				MasterIpv4CidrBlock:   pulumi.String(vars.ApiServerIpCidr),
			}),
			IpAllocationPolicy: container.ClusterIpAllocationPolicyPtrInput(
				// setting this is mandatory for shared vpc setup
				&container.ClusterIpAllocationPolicyArgs{
					ClusterSecondaryRangeName:  pulumi.String(locals.KubernetesPodSecondaryIpRangeName),
					ServicesSecondaryRangeName: pulumi.String(locals.KubernetesServiceSecondaryIpRangeName),
				}),
			MasterAuthorizedNetworksConfig: container.ClusterMasterAuthorizedNetworksConfigPtrInput(
				&container.ClusterMasterAuthorizedNetworksConfigArgs{
					CidrBlocks: container.ClusterMasterAuthorizedNetworksConfigCidrBlockArray{
						container.ClusterMasterAuthorizedNetworksConfigCidrBlockArgs{
							CidrBlock:   pulumi.String(vars.ClusterMasterAuthorizedNetworksCidrBlock),
							DisplayName: pulumi.String(vars.ClusterMasterAuthorizedNetworksCidrBlockDescription),
						}},
				}),
			ClusterAutoscaling: clusterAutoscalingArgs,
			//todo: disabling billing export temporarily
			//ResourceUsageExportConfig: container.ClusterResourceUsageExportConfigPtrInput(&container.ClusterResourceUsageExportConfigArgs{
			//	BigqueryDestination: container.ClusterResourceUsageExportConfigBigqueryDestinationArgs{
			//		DatasetId: pulumi.String(input.UsageMeteringDatasetId)},
			//	EnableNetworkEgressMetering:       pulumi.Bool(false),
			//	EnableResourceConsumptionMetering: pulumi.Bool(true),
			//}),
			LoggingConfig: container.ClusterLoggingConfigPtrInput(
				&container.ClusterLoggingConfigArgs{
					EnableComponents: pulumi.ToStringArray(locals.ContainerClusterLoggingComponentList),
				}),
		},
		pulumi.Provider(gcpProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add container cluster")
	}

	//export cluster attributes
	ctx.Export(outputs.CLUSTER_ENDPOINT, createdCluster.Endpoint)
	ctx.Export(outputs.CLUSTER_CA_DATA, createdCluster.MasterAuth.ClusterCaCertificate())

	return createdCluster, nil
}
