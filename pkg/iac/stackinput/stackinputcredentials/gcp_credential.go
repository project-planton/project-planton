package stackinputcredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	GcpProviderConfigKey  = "gcpProviderConfig"
	gcpProviderConfigYaml = "gcp-credential.yaml"
)

func AddGcpProviderConfig(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.GcpProviderConfig != "" {
		credentialContent, err := os.ReadFile(credentialOptions.GcpProviderConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.GcpProviderConfig)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[GcpProviderConfigKey] = credentialContentMap
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
