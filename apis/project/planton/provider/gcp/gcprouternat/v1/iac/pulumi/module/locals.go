package module

import (
	"strconv"

	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gcprouternatv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcprouternat/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps frequently used values (metadata, labels, credentials) handy for the module.
type Locals struct {
	GcpCredentialSpec *gcpcredentialv1.GcpCredentialSpec
	GcpRouterNat      *gcprouternatv1.GcpRouterNat
	GcpLabels         map[string]string
}

// initializeLocals populates the Locals struct from the stack input.
// It mirrors the pattern used in the gcp_router_nat module and applies the same Planton label strategy.
func initializeLocals(_ *pulumi.Context, stackInput *gcprouternatv1.GcpRouterNatStackInput) *Locals {
	locals := &Locals{}

	locals.GcpRouterNat = stackInput.Target

	// Standard Planton-wide labels for GCP resources
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: locals.GcpRouterNat.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpRouterNat.String(),
	}

	if locals.GcpRouterNat.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpRouterNat.Metadata.Org
	}

	if locals.GcpRouterNat.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpRouterNat.Metadata.Env
	}

	if locals.GcpRouterNat.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = locals.GcpRouterNat.Metadata.Id
	}

	locals.GcpCredentialSpec = stackInput.ProviderCredential

	return locals
}
