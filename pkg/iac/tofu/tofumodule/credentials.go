package tofumodule

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tofumodule/providercredentials"
	"gopkg.in/yaml.v3"
)

func GetCredentialEnvVars(stackInputYaml, fileCacheLoc string) ([]string, error) {
	stackInputContentMap := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(stackInputYaml), &stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal stack input yaml to map")
	}

	credentialEnvVars := map[string]string{}

	credentialEnvVars, err = providercredentials.AddAwsCredentialEnvVars(stackInputContentMap, credentialEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get AWS provider credentials")
	}

	credentialEnvVars, err = providercredentials.AddAzureCredentialEnvVars(stackInputContentMap, credentialEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add Azure provider credentials")
	}

	credentialEnvVars, err = providercredentials.AddGcpProviderConfigEnvVars(stackInputContentMap, credentialEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add GCP provider credentials")
	}

	credentialEnvVars, err = providercredentials.AddConfluentCredentialEnvVars(stackInputContentMap, credentialEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add Confluent provider credentials")
	}

	credentialEnvVars, err = providercredentials.AddKubernetesCredentialEnvVars(stackInputContentMap, credentialEnvVars, fileCacheLoc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Kubernetes provider credentials")
	}

	credentialEnvVars, err = providercredentials.AddMongodbAtlasCredentialEnvVars(stackInputContentMap, credentialEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get MongoDB Atlas provider credentials")
	}

	credentialEnvVars, err = providercredentials.AddSnowflakeCredentialEnvVars(stackInputContentMap, credentialEnvVars)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Snowflake provider credentials")
	}

	return mapToSlice(credentialEnvVars), nil
}

// mapToSlice converts a map of string to string into a slice of string slices by joining key-value pairs with an equals sign.
func mapToSlice(inputMap map[string]string) []string {
	var result []string
	for key, value := range inputMap {
		result = append(result, key+"="+value)
	}
	return result
}
