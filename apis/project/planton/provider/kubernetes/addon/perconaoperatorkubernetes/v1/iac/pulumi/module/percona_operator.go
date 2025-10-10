package module

import (
	"github.com/pkg/errors"
	perconaoperatorkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/perconaoperatorkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the Percona Operator for MongoDB Kubernetes add-on.
func Resources(ctx *pulumi.Context, stackInput *perconaoperatorkubernetesv1.PerconaOperatorKubernetesStackInput) error {
	// set up kubernetes provider from the supplied cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(
		ctx, stackInput.ProviderCredential, "kubernetes")
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

	// prepare helm values with resource limits from spec
	helmValues := pulumi.Map{
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
	}

	// deploy the operator via Helm
	_, err = helm.NewRelease(ctx, "percona-operator",
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
		return errors.Wrap(err, "failed to install percona-operator helm release")
	}

	// export stack outputs
	ctx.Export(OpNamespace, ns.Metadata.Name())

	return nil
}
