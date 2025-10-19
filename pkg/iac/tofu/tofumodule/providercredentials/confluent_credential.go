package providercredentials

import (
	"github.com/pkg/errors"
	confluent"github.com/project-planton/project-planton/apis/project/planton/provider/confluent"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddConfluentCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(confluentcredentialv1.ConfluentProviderConfig)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.ConfluentCredentialKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	credentialEnvVars["CONFLUENT_API_KEY"] = credentialSpec.ApiKey
	credentialEnvVars["CONFLUENT_API_SECRET"] = credentialSpec.ApiSecret

	return credentialEnvVars, nil
}
