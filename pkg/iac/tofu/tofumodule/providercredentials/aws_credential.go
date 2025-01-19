package providercredentials

import (
	"github.com/pkg/errors"
	awscredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/awscredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddAwsCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(awscredentialv1.AwsCredentialSpec)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.AwsCredentialKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	credentialEnvVars["AWS_REGION"] = credentialSpec.Region
	credentialEnvVars["AWS_ACCESS_KEY_ID"] = credentialSpec.AccessKeyId
	credentialEnvVars["AWS_SECRET_ACCESS_KEY"] = credentialSpec.SecretAccessKey

	return credentialEnvVars, nil
}
