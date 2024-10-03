package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	azureCredentialKey  = "azureCredential"
	azureCredentialYaml = "azure-credential.yaml"
)

func AddAzureCredential(stackInputContentMap map[string]interface{},
	stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.AzureCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.AzureCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.AzureCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[azureCredentialKey] = credentialContentMap
	}
	return stackInputContentMap, nil
}

func LoadAzureCredential(dir string) (string, error) {
	path := dir + "/" + azureCredentialYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
