package module

import (
	"github.com/pkg/errors"
	azurenatgatewayv1 "github.com/project-planton/project-planton/apis/project/planton/provider/azure/azurenatgateway/v1"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *azurenatgatewayv1.AzureNatGatewayStackInput) error {
	azureProviderConfig := stackInput.ProviderConfig
	//create azure provider using the credentials from the input
	_, err := azure.NewProvider(ctx,
		"azure",
		&azure.ProviderArgs{
			ClientId:       pulumi.String(azureProviderConfig.ClientId),
			ClientSecret:   pulumi.String(azureProviderConfig.ClientSecret),
			SubscriptionId: pulumi.String(azureProviderConfig.SubscriptionId),
			TenantId:       pulumi.String(azureProviderConfig.TenantId),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create azure provider")
	}
	return nil
}
