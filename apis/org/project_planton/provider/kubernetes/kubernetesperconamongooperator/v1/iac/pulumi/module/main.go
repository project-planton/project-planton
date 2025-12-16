package module

import (
	"github.com/pkg/errors"
	kubernetesperconamongooperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesperconamongooperator/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the Percona Operator for MongoDB Kubernetes add-on.
func Resources(ctx *pulumi.Context, stackInput *kubernetesperconamongooperatorv1.KubernetesPerconaMongoOperatorStackInput) error {
	// set up kubernetes provider from the supplied cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx, stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// use the configured chart version
	chartVersion := locals.HelmChartVersion

	// determine namespace - use from spec or default
	namespace := stackInput.Target.Spec.Namespace.GetValue()

	// conditionally create namespace resource based on create_namespace flag
	var createdNamespace *corev1.Namespace
	if stackInput.Target.Spec.CreateNamespace {
		createdNamespace, err = corev1.NewNamespace(ctx, namespace,
			&corev1.NamespaceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String(namespace),
				},
			},
			pulumi.Provider(kubeProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create namespace")
		}
	}

	// prepare helm values with resource limits from spec
	helmValues := pulumi.Map{
		// Enable cluster-wide mode to watch all namespaces
		"watchAllNamespaces": pulumi.Bool(true),
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

	// prepare helm release options
	helmOpts := []pulumi.ResourceOption{
		pulumi.Provider(kubeProvider),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}),
	}
	if createdNamespace != nil {
		helmOpts = append(helmOpts, pulumi.Parent(createdNamespace))
	}

	// deploy the operator via Helm
	_, err = helm.NewRelease(ctx, "percona-operator",
		&helm.ReleaseArgs{
			Name:            pulumi.String(locals.HelmChartName),
			Namespace:       pulumi.String(namespace),
			Chart:           pulumi.String(locals.HelmChartName),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(300),
			Values:          helmValues,
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(locals.HelmChartRepo),
			},
		},
		helmOpts...)
	if err != nil {
		return errors.Wrap(err, "failed to install percona-operator helm release")
	}

	// export stack outputs
	if createdNamespace != nil {
		ctx.Export(OpNamespace, createdNamespace.Metadata.Name())
	} else {
		ctx.Export(OpNamespace, pulumi.String(namespace))
	}

	return nil
}
