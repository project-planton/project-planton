package module

import (
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	AwsDynamodb *awsdynamodbv1.AwsDynamodb
	AwsTags     map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.AwsDynamodb = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsDynamodb.Metadata.Org,
		awstagkeys.Environment:  locals.AwsDynamodb.Metadata.Env,
		awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsEcsService.String(),
		awstagkeys.ResourceId:   locals.AwsDynamodb.Metadata.Id,
	}

	return locals
}
