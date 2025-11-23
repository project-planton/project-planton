package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// kubernetesElasticOperator installs the Elastic Cloud on Kubernetes operator.
func kubernetesElasticOperator(ctx *pulumi.Context, locals *Locals,
	k8sProvider *pulumikubernetes.Provider) error {

	// Get namespace from spec
	namespace := locals.KubernetesElasticOperator.Spec.Namespace.GetValue()
	if namespace == "" {
		namespace = vars.Namespace // fallback to default
	}

	// --------------------------------------------------------------------
	// 1. Namespace
	// --------------------------------------------------------------------
	ns, err := corev1.NewNamespace(ctx, namespace, &corev1.NamespaceArgs{
		Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
			Name:   pulumi.String(namespace),
			Labels: pulumi.ToStringMap(locals.KubeLabels),
		}),
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return errors.Wrap(err, "create namespace")
	}

	// --------------------------------------------------------------------
	// 2. Helm values – propagate Planton labels + optional resources.
	// --------------------------------------------------------------------
	values := pulumi.Map{
		"configKubernetes": pulumi.Map{
			"inherited_labels": pulumi.ToStringArray([]string{
				kuberneteslabelkeys.Resource,
				kuberneteslabelkeys.Organization,
				kuberneteslabelkeys.Environment,
				kuberneteslabelkeys.ResourceKind,
				kuberneteslabelkeys.ResourceId,
			}),
		},
	}

	if cr := locals.KubernetesElasticOperator.Spec.GetContainer().GetResources(); cr != nil {
		res := pulumi.Map{}
		if lim := cr.GetLimits(); lim != nil &&
			(lim.Cpu != "" || lim.Memory != "") {
			res["limits"] = pulumi.StringMap{
				"cpu":    pulumi.String(lim.Cpu),
				"memory": pulumi.String(lim.Memory),
			}
		}
		if req := cr.GetRequests(); req != nil &&
			(req.Cpu != "" || req.Memory != "") {
			res["requests"] = pulumi.StringMap{
				"cpu":    pulumi.String(req.Cpu),
				"memory": pulumi.String(req.Memory),
			}
		}
		if len(res) > 0 {
			values["resources"] = res
		}
	}

	// --------------------------------------------------------------------
	// 3. Helm release
	// --------------------------------------------------------------------
	_, err = helm.NewRelease(ctx, "kubernetes-elastic-operator", &helm.ReleaseArgs{
		Name:            pulumi.String(vars.HelmChartName),
		Namespace:       ns.Metadata.Name(),
		Chart:           pulumi.String(vars.HelmChartName),
		Version:         pulumi.String(vars.HelmChartVersion),
		RepositoryOpts:  helm.RepositoryOptsArgs{Repo: pulumi.String(vars.HelmChartRepo)},
		CreateNamespace: pulumi.Bool(false),
		Atomic:          pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		WaitForJobs:     pulumi.Bool(true),
		Timeout:         pulumi.Int(180),
		Values:          values,
	}, pulumi.Parent(ns),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "install helm chart")
	}

	// --------------------------------------------------------------------
	// 4. Stack outputs
	// --------------------------------------------------------------------
	ctx.Export(OpNamespace, ns.Metadata.Name())

	return nil
}
