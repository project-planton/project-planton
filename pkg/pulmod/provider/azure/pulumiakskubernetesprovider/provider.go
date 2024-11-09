package pulumiakskubernetesprovider

import (
	"github.com/pkg/errors"
	azurecredentialv1 "github.com/project-planton/project-planton/apis/go/project/planton/credential/azurecredential/v1"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/containerservice"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithAddedClusterWithAzureCredentials returns kubernetes provider for the added AKS cluster based on the azure provider
func GetWithAddedClusterWithAzureCredentials(ctx *pulumi.Context,
	addedAksCluster *containerservice.KubernetesCluster,
	azureCredentialSpec *azurecredentialv1.AzureCredentialSpec,
	dependencies []pulumi.Resource,
	providerName string) (*kubernetes.Provider, error) {

	clusterCaCert := addedAksCluster.KubeConfigs.ApplyT(
		func(kubeConfigs []containerservice.KubernetesClusterKubeConfig) string {
			return *kubeConfigs[0].ClusterCaCertificate
		})

	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig: pulumi.Sprintf(AzureExecPluginKubeConfigTemplate,
				addedAksCluster.Fqdn,
				clusterCaCert,
				azureCredentialSpec.ClientId,
				azureCredentialSpec.ClientSecret,
				azureCredentialSpec.TenantId,
			),
		}, pulumi.DependsOn(dependencies))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
