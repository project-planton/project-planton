package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/pkg/fileutil"
	"sigs.k8s.io/yaml"
)

const (
	AzureProviderConfigKey  = "provider_config"
	azureProviderConfigYaml = "azure-provider-config.yaml"
)

func AddAzureProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.AzureProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.AzureProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.AzureProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[AzureProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadAzureProviderConfig(dir string) (string, error) {
	path := dir + "/" + azureProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
