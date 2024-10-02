package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	confluentCredentialKey = "confluentCredential"
)

func AddConfluentCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.ConfluentCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.ConfluentCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.ConfluentCredential)
		}
		stackInputContentMap[confluentCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
