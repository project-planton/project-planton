package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	mongodbAtlasCredentialKey  = "mongodbAtlasCredential"
	mongodbAtlasCredentialYaml = "mongodb-atlas-credential.yaml"
)

func AddAtlasCredential(stackInputContentMap map[string]interface{},
	stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.MongodbAtlasCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.MongodbAtlasCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.MongodbAtlasCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[mongodbAtlasCredentialKey] = credentialContentMap
	}
	return stackInputContentMap, nil
}

func LoadMongodbAtlasCredential(dir string) (string, error) {
	path := dir + "/" + mongodbAtlasCredentialYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
