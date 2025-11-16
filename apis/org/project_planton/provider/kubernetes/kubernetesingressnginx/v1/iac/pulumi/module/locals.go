package module

import (
	"strconv"

	kubernetesingressnginxv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesingressnginx/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	KubernetesIngressNginx *kubernetesingressnginxv1.KubernetesIngressNginx
	Namespace              string
	ReleaseName            string
	ServiceName            string
	ServiceType            string
	ChartVersion           string
	Labels                 map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kubernetesingressnginxv1.KubernetesIngressNginxStackInput) *Locals {
	locals := &Locals{}

	locals.KubernetesIngressNginx = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesIngressNginx.String(),
	}

	if target.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	// Use fixed namespace for this addon
	locals.Namespace = vars.Namespace

	// Determine chart version
	locals.ChartVersion = target.Spec.ChartVersion
	if locals.ChartVersion == "" {
		locals.ChartVersion = vars.DefaultChartVersion
	}

	// Service names follow Helm chart defaults
	locals.ReleaseName = vars.HelmChartName
	locals.ServiceName = vars.HelmChartName + "-controller"
	locals.ServiceType = "LoadBalancer"

	// Export stack outputs
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))
	ctx.Export(OpReleaseName, pulumi.String(locals.ReleaseName))
	ctx.Export(OpServiceName, pulumi.String(locals.ServiceName))
	ctx.Export(OpServiceType, pulumi.String(locals.ServiceType))

	return locals
}
