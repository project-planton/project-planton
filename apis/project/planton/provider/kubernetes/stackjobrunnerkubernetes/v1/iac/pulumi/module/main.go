package module

import (
	"github.com/pkg/errors"
	stackjobrunnerkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/stackjobrunnerkubernetes/v1"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *stackjobrunnerkubernetesv1.StackJobRunnerKubernetesStackInput) error {
	//create kubernetes-provider from the credential in the stack-input
	_, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	return nil
}