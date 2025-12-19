package module

import (
	"github.com/pkg/errors"
	kubernetestektonoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetestektonoperator/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi entry-point.
func Resources(ctx *pulumi.Context,
	in *kubernetestektonoperatorv1.KubernetesTektonOperatorStackInput) error {

	locals := initializeLocals(ctx, in)

	k8sProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, in.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "setup kubernetes provider")
	}

	if err = tektonOperator(ctx, locals, k8sProvider); err != nil {
		return errors.Wrap(err, "deploy tekton operator")
	}

	return nil
}
