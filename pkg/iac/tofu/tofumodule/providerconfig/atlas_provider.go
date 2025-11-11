package providerconfig

import (
	"github.com/pkg/errors"
	atlasprovider "github.com/project-planton/project-planton/apis/org/project-planton/provider/atlas"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddAtlasProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	atlasProviderConfig := new(atlasprovider.AtlasProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.AtlasProviderConfigKey, atlasProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	providerConfigEnvVars["MONGODB_ATLAS_PUBLIC_KEY"] = atlasProviderConfig.PublicKey
	providerConfigEnvVars["MONGODB_ATLAS_PRIVATE_KEY"] = atlasProviderConfig.PrivateKey

	return providerConfigEnvVars, nil
}
