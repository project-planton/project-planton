package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	mongodbAtlasCredentialKey = "mongodbAtlasCredential"
)

func AddAtlasCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.MongodbAtlasCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.MongodbAtlasCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.MongodbAtlasCredential)
		}
		stackInputContentMap[mongodbAtlasCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
