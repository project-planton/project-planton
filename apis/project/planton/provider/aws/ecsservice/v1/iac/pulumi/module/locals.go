package module

import (
	ecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	EcsService *ecsservicev1.EcsService
	AwsTags    map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *ecsservicev1.EcsServiceStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.EcsService = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.EcsService.Metadata.Org,
		awstagkeys.Environment:  locals.EcsService.Metadata.Env,
		awstagkeys.ResourceKind: string(apiresourcekind.EcsServiceKind),
		awstagkeys.ResourceId:   locals.EcsService.Metadata.Id,
	}

	return locals
}
