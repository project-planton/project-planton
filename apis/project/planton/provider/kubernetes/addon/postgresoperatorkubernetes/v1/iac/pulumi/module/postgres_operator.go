package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// postgresOperator deploys the Zalando Postgres‑Operator via Helm.
func postgresOperator(ctx *pulumi.Context, locals *Locals, kubernetesProvider *pulumikubernetes.Provider) error {
	// 1. Namespace
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:   pulumi.String(vars.Namespace),
				Labels: pulumi.ToStringMap(locals.KubernetesLabels),
			}),
		},
		pulumi.Provider(kubernetesProvider),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// 2. Create backup Secret and ConfigMap if backup_config is specified
	backupConfigMapName, err := createBackupResources(
		ctx,
		locals.PostgresOperatorKubernetes.Spec.BackupConfig,
		createdNamespace.Metadata.Name().Elem(),
		kubernetesProvider,
		locals.KubernetesLabels,
	)
	if err != nil {
		return errors.Wrap(err, "failed to create backup resources")
	}

	// 3. Build Helm values with backup ConfigMap if configured
	helmValues := backupConfigMapName.ApplyT(func(cmName string) pulumi.Map {
		baseValues := pulumi.Map{
			"configKubernetes": pulumi.Map{
				"inherited_labels": pulumi.ToStringArray([]string{
					kuberneteslabelkeys.Resource,
					kuberneteslabelkeys.Organization,
					kuberneteslabelkeys.Environment,
					kuberneteslabelkeys.ResourceKind,
					kuberneteslabelkeys.ResourceId,
				}),
			},
		}

		// Add pod_environment_configmap if backup is configured
		if cmName != "" {
			baseValues["configKubernetes"].(pulumi.Map)["pod_environment_configmap"] = pulumi.String(cmName)
		}

		return baseValues
	}).(pulumi.MapOutput)

	// 4. Helm release
	_, err = helm.NewRelease(ctx,
		"postgres-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(vars.HelmChartVersion),
			RepositoryOpts:  helm.RepositoryOptsArgs{Repo: pulumi.String(vars.HelmChartRepo)},
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values:          helmValues,
		},
		pulumi.Parent(createdNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}

	// 5. Export stack‑output(s)
	ctx.Export(OpNamespace, createdNamespace.Metadata.Name())

	return nil
}
