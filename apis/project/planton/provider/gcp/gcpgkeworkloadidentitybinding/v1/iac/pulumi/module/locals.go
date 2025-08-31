package module

import (
	"strconv"

	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gcpgkeworkloadidentitybindingv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeworkloadidentitybinding/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps frequently‑used values (metadata, labels, credentials) handy for the module.
type Locals struct {
	GcpCredentialSpec             *gcpcredentialv1.GcpCredentialSpec
	GcpGkeWorkloadIdentityBinding *gcpgkeworkloadidentitybindingv1.GcpGkeWorkloadIdentityBinding
	GcpLabels                     map[string]string
}

// initializeLocals populates the Locals struct from the stack input.
// It mirrors the pattern used in the gcp_gke_cluster module and applies the same Planton label strategy.
func initializeLocals(_ *pulumi.Context, stackInput *gcpgkeworkloadidentitybindingv1.GcpGkeWorkloadIdentityBindingStackInput) *Locals {
	locals := &Locals{}

	locals.GcpGkeWorkloadIdentityBinding = stackInput.Target

	// Standard Planton‑wide labels for GCP resources
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: locals.GcpGkeWorkloadIdentityBinding.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpGkeWorkloadIdentityBinding.String(),
	}

	if locals.GcpGkeWorkloadIdentityBinding.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpGkeWorkloadIdentityBinding.Metadata.Org
	}

	if locals.GcpGkeWorkloadIdentityBinding.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpGkeWorkloadIdentityBinding.Metadata.Env
	}

	if locals.GcpGkeWorkloadIdentityBinding.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GcpGkeWorkloadIdentityBinding.Metadata.Id
	}

	locals.GcpCredentialSpec = stackInput.ProviderCredential

	return locals
}
