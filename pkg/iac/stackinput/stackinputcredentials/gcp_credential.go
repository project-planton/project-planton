package stackinputcredentials

import (
	"github.com/pkg/errors"
	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	gcpCredentialKey  = "gcpCredential"
	gcpCredentialYaml = "gcp-credential.yaml"
)

func AddGcpCredential(stackInputContentMap map[string]interface{},
	credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.GcpCredential != "" {
		credentialContent, err := os.ReadFile(credentialOptions.GcpCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.GcpCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[gcpCredentialKey] = credentialContentMap
	}
	return stackInputContentMap, nil
}

func LoadGcpCredential(dir string) (string, error) {
	path := dir + "/" + gcpCredentialYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}

func GetGcpCredential(stackInputContentMap map[string]interface{}) (*gcpcredentialv1.GcpCredentialSpec, error) {
	gcpCredential, ok := stackInputContentMap[gcpCredentialKey]
	if !ok {
		return nil, nil
	}

	gcpCredentialBytes, err := yaml.Marshal(gcpCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal gcp credential content")
	}

	gcpCredentialSpec := new(gcpcredentialv1.GcpCredentialSpec)
	err = yaml.Unmarshal(gcpCredentialBytes, gcpCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal gcp credential content")
	}

	return gcpCredentialSpec, nil
}
