package module

import (
	"strconv"

	"github.com/project-planton/project-planton/pkg/shared/cloudresourcekind"

	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	awskmskeyv1 "github.com/project-planton/project-planton/pkg/provider/aws/awskmskey/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsKmsKey *awskmskeyv1.AwsKmsKey
	AwsTags   map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awskmskeyv1.AwsKmsKeyStackInput) *Locals {
	locals := &Locals{}
	locals.AwsKmsKey = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsKmsKey.Metadata.Org,
		awstagkeys.Environment:  locals.AwsKmsKey.Metadata.Env,
		awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsKmsKey.String(),
		awstagkeys.ResourceId:   locals.AwsKmsKey.Metadata.Id,
	}

	return locals
}
