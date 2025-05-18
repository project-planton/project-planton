package module

import (
	"fmt"
	temporalkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/temporalkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

// Locals keeps all frequently-used values in one place – similar to a
// Terraform “locals {}” block.
type Locals struct {
	Namespace               string
	Labels                  map[string]string
	TemporalKubernetes      *temporalkubernetesv1.TemporalKubernetes
	FrontendServiceName     string
	UIServiceName           string
	FrontendEndpoint        string
	UIEndpoint              string
	PortForwardFrontendCmd  string
	PortForwardUICmd        string
	IngressFrontendHostname string
	IngressUIHostname       string
}

// initializeLocals builds the Locals struct and immediately exports the
// values required by TemporalKubernetesStackOutputs.
func initializeLocals(ctx *pulumi.Context,
	stackInput *temporalkubernetesv1.TemporalKubernetesStackInput) *Locals {

	locals := &Locals{}
	locals.TemporalKubernetes = stackInput.Target
	target := stackInput.Target

	// -------------------------------- labels ---------------------------------
	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_TemporalKubernetes.String(),
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

	// ------------------------------- namespace -------------------------------
	locals.Namespace = target.Metadata.Name
	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// ---------------------------- service names ------------------------------
	locals.FrontendServiceName = fmt.Sprintf("%s-frontend", target.Metadata.Name)
	locals.UIServiceName = fmt.Sprintf("%s-web", target.Metadata.Name)

	ctx.Export(OpFrontendService, pulumi.String(locals.FrontendServiceName))
	ctx.Export(OpUIService, pulumi.String(locals.UIServiceName))

	// --------------------------- cluster endpoints ---------------------------
	locals.FrontendEndpoint = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.FrontendServiceName, locals.Namespace, vars.FrontendPort)
	locals.UIEndpoint = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.UIServiceName, locals.Namespace, vars.UIPort)

	ctx.Export(OpFrontendEndpoint, pulumi.String(locals.FrontendEndpoint))
	ctx.Export(OpWebUiEndpoint, pulumi.String(locals.UIEndpoint))

	// --------------------------- port-forward cmds ---------------------------
	locals.PortForwardFrontendCmd = fmt.Sprintf(
		"kubectl port-forward -n %s service/%s 7233:7233",
		locals.Namespace, locals.FrontendServiceName)
	locals.PortForwardUICmd = fmt.Sprintf(
		"kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.UIServiceName)

	ctx.Export(OpPortForwardFrontendCommand, pulumi.String(locals.PortForwardFrontendCmd))
	ctx.Export(OpPortForwardUICommand, pulumi.String(locals.PortForwardUICmd))

	// ------------------------------- ingress ---------------------------------
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.Enabled &&
		target.Spec.Ingress.DnsDomain != "" {

		locals.IngressFrontendHostname = fmt.Sprintf("%s-frontend.%s",
			locals.Namespace, target.Spec.Ingress.DnsDomain)
		locals.IngressUIHostname = fmt.Sprintf("%s-ui.%s",
			locals.Namespace, target.Spec.Ingress.DnsDomain)

		ctx.Export(OpExternalFrontendHostname, pulumi.String(locals.IngressFrontendHostname))
		ctx.Export(OpExternalUIHostname, pulumi.String(locals.IngressUIHostname))
	}

	return locals
}
