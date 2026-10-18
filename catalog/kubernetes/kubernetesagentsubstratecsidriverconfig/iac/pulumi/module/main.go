package module

import (
	"github.com/pkg/errors"
	kubernetesagentsubstratecsidriverconfigv1alpha1 "github.com/plantonhq/planton/catalog/kubernetes/kubernetesagentsubstratecsidriverconfig/v1alpha1"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/manifestcr"
	"github.com/plantonhq/planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates one Agent Substrate CSIDriverConfig.
//
// The object is the manifest's projection (manifestcr), the same one the
// generated Terraform module applies: the driver's name, its Controller
// endpoint, its node socket and its TLS settings reach the cluster-scoped
// custom resource verbatim under their upstream keys, `name` (which the spec
// holds equal to the driver's name) becomes metadata.name, and the spec's
// `labels` and `annotations` become the object's own metadata under Planton's
// identity labels.
//
// No await: a CSIDriverConfig has no status. The API server and the node
// agents read it when an actor's external volume names a StorageClass whose
// provisioner is this driver. Terraform equivalent: kubectl_manifest without
// a wait.
func Resources(ctx *pulumi.Context, iacInput *kubernetesagentsubstratecsidriverconfigv1alpha1.KubernetesAgentSubstrateCsiDriverConfigIacInput) error {
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, iacInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	if _, err := manifestcr.Apply(ctx, iacInput.Target, pulumi.Provider(kubernetesProvider)); err != nil {
		return errors.Wrap(err, "failed to create csi driver config")
	}

	ctx.Export(OpDriverName, pulumi.String(iacInput.Target.GetSpec().GetDriverName()))
	return nil
}
