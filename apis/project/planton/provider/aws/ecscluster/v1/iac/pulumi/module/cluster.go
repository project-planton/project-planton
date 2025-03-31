package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func cluster(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	clusterName := locals.EcsCluster.Metadata.Name

	ecsCluster, err := ecs.NewCluster(ctx, locals.EcsCluster.Metadata.Name, &ecs.ClusterArgs{
		Name: pulumi.String(clusterName),
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to create ECS cluster: %s", clusterName))
	}

	ctx.Export(OpClusterName, ecsCluster.Name)
	ctx.Export(OpClusterArn, ecsCluster.Arn)

	return nil
}
