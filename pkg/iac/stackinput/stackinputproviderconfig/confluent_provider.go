package stackinputproviderconfig

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	ConfluentProviderConfigKey  = "confluentProviderConfig"
	confluentProviderConfigYaml = "confluent-provider-config.yaml"
)

func AddConfluentProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.ConfluentProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.ConfluentProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.ConfluentProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[ConfluentProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadConfluentProviderConfig(dir string) (string, error) {
	path := dir + "/" + confluentProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
