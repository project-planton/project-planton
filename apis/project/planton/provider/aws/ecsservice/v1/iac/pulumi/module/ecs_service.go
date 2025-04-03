package module

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// service creates and wires up the ECS Task Definition and ECS Service resources.
func service(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	originalSpec := locals.EcsService.Spec

	convertedSpec := LocalsEcsServiceSpec{
		ClusterArn:           originalSpec.ClusterArn,
		ServiceName:          locals.EcsService.Metadata.Name,
		ImageRepo:            originalSpec.Container.Image.Repo,
		ImageTag:             originalSpec.Container.Image.Tag,
		Port:                 originalSpec.Container.Port,
		Replicas:             int(originalSpec.Container.Replicas),
		Cpu:                  originalSpec.Container.Cpu,
		Memory:               originalSpec.Container.Memory,
		Subnets:              originalSpec.Network.Subnets,
		SecurityGroups:       originalSpec.Network.SecurityGroups,
		AssignPublicIp:       originalSpec.Network.AssignPublicIp,
		TaskExecutionRoleArn: originalSpec.Iam.TaskExecutionRoleArn,
		TaskRoleArn:          originalSpec.Iam.TaskRoleArn,
		EnvVariables:         originalSpec.Container.Env.Variables,
		EnvSecrets:           originalSpec.Container.Env.Secrets,
	}

	containerDefs, err := buildContainerDefinitions(&convertedSpec)
	if err != nil {
		return errors.Wrap(err, "failed to build container definitions JSON")
	}

	taskDef, err := ecs.NewTaskDefinition(ctx, locals.EcsService.Metadata.Name+"-taskdef", &ecs.TaskDefinitionArgs{
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

	ecsService, err := ecs.NewService(ctx, locals.EcsService.Metadata.Name+"-service", &ecs.ServiceArgs{
		Name:           pulumi.String(convertedSpec.ServiceName),
		Cluster:        pulumi.String(convertedSpec.ClusterArn),
		LaunchType:     pulumi.String("FARGATE"),
		DesiredCount:   pulumi.Int(convertedSpec.Replicas),
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
	ctx.Export(OpEcsClusterName, pulumi.String(convertedSpec.ClusterArn))
	ctx.Export(OpLoadBalancerDnsName, pulumi.String(""))
	ctx.Export(OpServiceUrl, pulumi.String(""))
	ctx.Export(OpServiceDiscoveryName, pulumi.String(""))

	return nil
}

// buildContainerDefinitions builds a JSON array of container definitions
// based on our local ECS service spec.
func buildContainerDefinitions(spec *LocalsEcsServiceSpec) (string, error) {
	var envVars []map[string]string

	for k, v := range spec.EnvVariables {
		envVars = append(envVars, map[string]string{
			"name":  k,
			"value": v,
		})
	}

	for k, v := range spec.EnvSecrets {
		envVars = append(envVars, map[string]string{
			"name":  k,
			"value": v,
		})
	}

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

// toPulumiStrings is a helper that converts a native string slice to a pulumi.StringInput slice.
func toPulumiStrings(input []string) pulumi.StringArray {
	output := make(pulumi.StringArray, len(input))
	for i, s := range input {
		output[i] = pulumi.String(s)
	}
	return output
}

// LocalsEcsServiceSpec is an internal struct that adapts the new EcsServiceSpec
// fields into something easier for building ECS resources.
type LocalsEcsServiceSpec struct {
	ClusterArn           string
	ServiceName          string
	ImageRepo            string
	ImageTag             string
	Port                 int32
	Replicas             int
	Cpu                  int32
	Memory               int32
	Subnets              []string
	SecurityGroups       []string
	AssignPublicIp       bool
	TaskExecutionRoleArn string
	TaskRoleArn          string
	EnvVariables         map[string]string
	EnvSecrets           map[string]string
}
