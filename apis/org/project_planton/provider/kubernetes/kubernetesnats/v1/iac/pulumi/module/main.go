package module

import (
	"github.com/pkg/errors"
	kubernetesnatsv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesnats/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the single entry-point consumed by the ProjectPlanton
// runtime.  It wires together noun-style helpers in a Terraform-like
// top-down order so the flow is easy for DevOps engineers to follow.
func Resources(ctx *pulumi.Context,
	stackInput *kubernetesnatsv1.KubernetesNatsStackInput) error {

	// ----------------------------- locals ---------------------------------
	locals := initializeLocals(ctx, stackInput)

	// ------------------------- kubernetes provider ------------------------
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// ------------------------------ namespace ----------------------------
	createdNamespace, err := namespace(ctx, locals, kubernetesProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// ----------------------------- secrets --------------------------------
	if err := tlsSecret(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create TLS secret")
	}

	// ------------------------------ helm ----------------------------------
	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to deploy NATS Helm chart")
	}

	// ----------------------------- ingress --------------------------------
	if err := ingress(ctx, locals, kubernetesProvider, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create external ingress")
	}

	return nil
}
