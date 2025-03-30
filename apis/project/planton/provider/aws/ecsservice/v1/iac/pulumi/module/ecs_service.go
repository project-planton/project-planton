package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ecsService creates the ECS service in the given clusterName with the provided TaskDefinition.
func ecsService(
	ctx *pulumi.Context,
	locals *Locals,
	provider *aws.Provider,
	clusterName string,
	taskDef *ecs.TaskDefinition,
) (*ecs.Service, error) {
	spec := locals.EcsService.Spec

	launchType := spec.LaunchType.String()
	if launchType == "ecs_launch_type_unspecified" {
		launchType = "FARGATE"
	}

	serviceArgs := &ecs.ServiceArgs{
		Name:           pulumi.String(locals.EcsService.Metadata.Name),
		Cluster:        pulumi.String(clusterName),
		TaskDefinition: taskDef.Arn,
		DesiredCount:   pulumi.Int(spec.DesiredCount),
		LaunchType:     pulumi.String(launchType),
		DeploymentCircuitBreaker: &ecs.ServiceDeploymentCircuitBreakerArgs{
			Enable:   pulumi.Bool(true),
			Rollback: pulumi.Bool(true),
		},
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			AssignPublicIp: pulumi.Bool(spec.Network.AssignPublicIp),
			SecurityGroups: pulumi.ToStringArray(spec.Network.SecurityGroupIds),
			Subnets:        pulumi.ToStringArray(spec.Network.SubnetIds),
		},
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}

	svc, err := ecs.NewService(ctx,
		fmt.Sprintf("%s-ecs-service", locals.EcsService.Metadata.Name),
		serviceArgs,
		pulumi.Provider(provider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ECS service")
	}

	return svc, nil
}
