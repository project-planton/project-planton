package pulumikubernetesprovider

import (
	"fmt"
	"github.com/pkg/errors"
	kubernetesprovider "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigkekubernetesprovider"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const emptyKubeConfigString = ""

var UnsupportedKubernetesProviderErr = errors.New("kubernetes provider is not supported")

// GetWithKubernetesProviderConfig returns kubernetes provider for the kubernetes cluster credential
func GetWithKubernetesProviderConfig(ctx *pulumi.Context,
	kubernetesProviderConfig *kubernetesprovider.KubernetesProviderConfig,
	providerName string) (*kubernetes.Provider, error) {

	if kubernetesProviderConfig == nil {
		provider, err := kubernetes.NewProvider(ctx,
			providerName,
			&kubernetes.ProviderArgs{
				EnableServerSideApply: pulumi.Bool(true),
			})
		if err != nil {
			return nil, errors.Wrap(err, "failed to get new provider")
		}
		return provider, nil
	}

	kubeConfigString := ""

	switch kubernetesProviderConfig.Provider {
	case kubernetesprovider.KubernetesProvider_aws_eks:
		kubeConfigString = awsEks(kubernetesProviderConfig)
	case kubernetesprovider.KubernetesProvider_azure_aks:
		kubeConfigString = azureAks(kubernetesProviderConfig)
	case kubernetesprovider.KubernetesProvider_digital_ocean_doks:
		kubeConfigString = digitalOceanDoks(kubernetesProviderConfig)
	case kubernetesprovider.KubernetesProvider_gcp_gke:
		kubeConfigString = gcpGke(kubernetesProviderConfig)
	default:
		return nil, UnsupportedKubernetesProviderErr
	}

	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig:            pulumi.String(kubeConfigString),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}

func awsEks(spec *kubernetesprovider.KubernetesProviderConfig) (kubeConfigString string) {
	if spec.AwsEks == nil {
		return emptyKubeConfigString
	}

	kubeConfigString = "coming-soon"

	return kubeConfigString
}

func azureAks(spec *kubernetesprovider.KubernetesProviderConfig) (kubeConfigString string) {
	if spec.AzureAks == nil {
		return emptyKubeConfigString
	}

	kubeConfigString = "coming-soon"

	return kubeConfigString
}

func digitalOceanDoks(spec *kubernetesprovider.KubernetesProviderConfig) (kubeConfigString string) {
	if spec.DigitalOceanDoks == nil {
		return emptyKubeConfigString
	}

	return spec.DigitalOceanDoks.KubeConfig
}

func gcpGke(spec *kubernetesprovider.KubernetesProviderConfig) (kubeConfigString string) {
	if spec.GcpGke == nil {
		return emptyKubeConfigString
	}

	c := spec.GcpGke

	kubeConfigString = fmt.Sprintf(pulumigkekubernetesprovider.GcpExecPluginKubeConfigTemplate,
		c.ClusterEndpoint,
		c.ClusterCaData,
		pulumigkekubernetesprovider.GcpExecPluginPath,
		c.ServiceAccountKeyBase64)

	return kubeConfigString
}
