package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// taskDefinition builds the ECS Task Definition from the EcsServiceSpec fields.
func taskDefinition(
	ctx *pulumi.Context,
	locals *Locals,
	provider *aws.Provider,
) (*ecs.TaskDefinition, error) {
	spec := locals.EcsService.Spec

	envMap := spec.EnvironmentVariables
	secretMap := spec.SecretVariables

	containerDef := fmt.Sprintf(`[
	{
		"name": "%s",
		"image": "%s",
		"portMappings": [
			{
				"containerPort": %d,
				"hostPort": %d
			}
		],
		"environment": %s,
		"secrets": %s,
		"essential": true
	}
]`,
		locals.EcsService.Metadata.Name,
		spec.ContainerImage,
		spec.ContainerPort,
		spec.ContainerPort,
		convertMapToEcsEnvFormat(envMap),
		convertMapToEcsSecretFormat(secretMap),
	)

	launchType := spec.LaunchType.String()
	if launchType == "ecs_launch_type_unspecified" {
		launchType = "FARGATE"
	}

	cpu := fmt.Sprintf("%d", spec.Cpu)
	memory := fmt.Sprintf("%d", spec.Memory)
	if spec.Cpu == 0 {
		cpu = "256"
	}
	if spec.Memory == 0 {
		memory = "512"
	}

	taskDef, err := ecs.NewTaskDefinition(ctx,
		fmt.Sprintf("%s-task-def", locals.EcsService.Metadata.Name),
		&ecs.TaskDefinitionArgs{
			Family:                  pulumi.String(fmt.Sprintf("%s-family", locals.EcsService.Metadata.Name)),
			Cpu:                     pulumi.String(cpu),
			Memory:                  pulumi.String(memory),
			NetworkMode:             pulumi.String("awsvpc"),
			RequiresCompatibilities: pulumi.StringArray{pulumi.String(launchType)},
			ContainerDefinitions:    pulumi.String(containerDef),
			Tags:                    pulumi.ToStringMap(locals.AwsTags),
		},
		pulumi.Provider(provider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ECS task definition")
	}

	return taskDef, nil
}
