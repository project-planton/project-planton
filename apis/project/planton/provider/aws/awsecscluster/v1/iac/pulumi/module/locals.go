package module

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"strconv"

	awsecsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecscluster/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsEcsCluster *awsecsclusterv1.AwsEcsCluster
	AwsTags       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsecsclusterv1.AwsEcsClusterStackInput) *Locals {
	locals := &Locals{}

	locals.AwsEcsCluster = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsEcsCluster.Metadata.Org,
		awstagkeys.Environment:  locals.AwsEcsCluster.Metadata.Env,
		awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsEcsCluster.String(),
		awstagkeys.ResourceId:   locals.AwsEcsCluster.Metadata.Id,
	}

	return locals
}
