package module

import (
	"strconv"

	awslambdav1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awslambda/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsLambda *awslambdav1.AWSLambda
	AwsTags   map[string]string
}

func initializeLocals(ctx *pulumi.Context, in *awslambdav1.AWSLambdaStackInput) *Locals {
	locals := &Locals{
		AwsLambda: in.Target,
	}

	if in.Target != nil {
		locals.AwsTags = map[string]string{
			awstagkeys.Resource:     strconv.FormatBool(true),
			awstagkeys.Organization: in.Target.Metadata.Org,
			awstagkeys.Environment:  in.Target.Metadata.Env,
			awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsLambda.String(),
			awstagkeys.ResourceId:   in.Target.Metadata.Id,
			awstagkeys.Name:         in.Target.Metadata.Name,
		}
	} else {
		locals.AwsTags = map[string]string{}
	}

	return locals
}
