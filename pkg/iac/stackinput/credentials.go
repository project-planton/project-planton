package stackinput

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func addCredentials(stackInputContentMap map[string]interface{},
	credentialOptions stackinputcredentials.StackInputCredentialOptions) (updatedStackInputContentMap map[string]interface{}, err error) {
	updatedStackInputContentMap, err = stackinputcredentials.AddAtlasCredential(stackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add atlas-credential")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddAwsCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add aws-credential")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddAzureCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add azure-credential")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddConfluentCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add confluent-credential")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddDockerCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add docker-credential")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddGcpCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add gcp-credential")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddKubernetesCluster(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add kubernetes-cluster")
	}
	updatedStackInputContentMap, err = stackinputcredentials.AddSnowflakeCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add snowflake-credential")
	}
	return updatedStackInputContentMap, nil
}
