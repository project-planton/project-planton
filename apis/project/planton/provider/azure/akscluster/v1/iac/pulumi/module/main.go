package module

import (
	"github.com/pkg/errors"
	aksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/azure/akscluster/v1"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *aksclusterv1.AksClusterStackInput) error {
	azureCredential := stackInput.AzureCredential
	//create azure provider using the credentials from the input
	_, err := azure.NewProvider(ctx,
		"azure",
		&azure.ProviderArgs{
			ClientId:       pulumi.String(azureCredential.ClientId),
			ClientSecret:   pulumi.String(azureCredential.ClientSecret),
			SubscriptionId: pulumi.String(azureCredential.SubscriptionId),
			TenantId:       pulumi.String(azureCredential.TenantId),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create azure provider")
	}
	return nil
}
