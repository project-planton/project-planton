package module

import (
	"github.com/pkg/errors"
	signozkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/signozkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *signozkubernetesv1.SignozKubernetesStackInput) error {
	//create kubernetes-provider from the credential in the stack-input
	_, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.ProviderCredential, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	return nil
}
