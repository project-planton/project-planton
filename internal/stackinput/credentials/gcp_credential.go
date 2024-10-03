package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
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
		stackInputContentMap[gcpCredentialKey] = string(credentialContent)
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
