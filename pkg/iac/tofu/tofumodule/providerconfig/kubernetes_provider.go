package providerconfig

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	kubernetesprovider "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigkekubernetesprovider"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	"os"
	"path/filepath"
)

func AddKubernetesProviderConfigEnvVars(stackInputContentMap map[string]interface{},
	providerConfigEnvVars map[string]string,
	fileCacheLoc string) (map[string]string, error) {
	kubernetesProviderConfig := new(kubernetesprovider.KubernetesProviderConfig)

	isProviderConfigLoaded, err := stackinput.LoadProviderConfig(stackInputContentMap,
		stackinputproviderconfig.KubernetesProviderConfigKey, kubernetesProviderConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider config from stack-input content")
	}

	//this means that the stack input does not contain the provider credential, so no environment variables will be added.
	if !isProviderConfigLoaded {
		return providerConfigEnvVars, nil
	}

	var kubeConfig string

	switch kubernetesProviderConfig.Provider {
	case kubernetesprovider.KubernetesProvider_gcp_gke:
		kubeConfig, err = buildGcpGkeKubeConfig(kubernetesProviderConfig.GcpGke)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kube-config for GCP GKE")
		}
	case kubernetesprovider.KubernetesProvider_aws_eks:
		kubeConfig, err = buildAwsEksKubeConfig(kubernetesProviderConfig.AwsEks)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kube-config for AWS EKS")
		}
	case kubernetesprovider.KubernetesProvider_azure_aks:
		kubeConfig, err = buildAzureAksKubeConfig(kubernetesProviderConfig.AzureAks)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build kube-config for Azure AKS")
		}
	}

	kubeConfigPath := filepath.Join(fileCacheLoc, uuid.New().String())
	if err := os.WriteFile(kubeConfigPath, []byte(kubeConfig), 0644); err != nil {
		return nil, errors.Wrap(err, "failed to write kube-config to file")
	}
	providerConfigEnvVars["KUBECONFIG"] = kubeConfigPath

	return providerConfigEnvVars, nil
}

func buildGcpGkeKubeConfig(c *kubernetesprovider.KubernetesProviderConfigGcpGke) (string, error) {
	kubeConfigString := ""

	kubeConfigString = fmt.Sprintf(pulumigkekubernetesprovider.GcpExecPluginKubeConfigTemplate,
		c.ClusterEndpoint,
		c.ClusterCaData,
		pulumigkekubernetesprovider.GcpExecPluginPath,
		c.ServiceAccountKeyBase64)

	return kubeConfigString, nil
}

func buildAwsEksKubeConfig(ekskubernetesProviderConfig *kubernetesprovider.KubernetesProviderConfigAwsEks) (string, error) {
	return "", nil
}

func buildAzureAksKubeConfig(akskubernetesProviderConfig *kubernetesprovider.KubernetesProviderConfigAzureAks) (string, error) {
	return "", nil
}
