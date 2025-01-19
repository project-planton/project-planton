package providercredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddSnowflakeCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec, err := stackinputcredentials.GetSnowflakeCredential(stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get snowflake credential spec from stack-input content")
	}

	credentialEnvVars["SNOWFLAKE_ACCOUNT"] = credentialSpec.Account
	credentialEnvVars["SNOWFLAKE_REGION"] = credentialSpec.Region
	credentialEnvVars["SNOWFLAKE_USERNAME"] = credentialSpec.Username
	credentialEnvVars["SNOWFLAKE_PASSWORD"] = credentialSpec.Password

	return credentialEnvVars, nil
}
