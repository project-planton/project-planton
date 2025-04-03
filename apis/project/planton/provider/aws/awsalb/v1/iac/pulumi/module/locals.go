package module

import (
	"strconv"

	awsalbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsalb/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals holds the AWS ALB resource definition from the stack input
// and a map of AWS tags to apply to resources.
type Locals struct {
	AwsAlb  *awsalbv1.AwsAlb
	AwsTags map[string]string
}

// initializeLocals is analogous to Terraform "locals." It reads
// values from AwsAlbStackInput to build a Locals instance.
func initializeLocals(ctx *pulumi.Context, stackInput *awsalbv1.AwsAlbStackInput) *Locals {
	locals := &Locals{}

	locals.AwsAlb = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsAlb.Metadata.Org,
		awstagkeys.Environment:  locals.AwsAlb.Metadata.Env,
		awstagkeys.ResourceKind: string(apiresourcekind.AwsAlbKind),
		awstagkeys.ResourceId:   locals.AwsAlb.Metadata.Id,
	}

	return locals
}
