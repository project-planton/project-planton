package providercredentials

import (
	"github.com/pkg/errors"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func AddKubernetesCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string) (map[string]string, error) {
	credentialSpec := new(kubernetesclustercredentialv1.KubernetesClusterCredentialSpec)

	isCredentialLoaded, err := stackinput.LoadCredential(stackInputContentMap,
		stackinputcredentials.KubernetesClusterKey, credentialSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credential spec from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isCredentialLoaded {
		return credentialEnvVars, nil
	}

	var kubeconfig string

	switch credentialSpec.Provider {
	case kubernetesclustercredentialv1.KubernetesProvider_gcp_gke:
		kubeconfig, err = buildGcpGkeKubeConfig(credentialSpec.GcpGke)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kubeconfig for GCP GKE")
		}
	case kubernetesclustercredentialv1.KubernetesProvider_aws_eks:
		kubeconfig, err = buildAwsEksKubeConfig(credentialSpec.AwsEks)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kubeconfig for AWS EKS")
		}
	case kubernetesclustercredentialv1.KubernetesProvider_azure_aks:
		kubeconfig, err = buildAzureAksKubeConfig(credentialSpec.AzureAks)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kubeconfig for Azure AKS")
		}
	}

	credentialEnvVars["KUBECONFIG"] = kubeconfig

	return credentialEnvVars, nil
}

func buildGcpGkeKubeConfig(gcpGkeCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialGcpGke) (string, error) {
	return "", nil
}

func buildAwsEksKubeConfig(eksCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialAwsEks) (string, error) {
	return "", nil
}

func buildAzureAksKubeConfig(aksCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialAzureAks) (string, error) {
	return "", nil
}
