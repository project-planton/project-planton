package module

import (
	"github.com/pkg/errors"
	strimzikafkaoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/addon/strimzikafkaoperator/v1"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// kafkaOperator installs the Strimzi Kafka Operator on the target cluster.
//
// The function:
//  1. Creates/labels the operator namespace.
//  2. Deploys the Helm chart (watch‑any‑namespace=true so one install can
//     manage topics/streams across all namespaces).
//  3. Exports the namespace name so other stacks can import it later.
func kafkaOperator(
	ctx *pulumi.Context,
	target *strimzikafkaoperatorv1.StrimziKafkaOperator,
	kubernetesProvider *pulumikubernetes.Provider,
) error {
	// ---------------------------------------------------------------------
	// 1. Namespace
	// ---------------------------------------------------------------------
	ns, err := corev1.NewNamespace(
		ctx,
		vars.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name: pulumi.String(vars.Namespace),
				// Basic labels – feel free to extend if you add a locals struct later.
				Labels: pulumi.StringMap{
					"app.kubernetes.io/managed-by": pulumi.String("project-planton"),
					"app.kubernetes.io/name":       pulumi.String(target.Metadata.Name),
				},
			}),
		},
		pulumi.Provider(kubernetesProvider),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace for Kafka operator")
	}

	// ---------------------------------------------------------------------
	// 2. Helm release
	// ---------------------------------------------------------------------
	_, err = helm.NewRelease(
		ctx,
		"strimzi-kafka-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       ns.Metadata.Name(),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(vars.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"watchAnyNamespace": pulumi.Bool(true),
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		pulumi.Parent(ns),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create Strimzi Helm release")
	}

	// ---------------------------------------------------------------------
	// 3. Stack output
	// ---------------------------------------------------------------------
	ctx.Export(OpNamespace, ns.Metadata.Name())

	return nil
}
