package stackinputproviderconfig

import (
	"os"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/pkg/fileutil"
	"sigs.k8s.io/yaml"
)

const (
	KubernetesProviderConfigKey  = "provider_config"
	kubernetesProviderConfigYaml = "kubernetes-provider-config.yaml"
)

func AddKubernetesProviderConfig(stackInputContentMap map[string]interface{},
	providerConfigOptions StackInputProviderConfigOptions) (map[string]interface{}, error) {
	if providerConfigOptions.KubernetesProviderConfig != "" {
		providerConfigContent, err := os.ReadFile(providerConfigOptions.KubernetesProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", providerConfigOptions.KubernetesProviderConfig)
		}
		var providerConfigContentMap map[string]interface{}
		err = yaml.Unmarshal(providerConfigContent, &providerConfigContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[KubernetesProviderConfigKey] = providerConfigContentMap
	}
	return stackInputContentMap, nil
}

func LoadKubernetesProviderConfig(dir string) (string, error) {
	path := dir + "/" + kubernetesProviderConfigYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
