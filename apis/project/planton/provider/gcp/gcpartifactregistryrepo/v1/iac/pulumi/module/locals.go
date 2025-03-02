package module

import (
	gcpartifactregistryrepov1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistryrepo/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpArtifactRegistryRepo *gcpartifactregistryrepov1.GcpArtifactRegistryRepo
	GcpLabels               map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpartifactregistryrepov1.GcpArtifactRegistryRepoStackInput) *Locals {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	//assign value for the locals variable to make it available across the project
	locals.GcpArtifactRegistryRepo = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceId:   locals.GcpArtifactRegistryRepo.Metadata.Id,
		gcplabelkeys.ResourceKind: "gcp_artifact_registry",
	}

	if locals.GcpArtifactRegistryRepo.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpArtifactRegistryRepo.Metadata.Org
	}

	return locals
}
