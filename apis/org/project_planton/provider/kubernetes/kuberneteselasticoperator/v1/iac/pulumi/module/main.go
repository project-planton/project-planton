package module

import (
	"github.com/pkg/errors"
	kuberneteselasticoperatorv1 "github.com/plantonhq/project-planton/apis/org/project_planton/provider/kubernetes/kuberneteselasticoperator/v1"
	"github.com/plantonhq/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi entry‑point.
func Resources(ctx *pulumi.Context,
	in *kuberneteselasticoperatorv1.KubernetesElasticOperatorStackInput) error {

	locals := initializeLocals(ctx, in)

	k8sProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, in.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "setup kubernetes provider")
	}

	if err = kubernetesElasticOperator(ctx, locals, k8sProvider); err != nil {
		return errors.Wrap(err, "deploy elastic operator")
	}

	return nil
}
