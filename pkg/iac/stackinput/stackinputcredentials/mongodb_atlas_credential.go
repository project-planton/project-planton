package stackinputcredentials

import (
	"github.com/pkg/errors"
	mongodbatlascredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/mongodbatlascredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	mongodbAtlasCredentialKey  = "mongodbAtlasCredential"
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

func GetMongodbAtlasCredential(stackInputContentMap map[string]interface{}) (*mongodbatlascredentialv1.MongodbAtlasCredentialSpec, error) {
	mongodbAtlasCredential, ok := stackInputContentMap[mongodbAtlasCredentialKey]
	if !ok {
		return nil, nil
	}

	mongodbAtlasCredentialBytes, err := yaml.Marshal(mongodbAtlasCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal mongodbAtlas credential content")
	}
	mongodbAtlasCredentialSpec := new(mongodbatlascredentialv1.MongodbAtlasCredentialSpec)
	err = yaml.Unmarshal(mongodbAtlasCredentialBytes, mongodbAtlasCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal mongodbAtlas credential content")
	}

	return mongodbAtlasCredentialSpec, nil
}
