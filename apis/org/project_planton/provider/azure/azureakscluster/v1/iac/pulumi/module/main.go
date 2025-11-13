package module

import (
	"fmt"

	"github.com/pkg/errors"
	azureaksclusterv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/azure/azureakscluster/v1"
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	azurenative "github.com/pulumi/pulumi-azure-native-sdk/sdk/v2/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *azureaksclusterv1.AzureAksClusterStackInput) error {
	azureProviderConfig := stackInput.ProviderConfig
	
	// Create Azure provider using the credentials from the input
	provider, err := azurenative.NewProvider(ctx,
		"azure",
		&azurenative.ProviderArgs{
			ClientId:       pulumi.String(azureProviderConfig.ClientId),
			ClientSecret:   pulumi.String(azureProviderConfig.ClientSecret),
			SubscriptionId: pulumi.String(azureProviderConfig.SubscriptionId),
			TenantId:       pulumi.String(azureProviderConfig.TenantId),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create azure provider")
	}

	// Get inputs
	target := stackInput.Target
	spec := target.Spec
	
	// Create resource group name from metadata
	resourceGroupName := fmt.Sprintf("rg-%s", target.Metadata.Name)
	
	// Create Resource Group
	resourceGroup, err := resources.NewResourceGroup(ctx, resourceGroupName, &resources.ResourceGroupArgs{
		ResourceGroupName: pulumi.String(resourceGroupName),
		Location:          pulumi.String(spec.Region),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create resource group")
	}

	// Determine network plugin
	networkPlugin := "azure"
	if spec.NetworkPlugin == azureaksclusterv1.AzureAksClusterNetworkPlugin_KUBENET {
		networkPlugin = "kubenet"
	}

	// Build authorized IP ranges
	var authorizedIpRanges pulumi.StringArray
	for _, cidr := range spec.AuthorizedIpRanges {
		authorizedIpRanges = append(authorizedIpRanges, pulumi.String(cidr))
	}

	// Determine AAD RBAC configuration
	aadEnabled := !spec.DisableAzureAdRbac

	// Build AKS cluster arguments
	aksClusterArgs := &containerservice.ManagedClusterArgs{
		ResourceGroupName: resourceGroup.Name,
		ResourceName:      pulumi.String(target.Metadata.Name),
		Location:          pulumi.String(spec.Region),
		
		// Kubernetes version
		KubernetesVersion: pulumi.String(spec.KubernetesVersion),
		
		// DNS prefix
		DnsPrefix: pulumi.Sprintf("%s-dns", target.Metadata.Name),
		
		// Identity - Use System-Assigned Managed Identity
		Identity: &containerservice.ManagedClusterIdentityArgs{
			Type: pulumi.String("SystemAssigned"),
		},
		
		// Network Profile
		NetworkProfile: &containerservice.ContainerServiceNetworkProfileArgs{
			NetworkPlugin: pulumi.String(networkPlugin),
			LoadBalancerSku: pulumi.String("standard"),
			ServiceCidr: pulumi.String("10.0.0.0/16"),
			DnsServiceIp: pulumi.String("10.0.0.10"),
		},
		
		// Default node pool (system pool)
		AgentPoolProfiles: containerservice.ManagedClusterAgentPoolProfileArray{
			&containerservice.ManagedClusterAgentPoolProfileArgs{
				Name:              pulumi.String("system"),
				Count:             pulumi.Int(3),
				VmSize:            pulumi.String("Standard_D2s_v3"),
				Mode:              pulumi.String("System"),
				OsType:            pulumi.String("Linux"),
				VnetSubnetID:      pulumi.String(spec.VnetSubnetId.GetValue()),
				EnableAutoScaling: pulumi.Bool(true),
				MinCount:          pulumi.Int(3),
				MaxCount:          pulumi.Int(5),
			},
		},
		
		// API Server Access Profile
		ApiServerAccessProfile: &containerservice.ManagedClusterAPIServerAccessProfileArgs{
			EnablePrivateCluster:       pulumi.Bool(spec.PrivateClusterEnabled),
			AuthorizedIPRanges:         authorizedIpRanges,
			EnablePrivateClusterPublicFQDN: pulumi.Bool(!spec.PrivateClusterEnabled),
		},
		
		// Azure AD Integration
		AadProfile: &containerservice.ManagedClusterAADProfileArgs{
			Managed:           pulumi.Bool(aadEnabled),
			EnableAzureRBAC:   pulumi.Bool(aadEnabled),
		},
		
		// Add-ons
		AddonProfiles: containerservice.ManagedClusterAddonProfileMap{},
	}

	// Add monitoring addon if Log Analytics workspace is provided
	if spec.LogAnalyticsWorkspaceId != "" {
		aksClusterArgs.AddonProfiles["omsagent"] = &containerservice.ManagedClusterAddonProfileArgs{
			Enabled: pulumi.Bool(true),
			Config: pulumi.StringMap{
				"logAnalyticsWorkspaceResourceID": pulumi.String(spec.LogAnalyticsWorkspaceId),
			},
		}
	}

	// Create AKS cluster
	aksCluster, err := containerservice.NewManagedCluster(ctx, target.Metadata.Name, aksClusterArgs, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create AKS cluster")
	}

	// Get kubeconfig
	listCredentials := containerservice.ListManagedClusterUserCredentialsOutput(ctx, containerservice.ListManagedClusterUserCredentialsOutputArgs{
		ResourceGroupName: resourceGroup.Name,
		ResourceName:      aksCluster.Name,
	})

	kubeconfig := listCredentials.Kubeconfigs().Index(pulumi.Int(0)).Value()

	// Export outputs aligned to AzureAksClusterStackOutputs
	ctx.Export(OpApiServerEndpoint, aksCluster.Fqdn)
	ctx.Export(OpClusterResourceId, aksCluster.ID())
	ctx.Export(OpClusterKubeconfig, kubeconfig)
	ctx.Export(OpManagedIdentityPrincipalId, aksCluster.IdentityProfile.MapIndex(pulumi.String("kubeletidentity")).ObjectID())

	return nil
}
