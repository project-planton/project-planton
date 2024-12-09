package module

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"time"
)

func rdsInstance(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider, createdSecurityGroup *ec2.SecurityGroup) (*rds.Instance, error) {
	rdsInstanceArgs := &rds.InstanceArgs{
		Identifier:                       pulumi.String(locals.AwsRdsInstance.Metadata.Id),
		DbName:                           pulumi.String(locals.AwsRdsInstance.Spec.DbName),
		Port:                             pulumi.Int(locals.AwsRdsInstance.Spec.Port),
		CharacterSetName:                 pulumi.String(locals.AwsRdsInstance.Spec.CharacterSetName),
		InstanceClass:                    pulumi.String(locals.AwsRdsInstance.Spec.InstanceClass),
		MaxAllocatedStorage:              pulumi.Int(locals.AwsRdsInstance.Spec.MaxAllocatedStorage),
		StorageEncrypted:                 pulumi.Bool(locals.AwsRdsInstance.Spec.StorageEncrypted),
		KmsKeyId:                         pulumi.String(locals.AwsRdsInstance.Spec.KmsKeyId),
		MultiAz:                          pulumi.Bool(locals.AwsRdsInstance.Spec.IsMultiAz),
		CaCertIdentifier:                 pulumi.String(locals.AwsRdsInstance.Spec.CaCertIdentifier),
		LicenseModel:                     pulumi.String(locals.AwsRdsInstance.Spec.LicenseModel),
		StorageType:                      pulumi.String(locals.AwsRdsInstance.Spec.StorageType),
		Iops:                             pulumi.Int(locals.AwsRdsInstance.Spec.Iops),
		PubliclyAccessible:               pulumi.Bool(locals.AwsRdsInstance.Spec.IsPubliclyAccessible),
		SnapshotIdentifier:               pulumi.String(locals.AwsRdsInstance.Spec.SnapshotIdentifier),
		AllowMajorVersionUpgrade:         pulumi.Bool(locals.AwsRdsInstance.Spec.AllowMajorVersionUpgrade),
		AutoMinorVersionUpgrade:          pulumi.Bool(locals.AwsRdsInstance.Spec.AutoMinorVersionUpgrade),
		ApplyImmediately:                 pulumi.Bool(locals.AwsRdsInstance.Spec.ApplyImmediately),
		MaintenanceWindow:                pulumi.String(locals.AwsRdsInstance.Spec.MaintenanceWindow),
		CopyTagsToSnapshot:               pulumi.Bool(locals.AwsRdsInstance.Spec.CopyTagsToSnapshot),
		BackupRetentionPeriod:            pulumi.Int(locals.AwsRdsInstance.Spec.BackupRetentionPeriod),
		BackupWindow:                     pulumi.String(locals.AwsRdsInstance.Spec.BackupWindow),
		DeletionProtection:               pulumi.Bool(locals.AwsRdsInstance.Spec.DeletionProtection),
		SkipFinalSnapshot:                pulumi.Bool(locals.AwsRdsInstance.Spec.SkipFinalSnapshot),
		Timezone:                         pulumi.String(locals.AwsRdsInstance.Spec.Timezone),
		IamDatabaseAuthenticationEnabled: pulumi.Bool(locals.AwsRdsInstance.Spec.IamDatabaseAuthenticationEnabled),
		EnabledCloudwatchLogsExports:     pulumi.ToStringArray(locals.AwsRdsInstance.Spec.EnabledCloudwatchLogsExports),
		Username:                         pulumi.String(locals.AwsRdsInstance.Spec.Username),
		Tags:                             pulumi.ToStringMap(locals.Labels),
	}

	if len(locals.AwsRdsInstance.Spec.SubnetIds) > 0 && locals.AwsRdsInstance.Spec.DbSubnetGroupName == "" {
		createdSubnetGroup, err := subnetGroup(ctx, locals, awsProvider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create subnet group")
		}
		rdsInstanceArgs.DbSubnetGroupName = createdSubnetGroup.Name
	}
	if locals.AwsRdsInstance.Spec.DbSubnetGroupName != "" {
		rdsInstanceArgs.DbSubnetGroupName = pulumi.String(locals.AwsRdsInstance.Spec.DbSubnetGroupName)
	}

	if locals.AwsRdsInstance.Spec.ParameterGroupName == "" {
		createdParameterGroup, err := parameterGroup(ctx, locals, awsProvider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create parameter group")
		}
		rdsInstanceArgs.ParameterGroupName = createdParameterGroup.Name
	} else {
		rdsInstanceArgs.ParameterGroupName = pulumi.String(locals.AwsRdsInstance.Spec.ParameterGroupName)
	}

	if locals.AwsRdsInstance.Spec.OptionGroupName == "" {
		createdOptionGroup, err := optionGroup(ctx, locals, awsProvider)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create option group")
		}
		rdsInstanceArgs.OptionGroupName = createdOptionGroup.Name
	} else {
		rdsInstanceArgs.OptionGroupName = pulumi.String(locals.AwsRdsInstance.Spec.OptionGroupName)
	}

	manageMasterUserPassword := locals.AwsRdsInstance.Spec.ManageMasterUserPassword
	if locals.AwsRdsInstance.Spec.ReplicateSourceDb == "" {
		rdsInstanceArgs.Engine = pulumi.String(locals.AwsRdsInstance.Spec.Engine)
		rdsInstanceArgs.EngineVersion = pulumi.String(locals.AwsRdsInstance.Spec.EngineVersion)
		rdsInstanceArgs.AllocatedStorage = pulumi.Int(locals.AwsRdsInstance.Spec.AllocatedStorage)
		if manageMasterUserPassword {
			rdsInstanceArgs.ManageMasterUserPassword = pulumi.Bool(manageMasterUserPassword)
			if locals.AwsRdsInstance.Spec.MasterUserSecretKmsKeyId != "" {
				rdsInstanceArgs.MasterUserSecretKmsKeyId = pulumi.String(locals.AwsRdsInstance.Spec.MasterUserSecretKmsKeyId)
			}
		} else {
			rdsInstanceArgs.Password = pulumi.String(locals.AwsRdsInstance.Spec.Password)
		}
	} else {
		rdsInstanceArgs.ReplicateSourceDb = pulumi.String(locals.AwsRdsInstance.Spec.ReplicateSourceDb)
	}

	if !locals.AwsRdsInstance.Spec.IsMultiAz {
		rdsInstanceArgs.AvailabilityZone = pulumi.String(locals.AwsRdsInstance.Spec.AvailabilityZone)
	}

	if locals.AwsRdsInstance.Spec.StorageType == "gp3" {
		rdsInstanceArgs.StorageThroughput = pulumi.Int(locals.AwsRdsInstance.Spec.StorageThroughput)
	}

	if !locals.AwsRdsInstance.Spec.SkipFinalSnapshot {
		entropy := ulid.Monotonic(rand.Reader, 0)
		ulidValue := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
		rdsInstanceArgs.FinalSnapshotIdentifier = pulumi.Sprintf("%s-%s", locals.AwsRdsInstance.Metadata.Id, ulidValue)
	}

	performanceInsightsEnabled := false
	performanceInsightsKmsKeyId := ""
	performanceInsightsRetentionPeriod := 7
	if locals.AwsRdsInstance.Spec.PerformanceInsights != nil {
		performanceInsightsEnabled = locals.AwsRdsInstance.Spec.PerformanceInsights.IsEnabled
		performanceInsightsKmsKeyId = locals.AwsRdsInstance.Spec.PerformanceInsights.KmsKeyId
		performanceInsightsRetentionPeriod = int(locals.AwsRdsInstance.Spec.PerformanceInsights.RetentionPeriod)
	}

	if performanceInsightsEnabled {
		rdsInstanceArgs.PerformanceInsightsEnabled = pulumi.Bool(performanceInsightsEnabled)
		rdsInstanceArgs.PerformanceInsightsKmsKeyId = pulumi.String(performanceInsightsKmsKeyId)
		rdsInstanceArgs.PerformanceInsightsRetentionPeriod = pulumi.Int(performanceInsightsRetentionPeriod)
	}

	if locals.AwsRdsInstance.Spec.Monitoring != nil {
		rdsInstanceArgs.MonitoringInterval = pulumi.Int(locals.AwsRdsInstance.Spec.Monitoring.MonitoringInterval)
		rdsInstanceArgs.MonitoringRoleArn = pulumi.String(locals.AwsRdsInstance.Spec.Monitoring.MonitoringRoleArn)
	}

	if locals.AwsRdsInstance.Spec.SnapshotIdentifier == "" {
		restoreInTime := &rds.InstanceRestoreToPointInTimeArgs{}
		if locals.AwsRdsInstance.Spec.RestoreToPointInTime != nil {
			restoreInTime = &rds.InstanceRestoreToPointInTimeArgs{
				RestoreTime:                         pulumi.String(locals.AwsRdsInstance.Spec.RestoreToPointInTime.RestoreTime),
				SourceDbInstanceAutomatedBackupsArn: pulumi.String(locals.AwsRdsInstance.Spec.RestoreToPointInTime.SourceDbInstanceAutomatedBackupsArn),
				SourceDbInstanceIdentifier:          pulumi.String(locals.AwsRdsInstance.Spec.RestoreToPointInTime.SourceDbInstanceIdentifier),
				SourceDbiResourceId:                 pulumi.String(locals.AwsRdsInstance.Spec.RestoreToPointInTime.SourceDbiResourceId),
				UseLatestRestorableTime:             pulumi.Bool(locals.AwsRdsInstance.Spec.RestoreToPointInTime.UseLatestRestorableTime),
			}
			rdsInstanceArgs.RestoreToPointInTime = restoreInTime
		}
	}

	vpcSecurityGroupIds := pulumi.ToStringArray(locals.AwsRdsInstance.Spec.AssociateSecurityGroupIds)
	vpcSecurityGroupIds = append(vpcSecurityGroupIds, createdSecurityGroup.ID())

	rdsInstanceArgs.VpcSecurityGroupIds = vpcSecurityGroupIds

	// Create RDS Instance
	rdsInstance, err := rds.NewInstance(ctx, "rdsInstance", rdsInstanceArgs,
		pulumi.Provider(awsProvider), pulumi.DependsOn([]pulumi.Resource{createdSecurityGroup}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rds instance")
	}

	return rdsInstance, nil
}
