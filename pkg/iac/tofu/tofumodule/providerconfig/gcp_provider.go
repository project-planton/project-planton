package providerconfig

import (
	"encoding/base64"
	"github.com/pkg/errors"
	gcpprovider "github.com/plantonhq/project-planton/apis/org/project_planton/provider/gcp"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func AddGcpProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string) (map[string]string, error) {
	gcpProviderConfig := new(gcpprovider.GcpProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.GcpProviderConfigKey, gcpProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	serviceAccountKey, err := base64.StdEncoding.DecodeString(gcpProviderConfig.ServiceAccountKeyBase64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode service account key from base64")
	}

	providerConfigEnvVars["GOOGLE_CREDENTIALS"] = string(serviceAccountKey)

	return providerConfigEnvVars, nil
}
