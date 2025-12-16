package module

import (
	"github.com/pkg/errors"
	kubernetesaltinityoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesaltinityoperator/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the Altinity ClickHouse Operator Kubernetes add-on.
func Resources(ctx *pulumi.Context, stackInput *kubernetesaltinityoperatorv1.KubernetesAltinityOperatorStackInput) error {
	// set up kubernetes provider from the supplied cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// initialize local values with computed data transformations
	locals := newLocals(stackInput)

	// conditionally create namespace if requested
	var nsName pulumi.StringInput
	var helmReleaseOpts []pulumi.ResourceOption

	if stackInput.Target.Spec.CreateNamespace {
		// create dedicated namespace
		ns, err := corev1.NewNamespace(ctx, locals.Namespace,
			&corev1.NamespaceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String(locals.Namespace),
				},
			},
			pulumi.Provider(kubeProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create namespace")
		}
		nsName = ns.Metadata.Name().Elem()
		helmReleaseOpts = []pulumi.ResourceOption{
			pulumi.Provider(kubeProvider),
			pulumi.Parent(ns),
			pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
		}
	} else {
		// use existing namespace name directly
		nsName = pulumi.String(locals.Namespace)
		helmReleaseOpts = []pulumi.ResourceOption{
			pulumi.Provider(kubeProvider),
			pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
		}
	}

	// deploy the operator via Helm
	_, err = helm.NewRelease(ctx, "kubernetes-altinity-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName),
			Namespace:       nsName,
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(vars.HelmChartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(300),
			Values:          locals.HelmValues,
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		helmReleaseOpts...)
	if err != nil {
		return errors.Wrap(err, "failed to install kubernetes-altinity-operator helm release")
	}

	// export stack outputs
	ctx.Export(OpNamespace, nsName)

	return nil
}
