package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project-planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	pulumiyaml "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// solrOperator installs the Solr Operator in the Kubernetes cluster using Helm.
// It creates the necessary namespace, applies CRD resources, and deploys the Helm chart.
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
// 1. Creates a namespace for the Solr Operator and labels it with metadata from locals.
// 2. Applies the Solr Operator CRDs by downloading and adding the manifest file.
// 3. Deploys the Solr Operator Helm chart into the created namespace with specific values.
// 4. Uses Helm chart repository and version specified in the vars package.
// 5. Handles errors and returns any errors encountered during the namespace creation, CRD application, or Helm release deployment.
func solrOperator(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider *pulumikubernetes.Provider) error {
	//create namespace resource
	createdNamespace, err := corev1.NewNamespace(ctx,
		vars.SolrOperator.Namespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.SolrOperator.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create namespace")
	}

	//create solr-operator crd resources
	createdCrdsManifestFile, err := pulumiyaml.NewConfigFile(ctx, "solr-operator-crds",
		&pulumiyaml.ConfigFileArgs{
			File: vars.SolrOperator.CrdManifestDownloadUrl,
		}, pulumi.Provider(kubernetesProvider),
	)
	if err != nil {
		return errors.Wrap(err, "failed to add solr-operator crds manifest")
	}

	//create helm-release
	_, err = helm.NewRelease(ctx, "solr-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.SolrOperator.HelmChartName),
			Namespace:       createdNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.SolrOperator.HelmChartName),
			Version:         pulumi.String(vars.SolrOperator.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values:          pulumi.Map{},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.SolrOperator.HelmChartRepo),
			},
		}, pulumi.Parent(createdNamespace),
		pulumi.DependsOn([]pulumi.Resource{createdCrdsManifestFile}),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create helm release")
	}
	return nil
}
