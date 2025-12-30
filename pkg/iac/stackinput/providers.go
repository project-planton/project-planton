package stackinput

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	"github.com/project-planton/project-planton/pkg/protobufyaml"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

func addProviderConfigs(stackInputContentMap map[string]interface{},
	providerConfigOptions stackinputproviderconfig.StackInputProviderConfigOptions) (updatedStackInputContentMap map[string]interface{}, err error) {
	updatedStackInputContentMap, err = stackinputproviderconfig.AddAtlasProviderConfig(stackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add atlas-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddAuth0ProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add auth0-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddAwsProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add aws-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddAzureProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add azure-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddCloudflareProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add cloudflare-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddConfluentProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add confluent-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddGcpProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add gcp-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddKubernetesProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add kubernetes-provider-config")
	}
	updatedStackInputContentMap, err = stackinputproviderconfig.AddSnowflakeProviderConfig(updatedStackInputContentMap, providerConfigOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add snowflake-provider-config")
	}
	return updatedStackInputContentMap, nil
}

func LoadProviderConfig(stackInputContentMap map[string]interface{}, providerConfigKey string,
	providerConfigObject proto.Message) (isProviderConfigLoaded bool, err error) {
	providerConfigYaml, ok := stackInputContentMap[providerConfigKey]
	if !ok {
		return false, nil
	}

	providerConfigBytes, err := yaml.Marshal(providerConfigYaml)
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal provider config yaml content")
	}

	err = protobufyaml.LoadYamlBytes(providerConfigBytes, providerConfigObject)
	if err != nil {
		return false, errors.Wrap(err, "failed to load yaml bytes into provider config")
	}

	return true, nil
}
