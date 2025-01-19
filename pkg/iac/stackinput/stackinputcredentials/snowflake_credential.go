package stackinputcredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	SnowflakeCredentialKey  = "snowflakeCredential"
	snowflakeCredentialYaml = "snowflake_credential.yaml"
)

func AddSnowflakeCredential(stackInputContentMap map[string]interface{}, credentialOptions StackInputCredentialOptions) (map[string]interface{}, error) {
	if credentialOptions.SnowflakeCredential != "" {
		credentialContent, err := os.ReadFile(credentialOptions.SnowflakeCredential)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read file: %s", credentialOptions.SnowflakeCredential)
		}
		var credentialContentMap map[string]interface{}
		err = yaml.Unmarshal(credentialContent, &credentialContentMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal target manifest file")
		}
		stackInputContentMap[SnowflakeCredentialKey] = credentialContentMap
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
