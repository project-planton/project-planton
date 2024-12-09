package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func rdsClusterInstance(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider, createdRdsCluster *rds.Cluster) ([]*rds.ClusterInstance, error) {
	clusterInstanceArgs := &rds.ClusterInstanceArgs{
		ClusterIdentifier:          createdRdsCluster.ID(),
		InstanceClass:              pulumi.String(locals.AwsRdsCluster.Spec.InstanceType),
		DbSubnetGroupName:          createdRdsCluster.DbSubnetGroupName,
		PubliclyAccessible:         pulumi.Bool(locals.AwsRdsCluster.Spec.IsPubliclyAccessible),
		Tags:                       pulumi.ToStringMap(locals.Labels),
		Engine:                     createdRdsCluster.Engine,
		EngineVersion:              createdRdsCluster.EngineVersion,
		AutoMinorVersionUpgrade:    pulumi.Bool(true),
		ApplyImmediately:           pulumi.Bool(true),
		PreferredMaintenanceWindow: pulumi.String(locals.AwsRdsCluster.Spec.MaintenanceWindow),
		PreferredBackupWindow:      pulumi.String(locals.AwsRdsCluster.Spec.BackupWindow),
		CopyTagsToSnapshot:         pulumi.Bool(false),
		CaCertIdentifier:           pulumi.String(locals.AwsRdsCluster.Spec.CaCertIdentifier),
	}

	if locals.AwsRdsCluster.Spec.Serverlessv2ScalingConfiguration != nil {
		clusterInstanceArgs.InstanceClass = pulumi.String("db.serverless")
	}

	if locals.AwsRdsCluster.Spec.EnhancedMonitoringRoleEnabled {
		enhancedMonitoringIamRole, err := enhancedMonitoring(ctx, locals, awsProvider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create enhanced monitoring iam role")
		}
		clusterInstanceArgs.MonitoringRoleArn = enhancedMonitoringIamRole.Arn
		clusterInstanceArgs.MonitoringInterval = pulumi.Int(locals.AwsRdsCluster.Spec.RdsMonitoringInterval)
	}

	clusterInstanceArgs.PerformanceInsightsEnabled = pulumi.Bool(locals.AwsRdsCluster.Spec.IsPerformanceInsightsEnabled)
	if locals.AwsRdsCluster.Spec.IsPerformanceInsightsEnabled {
		clusterInstanceArgs.PerformanceInsightsKmsKeyId = pulumi.String(locals.AwsRdsCluster.Spec.PerformanceInsightsKmsKeyId)
	}

	var rdsClusterInstances []*rds.ClusterInstance
	for i := 0; i < int(locals.AwsRdsCluster.Spec.ClusterSize); i++ {
		clusterInstanceIdentifier := fmt.Sprintf("%s-%d", locals.AwsRdsCluster.Metadata.Id, i+1)
		clusterInstanceArgsCopy := *clusterInstanceArgs
		clusterInstanceArgsCopy.Identifier = pulumi.String(clusterInstanceIdentifier)
		// Create RDS Cluster
		createdRdsClusterInstance, err := rds.NewClusterInstance(ctx, clusterInstanceIdentifier,
			&clusterInstanceArgsCopy,
			pulumi.Provider(awsProvider), pulumi.Parent(createdRdsCluster), pulumi.IgnoreChanges([]string{
				"engine_version",
			}))
		if err != nil {
			return nil, errors.Wrap(err, "failed to create rds cluster instance")
		}
		rdsClusterInstances = append(rdsClusterInstances, createdRdsClusterInstance)
	}
	return rdsClusterInstances, nil
}
