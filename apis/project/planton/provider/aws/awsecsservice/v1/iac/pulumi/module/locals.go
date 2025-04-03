package module

import (
	"strconv"

	awsecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecsservice/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals defines the local variables/fields used throughout our Pulumi module.
// This struct includes the AWS ECS Service resource definition from the stack input
// and a map of AWS tags we'll apply to resources.
type Locals struct {
	AwsEcsService *awsecsservicev1.AwsEcsService
	AwsTags       map[string]string
}

// initializeLocals pulls values from the stack input (AwsEcsServiceStackInput)
// and populates the Locals struct. We mimic Terraform "locals" by storing
// precomputed values or referencing resource fields directly.
func initializeLocals(ctx *pulumi.Context, stackInput *awsecsservicev1.AwsEcsServiceStackInput) *Locals {
	locals := &Locals{}

	locals.AwsEcsService = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsEcsService.Metadata.Org,
		awstagkeys.Environment:  locals.AwsEcsService.Metadata.Env,
		awstagkeys.ResourceKind: string(apiresourcekind.AwsEcsServiceKind),
		awstagkeys.ResourceId:   locals.AwsEcsService.Metadata.Id,
	}

	return locals
}
