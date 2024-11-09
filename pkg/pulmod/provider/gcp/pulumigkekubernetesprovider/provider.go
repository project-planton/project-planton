package pulumigkekubernetesprovider

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithCreatedGkeClusterAndCreatedGsaKey returns kubernetes provider for the added container cluster based on the google provider
// the provider creation should depend on the readiness of the node-pools
func GetWithCreatedGkeClusterAndCreatedGsaKey(ctx *pulumi.Context,
	createdServiceAccountKey *serviceaccount.Key,
	createdGkeCluster *container.Cluster,
	dependencies []pulumi.Resource, providerName string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig: pulumi.Sprintf(GcpExecPluginKubeConfigTemplate,
				createdGkeCluster.Endpoint,
				createdGkeCluster.MasterAuth.ClusterCaCertificate().Elem(),
				GcpExecPluginPath,
				createdServiceAccountKey.PrivateKey),
		}, pulumi.DependsOn(dependencies))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
