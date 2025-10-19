package providercredentials

import (
	"github.com/pkg/errors"
	azure"github.com/project-planton/project-planton/apis/project/planton/provider/azure"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddAzureCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(azure.AzureProviderConfig)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.AzureCredentialKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	credentialEnvVars["ARM_CLIENT_ID"] = credentialSpec.ClientId
	credentialEnvVars["ARM_CLIENT_SECRET"] = credentialSpec.ClientSecret
	credentialEnvVars["ARM_TENANT_ID"] = credentialSpec.TenantId
	credentialEnvVars["ARM_SUBSCRIPTION_ID"] = credentialSpec.SubscriptionId

	return credentialEnvVars, nil
}
