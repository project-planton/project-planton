package pulumiekskubernetesprovider

import (
	"github.com/pkg/errors"
	awscredentialv1 "github.com/project-planton/project-planton/apis/go/project/planton/credential/awscredential/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetWithCreatedEksClusterWithAwsCredentials returns kubernetes provider for the added eks cluster based on the aws provider
func GetWithCreatedEksClusterWithAwsCredentials(ctx *pulumi.Context, createdEksCluster *eks.Cluster,
	awsCredentialSpec *awscredentialv1.AwsCredentialSpec,
	dependencies []pulumi.Resource, providerName string) (*kubernetes.Provider, error) {
	provider, err := kubernetes.NewProvider(ctx,
		providerName,
		&kubernetes.ProviderArgs{
			EnableServerSideApply: pulumi.Bool(true),
			Kubeconfig: pulumi.Sprintf(AwsExecPluginKubeConfigTemplate,
				createdEksCluster.Endpoint,
				createdEksCluster.CertificateAuthority.Data().Elem(),
				awsCredentialSpec.AccessKeyId,
				awsCredentialSpec.SecretAccessKey,
				awsCredentialSpec.Region,
			),
		}, pulumi.DependsOn(dependencies))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new provider")
	}
	return provider, nil
}
