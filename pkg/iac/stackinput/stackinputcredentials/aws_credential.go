package stackinputcredentials

import (
	"github.com/pkg/errors"
	awscredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/awscredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	awsCredentialKey  = "awsCredential"
	awsCredentialYaml = "aws-credential.yaml"
)

func AddAwsCredential(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.AwsCredential != "" {
		credentialContent, err := os.ReadFile(credentialOptions.AwsCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.AwsCredential)
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

func GetAwsCredential(stackInputContentMap map[string]interface{}) (*awscredentialv1.AwsCredentialSpec, error) {
	awsCredential, ok := stackInputContentMap[awsCredentialKey]
	if !ok {
		return nil, nil
	}

	awsCredentialBytes, err := yaml.Marshal(awsCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal aws credential content")
	}

	var awsCredentialSpec awscredentialv1.AwsCredentialSpec
	err = yaml.Unmarshal(awsCredentialBytes, &awsCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal aws credential content")
	}

	return &awsCredentialSpec, nil
}
