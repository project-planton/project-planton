package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	dockerCredentialKey = "dockerCredential"
)

func AddDockerCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.DockerCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.DockerCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.DockerCredential)
		}
		stackInputContentMap[dockerCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
