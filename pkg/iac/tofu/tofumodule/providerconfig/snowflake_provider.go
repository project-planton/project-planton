package providerconfig

import (
	"github.com/pkg/errors"
	snowflake "github.com/plantonhq/project-planton/apis/org/project_planton/provider/snowflake"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddSnowflakeProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	snowflakeProviderConfig := new(snowflake.SnowflakeProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.SnowflakeProviderConfigKey, snowflakeProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	providerConfigEnvVars["SNOWFLAKE_ACCOUNT"] = snowflakeProviderConfig.Account
	providerConfigEnvVars["SNOWFLAKE_REGION"] = snowflakeProviderConfig.Region
	providerConfigEnvVars["SNOWFLAKE_USERNAME"] = snowflakeProviderConfig.Username
	providerConfigEnvVars["SNOWFLAKE_PASSWORD"] = snowflakeProviderConfig.Password

	return providerConfigEnvVars, nil
}
