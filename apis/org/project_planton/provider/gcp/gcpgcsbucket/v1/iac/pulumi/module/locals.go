package module

import (
	"strconv"
	"strings"

	gcpgcsbucketv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpgcsbucket/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	GcpLabels    map[string]string
	GcpGcsBucket *gcpgcsbucketv1.GcpGcsBucket
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpgcsbucketv1.GcpGcsBucketStackInput) *Locals {
	locals := &Locals{}

	locals.GcpGcsBucket = stackInput.Target

	target := stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: strings.ToLower(cloudresourcekind.CloudResourceKind_GcpGcsBucket.String()),
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
