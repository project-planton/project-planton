package module

import (
	"strconv"

	ecsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecscluster/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	EcsCluster *ecsclusterv1.EcsCluster
	AwsTags    map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *ecsclusterv1.EcsClusterStackInput) *Locals {
	locals := &Locals{}

	locals.EcsCluster = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.EcsCluster.Metadata.Org,
		awstagkeys.Environment:  locals.EcsCluster.Metadata.Env,
		awstagkeys.ResourceKind: string(apiresourcekind.EcsClusterKind),
		awstagkeys.ResourceId:   locals.EcsCluster.Metadata.Id,
	}

	return locals
}
