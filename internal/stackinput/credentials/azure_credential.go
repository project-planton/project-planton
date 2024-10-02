package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	azureCredentialKey = "azureCredential"
)

func AddAzureCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.AzureCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.AzureCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.AzureCredential)
		}
		stackInputContentMap[azureCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
