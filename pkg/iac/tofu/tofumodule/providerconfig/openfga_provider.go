package providerconfig

import (
	"github.com/pkg/errors"
	openfgaprovider "github.com/plantonhq/project-planton/apis/org/project_planton/provider/openfga"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddOpenFgaProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	openfgaProviderConfig := new(openfgaprovider.OpenFgaProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.OpenFgaProviderConfigKey, openfgaProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	// This means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	// Required field
	providerConfigEnvVars["FGA_API_URL"] = openfgaProviderConfig.ApiUrl

	// Optional fields - only set if they have values
	if openfgaProviderConfig.ApiToken != "" {
		providerConfigEnvVars["FGA_API_TOKEN"] = openfgaProviderConfig.ApiToken
	}
	if openfgaProviderConfig.ClientId != "" {
		providerConfigEnvVars["FGA_CLIENT_ID"] = openfgaProviderConfig.ClientId
	}
	if openfgaProviderConfig.ClientSecret != "" {
		providerConfigEnvVars["FGA_CLIENT_SECRET"] = openfgaProviderConfig.ClientSecret
	}
	if openfgaProviderConfig.ApiTokenIssuer != "" {
		providerConfigEnvVars["FGA_API_TOKEN_ISSUER"] = openfgaProviderConfig.ApiTokenIssuer
	}
	if openfgaProviderConfig.ApiScopes != "" {
		providerConfigEnvVars["FGA_API_SCOPES"] = openfgaProviderConfig.ApiScopes
	}
	if openfgaProviderConfig.ApiAudience != "" {
		providerConfigEnvVars["FGA_API_AUDIENCE"] = openfgaProviderConfig.ApiAudience
	}

	return providerConfigEnvVars, nil
}
