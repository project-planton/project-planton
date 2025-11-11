package module

import (
	"strconv"

	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	awsdynamodbv1 "github.com/project-planton/project-planton/pkg/provider/aws/awsdynamodb/v1"
	"github.com/project-planton/project-planton/pkg/shared/cloudresourcekind"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	Target  *awsdynamodbv1.AwsDynamodb
	Spec    *awsdynamodbv1.AwsDynamodbSpec
	AwsTags map[string]string
}

func initializeLocals(ctx *pulumi.Context, in *awsdynamodbv1.AwsDynamodbStackInput) *Locals {
	locals := &Locals{}
	locals.Target = in.Target
	locals.Spec = in.Target.Spec

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.Target.Metadata.Org,
		awstagkeys.Environment:  locals.Target.Metadata.Env,
		awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsDynamodb.String(),
		awstagkeys.ResourceId:   locals.Target.Metadata.Id,
	}
	return locals
}
