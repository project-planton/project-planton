package providerconfig

import (
	"github.com/pkg/errors"
	aws "github.com/project-planton/project-planton/apis/org/project_planton/provider/aws"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddAwsProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	awsProviderConfig := new(aws.AwsProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.AwsProviderConfigKey, awsProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider config, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	providerConfigEnvVars["AWS_REGION"] = awsProviderConfig.GetRegion()
	providerConfigEnvVars["AWS_ACCESS_KEY_ID"] = awsProviderConfig.AccessKeyId
	providerConfigEnvVars["AWS_SECRET_ACCESS_KEY"] = awsProviderConfig.SecretAccessKey

	return providerConfigEnvVars, nil
}
