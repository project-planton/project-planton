package module

import (
	"github.com/pkg/errors"
	altinityoperatorkubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the Altinity ClickHouse Operator Kubernetes add-on.
func Resources(ctx *pulumi.Context, stackInput *altinityoperatorkubernetesv1.AltinityOperatorKubernetesStackInput) error {
	// set up kubernetes provider from the supplied cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// use the configured chart version
	chartVersion := vars.HelmChartVersion

	// determine namespace - use from spec or default
	namespace := stackInput.Target.Spec.Namespace
	if namespace == "" {
		namespace = vars.DefaultNamespace
	}

	// create dedicated namespace
	ns, err := corev1.NewNamespace(ctx, namespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(namespace),
			},
		},
		pulumi.Provider(kubeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// prepare helm values with CRD installation enabled and resource limits from spec
	helmValues := pulumi.Map{
		"operator": pulumi.Map{
			"createCRD": pulumi.Bool(true),
			"resources": pulumi.Map{
				"limits": pulumi.Map{
					"cpu":    pulumi.String(stackInput.Target.Spec.Container.Resources.Limits.Cpu),
					"memory": pulumi.String(stackInput.Target.Spec.Container.Resources.Limits.Memory),
				},
				"requests": pulumi.Map{
					"cpu":    pulumi.String(stackInput.Target.Spec.Container.Resources.Requests.Cpu),
					"memory": pulumi.String(stackInput.Target.Spec.Container.Resources.Requests.Memory),
				},
			},
		},
		// Configure operator to watch all namespaces cluster-wide
		// Use regex pattern ".*" to match all namespaces (empty array watches only installation namespace)
		"configs": pulumi.Map{
			"files": pulumi.Map{
				"config.yaml": pulumi.Map{
					"watch": pulumi.Map{
						"namespaces": pulumi.Array{
							pulumi.String(".*"),
						},
					},
				},
			},
		},
	}

	// deploy the operator via Helm
	_, err = helm.NewRelease(ctx, "altinity-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       ns.Metadata.Name(),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(300),
			Values:          helmValues,
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to install altinity-operator helm release")
	}

	// export stack outputs
	ctx.Export(OpNamespace, ns.Metadata.Name())

	return nil
}
