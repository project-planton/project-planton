package stackinputcredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	MongodbAtlasCredentialKey  = "mongodbAtlasCredential"
	mongodbAtlasCredentialYaml = "mongodb-atlas-credential.yaml"
)

func AddAtlasCredential(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.MongodbAtlasCredential != "" {
		credentialContent, err := os.ReadFile(credentialOptions.MongodbAtlasCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.MongodbAtlasCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[MongodbAtlasCredentialKey] = credentialContentMap
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
