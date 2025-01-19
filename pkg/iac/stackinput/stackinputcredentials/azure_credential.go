package stackinputcredentials

import (
	"github.com/pkg/errors"
	azurecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/azurecredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	azureCredentialKey  = "azureCredential"
	azureCredentialYaml = "azure-credential.yaml"
)

func AddAzureCredential(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.AzureCredential != "" {
		credentialContent, err := os.ReadFile(credentialOptions.AzureCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.AzureCredential)
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

func GetAzureCredential(stackInputContentMap map[string]interface{}) (*azurecredentialv1.AzureCredentialSpec, error) {
	azureCredential, ok := stackInputContentMap[azureCredentialKey]
	if !ok {
		return nil, nil
	}

	azureCredentialBytes, err := yaml.Marshal(azureCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal azure credential content")
	}

	azureCredentialSpec := new(azurecredentialv1.AzureCredentialSpec)
	err = yaml.Unmarshal(azureCredentialBytes, azureCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal azure credential content")
	}

	return azureCredentialSpec, nil
}
