package providercredentials

import (
	"github.com/pkg/errors"
	mongodbatlascredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/mongodbatlascredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddMongodbAtlasCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(mongodbatlascredentialv1.MongodbAtlasCredentialSpec)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.MongodbAtlasCredentialKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	credentialEnvVars["MONGODB_ATLAS_PUBLIC_KEY"] = credentialSpec.PublicKey
	credentialEnvVars["MONGODB_ATLAS_PRIVATE_KEY"] = credentialSpec.PrivateKey

	return credentialEnvVars, nil
}
