package credentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	gcpCredentialKey  = "gcpCredential"
	gcpCredentialYaml = "gcp-credential.yaml"
)

func AddGcpCredential(stackInputContentMap map[string]interface{},
	stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.GcpCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.GcpCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.GcpCredential)
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
