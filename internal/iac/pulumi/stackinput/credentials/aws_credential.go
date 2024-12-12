package credentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	awsCredentialKey  = "awsCredential"
	awsCredentialYaml = "aws-credential.yaml"
)

func AddAwsCredential(stackInputContentMap map[string]interface{},
	stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.AwsCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.AwsCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.AwsCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[awsCredentialKey] = credentialContentMap
	}
	return stackInputContentMap, nil
}

func LoadAwsCredential(dir string) (string, error) {
	path := dir + "/" + awsCredentialYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
