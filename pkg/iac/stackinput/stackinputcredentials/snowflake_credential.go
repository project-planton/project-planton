package stackinputcredentials

import (
	"github.com/pkg/errors"
	snowflakecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/snowflakecredential/v1"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	snowflakeCredentialKey  = "snowflakeCredential"
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
		stackInputContentMap[snowflakeCredentialKey] = credentialContentMap
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

func GetSnowflakeCredential(stackInputContentMap map[string]interface{}) (*snowflakecredentialv1.SnowflakeCredentialSpec, error) {
	snowflakeCredential, ok := stackInputContentMap[snowflakeCredentialKey]
	if !ok {
		return nil, nil
	}

	snowflakeCredentialBytes, err := yaml.Marshal(snowflakeCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal snowflake credential content")
	}
	snowflakeCredentialSpec := new(snowflakecredentialv1.SnowflakeCredentialSpec)
	err = yaml.Unmarshal(snowflakeCredentialBytes, snowflakeCredentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal snowflake credential content")
	}

	return snowflakeCredentialSpec, nil
}
