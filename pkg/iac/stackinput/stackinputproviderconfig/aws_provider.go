package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
)

const (
	AwsProviderConfigKey  = "provider_config"
	awsProviderConfigYaml = "aws-provider-config.yaml"
)

func AddAwsProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.AwsProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.AwsProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.AwsProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[AwsProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadAwsProviderConfig(dir string) (string, error) {
	path := dir + "/" + awsProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
