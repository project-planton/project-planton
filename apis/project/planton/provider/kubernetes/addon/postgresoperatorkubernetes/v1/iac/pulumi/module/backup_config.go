package module

import (
	"fmt"

	"github.com/pkg/errors"
	postgresoperatorkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/postgresoperatorkubernetes/v1"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	backupSecretName    = "r2-postgres-backup-credentials"
	backupConfigMapName = "postgres-pod-backup-config"
)

// createBackupResources creates the Secret and ConfigMap for PostgreSQL backups using Cloudflare R2.
// Returns the ConfigMap name if backup is configured, empty string otherwise.
func createBackupResources(
	ctx *pulumi.Context,
	backupConfig *postgresoperatorkubernetesv1.PostgresOperatorKubernetesBackupConfig,
	namespace pulumi.StringInput,
	kubernetesProvider *pulumikubernetes.Provider,
	labels map[string]string,
) (pulumi.StringOutput, error) {
	// If no backup config provided, return empty
	if backupConfig == nil {
		return pulumi.String("").ToStringOutput(), nil
	}

	r2Config := backupConfig.R2Config
	if r2Config == nil {
		return pulumi.String("").ToStringOutput(), errors.New("backup_config.r2_config is required when backup_config is specified")
	}

	// 1. Create Secret for R2 credentials
	createdSecret, err := corev1.NewSecret(ctx,
		backupSecretName,
		&corev1.SecretArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:      pulumi.String(backupSecretName),
				Namespace: namespace,
				Labels:    pulumi.ToStringMap(labels),
			}),
			Type: pulumi.String("Opaque"),
			StringData: pulumi.StringMap{
				"AWS_ACCESS_KEY_ID":     pulumi.String(r2Config.AccessKeyId),
				"AWS_SECRET_ACCESS_KEY": pulumi.String(r2Config.SecretAccessKey),
			},
		},
		pulumi.Provider(kubernetesProvider),
	)
	if err != nil {
		return pulumi.String("").ToStringOutput(), errors.Wrap(err, "failed to create backup credentials secret")
	}

	// 2. Build R2 endpoint URL
	r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2Config.CloudflareAccountId)

	// 3. Build S3 prefix (default or custom)
	s3Prefix := "backups/$(SCOPE)/$(PGVERSION)"
	if backupConfig.S3PrefixTemplate != "" {
		s3Prefix = backupConfig.S3PrefixTemplate
	}

	// Full WALG_S3_PREFIX with bucket and path
	walgS3Prefix := fmt.Sprintf("s3://%s/%s", r2Config.BucketName, s3Prefix)

	// 4. Build ConfigMap data
	configMapData := pulumi.StringMap{
		// WAL-G flags (default to true if not explicitly disabled)
		"USE_WALG_BACKUP":        pulumi.String(boolToString(backupConfig.EnableWalGBackup, true)),
		"USE_WALG_RESTORE":       pulumi.String(boolToString(backupConfig.EnableWalGRestore, true)),
		"CLONE_USE_WALG_RESTORE": pulumi.String(boolToString(backupConfig.EnableCloneWalGRestore, true)),

		// S3/R2 configuration
		"WALG_S3_PREFIX":       pulumi.String(walgS3Prefix),
		"AWS_ENDPOINT":         pulumi.String(r2Endpoint),
		"AWS_REGION":           pulumi.String("auto"), // R2 uses "auto" region
		"AWS_FORCE_PATH_STYLE": pulumi.String("true"), // Required for R2

		// Backup schedule
		"BACKUP_SCHEDULE": pulumi.String(backupConfig.BackupSchedule),

		// Credentials (reference the Secret - Zalando operator will mount them)
		// We don't include credentials in ConfigMap; they come from the Secret
		// The Secret will be mounted at /run/etc/wal-e.d/env by Zalando operator
		"AWS_ACCESS_KEY_ID":     pulumi.String(r2Config.AccessKeyId),
		"AWS_SECRET_ACCESS_KEY": pulumi.String(r2Config.SecretAccessKey),
	}

	// 5. Create ConfigMap
	createdConfigMap, err := corev1.NewConfigMap(ctx,
		backupConfigMapName,
		&corev1.ConfigMapArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:      pulumi.String(backupConfigMapName),
				Namespace: namespace,
				Labels:    pulumi.ToStringMap(labels),
			}),
			Data: configMapData,
		},
		pulumi.Provider(kubernetesProvider),
		pulumi.DependsOn([]pulumi.Resource{createdSecret}),
	)
	if err != nil {
		return pulumi.String("").ToStringOutput(), errors.Wrap(err, "failed to create backup config map")
	}

	// Return the ConfigMap name (namespace/name format for Zalando operator)
	return pulumi.Sprintf("%s/%s", namespace, createdConfigMap.Metadata.Name().Elem()), nil
}

// boolToString converts a bool to "true"/"false" string, with a default value when the bool is false.
func boolToString(value bool, defaultWhenFalse bool) string {
	if value {
		return "true"
	}
	if defaultWhenFalse {
		return "true"
	}
	return "false"
}
