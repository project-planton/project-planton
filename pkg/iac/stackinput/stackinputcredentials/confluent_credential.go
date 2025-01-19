package stackinputcredentials

import (
	"github.com/pkg/errors"
	confluentcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/confluentcredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	confluentCredentialKey  = "confluentCredential"
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

func GetConfluentCredential(stackInputContentMap map[string]interface{}) (*confluentcredentialv1.ConfluentCredentialSpec, error) {
	confluentCredential, ok := stackInputContentMap[confluentCredentialKey]
	if !ok {
		return nil, nil
	}

	confluentCredentialBytes, err := yaml.Marshal(confluentCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal confluent credential content")
	}

	confluentCredentialSpec := new(confluentcredentialv1.ConfluentCredentialSpec)

	err = yaml.Unmarshal(confluentCredentialBytes, confluentCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal confluent credential content")
	}

	return confluentCredentialSpec, nil
}
