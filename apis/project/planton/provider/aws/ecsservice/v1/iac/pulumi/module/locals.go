package module

import (
	"strconv"

	ecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals defines the local variables/fields used throughout our Pulumi module.
// This struct includes the ECS Service resource definition from the stack input
// and a map of AWS tags we'll apply to resources.
type Locals struct {
	EcsService *ecsservicev1.EcsService
	AwsTags    map[string]string
}

// initializeLocals pulls values from the stack input (EcsServiceStackInput)
// and populates the Locals struct. We mimic Terraform "locals" by storing
// precomputed values or referencing resource fields directly.
func initializeLocals(ctx *pulumi.Context, stackInput *ecsservicev1.EcsServiceStackInput) *Locals {
	locals := &Locals{}

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
