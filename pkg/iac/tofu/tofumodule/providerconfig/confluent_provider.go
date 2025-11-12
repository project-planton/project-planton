package providerconfig

import (
	"github.com/pkg/errors"
	confluent "github.com/project-planton/project-planton/apis/org/project_planton/provider/confluent"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddConfluentProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	confluentProviderConfig := new(confluent.ConfluentProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.ConfluentProviderConfigKey, confluentProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	providerConfigEnvVars["CONFLUENT_API_KEY"] = confluentProviderConfig.ApiKey
	providerConfigEnvVars["CONFLUENT_API_SECRET"] = confluentProviderConfig.ApiSecret

	return providerConfigEnvVars, nil
}
