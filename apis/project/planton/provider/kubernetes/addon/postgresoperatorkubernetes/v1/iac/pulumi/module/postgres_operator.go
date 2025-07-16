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

	// 2. Helm release
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
			Values: pulumi.Map{
				"configKubernetes": pulumi.Map{
					"inherited_labels": pulumi.ToStringArray([]string{
						kuberneteslabelkeys.Resource,
						kuberneteslabelkeys.Organization,
						kuberneteslabelkeys.Environment,
						kuberneteslabelkeys.ResourceKind,
						kuberneteslabelkeys.ResourceId,
					}),
				},
			},
		},
		pulumi.Parent(createdNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}

	// 3. Export stack‑output(s)
	ctx.Export(OpNamespace, createdNamespace.Metadata.Name())

	return nil
}
