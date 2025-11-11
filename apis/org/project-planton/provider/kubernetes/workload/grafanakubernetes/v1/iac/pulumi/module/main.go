package module

import (
	"github.com/pkg/errors"
	grafanakubernetesv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/kubernetes/workload/grafanakubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *grafanakubernetesv1.GrafanaKubernetesStackInput) error {
	//create kubernetes-provider from the credential in the stack-input
	_, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	return nil
}
