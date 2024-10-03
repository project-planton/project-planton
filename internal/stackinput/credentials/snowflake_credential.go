package credentials

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
	"os"
)

const (
	snowflakeCredentialKey  = "snowflakeCredential"
	snowflakeCredentialYaml = "snowflake_credential.yaml"
)

func AddSnowflakeCredential(stackInputContentMap map[string]interface{}, stackInputOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if stackInputOptions.SnowflakeCredential != "" {
		credentialContent, err := os.ReadFile(stackInputOptions.SnowflakeCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", stackInputOptions.SnowflakeCredential)
		}
		stackInputContentMap[snowflakeCredentialKey] = string(credentialContent)
	}
	return stackInputContentMap, nil
}

func LoadSnowflakeCredential(dir string) (string, error) {
	path := dir + "/" + snowflakeCredentialYaml
	isExists, err := fileutil.IsExists(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check file: %s", path)
	}
	if !isExists {
		return "", nil
	}
	return path, nil
}
