package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1/iac/pulumi/module/outputs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func autoScale(ctx *pulumi.Context, locals *Locals,
	awsProvider *aws.Provider, createdDynamodbTable *dynamodb.Table) error {
	awsDynamodb := locals.AwsDynamodb
	enableAutoScale := false
	if awsDynamodb.Spec.AutoScale != nil {
		enableAutoScale = awsDynamodb.Spec.AutoScale.IsEnabled
	}

	if !enableAutoScale || awsDynamodb.Spec.BillingMode != "PROVISIONED" {
		return nil
	}

	//read capacity
	autoScaleMinReadCapacity := 5
	autoScaleMaxReadCapacity := 20
	autoScaleReadTarget := 50.0

	if awsDynamodb.Spec.AutoScale != nil &&
		awsDynamodb.Spec.AutoScale.ReadCapacity != nil {
		autoScaleMinReadCapacity = int(awsDynamodb.Spec.AutoScale.ReadCapacity.MinCapacity)
		autoScaleMaxReadCapacity = int(awsDynamodb.Spec.AutoScale.ReadCapacity.MaxCapacity)
		autoScaleReadTarget = awsDynamodb.Spec.AutoScale.ReadCapacity.TargetUtilization
	}

	//write capacity
	autoScaleMinWriteCapacity := 5
	autoScaleMaxWriteCapacity := 20
	autoScaleWriteTarget := 50.0

	if awsDynamodb.Spec.AutoScale != nil &&
		awsDynamodb.Spec.AutoScale.WriteCapacity != nil {
		autoScaleMinWriteCapacity = int(awsDynamodb.Spec.AutoScale.WriteCapacity.MinCapacity)
		autoScaleMaxWriteCapacity = int(awsDynamodb.Spec.AutoScale.WriteCapacity.MaxCapacity)
		autoScaleWriteTarget = awsDynamodb.Spec.AutoScale.WriteCapacity.TargetUtilization
	}

	//index read capacity
	autoScaleMinIndexReadCapacity := 5
	autoScaleMaxIndexReadCapacity := 20
	autoScaleIndexReadTarget := 50.0

	if awsDynamodb.Spec.AutoScale != nil &&
		awsDynamodb.Spec.AutoScale.ReadCapacity != nil {
		autoScaleMinIndexReadCapacity = int(awsDynamodb.Spec.AutoScale.ReadCapacity.MinCapacity)
		autoScaleMaxIndexReadCapacity = int(awsDynamodb.Spec.AutoScale.ReadCapacity.MaxCapacity)
		autoScaleIndexReadTarget = awsDynamodb.Spec.AutoScale.ReadCapacity.TargetUtilization
	}

	//index write capacity
	autoScaleMinIndexWriteCapacity := 5
	autoScaleMaxIndexWriteCapacity := 20
	autoScaleIndexWriteTarget := 50.0

	if awsDynamodb.Spec.AutoScale != nil &&
		awsDynamodb.Spec.AutoScale.WriteCapacity != nil {
		autoScaleMinIndexWriteCapacity = int(awsDynamodb.Spec.AutoScale.WriteCapacity.MinCapacity)
		autoScaleMaxIndexWriteCapacity = int(awsDynamodb.Spec.AutoScale.WriteCapacity.MaxCapacity)
		autoScaleIndexWriteTarget = awsDynamodb.Spec.AutoScale.WriteCapacity.TargetUtilization
	}

	readTarget, err := appautoscaling.NewTarget(ctx, "readTarget", &appautoscaling.TargetArgs{
		MaxCapacity:       pulumi.Int(autoScaleMaxReadCapacity),
		MinCapacity:       pulumi.Int(autoScaleMinReadCapacity),
		ResourceId:        pulumi.String("table/" + awsDynamodb.Metadata.Name),
		ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
		ServiceNamespace:  pulumi.String("dynamodb"),
		Tags:              pulumi.ToStringMap(locals.Labels),
	}, pulumi.Provider(awsProvider),
		pulumi.Parent(createdDynamodbTable),
		pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable}))
	if err != nil {
		return errors.Wrap(err, "failed to create read target auto scaling resources")
	}

	readPolicy, err := appautoscaling.NewPolicy(ctx, "readPolicy", &appautoscaling.PolicyArgs{
		Name:              pulumi.Sprintf("DynamoDBReadCapacityUtilization:%s", readTarget.ID().ElementType()),
		PolicyType:        pulumi.String("TargetTrackingScaling"),
		ResourceId:        readTarget.ResourceId,
		ScalableDimension: readTarget.ScalableDimension,
		ServiceNamespace:  readTarget.ServiceNamespace,
		TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
			PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
				PredefinedMetricType: pulumi.String("DynamoDBReadCapacityUtilization"),
			},
			TargetValue: pulumi.Float64(autoScaleReadTarget),
		},
	}, pulumi.Provider(awsProvider),
		pulumi.Parent(readTarget),
		pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable, readTarget}))
	if err != nil {
		return errors.Wrap(err, "failed to create read policy")
	}

	ctx.Export(outputs.AutoscalingReadPolicyArn, readPolicy.Arn)

	indexReadPolicyArnList := pulumi.StringArray{}
	for _, index := range awsDynamodb.Spec.GlobalSecondaryIndexes {
		indexTarget, err := appautoscaling.NewTarget(ctx, fmt.Sprintf("readTargetIndex-%s", index.Name), &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(autoScaleMaxIndexReadCapacity),
			MinCapacity:       pulumi.Int(autoScaleMinIndexReadCapacity),
			ResourceId:        pulumi.String(fmt.Sprintf("table/%s/index/%s", awsDynamodb.Metadata.Name, index.Name)),
			ScalableDimension: pulumi.String("dynamodb:index:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
			Tags:              pulumi.ToStringMap(locals.Labels),
		}, pulumi.Provider(awsProvider),
			pulumi.Parent(readTarget),
			pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable, readTarget}))
		if err != nil {
			return errors.Wrap(err, "failed to create read target index auto scaling resources")
		}

		// Create a Scaling Policy
		indexPolicy, err := appautoscaling.NewPolicy(ctx, fmt.Sprintf("readPolicyIndex-%s", index.Name), &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			ResourceId:        indexTarget.ResourceId,
			ScalableDimension: indexTarget.ScalableDimension,
			ServiceNamespace:  indexTarget.ServiceNamespace,
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("DynamoDBReadCapacityUtilization"),
				},
				TargetValue: pulumi.Float64(autoScaleIndexReadTarget),
			},
		}, pulumi.Provider(awsProvider),
			pulumi.Parent(indexTarget),
			pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable, indexTarget}))
		if err != nil {
			return errors.Wrap(err, "failed to create read policy index auto scaling resources")
		}
		indexReadPolicyArnList = append(indexReadPolicyArnList, indexPolicy.Arn)
	}

	ctx.Export(outputs.AutoscalingIndexReadPolicyArnList, indexReadPolicyArnList)

	writeTarget, err := appautoscaling.NewTarget(ctx, "writeTarget", &appautoscaling.TargetArgs{
		MaxCapacity:       pulumi.Int(autoScaleMaxWriteCapacity),
		MinCapacity:       pulumi.Int(autoScaleMinWriteCapacity),
		ResourceId:        pulumi.String("table/" + awsDynamodb.Metadata.Name),
		ScalableDimension: pulumi.String("dynamodb:table:WriteCapacityUnits"),
		ServiceNamespace:  pulumi.String("dynamodb"),
		Tags:              pulumi.ToStringMap(locals.Labels),
	}, pulumi.Provider(awsProvider),
		pulumi.Parent(createdDynamodbTable),
		pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable}))
	if err != nil {
		return errors.Wrap(err, "failed to create write target auto scaling resources")
	}

	writePolicy, err := appautoscaling.NewPolicy(ctx, "writePolicy", &appautoscaling.PolicyArgs{
		Name:              pulumi.Sprintf("DynamoDBWriteCapacityUtilization:%s", writeTarget.ID().ElementType()),
		PolicyType:        pulumi.String("TargetTrackingScaling"),
		ResourceId:        writeTarget.ResourceId,
		ScalableDimension: writeTarget.ScalableDimension,
		ServiceNamespace:  writeTarget.ServiceNamespace,
		TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
			PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
				PredefinedMetricType: pulumi.String("DynamoDBWriteCapacityUtilization"),
			},
			TargetValue: pulumi.Float64(autoScaleWriteTarget),
		},
	}, pulumi.Provider(awsProvider),
		pulumi.Parent(writeTarget),
		pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable, writeTarget}))
	if err != nil {
		return errors.Wrap(err, "failed to create write policy")
	}

	ctx.Export(outputs.AutoscalingWritePolicyArn, writePolicy.Arn)

	indexWritePolicyArnList := pulumi.StringArray{}
	for _, index := range awsDynamodb.Spec.GlobalSecondaryIndexes {
		indexTarget, err := appautoscaling.NewTarget(ctx, fmt.Sprintf("writeTargetIndex-%s", index.Name), &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(autoScaleMaxIndexWriteCapacity),
			MinCapacity:       pulumi.Int(autoScaleMinIndexWriteCapacity),
			ResourceId:        pulumi.String(fmt.Sprintf("table/%s/index/%s", awsDynamodb.Metadata.Name, index.Name)),
			ScalableDimension: pulumi.String("dynamodb:index:WriteCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
			Tags:              pulumi.ToStringMap(locals.Labels),
		}, pulumi.Provider(awsProvider),
			pulumi.Parent(writeTarget),
			pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable, writeTarget}))
		if err != nil {
			return errors.Wrap(err, "failed to create write target index auto scaling resources")
		}

		// Create a Scaling Policy
		indexPolicy, err := appautoscaling.NewPolicy(ctx, fmt.Sprintf("writePolicyIndex-%s", index.Name), &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			ResourceId:        indexTarget.ResourceId,
			ScalableDimension: indexTarget.ScalableDimension,
			ServiceNamespace:  indexTarget.ServiceNamespace,
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("DynamoDBWriteCapacityUtilization"),
				},
				TargetValue: pulumi.Float64(autoScaleIndexWriteTarget),
			},
		}, pulumi.Provider(awsProvider),
			pulumi.Parent(indexTarget),
			pulumi.DependsOn([]pulumi.Resource{createdDynamodbTable, indexTarget}))
		if err != nil {
			return errors.Wrap(err, "failed to create write policy index auto scaling resources")
		}
		indexWritePolicyArnList = append(indexWritePolicyArnList, indexPolicy.Arn)
	}
	ctx.Export(outputs.AutoscalingIndexWritePolicyArnList, indexWritePolicyArnList)
	return nil
}
