package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"sigs.k8s.io/yaml"
)

const (
	AtlasProviderConfigKey  = "provider_config"
	atlasProviderConfigYaml = "atlas-provider-config.yaml"
)

func AddAtlasProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.AtlasProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.AtlasProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.AtlasProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[AtlasProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadAtlasProviderConfig(dir string) (string, error) {
	path := dir + "/" + atlasProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
