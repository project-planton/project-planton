package pulumikubernetesprovider

import (
	"fmt"
	"github.com/pkg/errors"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigkekubernetesprovider"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const emptyKubeConfigString = ""

var UnsupportedKubernetesProviderErr = errors.New("kubernetes provider is not supported")

// GetWithKubernetesClusterCredential returns kubernetes provider for the kubernetes cluster credential
func GetWithKubernetesClusterCredential(ctx *pulumi.Context,
	kubernetesClusterCredentialSpec *kubernetesclustercredentialv1.KubernetesClusterCredentialSpec,
	providerName string) (*kubernetes.Provider, error) {

	if kubernetesClusterCredentialSpec == nil {
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

	switch kubernetesClusterCredentialSpec.Provider {
	case kubernetesclustercredentialv1.KubernetesProvider_aws_eks:
		kubeConfigString = awsEks(kubernetesClusterCredentialSpec)
	case kubernetesclustercredentialv1.KubernetesProvider_azure_aks:
		kubeConfigString = azureAks(kubernetesClusterCredentialSpec)
	case kubernetesclustercredentialv1.KubernetesProvider_digital_ocean_doks:
		kubeConfigString = digitalOceanDoks(kubernetesClusterCredentialSpec)
	case kubernetesclustercredentialv1.KubernetesProvider_gcp_gke:
		kubeConfigString = gcpGke(kubernetesClusterCredentialSpec)
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

func awsEks(spec *kubernetesclustercredentialv1.KubernetesClusterCredentialSpec) (kubeConfigString string) {
	if spec.AwsEks == nil {
		return emptyKubeConfigString
	}

	kubeConfigString = "coming-soon"

	return kubeConfigString
}

func azureAks(spec *kubernetesclustercredentialv1.KubernetesClusterCredentialSpec) (kubeConfigString string) {
	if spec.AzureAks == nil {
		return emptyKubeConfigString
	}

	kubeConfigString = "coming-soon"

	return kubeConfigString
}

func digitalOceanDoks(spec *kubernetesclustercredentialv1.KubernetesClusterCredentialSpec) (kubeConfigString string) {
	if spec.DigitalOceanDoks == nil {
		return emptyKubeConfigString
	}

	return spec.DigitalOceanDoks.KubeConfig
}

func gcpGke(spec *kubernetesclustercredentialv1.KubernetesClusterCredentialSpec) (kubeConfigString string) {
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
