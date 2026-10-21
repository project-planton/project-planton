package module

import (
	"github.com/pkg/errors"
	kubernetesagentsubstratesandboxconfigv1alpha1 "github.com/plantonhq/planton/catalog/kubernetes/kubernetesagentsubstratesandboxconfig/v1alpha1"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/manifestcr"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates one Agent Substrate SandboxConfig.
//
// The object is the manifest's projection (manifestcr), the same one the
// generated Terraform module applies: the class, the default version and the
// versions reach the cluster-scoped custom resource verbatim under their
// upstream keys (each architecture's asset files written as the bare map
// upstream defines), `name` becomes metadata.name, and the spec's `labels`
// and `annotations` become the object's own metadata under Planton's identity
// labels -- the class-default annotation included.
//
// No await: a SandboxConfig has no status; the cluster's admission policy
// accepts or refuses it at apply, and the API server and node agents read it
// when a template names it. Destroy waits while a worker pool still runs the
// config (Substrate's protection finalizer), which a pool that references it
// releases first because the reference orders it after the config. Terraform
// equivalent: kubectl_manifest without a wait.
func Resources(ctx *pulumi.Context, iacInput *kubernetesagentsubstratesandboxconfigv1alpha1.KubernetesAgentSubstrateSandboxConfigIacInput) error {
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, iacInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	obj, err := manifestcr.Apply(ctx, iacInput.Target, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create sandbox config")
	}

	ctx.Export(OpSandboxConfigName, pulumi.String(obj.Name))
	ctx.Export(OpSandboxClass, pulumi.String(iacInput.Target.GetSpec().GetSandboxClass()))
	return nil
}
