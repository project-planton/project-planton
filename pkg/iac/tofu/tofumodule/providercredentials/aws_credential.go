package providercredentials

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddAwsCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec, err := stackinputcredentials.GetAwsCredential(stackInputContentMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get aws credential spec from stack-input content")
	}

	credentialEnvVars["AWS_REGION"] = credentialSpec.Region
	credentialEnvVars["AWS_ACCESS_KEY_ID"] = credentialSpec.AccessKeyId
	credentialEnvVars["AWS_SECRET_ACCESS_KEY"] = credentialSpec.SecretAccessKey

	return credentialEnvVars, nil
}
