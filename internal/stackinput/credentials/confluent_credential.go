package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
	"os"
)

const (
	confluentCredentialKey  = "confluentCredential"
	confluentCredentialYaml = "confluent-credential.yaml"
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

func LoadConfluentCredential(dir string) (string, error) {
	path := dir + "/" + confluentCredentialYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
