package providercredentials

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigkekubernetesprovider"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
	"os"
	"path/filepath"
)

func AddKubernetesCredentialEnvVars(stackInputContentMap map[string]interface{},
	credentialEnvVars map[string]string,
	fileCacheLoc string) (map[string]string, error) {
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

	var kubeConfig string

	switch credentialSpec.Provider {
	case kubernetesclustercredentialv1.KubernetesProvider_gcp_gke:
		kubeConfig, err = buildGcpGkeKubeConfig(credentialSpec.GcpGke)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kube-config for GCP GKE")
		}
	case kubernetesclustercredentialv1.KubernetesProvider_aws_eks:
		kubeConfig, err = buildAwsEksKubeConfig(credentialSpec.AwsEks)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kube-config for AWS EKS")
		}
	case kubernetesclustercredentialv1.KubernetesProvider_azure_aks:
		kubeConfig, err = buildAzureAksKubeConfig(credentialSpec.AzureAks)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kube-config for Azure AKS")
		}
	}

	kubeConfigPath := filepath.Join(fileCacheLoc, uuid.New().String())
	if err := os.WriteFile(kubeConfigPath, []byte(kubeConfig), 0644); err != nil {
		return nil, errors.Wrap(err, "failed to write kube-config to file")
	}
	credentialEnvVars["KUBECONFIG"] = kubeConfigPath

	return credentialEnvVars, nil
}

func buildGcpGkeKubeConfig(c *kubernetesclustercredentialv1.KubernetesClusterCredentialGcpGke) (string, error) {
	kubeConfigString := ""

	kubeConfigString = fmt.Sprintf(pulumigkekubernetesprovider.GcpExecPluginKubeConfigTemplate,
		c.ClusterEndpoint,
		c.ClusterCaData,
		pulumigkekubernetesprovider.GcpExecPluginPath,
		c.ServiceAccountKeyBase64)

	return kubeConfigString, nil
}

func buildAwsEksKubeConfig(eksCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialAwsEks) (string, error) {
	return "", nil
}

func buildAzureAksKubeConfig(aksCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialAzureAks) (string, error) {
	return "", nil
}
