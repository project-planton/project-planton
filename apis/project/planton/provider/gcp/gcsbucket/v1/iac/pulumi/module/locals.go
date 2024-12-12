package module

import (
	gcsbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcsbucket/v1"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpLabels map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcsbucketv1.GcsBucketStackInput) *Locals {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	gcsBucket := stackInput.Target
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceId:   gcsBucket.Metadata.Id,
		gcplabelkeys.ResourceKind: "gcs_bucket",
	}

	if gcsBucket.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = gcsBucket.Metadata.Org
	}

	return locals
}
