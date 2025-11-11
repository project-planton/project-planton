package module

import (
	"github.com/pkg/errors"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/kubernetes/workload/postgreskubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	zalandov1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/zalandooperator/kubernetes/acid/v1"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *postgreskubernetesv1.PostgresKubernetesStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:   pulumi.String(locals.Namespace),
				Labels: pulumi.ToStringMap(locals.Labels),
			}),
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	// Build restore configuration (standby block + STANDBY_* env vars)
	// Pass operator-level bucket as fallback (if available from stack input)
	// TODO: Extract operator bucket from stackInput if operator config is available
	operatorBucket := ""
	var restoreConfig *postgreskubernetesv1.PostgresKubernetesRestoreConfig
	if locals.PostgresKubernetes.Spec.BackupConfig != nil {
		restoreConfig = locals.PostgresKubernetes.Spec.BackupConfig.Restore
	}
	standbyBlock, standbyEnvVars, err := buildRestoreConfig(restoreConfig, operatorBucket)
	if err != nil {
		return errors.Wrap(err, "failed to build restore configuration")
	}

	// Build backup environment variables (existing function)
	backupEnvVars := buildBackupEnvVars(locals.PostgresKubernetes.Spec.BackupConfig, locals.PostgresKubernetes.Metadata.Name)

	// Merge backup and restore environment variables
	var allEnvVars pulumi.MapArrayInput
	if standbyEnvVars != nil && backupEnvVars != nil {
		// Both sets of env vars exist, merge them
		backupArray, ok := backupEnvVars.(pulumi.MapArray)
		if ok {
			allEnvVars = pulumi.MapArray(append(standbyEnvVars, backupArray...))
		} else {
			allEnvVars = pulumi.MapArray(standbyEnvVars)
		}
	} else if standbyEnvVars != nil {
		// Only standby env vars
		allEnvVars = pulumi.MapArray(standbyEnvVars)
	} else {
		// Only backup env vars (or none)
		allEnvVars = backupEnvVars
	}

	//create zalando postgresql resource
	postgresqlArgs := &zalandov1.PostgresqlArgs{
		Metadata: metav1.ObjectMetaArgs{
			// for zolando operator the name is required to be always prefixed by teamId
			// a kubernetes service with the same name is created by the operator
			Name:      pulumi.Sprintf("%s-%s", vars.TeamId, locals.PostgresKubernetes.Metadata.Name),
			Namespace: createdNamespace.Metadata.Name(),
			Labels:    pulumi.ToStringMap(locals.Labels),
		},
		Spec: zalandov1.PostgresqlSpecArgs{
			NumberOfInstances: pulumi.Int(locals.PostgresKubernetes.Spec.Container.Replicas),
			Patroni:           zalandov1.PostgresqlSpecPatroniArgs{},
			PodAnnotations: pulumi.ToStringMap(map[string]string{
				"postgres-cluster-id": locals.PostgresKubernetes.Metadata.Name,
			}),
			Postgresql: zalandov1.PostgresqlSpecPostgresqlArgs{
				Version: pulumi.String(vars.PostgresVersion),
				Parameters: pulumi.StringMap{
					"max_connections": pulumi.String("200"),
				},
			},
			Resources: zalandov1.PostgresqlSpecResourcesArgs{
				Limits: zalandov1.PostgresqlSpecResourcesLimitsArgs{
					Cpu:    pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Limits.Cpu),
					Memory: pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Limits.Memory),
				},
				Requests: zalandov1.PostgresqlSpecResourcesRequestsArgs{
					Cpu:    pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Requests.Cpu),
					Memory: pulumi.String(locals.PostgresKubernetes.Spec.Container.Resources.Requests.Memory),
				},
			},
			TeamId: pulumi.String(vars.TeamId),
			Volume: zalandov1.PostgresqlSpecVolumeArgs{
				Size: pulumi.String(locals.PostgresKubernetes.Spec.Container.DiskSize),
			},
			// Add standby block if restore is enabled (for disaster recovery)
			Standby: standbyBlock,
			// Merge backup and restore environment variables
			Env: allEnvVars,
		},
	}

	_, err = zalandov1.NewPostgresql(ctx,
		"database",
		postgresqlArgs,
		pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create postgresql")
	}

	if locals.PostgresKubernetes.Spec.Ingress == nil ||
		!locals.PostgresKubernetes.Spec.Ingress.Enabled ||
		locals.PostgresKubernetes.Spec.Ingress.Hostname == "" {
		//if ingress is not enabled, no load-balancer resource is required. so just exit the function.
		return nil
	}

	if err := ingress(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create ingress")
	}
	return nil
}
