package tofumodule

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tofumodule/providerconfig"
	"gopkg.in/yaml.v3"
)

func GetProviderConfigEnvVars(stackInputYaml, fileCacheLoc string) ([]string, error) {
	stackInputContentMap := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(stackInputYaml), &stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal stack input yaml to map")
	}

	providerConfigEnvVars := map[string]string{}

	// providerConfigEnvVars, err = providercredentials.AddAwsProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get AWS provider config")
	}

	providerConfigEnvVars, err = providerconfig.AddAzureProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add Azure provider config")
	}

	providerConfigEnvVars, err = providerconfig.AddGcpProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add GCP provider config")
	}

	providerConfigEnvVars, err = providerconfig.AddConfluentProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add Confluent provider config")
	}

	providerConfigEnvVars, err = providerconfig.AddKubernetesProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars, fileCacheLoc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Kubernetes provider config")
	}

	providerConfigEnvVars, err = providerconfig.AddAtlasProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get MongoDB Atlas provider config")
	}

	providerConfigEnvVars, err = providerconfig.AddSnowflakeProviderConfigEnvVars(stackInputContentMap, providerConfigEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Snowflake provider config")
	}

	return mapToSlice(providerConfigEnvVars), nil
}

// mapToSlice converts a map of string to string into a slice of string slices by joining key-value pairs with an equals sign.
func mapToSlice(inputMap map[string]string) []string {
	var result []string
	for key, value := range inputMap {
		result = append(result, key+"="+value)
	}
	return result
}
