package backendcredentials

import (
	"github.com/pkg/errors"
	azurecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/azurecredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddAzureCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(azurecredentialv1.AzureCredentialSpec)

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
