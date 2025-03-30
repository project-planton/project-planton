package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ecsCluster returns the ECS cluster name. If spec.ClusterName is empty, it creates a new cluster.
func ecsCluster(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*ecs.Cluster, error) {
	newCluster, err := ecs.NewCluster(ctx,
		fmt.Sprintf("%s-cluster", locals.EcsService.Metadata.Name),
		&ecs.ClusterArgs{
			Name: pulumi.String(fmt.Sprintf("%s-default-cluster", locals.EcsService.Metadata.Name)),
		},
		pulumi.Provider(provider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create default ECS cluster")
	}

	// Return the cluster name
	return newCluster, nil
}
