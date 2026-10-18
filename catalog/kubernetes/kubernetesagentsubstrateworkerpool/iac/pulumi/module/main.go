package module

import (
	"github.com/pkg/errors"
	kubernetesagentsubstrateworkerpoolv1alpha1 "github.com/plantonhq/planton/catalog/kubernetes/kubernetesagentsubstrateworkerpool/v1alpha1"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/manifestcr"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates one Agent Substrate WorkerPool.
//
// The object is the manifest's projection (manifestcr), the same one the
// generated Terraform module applies: the size, the worker image (resolved
// from the control plane's output when referenced), the worker pod template
// and the sandbox class reach the custom resource verbatim under their
// upstream keys, `namespace` becomes metadata.namespace, and the spec's
// `labels` and `annotations` become the object's own metadata under Planton's
// identity labels -- the labels actor templates select the pool by.
//
// No await: Substrate's controller renders the worker Deployment after the
// object lands, and only once the pool's class resolves to a SandboxConfig;
// the workers then register at their own pace. The E2E verifier asserts the
// workers are ready and registered with the API. Terraform equivalent:
// kubectl_manifest without a wait.
func Resources(ctx *pulumi.Context, iacInput *kubernetesagentsubstrateworkerpoolv1alpha1.KubernetesAgentSubstrateWorkerPoolIacInput) error {
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, iacInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	obj, err := manifestcr.Apply(ctx, iacInput.Target, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create worker pool")
	}

	ctx.Export(OpWorkerPoolName, pulumi.String(obj.Name))
	ctx.Export(OpNamespace, pulumi.String(obj.Namespace))
	return nil
}
