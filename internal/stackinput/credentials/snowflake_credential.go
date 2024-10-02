package credentials

import (
	"github.com/pkg/errors"
	"os"
)

const (
	snowflakeCredentialKey = "snowflakeCredential"
)

func AddSnowflakeCredential(stackInputContentMap map[string]string, stackInputOptions StackInputCredentialOptions) (map[string]string, error) {
	if stackInputOptions.SnowflakeCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.SnowflakeCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.SnowflakeCredential)
		}
		stackInputContentMap[snowflakeCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}
