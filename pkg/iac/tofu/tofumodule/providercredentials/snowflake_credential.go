package providercredentials

import (
	"github.com/pkg/errors"
	snowflake"github.com/project-planton/project-planton/apis/project/planton/provider/snowflake"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddSnowflakeCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(snowflakecredentialv1.SnowflakeProviderConfig)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.SnowflakeCredentialKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	credentialEnvVars["SNOWFLAKE_ACCOUNT"] = credentialSpec.Account
	credentialEnvVars["SNOWFLAKE_REGION"] = credentialSpec.Region
	credentialEnvVars["SNOWFLAKE_USERNAME"] = credentialSpec.Username
	credentialEnvVars["SNOWFLAKE_PASSWORD"] = credentialSpec.Password

	return credentialEnvVars, nil
}
