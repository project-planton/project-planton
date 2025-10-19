package module

import (
	"strconv"

	gcprouternatv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcprouternat/v1"
	gcpprovider "github.com/project-planton/project-planton/apis/project/planton/provider/gcp"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals collects frequently used input values and derived labels.
type Locals struct {
	GcpProviderConfig *gcpprovider.GcpProviderConfig
	GcpRouterNat      *gcprouternatv1.GcpRouterNat
	GcpLabels         map[string]string
}

// initializeLocals converts the stack‑input into a struct that is easy to reference
// (mirrors the Terraform “locals” pattern).
func initializeLocals(_ *pulumi.Context, stackInput *gcprouternatv1.GcpRouterNatStackInput) *Locals {
	target := stackInput.Target

	labels := map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_GcpRouterNat.String(),
	}

	if target.Metadata.Org != "" {
		labels[gcplabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		labels[gcplabelkeys.Environment] = target.Metadata.Env
	}

	if target.Metadata.Id != "" {
		labels[gcplabelkeys.ResourceId] = target.Metadata.Id
	}

	return &Locals{
		GcpProviderConfig: stackInput.ProviderConfig,
		GcpRouterNat:      target,
		GcpLabels:         labels,
	}
}
