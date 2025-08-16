package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func cluster(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	clusterName := locals.AwsEcsCluster.Metadata.Name

	args := &ecs.ClusterArgs{
		Name: pulumi.String(clusterName),
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}

	// Optional: enable CloudWatch Container Insights
	if locals.AwsEcsCluster.Spec.EnableContainerInsights {
		args.Settings = ecs.ClusterSettingArray{
			&ecs.ClusterSettingArgs{
				Name:  pulumi.String("containerInsights"),
				Value: pulumi.String("enabled"),
			},
		}
	}

	// Optional: enable ECS Exec
	if locals.AwsEcsCluster.Spec.EnableExecuteCommand {
		args.Configuration = &ecs.ClusterConfigurationArgs{
			ExecuteCommandConfiguration: &ecs.ClusterConfigurationExecuteCommandConfigurationArgs{
				Logging: pulumi.String("DEFAULT"),
			},
		}
	}

	awsEcsCluster, err := ecs.NewCluster(ctx, locals.AwsEcsCluster.Metadata.Name, args, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to create ECS cluster: %s", clusterName))
	}

	// Optional: capacity providers (managed via separate resource)
	if len(locals.AwsEcsCluster.Spec.CapacityProviders) > 0 {
		var cps pulumi.StringArray
		for _, cp := range locals.AwsEcsCluster.Spec.CapacityProviders {
			cps = append(cps, pulumi.String(cp))
		}
		_, err := ecs.NewClusterCapacityProviders(ctx,
			fmt.Sprintf("%s-capacity-providers", locals.AwsEcsCluster.Metadata.Name),
			&ecs.ClusterCapacityProvidersArgs{
				ClusterName:       awsEcsCluster.Name,
				CapacityProviders: cps,
			}, pulumi.Provider(provider), pulumi.Parent(awsEcsCluster))
		if err != nil {
			return errors.Wrap(err, "unable to attach capacity providers")
		}
	}

	ctx.Export(OpClusterName, awsEcsCluster.Name)
	ctx.Export(OpClusterArn, awsEcsCluster.Arn)
	// Export capacity providers if configured
	if len(locals.AwsEcsCluster.Spec.CapacityProviders) > 0 {
		for i, cp := range locals.AwsEcsCluster.Spec.CapacityProviders {
			ctx.Export(fmt.Sprintf("%s.%d", OpClusterCapacityProviders, i), pulumi.String(cp))
		}
	}

	return nil
}
