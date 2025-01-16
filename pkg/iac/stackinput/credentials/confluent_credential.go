package credentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	confluentCredentialKey  = "confluentCredential"
	confluentCredentialYaml = "confluent-credential.yaml"
)

func AddConfluentCredential(stackInputContentMap map[string]interface{},
	stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.ConfluentCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.ConfluentCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.ConfluentCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[confluentCredentialKey] = credentialContentMap
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
