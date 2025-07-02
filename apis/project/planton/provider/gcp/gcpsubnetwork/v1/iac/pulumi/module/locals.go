package module

import (
	"strconv"

	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gcpsubnetworkv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpsubnetwork/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps frequently‑used values (metadata, labels, credentials) handy for the module.
type Locals struct {
	GcpCredentialSpec *gcpcredentialv1.GcpCredentialSpec
	GcpSubnetwork            *gcpsubnetworkv1.GcpSubnetwork
	GcpLabels         map[string]string
}

// initializeLocals populates the Locals struct from the stack input.
// It mirrors the pattern used in the gcp_gke_cluster module and applies the same Planton label strategy.
func initializeLocals(_ *pulumi.Context, stackInput *gcpsubnetworkv1.GcpSubnetworkStackInput) *Locals {
	locals := &Locals{}

	locals.GcpSubnetwork = stackInput.Target

	// Standard Planton‑wide labels for GCP resources
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: locals.GcpSubnetwork.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpSubnetwork.String(),
	}

	if locals.GcpSubnetwork.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpSubnetwork.Metadata.Org
	}

	if locals.GcpSubnetwork.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpSubnetwork.Metadata.Env
	}

	if locals.GcpSubnetwork.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GcpSubnetwork.Metadata.Id
	}

	locals.GcpCredentialSpec = stackInput.ProviderCredential

	return locals
}
