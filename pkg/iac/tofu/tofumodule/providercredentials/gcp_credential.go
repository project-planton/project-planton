package providercredentials

import (
	"encoding/base64"
	"github.com/pkg/errors"
	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddGcpCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(gcpcredentialv1.GcpCredentialSpec)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.GcpCredentialKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	serviceAccountKey, err := base64.StdEncoding.DecodeString(credentialSpec.ServiceAccountKeyBase64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode service account key from base64")
	}

	credentialEnvVars["GOOGLE_CREDENTIALS"] = string(serviceAccountKey)

	return credentialEnvVars, nil
}
