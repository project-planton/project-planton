package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func appAutoscaling(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider, createdRdsCluster *rds.Cluster) error {
	var isAutoScalingEnabled = false
	if locals.AwsRdsCluster.Spec.AutoScaling != nil {
		isAutoScalingEnabled = locals.AwsRdsCluster.Spec.AutoScaling.IsEnabled
	}

	if !isAutoScalingEnabled {
		return nil
	}

	autoScalingTarget := &appautoscaling.TargetArgs{
		ResourceId:        pulumi.Sprintf("cluster:%s", locals.AwsRdsCluster.Metadata.Id),
		ScalableDimension: pulumi.String("rds:cluster:ReadReplicaCount"),
		ServiceNamespace:  pulumi.String("rds"),
		Tags:              pulumi.ToStringMap(locals.Labels),
	}

	minCapacity := 1
	if locals.AwsRdsCluster.Spec.AutoScaling.MinCapacity > 0 {
		minCapacity = int(locals.AwsRdsCluster.Spec.AutoScaling.MinCapacity)
	}
	autoScalingTarget.MinCapacity = pulumi.Int(minCapacity)

	maxCapacity := 5
	if locals.AwsRdsCluster.Spec.AutoScaling.MaxCapacity > 0 {
		maxCapacity = int(locals.AwsRdsCluster.Spec.AutoScaling.MaxCapacity)
	}
	autoScalingTarget.MaxCapacity = pulumi.Int(maxCapacity)

	replicasTarget, err := appautoscaling.NewTarget(ctx, "replicas", autoScalingTarget, pulumi.Provider(awsProvider),
		pulumi.Parent(createdRdsCluster),
		pulumi.DependsOn([]pulumi.Resource{createdRdsCluster}))
	if err != nil {
		return errors.Wrap(err, "failed to create replicas auto scaling resources")
	}

	autoScalingTargetPolicy := &appautoscaling.PolicyArgs{
		Name:              pulumi.String(locals.AwsRdsCluster.Metadata.Id),
		ResourceId:        replicasTarget.ResourceId,
		ScalableDimension: replicasTarget.ScalableDimension,
		ServiceNamespace:  replicasTarget.ServiceNamespace,
	}

	policyType := "TargetTrackingScaling"
	if locals.AwsRdsCluster.Spec.AutoScaling.PolicyType != "" {
		policyType = locals.AwsRdsCluster.Spec.AutoScaling.PolicyType
	}
	autoScalingTargetPolicy.PolicyType = pulumi.String(policyType)

	targetMetrics := "RDSReaderAverageCPUUtilization"
	if locals.AwsRdsCluster.Spec.AutoScaling.TargetMetrics != "" {
		targetMetrics = locals.AwsRdsCluster.Spec.AutoScaling.TargetMetrics
	}

	targetValue := 75.0
	if locals.AwsRdsCluster.Spec.AutoScaling.TargetValue > 0 {
		targetValue = locals.AwsRdsCluster.Spec.AutoScaling.TargetValue
	}

	scaleInCoolDown := 300
	if locals.AwsRdsCluster.Spec.AutoScaling.ScaleInCooldown > 0 {
		scaleInCoolDown = int(locals.AwsRdsCluster.Spec.AutoScaling.ScaleInCooldown)
	}

	scaleOutCoolDown := 300
	if locals.AwsRdsCluster.Spec.AutoScaling.ScaleOutCooldown > 0 {
		scaleOutCoolDown = int(locals.AwsRdsCluster.Spec.AutoScaling.ScaleOutCooldown)
	}

	targetTrackingScalingPolicyConfiguration := &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
		PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
			PredefinedMetricType: pulumi.String(targetMetrics),
		},
		DisableScaleIn:   pulumi.Bool(false),
		TargetValue:      pulumi.Float64(targetValue),
		ScaleInCooldown:  pulumi.Int(scaleInCoolDown),
		ScaleOutCooldown: pulumi.Int(scaleOutCoolDown),
	}

	autoScalingTargetPolicy.TargetTrackingScalingPolicyConfiguration = targetTrackingScalingPolicyConfiguration

	_, err = appautoscaling.NewPolicy(ctx, "replicas-policy", autoScalingTargetPolicy, pulumi.Provider(awsProvider),
		pulumi.Parent(replicasTarget),
		pulumi.DependsOn([]pulumi.Resource{createdRdsCluster, replicasTarget}))
	if err != nil {
		return errors.Wrap(err, "failed to create replicas policy")
	}
	return nil
}
