package stackinputproviderconfig

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	SnowflakeProviderConfigKey  = "snowflakeProviderConfig"
	snowflakeProviderConfigYaml = "snowflake-provider-config.yaml"
)

func AddSnowflakeProviderConfig(stackInputContentMap map[string]interface{}, providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.SnowflakeProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.SnowflakeProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.SnowflakeProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[SnowflakeProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadSnowflakeProviderConfig(dir string) (string, error) {
	path := dir + "/" + snowflakeProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
