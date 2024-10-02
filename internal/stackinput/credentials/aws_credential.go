package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	awsCredentialKey = "awsCredential"
)

func AddAwsCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.AwsCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.AwsCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.AwsCredential)
		}
		stackInputContentMap[awsCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
