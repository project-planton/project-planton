package module

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	ecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func service(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	originalSpec := locals.EcsService.Spec

	convertedSpec := LocalsEcsServiceSpec{
		ClusterName:          originalSpec.ClusterName,
		ServiceName:          originalSpec.ServiceName,
		Image:                originalSpec.Image,
		ContainerPort:        originalSpec.ContainerPort,
		DesiredCount:         int(originalSpec.DesiredCount),
		Cpu:                  originalSpec.Cpu,
		Memory:               originalSpec.Memory,
		Subnets:              originalSpec.Subnets,
		SecurityGroups:       originalSpec.SecurityGroups,
		AssignPublicIp:       originalSpec.AssignPublicIp,
		TaskExecutionRoleArn: originalSpec.TaskExecutionRoleArn,
		TaskRoleArn:          originalSpec.TaskRoleArn,
		Environment:          convertEnvironment(originalSpec.Environment),
	}

	serviceName := convertedSpec.ServiceName
	clusterName := convertedSpec.ClusterName

	containerDefs, err := buildContainerDefinitions(&convertedSpec)
	if err != nil {
		return errors.Wrap(err, "failed to build container definitions JSON")
	}

	taskDef, err := ecs.NewTaskDefinition(ctx, locals.EcsService.Metadata.Name+"-taskdef", &ecs.TaskDefinitionArgs{
		Family:                  pulumi.String(serviceName),
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

	ecsService, err := ecs.NewService(ctx, locals.EcsService.Metadata.Name+"-service", &ecs.ServiceArgs{
		Name:           pulumi.String(serviceName),
		Cluster:        pulumi.String(clusterName),
		LaunchType:     pulumi.String("FARGATE"),
		DesiredCount:   pulumi.Int(convertedSpec.DesiredCount),
		TaskDefinition: taskDef.Arn,
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			AssignPublicIp: pulumi.Bool(convertedSpec.AssignPublicIp),
			Subnets:        toPulumiStrings(convertedSpec.Subnets),
			SecurityGroups: toPulumiStrings(convertedSpec.SecurityGroups),
		},
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS service")
	}

	ctx.Export(OpEcsServiceName, ecsService.Name)
	ctx.Export(OpEcsClusterName, pulumi.String(clusterName))
	ctx.Export(OpLoadBalancerDnsName, pulumi.String(""))
	ctx.Export(OpServiceUrl, pulumi.String(""))
	ctx.Export(OpServiceDiscoveryName, pulumi.String(""))

	return nil
}

func buildContainerDefinitions(spec *LocalsEcsServiceSpec) (string, error) {
	var envVars []map[string]string
	for _, env := range spec.Environment {
		envVars = append(envVars, map[string]string{
			"name":  env.Name,
			"value": env.Value,
		})
	}

	container := map[string]interface{}{
		"name":        spec.ServiceName,
		"image":       spec.Image,
		"essential":   true,
		"environment": envVars,
	}

	if spec.ContainerPort != 0 {
		container["portMappings"] = []map[string]int32{
			{
				"containerPort": spec.ContainerPort,
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
func convertEnvironment(envList []*ecsservicev1.EcsServiceSpec_EnvironmentVar) []EcsEnvironmentVar {
	var converted []EcsEnvironmentVar
	for _, e := range envList {
		converted = append(converted, EcsEnvironmentVar{
			Name:  e.Name,
			Value: e.Value,
		})
	}
	return converted
}

type LocalsEcsServiceSpec struct {
	ClusterName          string
	ServiceName          string
	Image                string
	ContainerPort        int32
	DesiredCount         int
	Cpu                  int32
	Memory               int32
	Subnets              []string
	SecurityGroups       []string
	AssignPublicIp       bool
	TaskExecutionRoleArn string
	TaskRoleArn          string
	Environment          []EcsEnvironmentVar
}

type EcsEnvironmentVar struct {
	Name  string
	Value string
}
