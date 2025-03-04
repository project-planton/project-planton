package module

import (
	gcpdnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpdnszone/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpDnsZone *gcpdnszonev1.GcpDnsZone
	GcpLabels  map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpdnszonev1.GcpDnsZoneStackInput) *Locals {
	locals := &Locals{}

	locals.GcpDnsZone = stackInput.Target

	target := stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: string(apiresourcekind.GcpDnsZoneKind),
	}

	if target.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = target.Metadata.Env
	}

	return locals
}
