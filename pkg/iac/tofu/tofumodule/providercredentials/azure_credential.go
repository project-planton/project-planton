package providercredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddAzureCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec, err := stackinputcredentials.GetAzureCredential(stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get azure credential spec from stack-input content")
	}

	credentialEnvVars["ARM_CLIENT_ID"] = credentialSpec.ClientId
	credentialEnvVars["ARM_CLIENT_SECRET"] = credentialSpec.ClientSecret
	credentialEnvVars["ARM_TENANT_ID"] = credentialSpec.TenantId
	credentialEnvVars["ARM_SUBSCRIPTION_ID"] = credentialSpec.SubscriptionId

	return credentialEnvVars, nil
}
