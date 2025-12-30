package providerconfig

import (
	"github.com/pkg/errors"
	auth0provider "github.com/project-planton/project-planton/apis/org/project_planton/provider/auth0"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddAuth0ProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	auth0ProviderConfig := new(auth0provider.Auth0ProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.Auth0ProviderConfigKey, auth0ProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	providerConfigEnvVars["AUTH0_DOMAIN"] = auth0ProviderConfig.Domain
	providerConfigEnvVars["AUTH0_CLIENT_ID"] = auth0ProviderConfig.ClientId
	providerConfigEnvVars["AUTH0_CLIENT_SECRET"] = auth0ProviderConfig.ClientSecret

	return providerConfigEnvVars, nil
}
