package stackinputcredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	ConfluentCredentialKey  = "confluentCredential"
	confluentCredentialYaml = "confluent-credential.yaml"
)

func AddConfluentCredential(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.ConfluentCredential != "" {
		credentialContent, err := os.ReadFile(credentialOptions.ConfluentCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.ConfluentCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[ConfluentCredentialKey] = credentialContentMap
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
