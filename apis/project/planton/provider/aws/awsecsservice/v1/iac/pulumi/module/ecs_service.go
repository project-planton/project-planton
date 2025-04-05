package module

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// service creates and wires up the ECS Task Definition and AWS ECS Service resources.
func service(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	spec := locals.AwsEcsService.Spec
	serviceName := locals.AwsEcsService.Metadata.Name

	containerDefs, err := buildContainerDefinitions(
		serviceName,
		spec.Container.Image.Repo,
		spec.Container.Image.Tag,
		spec.Container.Port,
	)
	if err != nil {
		return errors.Wrap(err, "failed to build container definitions JSON")
	}

	taskDefinitionArgs := &ecs.TaskDefinitionArgs{
		Family:                  pulumi.String(serviceName),
		RequiresCompatibilities: pulumi.StringArray{pulumi.String("FARGATE")},
		Cpu:                     pulumi.String(fmt.Sprintf("%d", spec.Container.Cpu)),
		Memory:                  pulumi.String(fmt.Sprintf("%d", spec.Container.Memory)),
		NetworkMode:             pulumi.String("awsvpc"),
		ContainerDefinitions:    pulumi.String(containerDefs),
		Tags:                    pulumi.ToStringMap(locals.AwsTags),
	}

	if spec.Iam != nil {
		taskDefinitionArgs.ExecutionRoleArn = pulumi.String(spec.Iam.GetTaskExecutionRoleArn())
		taskDefinitionArgs.TaskRoleArn = pulumi.String(spec.Iam.GetTaskRoleArn())
	}

	taskDef, err := ecs.NewTaskDefinition(ctx,
		serviceName+"-taskdef",
		taskDefinitionArgs,
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS task definition")
	}

	// Build the service arguments
	serviceArgs := &ecs.ServiceArgs{
		Name:           pulumi.String(serviceName),
		Cluster:        pulumi.String(spec.ClusterArn),
		LaunchType:     pulumi.String("FARGATE"),
		DesiredCount:   pulumi.Int(int(spec.Container.Replicas)),
		TaskDefinition: taskDef.Arn,
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			Subnets:        toPulumiStrings(spec.Network.Subnets),
			SecurityGroups: toPulumiStrings(spec.Network.SecurityGroups),
		},
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}

	// If ALB is enabled and we have a non-zero container port, create a Target Group and attach
	if spec.Alb.GetEnabled() && spec.Container.Port != 0 {
		if len(spec.Network.Subnets) == 0 {
			return errors.New("at least one subnet is required for ALB usage")
		}
		firstSubnetID := spec.Network.Subnets[0]

		subnetLookup, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Id: &firstSubnetID,
		}, nil)
		if err != nil {
			return errors.Wrap(err, "failed to lookup first subnet (needed for ALB target group)")
		}

		targetGroup, err := lb.NewTargetGroup(ctx, serviceName+"-tg", &lb.TargetGroupArgs{
			Port:        pulumi.Int(int(spec.Container.Port)),
			Protocol:    pulumi.String("HTTP"),
			TargetType:  pulumi.String("ip"),
			VpcId:       pulumi.String(subnetLookup.VpcId),
			HealthCheck: &lb.TargetGroupHealthCheckArgs{Path: pulumi.String("/")},
			Tags:        pulumi.ToStringMap(locals.AwsTags),
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrap(err, "failed to create ALB target group")
		}

		serviceArgs.LoadBalancers = ecs.ServiceLoadBalancerArray{
			&ecs.ServiceLoadBalancerArgs{
				TargetGroupArn: targetGroup.Arn,
				ContainerName:  pulumi.String(serviceName),
				ContainerPort:  pulumi.Int(int(spec.Container.Port)),
			},
		}
	}

	awsEcsService, err := ecs.NewService(ctx, serviceName+"-service", serviceArgs, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS service")
	}

	// Export relevant outputs
	ctx.Export(OpAwsEcsServiceName, awsEcsService.Name)
	ctx.Export(OpEcsClusterName, pulumi.String(spec.ClusterArn))
	ctx.Export(OpLoadBalancerDnsName, pulumi.String(""))
	ctx.Export(OpServiceUrl, pulumi.String(""))
	ctx.Export(OpServiceDiscoveryName, pulumi.String(""))

	return nil
}

// buildContainerDefinitions constructs a JSON array of container definitions.
func buildContainerDefinitions(serviceName, repo, tag string, port int32) (string, error) {
	var envVars []map[string]string // left empty for brevity

	container := map[string]interface{}{
		"name":        serviceName,
		"image":       fmt.Sprintf("%s:%s", repo, tag),
		"essential":   true,
		"environment": envVars,
	}

	if port != 0 {
		container["portMappings"] = []map[string]int32{
			{
				"containerPort": port,
			},
		}
	}

	containerDefinitions := []interface{}{container}
	encoded, err := json.Marshal(containerDefinitions)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

// toPulumiStrings is a helper that converts a native string slice to a pulumi.StringInput slice.
func toPulumiStrings(input []string) pulumi.StringArray {
	output := make(pulumi.StringArray, len(input))
	for i, s := range input {
		output[i] = pulumi.String(s)
	}
	return output
}
