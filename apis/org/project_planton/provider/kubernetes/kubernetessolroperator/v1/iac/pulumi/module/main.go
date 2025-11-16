package module

import (
	"github.com/pkg/errors"
	kubernetessolroperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetessolroperator/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	pulumiyaml "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the Apache Solr Operator Kubernetes add‑on.
func Resources(ctx *pulumi.Context, stackInput *kubernetessolroperatorv1.KubernetesSolrOperatorStackInput) error {
	// set up kubernetes provider from the supplied cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// always install the stable chart version for now
	var chartVersion = vars.DefaultStableVersion

	// create dedicated namespace
	ns, err := corev1.NewNamespace(ctx, vars.Namespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(vars.Namespace),
			},
		},
		pulumi.Provider(kubeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// apply CRDs required by the operator
	crds, err := pulumiyaml.NewConfigFile(ctx, "solr-operator-crds",
		&pulumiyaml.ConfigFileArgs{
			File: vars.CrdManifestDownloadURL,
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns))
	if err != nil {
		return errors.Wrap(err, "failed to apply CRDs")
	}

	// deploy the operator via Helm
	_, err = helm.NewRelease(ctx, "solr-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       ns.Metadata.Name(),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values:          pulumi.Map{}, // no extra values at this time
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns),
		pulumi.DependsOn([]pulumi.Resource{crds}),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to install solr‑operator helm release")
	}

	// export stack outputs
	ctx.Export(OpNamespace, ns.Metadata.Name())

	return nil
}
