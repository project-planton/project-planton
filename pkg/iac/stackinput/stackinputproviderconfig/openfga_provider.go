package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/pkg/fileutil"
	"sigs.k8s.io/yaml"
)

const (
	OpenFgaProviderConfigKey  = "provider_config"
	openfgaProviderConfigYaml = "openfga-provider-config.yaml"
)

func AddOpenFgaProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.OpenFgaProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.OpenFgaProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.OpenFgaProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[OpenFgaProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadOpenFgaProviderConfig(dir string) (string, error) {
	path := dir + "/" + openfgaProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
