package module

import (
	"fmt"

	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/postgreskubernetes/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// buildBackupEnvVars creates environment variable overrides for per-database backup configuration.
// These environment variables override the operator-level backup settings for this specific database.
// If backupConfig is nil, returns nil (database inherits operator-level settings).
func buildBackupEnvVars(backupConfig *postgreskubernetesv1.PostgresKubernetesBackupConfig, databaseName string) pulumi.MapArrayInput {
	if backupConfig == nil {
		return nil
	}

	var envVars []pulumi.MapInput

	// Override S3 prefix if specified
	if backupConfig.S3Prefix != "" {
		envVars = append(envVars, pulumi.Map{
			"name":  pulumi.String("WALG_S3_PREFIX"),
			"value": pulumi.String(fmt.Sprintf("s3://%s", backupConfig.S3Prefix)),
		})
	}

	// Override backup schedule if specified
	if backupConfig.BackupSchedule != "" {
		envVars = append(envVars, pulumi.Map{
			"name":  pulumi.String("BACKUP_SCHEDULE"),
			"value": pulumi.String(backupConfig.BackupSchedule),
		})
	}

	// Override enable_backup if specified
	if backupConfig.EnableBackup != nil {
		envVars = append(envVars, pulumi.Map{
			"name":  pulumi.String("USE_WALG_BACKUP"),
			"value": pulumi.String(boolToString(*backupConfig.EnableBackup)),
		})
	}

	// Override enable_restore if specified
	if backupConfig.EnableRestore != nil {
		envVars = append(envVars, pulumi.Map{
			"name":  pulumi.String("USE_WALG_RESTORE"),
			"value": pulumi.String(boolToString(*backupConfig.EnableRestore)),
		})
	}

	// Override enable_clone if specified
	if backupConfig.EnableClone != nil {
		envVars = append(envVars, pulumi.Map{
			"name":  pulumi.String("CLONE_USE_WALG_RESTORE"),
			"value": pulumi.String(boolToString(*backupConfig.EnableClone)),
		})
	}

	// If no overrides specified, return nil (inherit operator settings)
	if len(envVars) == 0 {
		return nil
	}

	return pulumi.MapArray(envVars)
}

// boolToString converts a bool to "true" or "false" string
func boolToString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
