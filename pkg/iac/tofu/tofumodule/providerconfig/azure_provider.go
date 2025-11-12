package providerconfig

import (
	"github.com/pkg/errors"
	azure "github.com/project-planton/project-planton/apis/org/project_planton/provider/azure"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddAzureProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	azureProviderConfig := new(azure.AzureProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.AzureProviderConfigKey, azureProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	providerConfigEnvVars["ARM_CLIENT_ID"] = azureProviderConfig.ClientId
	providerConfigEnvVars["ARM_CLIENT_SECRET"] = azureProviderConfig.ClientSecret
	providerConfigEnvVars["ARM_TENANT_ID"] = azureProviderConfig.TenantId
	providerConfigEnvVars["ARM_SUBSCRIPTION_ID"] = azureProviderConfig.SubscriptionId

	return providerConfigEnvVars, nil
}
