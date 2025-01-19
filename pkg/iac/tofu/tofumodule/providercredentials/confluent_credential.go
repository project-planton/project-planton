package providercredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddConfluentCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec, err := stackinputcredentials.GetConfluentCredential(stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get confluent credential spec from stack-input content")
	}

	credentialEnvVars["CONFLUENT_API_KEY"] = credentialSpec.ApiKey
	credentialEnvVars["CONFLUENT_API_SECRET"] = credentialSpec.ApiSecret

	return credentialEnvVars, nil
}
