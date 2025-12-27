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
	// Initialize locals with computed values
	locals := initializeLocals(ctx, stackInput)

	// Set up kubernetes provider from the supplied cluster credential
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// --------------------------------------------------------------------
	// 1. Namespace - conditionally create based on create_namespace flag
	// --------------------------------------------------------------------
	if stackInput.Target.Spec.CreateNamespace {
		_, err := corev1.NewNamespace(ctx, locals.Namespace,
			&corev1.NamespaceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:   pulumi.String(locals.Namespace),
					Labels: pulumi.ToStringMap(locals.Labels),
					// CRITICAL: Background Deletion Propagation Policy
					//
					// This annotation prevents namespace deletion from timing out during `pulumi destroy`.
					//
					// Problem: By default, Pulumi uses "Foreground" cascading deletion for namespaces.
					// Kubernetes adds a `foregroundDeletion` finalizer and waits for all resources inside
					// the namespace to be deleted before removing the namespace itself. However, if the
					// Helm release or CRDs are being deleted concurrently, there can be race conditions
					// where finalizers on child resources (like operator-managed CRs) prevent timely cleanup.
					//
					// Solution: Using "background" propagation policy causes Kubernetes to delete the
					// namespace object immediately. The namespace controller then asynchronously cleans up
					// all resources within the namespace. This avoids blocking on child resource finalizers.
					//
					// Reference: https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/
					Annotations: pulumi.StringMap{
						"pulumi.com/deletionPropagationPolicy": pulumi.String("background"),
					},
				},
			},
			pulumi.Provider(kubernetesProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create namespace")
		}
	}

	// --------------------------------------------------------------------
	// 2. Apply CRDs required by the operator
	// Uses computed CrdsResourceName to avoid conflicts when multiple
	// instances share a namespace.
	// --------------------------------------------------------------------
	//
	// CRITICAL: Background Deletion Propagation Policy for CRDs
	//
	// The transformation below injects `pulumi.com/deletionPropagationPolicy: background`
	// into all CRD resources loaded from the remote manifest.
	//
	// Problem: CustomResourceDefinitions (CRDs) have a built-in protection mechanism where
	// they cannot be deleted while CustomResources (CRs) of that type still exist. During
	// `pulumi destroy` with foreground deletion, this can cause timeouts if:
	//   1. CRs exist in other namespaces that Pulumi isn't managing
	//   2. The operator's reconciliation loop recreates CRs during deletion
	//   3. Finalizers on CRs prevent timely cleanup
	//
	// Solution: Using "background" propagation allows the CRD deletion to proceed immediately,
	// with Kubernetes handling orphaned CRs asynchronously. This is safe for operator teardown
	// scenarios where the operator itself is being removed.
	//
	// Reference: https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/
	crds, err := pulumiyaml.NewConfigFile(ctx, locals.CrdsResourceName,
		&pulumiyaml.ConfigFileArgs{
			File: locals.CrdManifestURL,
			Transformations: []pulumiyaml.Transformation{
				// Inject background deletion policy annotation into all resources
				func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
					metadata, ok := state["metadata"].(map[string]interface{})
					if !ok {
						metadata = make(map[string]interface{})
						state["metadata"] = metadata
					}
					annotations, ok := metadata["annotations"].(map[string]interface{})
					if !ok {
						annotations = make(map[string]interface{})
						metadata["annotations"] = annotations
					}
					annotations["pulumi.com/deletionPropagationPolicy"] = "background"
				},
			},
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "failed to apply CRDs")
	}

	// --------------------------------------------------------------------
	// 3. Deploy the operator via Helm
	// Uses computed HelmReleaseName to avoid conflicts when multiple
	// instances share a namespace.
	// --------------------------------------------------------------------
	_, err = helm.NewRelease(ctx, locals.HelmReleaseName,
		&helm.ReleaseArgs{
			Name:            pulumi.String(locals.HelmReleaseName),
			Namespace:       pulumi.String(locals.Namespace),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(locals.ChartVersion),
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
		pulumi.Provider(kubernetesProvider),
		pulumi.DependsOn([]pulumi.Resource{crds}),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to install solr‑operator helm release")
	}

	return nil
}
