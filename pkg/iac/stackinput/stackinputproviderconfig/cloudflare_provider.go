package stackinputproviderconfig

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	CloudflareProviderConfigKey  = "cloudflareProviderConfig"
	cloudflareProviderConfigYaml = "cloudflare-provider-config.yaml"
)

func AddCloudflareProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.CloudflareProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.CloudflareProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.CloudflareProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[CloudflareProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadCloudflareProviderConfig(dir string) (string, error) {
	path := dir + "/" + cloudflareProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
