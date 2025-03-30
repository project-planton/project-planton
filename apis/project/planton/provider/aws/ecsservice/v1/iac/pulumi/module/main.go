package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	v1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
)

// Resources is the main entry point for setting up ECS cluster, task definition, and service.
func Resources(ctx *pulumi.Context, stackInput *v1.EcsServiceStackInput) error {
	locals := initializeLocals(ctx, stackInput)
	awsCredential := stackInput.ProviderCredential

	var provider *aws.Provider
	var err error

	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx,
			"ecsServiceProvider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with default credentials")
		}
	} else {
		provider, err = aws.NewProvider(ctx,
			"ecsServiceProvider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with explicit credentials")
		}
	}

	if locals.EcsService.Spec.ClusterName == "" {
		// 1. Create or reference the ECS cluster
		createdEcsCluster, err := ecsCluster(ctx, locals, provider)
		if err != nil {
			return errors.Wrap(err, "failed to handle ECS cluster")
		}
	}

	// 2. Create the ECS task definition
	taskDef, err := taskDefinition(ctx, locals, provider, clusterName)
	if err != nil {
		return errors.Wrap(err, "failed to create ECS task definition")
	}

	// 3. Create the ECS service
	ecsSvc, err := ecsService(ctx, locals, provider, clusterName, taskDef)
	if err != nil {
		return errors.Wrap(err, "failed to create ECS service")
	}

	// Export cluster name, service name
	ctx.Export("ecs_cluster_name", pulumi.String(clusterName))
	ctx.Export("ecs_service_name", ecsSvc.Name)

	return nil
}
