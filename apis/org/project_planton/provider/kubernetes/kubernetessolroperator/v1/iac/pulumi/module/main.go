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

	// Get namespace from spec
	namespace := stackInput.Target.Spec.Namespace.GetValue()
	if namespace == "" {
		namespace = vars.Namespace // fallback to default
	}

	// --------------------------------------------------------------------
	// 1. Namespace - conditionally create based on create_namespace flag
	// --------------------------------------------------------------------
	var ns *corev1.Namespace
	var namespaceOutput pulumi.StringInput

	if stackInput.Target.Spec.CreateNamespace {
		// Create the namespace
		ns, err = corev1.NewNamespace(ctx, namespace,
			&corev1.NamespaceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String(namespace),
				},
			},
			pulumi.Provider(kubeProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create namespace")
		}
		namespaceOutput = ns.Metadata.Name().Elem()
	} else {
		// Use existing namespace - just reference the name
		namespaceOutput = pulumi.String(namespace)
	}

	// --------------------------------------------------------------------
	// 2. Apply CRDs required by the operator
	// --------------------------------------------------------------------
	crdsOpts := []pulumi.ResourceOption{
		pulumi.Provider(kubeProvider),
	}
	// Only set parent if we created the namespace
	if ns != nil {
		crdsOpts = append(crdsOpts, pulumi.Parent(ns))
	}

	crds, err := pulumiyaml.NewConfigFile(ctx, "solr-operator-crds",
		&pulumiyaml.ConfigFileArgs{
			File: vars.CrdManifestDownloadURL,
		},
		crdsOpts...)
	if err != nil {
		return errors.Wrap(err, "failed to apply CRDs")
	}

	// --------------------------------------------------------------------
	// 3. Deploy the operator via Helm
	// --------------------------------------------------------------------
	helmReleaseOpts := []pulumi.ResourceOption{
		pulumi.Provider(kubeProvider),
		pulumi.DependsOn([]pulumi.Resource{crds}),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
	}
	// Only set parent if we created the namespace
	if ns != nil {
		helmReleaseOpts = append(helmReleaseOpts, pulumi.Parent(ns))
	}

	_, err = helm.NewRelease(ctx, "solr-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       namespaceOutput,
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
		helmReleaseOpts...)
	if err != nil {
		return errors.Wrap(err, "failed to install solr‑operator helm release")
	}

	// --------------------------------------------------------------------
	// 4. Export stack outputs
	// --------------------------------------------------------------------
	ctx.Export(OpNamespace, namespaceOutput)

	return nil
}
