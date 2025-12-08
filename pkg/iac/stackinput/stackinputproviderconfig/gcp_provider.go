package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"sigs.k8s.io/yaml"
)

const (
	GcpProviderConfigKey  = "provider_config"
	gcpProviderConfigYaml = "gcp-provider-config.yaml"
)

func AddGcpProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.GcpProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.GcpProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.GcpProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[GcpProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadGcpProviderConfig(dir string) (string, error) {
	path := dir + "/" + gcpProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
