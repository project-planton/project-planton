package providercredentials

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddGcpCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec, err := stackinputcredentials.GetGcpCredential(stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gcp credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if credentialSpec == nil {
		return credentialEnvVars, nil
	}
	
	serviceAccountKey, err := base64.StdEncoding.DecodeString(credentialSpec.ServiceAccountKeyBase64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode service account key from base64")
	}

	credentialEnvVars["GOOGLE_CREDENTIAL"] = string(serviceAccountKey)

	return credentialEnvVars, nil
}
