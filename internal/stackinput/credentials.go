package stackinput

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/stackinput/credentials"
)

func addCredentials(stackInputContentMap map[string]string,
	credentialOptions credentials.StackInputCredentialOptions) (updatedStackInputContentMap map[string]string, err error) {
	updatedStackInputContentMap, err = credentials.AddAtlasCredential(stackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add atlas-credential")
	}
	updatedStackInputContentMap, err = credentials.AddAwsCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add aws-credential")
	}
	updatedStackInputContentMap, err = credentials.AddAzureCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add azure-credential")
	}
	updatedStackInputContentMap, err = credentials.AddConfluentCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add confluent-credential")
	}
	updatedStackInputContentMap, err = credentials.AddDockerCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add docker-credential")
	}
	updatedStackInputContentMap, err = credentials.AddGcpCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add gcp-credential")
	}
	updatedStackInputContentMap, err = credentials.AddKubernetesCluster(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add kubernetes-cluster")
	}
	updatedStackInputContentMap, err = credentials.AddSnowflakeCredential(updatedStackInputContentMap, credentialOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add snowflake-credential")
	}
	return updatedStackInputContentMap, nil
}
