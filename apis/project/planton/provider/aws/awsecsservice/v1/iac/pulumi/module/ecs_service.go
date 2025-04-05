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
	originalSpec := locals.AwsEcsService.Spec

	convertedSpec := LocalsAwsEcsServiceSpec{
		ClusterArn:     originalSpec.ClusterArn,
		ServiceName:    locals.AwsEcsService.Metadata.Name,
		ImageRepo:      originalSpec.Container.Image.Repo,
		ImageTag:       originalSpec.Container.Image.Tag,
		Port:           originalSpec.Container.Port,
		Replicas:       int(originalSpec.Container.Replicas),
		Cpu:            originalSpec.Container.Cpu,
		Memory:         originalSpec.Container.Memory,
		Subnets:        originalSpec.Network.Subnets,
		SecurityGroups: originalSpec.Network.SecurityGroups,
	}

	if originalSpec.Iam != nil {
		convertedSpec.TaskExecutionRoleArn = originalSpec.Iam.TaskExecutionRoleArn
		convertedSpec.TaskRoleArn = originalSpec.Iam.TaskRoleArn
	}

	if originalSpec.Alb != nil {
		convertedSpec.AlbEnabled = originalSpec.Alb.Enabled
		convertedSpec.AlbArn = originalSpec.Alb.Arn
		convertedSpec.AlbRoutingType = originalSpec.Alb.RoutingType.String()
		convertedSpec.AlbPath = originalSpec.Alb.Path
		convertedSpec.AlbHostname = originalSpec.Alb.Hostname
	}

	containerDefs, err := buildContainerDefinitions(&convertedSpec)
	if err != nil {
		return errors.Wrap(err, "failed to build container definitions JSON")
	}

	taskDef, err := ecs.NewTaskDefinition(ctx, convertedSpec.ServiceName+"-taskdef", &ecs.TaskDefinitionArgs{
		Family:                  pulumi.String(convertedSpec.ServiceName),
		RequiresCompatibilities: pulumi.StringArray{pulumi.String("FARGATE")},
		Cpu:                     pulumi.String(fmt.Sprintf("%d", convertedSpec.Cpu)),
		Memory:                  pulumi.String(fmt.Sprintf("%d", convertedSpec.Memory)),
		NetworkMode:             pulumi.String("awsvpc"),
		ExecutionRoleArn:        pulumi.String(convertedSpec.TaskExecutionRoleArn),
		TaskRoleArn:             pulumi.String(convertedSpec.TaskRoleArn),
		ContainerDefinitions:    pulumi.String(containerDefs),
		Tags:                    pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS task definition")
	}

	serviceArgs := &ecs.ServiceArgs{
		Name:           pulumi.String(convertedSpec.ServiceName),
		Cluster:        pulumi.String(convertedSpec.ClusterArn),
		LaunchType:     pulumi.String("FARGATE"),
		DesiredCount:   pulumi.Int(convertedSpec.Replicas),
		TaskDefinition: taskDef.Arn,
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			Subnets:        toPulumiStrings(convertedSpec.Subnets),
			SecurityGroups: toPulumiStrings(convertedSpec.SecurityGroups),
		},
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}

	if convertedSpec.AlbEnabled && convertedSpec.Port != 0 {
		if len(convertedSpec.Subnets) == 0 {
			return errors.New("at least one subnet is required for ALB usage")
		}
		firstSubnetID := convertedSpec.Subnets[0]

		subnetLookup, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Id: &firstSubnetID,
		}, nil)
		if err != nil {
			return errors.Wrap(err, "failed to lookup first subnet (needed for ALB target group)")
		}

		targetGroup, err := lb.NewTargetGroup(ctx, convertedSpec.ServiceName+"-tg", &lb.TargetGroupArgs{
			Port:       pulumi.Int(int(convertedSpec.Port)),
			Protocol:   pulumi.String("HTTP"),
			TargetType: pulumi.String("ip"),
			VpcId:      pulumi.String(subnetLookup.VpcId),
			HealthCheck: &lb.TargetGroupHealthCheckArgs{
				Path: pulumi.String("/"),
			},
			Tags: pulumi.ToStringMap(locals.AwsTags),
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrap(err, "failed to create ALB target group")
		}

		serviceArgs.LoadBalancers = ecs.ServiceLoadBalancerArray{
			&ecs.ServiceLoadBalancerArgs{
				TargetGroupArn: targetGroup.Arn,
				ContainerName:  pulumi.String(convertedSpec.ServiceName),
				ContainerPort:  pulumi.Int(int(convertedSpec.Port)),
			},
		}
	}

	awsEcsService, err := ecs.NewService(ctx, convertedSpec.ServiceName+"-service", serviceArgs, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS service")
	}

	ctx.Export(OpAwsEcsServiceName, awsEcsService.Name)
	ctx.Export(OpEcsClusterName, pulumi.String(convertedSpec.ClusterArn))
	ctx.Export(OpLoadBalancerDnsName, pulumi.String(""))
	ctx.Export(OpServiceUrl, pulumi.String(""))
	ctx.Export(OpServiceDiscoveryName, pulumi.String(""))

	return nil
}

func buildContainerDefinitions(spec *LocalsAwsEcsServiceSpec) (string, error) {
	var envVars []map[string]string

	container := map[string]interface{}{
		"name":        spec.ServiceName,
		"image":       fmt.Sprintf("%s:%s", spec.ImageRepo, spec.ImageTag),
		"essential":   true,
		"environment": envVars,
	}

	if spec.Port != 0 {
		container["portMappings"] = []map[string]int32{
			{
				"containerPort": spec.Port,
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

func toPulumiStrings(input []string) pulumi.StringArray {
	output := make(pulumi.StringArray, len(input))
	for i, s := range input {
		output[i] = pulumi.String(s)
	}
	return output
}
