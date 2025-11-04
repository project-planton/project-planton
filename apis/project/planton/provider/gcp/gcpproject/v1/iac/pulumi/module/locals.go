package module

import (
	"strconv"
	"strings"

	gcpprojectv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpproject/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	GcpProject *gcpprojectv1.GcpProject
	GcpLabels  map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpprojectv1.GcpProjectStackInput) *Locals {
	locals := &Locals{}
	locals.GcpProject = stackInput.Target

	target := stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: strings.ToLower(cloudresourcekind.CloudResourceKind_GcpProject.String()),
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
