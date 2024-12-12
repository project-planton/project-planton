package module

import (
	gcpartifactregistryv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistry/v1"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpArtifactRegistry *gcpartifactregistryv1.GcpArtifactRegistry
	GcpLabels           map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpartifactregistryv1.GcpArtifactRegistryStackInput) *Locals {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	//assign value for the locals variable to make it available across the project
	locals.GcpArtifactRegistry = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceId:   locals.GcpArtifactRegistry.Metadata.Id,
		gcplabelkeys.ResourceKind: "gcp_artifact_registry",
	}

	if locals.GcpArtifactRegistry.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpArtifactRegistry.Metadata.Org
	}

	return locals
}
