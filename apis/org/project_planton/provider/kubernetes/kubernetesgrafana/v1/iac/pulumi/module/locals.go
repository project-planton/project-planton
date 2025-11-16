package module

import (
	"fmt"
	"strconv"

	kubernetesgrafanav1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesgrafana/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	IngressExternalHostname  string
	IngressInternalHostname  string
	KubePortForwardCommand   string
	KubeServiceFqdn          string
	KubeServiceName          string
	Namespace                string
	KubernetesGrafana        *kubernetesgrafanav1.KubernetesGrafana
	GrafanaPodSelectorLabels map[string]string
	Labels                   map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kubernetesgrafanav1.KubernetesGrafanaStackInput) *Locals {
	locals := &Locals{}

	locals.KubernetesGrafana = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesGrafana.String(),
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

	// Priority order:
	// 1. Default: metadata.name
	// 2. Override with custom label if provided
	// 3. Override with stackInput if provided

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if stackInput.KubernetesNamespace != "" {
		locals.Namespace = stackInput.KubernetesNamespace
	}

	ctx.Export(Namespace, pulumi.String(locals.Namespace))

	locals.GrafanaPodSelectorLabels = map[string]string{
		"app.kubernetes.io/name":     "grafana",
		"app.kubernetes.io/instance": target.Metadata.Name,
	}

	locals.KubeServiceName = fmt.Sprintf("%s-grafana", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:80",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.Enabled {
		return locals
	}

	// Calculate external hostname
	if target.Spec.Ingress.DnsDomain != "" {
		locals.IngressExternalHostname = fmt.Sprintf("https://grafana-%s.%s",
			target.Metadata.Name, target.Spec.Ingress.DnsDomain)
		ctx.Export(ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	}

	// Calculate internal hostname
	if target.Spec.Ingress.DnsDomain != "" {
		locals.IngressInternalHostname = fmt.Sprintf("https://grafana-%s-internal.%s",
			target.Metadata.Name, target.Spec.Ingress.DnsDomain)
		ctx.Export(InternalHostname, pulumi.String(locals.IngressInternalHostname))
	}

	return locals
}
