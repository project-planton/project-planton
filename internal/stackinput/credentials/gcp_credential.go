package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	gcpCredentialKey = "gcpCredential"
)

func AddGcpCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.GcpCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.GcpCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.GcpCredential)
		}
		stackInputContentMap[gcpCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
