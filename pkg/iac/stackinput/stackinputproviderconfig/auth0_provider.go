package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/pkg/fileutil"
	"sigs.k8s.io/yaml"
)

const (
	Auth0ProviderConfigKey  = "provider_config"
	auth0ProviderConfigYaml = "auth0-provider-config.yaml"
)

func AddAuth0ProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.Auth0ProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.Auth0ProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.Auth0ProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[Auth0ProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadAuth0ProviderConfig(dir string) (string, error) {
	path := dir + "/" + auth0ProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
