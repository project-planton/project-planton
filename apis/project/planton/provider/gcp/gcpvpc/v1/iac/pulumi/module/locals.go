package module

import (
	"strconv"

	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gcpvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpvpc/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps frequently‑used values (metadata, labels, credentials) handy for the module.
type Locals struct {
	GcpCredentialSpec *gcpcredentialv1.GcpCredentialSpec
	GcpVpc            *gcpvpcv1.GcpVpc
	GcpLabels         map[string]string
}

// initializeLocals populates the Locals struct from the stack input.
// It mirrors the pattern used in the gcp_gke_cluster module and applies the same Planton label strategy.
func initializeLocals(_ *pulumi.Context, stackInput *gcpvpcv1.GcpVpcStackInput) *Locals {
	locals := &Locals{}

	locals.GcpVpc = stackInput.Target

	// Standard Planton‑wide labels for GCP resources
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: locals.GcpVpc.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpVpc.String(),
	}

	if locals.GcpVpc.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpVpc.Metadata.Org
	}

	if locals.GcpVpc.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpVpc.Metadata.Env
	}

	if locals.GcpVpc.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GcpVpc.Metadata.Id
	}

	locals.GcpCredentialSpec = stackInput.ProviderCredential

	return locals
}
