package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func rdsCluster(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider, createdSecurityGroup *ec2.SecurityGroup) (*rds.Cluster, error) {
	clusterArgs := &rds.ClusterArgs{
		ClusterIdentifier:                pulumi.String(locals.AwsRdsCluster.Metadata.Id),
		DatabaseName:                     pulumi.String(locals.AwsRdsCluster.Spec.DatabaseName),
		PreferredMaintenanceWindow:       pulumi.String(locals.AwsRdsCluster.Spec.MaintenanceWindow),
		NetworkType:                      pulumi.String("IPV4"),
		IamDatabaseAuthenticationEnabled: pulumi.Bool(locals.AwsRdsCluster.Spec.IamDatabaseAuthenticationEnabled),
		Tags:                             pulumi.ToStringMap(locals.Labels),
		Engine:                           pulumi.String(locals.AwsRdsCluster.Spec.Engine),
		BackupRetentionPeriod:            pulumi.Int(5),
		EngineVersion:                    pulumi.String(locals.AwsRdsCluster.Spec.EngineVersion),
		AllowMajorVersionUpgrade:         pulumi.Bool(locals.AwsRdsCluster.Spec.AllowMajorVersionUpgrade),
		EngineMode:                       pulumi.String(locals.AwsRdsCluster.Spec.EngineMode),
		Port:                             pulumi.Int(locals.AwsRdsCluster.Spec.DatabasePort),
		PreferredBackupWindow:            pulumi.String(locals.AwsRdsCluster.Spec.BackupWindow),
		CopyTagsToSnapshot:               pulumi.Bool(false),
		ApplyImmediately:                 pulumi.Bool(true),
		EnabledCloudwatchLogsExports:     pulumi.ToStringArray(locals.AwsRdsCluster.Spec.EnabledCloudwatchLogsExports),
		DeletionProtection:               pulumi.Bool(locals.AwsRdsCluster.Spec.DeletionProtection),
		MasterUsername:                   pulumi.String(locals.AwsRdsCluster.Spec.MasterUser),
	}

	if locals.AwsRdsCluster.Spec.DatabasePort > 0 {
		clusterArgs.Port = pulumi.Int(locals.AwsRdsCluster.Spec.DatabasePort)
	}

	if locals.AwsRdsCluster.Spec.RetentionPeriod > 0 {
		clusterArgs.BackupRetentionPeriod = pulumi.Int(locals.AwsRdsCluster.Spec.RetentionPeriod)
	}

	if locals.AwsRdsCluster.Spec.ManageMasterUserPassword {
		clusterArgs.ManageMasterUserPassword = pulumi.Bool(true)
		if locals.AwsRdsCluster.Spec.MasterUserSecretKmsKeyId != "" {
			clusterArgs.MasterUserSecretKmsKeyId = pulumi.String(locals.AwsRdsCluster.Spec.MasterUserSecretKmsKeyId)
		}
	} else {
		clusterArgs.MasterPassword = pulumi.String(locals.AwsRdsCluster.Spec.MasterPassword)
	}

	clusterArgs.SkipFinalSnapshot = pulumi.Bool(locals.AwsRdsCluster.Spec.SkipFinalSnapshot)
	if !locals.AwsRdsCluster.Spec.SkipFinalSnapshot {
		clusterArgs.FinalSnapshotIdentifier = pulumi.Sprintf("%s-final-snapshot", locals.AwsRdsCluster.Metadata.Id)
	}

	if locals.AwsRdsCluster.Spec.EngineMode != "serverless" {
		clusterArgs.StorageEncrypted = pulumi.Bool(locals.AwsRdsCluster.Spec.StorageEncrypted)
		if locals.AwsRdsCluster.Spec.StorageEncrypted {
			clusterArgs.KmsKeyId = pulumi.String(locals.AwsRdsCluster.Spec.StorageKmsKeyArn)
		}
	}

	if locals.AwsRdsCluster.Spec.ScalingConfiguration != nil {
		maxCapacity := locals.AwsRdsCluster.Spec.ScalingConfiguration.MaxCapacity
		if maxCapacity == 0 {
			maxCapacity = 16
		}

		minCapacity := locals.AwsRdsCluster.Spec.ScalingConfiguration.MinCapacity
		if minCapacity == 0 {
			minCapacity = 2
		}

		secondsUntilAutoPause := locals.AwsRdsCluster.Spec.ScalingConfiguration.SecondsUntilAutoPause
		if secondsUntilAutoPause == 0 {
			secondsUntilAutoPause = 300
		}

		timeoutAction := locals.AwsRdsCluster.Spec.ScalingConfiguration.TimeoutAction
		if timeoutAction == "" {
			timeoutAction = "RollbackCapacityChange"
		}

		clusterArgs.ScalingConfiguration = &rds.ClusterScalingConfigurationArgs{
			AutoPause:             pulumi.Bool(locals.AwsRdsCluster.Spec.ScalingConfiguration.AutoPause),
			MaxCapacity:           pulumi.Int(maxCapacity),
			MinCapacity:           pulumi.Int(minCapacity),
			SecondsUntilAutoPause: pulumi.Int(secondsUntilAutoPause),
			TimeoutAction:         pulumi.String(timeoutAction),
		}
	}

	if locals.AwsRdsCluster.Spec.Serverlessv2ScalingConfiguration != nil {
		clusterArgs.Serverlessv2ScalingConfiguration = &rds.ClusterServerlessv2ScalingConfigurationArgs{
			MaxCapacity: pulumi.Float64(locals.AwsRdsCluster.Spec.Serverlessv2ScalingConfiguration.MaxCapacity),
			MinCapacity: pulumi.Float64(locals.AwsRdsCluster.Spec.Serverlessv2ScalingConfiguration.MinCapacity),
		}
	}

	vpcSecurityGroupIds := pulumi.ToStringArray(locals.AwsRdsCluster.Spec.AssociateSecurityGroupIds)
	vpcSecurityGroupIds = append(vpcSecurityGroupIds, createdSecurityGroup.ID())

	clusterArgs.VpcSecurityGroupIds = vpcSecurityGroupIds

	if len(locals.AwsRdsCluster.Spec.SubnetIds) > 0 && locals.AwsRdsCluster.Spec.DbSubnetGroupName == "" {
		createdSubnetGroup, err := subnetGroup(ctx, locals, awsProvider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create subnet group")
		}
		clusterArgs.DbSubnetGroupName = createdSubnetGroup.Name
	}
	if locals.AwsRdsCluster.Spec.DbSubnetGroupName != "" {
		clusterArgs.DbSubnetGroupName = pulumi.String(locals.AwsRdsCluster.Spec.DbSubnetGroupName)
	}

	if locals.AwsRdsCluster.Spec.ClusterParameterGroupName == "" {
		createdClusterParameterGroup, err := clusterParameterGroup(ctx, locals, awsProvider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create cluster parameter group")
		}
		clusterArgs.DbClusterParameterGroupName = createdClusterParameterGroup.Name
	} else {
		clusterArgs.DbClusterParameterGroupName = pulumi.String(locals.AwsRdsCluster.Spec.ClusterParameterGroupName)
	}

	clusterArgs.SnapshotIdentifier = pulumi.String(locals.AwsRdsCluster.Spec.SnapshotIdentifier)

	// Create RDS Cluster
	createdRdsCluster, err := rds.NewCluster(ctx, locals.AwsRdsCluster.Metadata.Id, clusterArgs, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create regional rds cluster")
	}

	return createdRdsCluster, nil
}
