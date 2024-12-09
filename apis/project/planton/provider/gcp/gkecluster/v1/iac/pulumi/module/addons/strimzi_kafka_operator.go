package addons

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/vars"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StrimziKafkaOperator installs the Strimzi Kafka Operator in the Kubernetes cluster using Helm.
// It creates the necessary namespace and deploys the Helm chart with specific values.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - kubernetesProvider: The Kubernetes provider for Pulumi.
//
// Returns:
// - error: An error object if there is any issue during the installation.
//
// The function performs the following steps:
// 1. Creates a namespace for the Strimzi Kafka Operator and labels it with metadata from locals.
// 2. Deploys the Strimzi Kafka Operator Helm chart into the created namespace with specific values, including enabling watching any namespace.
// 3. Uses Helm chart repository and version specified in the vars package.
// 4. Handles errors and returns any errors encountered during the namespace creation or Helm release deployment.
func StrimziKafkaOperator(ctx *pulumi.Context, locals *localz.Locals,
	kubernetesProvider *pulumikubernetes.Provider) error {
	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.StrimziKafkaOperator.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.StrimziKafkaOperator.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create namespace")
	}

	//create helm-release
	_, err = helm.NewRelease(ctx, "strimzi-kafka-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.StrimziKafkaOperator.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.StrimziKafkaOperator.HelmChartName),
			Version:         pulumi.String(vars.StrimziKafkaOperator.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"watchAnyNamespace": pulumi.Bool(true),
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.StrimziKafkaOperator.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}
	return nil
}
